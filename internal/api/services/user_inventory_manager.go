package services

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go-rust-drop/config"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserInventoryManager struct {
	userRepository repositories.UserRepository
}

func NewUserInventoryManager(ur repositories.UserRepository) UserInventoryManager {
	return UserInventoryManager{
		userRepository: ur,
	}
}

func (uis UserInventoryManager) GetInventoryForUser(userUUID string) (models.InventoryData, *utils.Errors) {
	var (
		err          error
		userAuth     models.UserAuthSteam
		inventory    models.InventoryData
		errorHandler *utils.Errors
	)

	userAuth, err = uis.userRepository.GetUserAuthByUserUUID(userUUID)
	if err != nil {
		return inventory, utils.NewErrors(http.StatusNotFound, "User not found", err)
	}

	inventory, errorHandler = uis.getInventory(userAuth.SteamUserID, config.SetSteamSettings())
	if errorHandler != nil {
		return inventory, errorHandler
	}

	return inventory, errorHandler
}

func (uis UserInventoryManager) getInventory(steamID string, settings config.SteamSettings) (models.InventoryData, *utils.Errors) {
	var (
		err          error
		inventory    models.InventoryData
		resp         *http.Response
		response     map[string]interface{}
		errorHandler *utils.Errors
		data         models.InventoryData
	)

	client := &http.Client{}
	endpoint := fmt.Sprintf(
		"%s/%s/%d/%d?api_key=%s",
		settings.SteamAPIs.Url,
		steamID,
		&settings.GameInventory.AppID,
		settings.GameInventory.ContextID,
		settings.SteamAPIs.APIKey,
	)

	resp, err = client.Get(endpoint)
	if err != nil {
		return inventory, utils.NewErrors(http.StatusInternalServerError, "Error getting inventory", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return inventory, utils.NewErrors(resp.StatusCode, "Error getting inventory", errors.New("Status code: "+strconv.Itoa(resp.StatusCode)))
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return inventory, utils.NewErrors(http.StatusInternalServerError, "Error getting inventory", err)
	}

	data, errorHandler = uis.mapResponseToAssetData(response, settings)
	if errorHandler != nil {
		return data, errorHandler
	}

	return data, nil
}

func (uis UserInventoryManager) mapResponseToAssetData(
	response map[string]interface{},
	settings config.SteamSettings,
) (models.InventoryData, *utils.Errors) {
	var (
		parsedAssets []models.AssetData
		errorHandler *utils.Errors
		allDetails   map[string]interface{}
	)

	allDetails, errorHandler = getDetailsForAllItems(settings)
	if errorHandler != nil || allDetails == nil {
		return models.InventoryData{}, errorHandler
	}

	itemsDetails := filterDetailsByDescriptions(allDetails, response["descriptions"].([]interface{}))

	assets := response["assets"].([]interface{})
	for _, asset := range assets {
		assetMap := asset.(map[string]interface{})
		classID := assetMap["classid"].(string)
		description := findDescriptionByClassID(response["descriptions"].([]interface{}), classID)

		amount, _ := strconv.Atoi(assetMap["amount"].(string))
		for i := 0; i < amount; i++ {
			assetID := fmt.Sprintf("%s%d", assetMap["assetid"].(string), i)
			marketHashName := description["market_hash_name"].(string)
			nameColor := description["name_color"].(string)
			tradable, _ := strconv.Atoi(description["tradable"].(string))
			marketable, _ := strconv.Atoi(description["marketable"].(string))

			if tradable == 0 || marketable == 0 {
				continue
			}

			backgroundColor := description["background_color"].(string)
			itemDetails := itemsDetails[marketHashName].(map[string]interface{})

			iconURL := itemDetails["image"].(string)
			iconURLLarge := iconURL
			price, _ := strconv.Atoi(fmt.Sprintf("%.0f", itemDetails["prices"].(map[string]interface{})["safe"].(float64)*100))

			parsedAssets = append(parsedAssets, models.AssetData{
				AssetID:         assetID,
				Amount:          uint(amount),
				Name:            description["name"].(string),
				ClassID:         classID,
				MarketHashName:  marketHashName,
				NameColor:       nameColor,
				Tradable:        tradable,
				BackgroundColor: backgroundColor,
				IconURL:         iconURL,
				IconURLLarge:    iconURLLarge,
				Price:           uint(price),
			})
		}
	}
	totalInventoryCount := int(response["total_inventory_count"].(float64))

	return models.InventoryData{
		AssetData:           parsedAssets,
		TotalInventoryCount: totalInventoryCount,
	}, nil
}

func filterDetailsByDescriptions(allDetails map[string]interface{}, descriptions []interface{}) map[string]interface{} {
	itemsDetails := make(map[string]interface{})
	for _, desc := range descriptions {
		marketHashName := desc.(map[string]interface{})["market_hash_name"].(string)
		if _, ok := allDetails[marketHashName]; ok {
			itemsDetails[marketHashName] = allDetails[marketHashName]
		}
	}
	return itemsDetails
}

func findDescriptionByClassID(descriptions []interface{}, classID string) map[string]interface{} {
	for _, desc := range descriptions {
		descMap := desc.(map[string]interface{})
		if descMap["classid"].(string) == classID {
			return descMap
		}
	}
	return nil
}

func getDetailsForAllItems(settings config.SteamSettings) (map[string]interface{}, *utils.Errors) {
	client := &http.Client{}
	endpoint := fmt.Sprintf(
		"%s/%d?api_key=%s",
		"https://api.steamapis.com/market/items/",
		settings.GameInventory.AppID,
		settings.SteamAPIs.APIKey,
	)
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, utils.NewErrors(http.StatusInternalServerError, "Error getting items details", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Cannot close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, utils.NewErrors(resp.StatusCode, "Error getting items details", errors.New(strconv.Itoa(resp.StatusCode)))
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, utils.NewErrors(http.StatusInternalServerError, "Error getting items details", err)
	}

	data := response["data"].(map[string]interface{})

	return data, nil
}
