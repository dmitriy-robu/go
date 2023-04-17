package services

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go-rust-drop/config"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserInventoryService struct {
	userRepo repositories.UserRepository
}

func (uis UserInventoryService) GetInventoryForUser(userUUID *string) (inventory *models.InventoryData, err error) {
	userAuth, err := uis.userRepo.GetUserAuthByUserUUID(*userUUID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user auth")
	}

	inventory, err = uis.getInventory(userAuth.SteamUserID, config.SetSteamSettings())
	if err != nil {
		return nil, errors.Wrap(err, "Error getting inventory")
	}

	return inventory, nil
}

func (uis UserInventoryService) getInventory(steamID *string, settings config.SteamSettings) (*models.InventoryData, error) {
	var err error

	client := &http.Client{}
	endpoint := fmt.Sprintf(
		"%s/%s/%d/%d?api_key=%s",
		settings.SteamAPIs.Url,
		*steamID,
		&settings.GameInventory.AppID,
		settings.GameInventory.ContextID,
		settings.SteamAPIs.APIKey,
	)

	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting inventory")
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Error getting inventory: %s", resp.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "Error decoding response")
	}

	return uis.mapResponseToAssetData(response, settings)
}

func (uis UserInventoryService) mapResponseToAssetData(
	response map[string]interface{},
	settings config.SteamSettings,
) (*models.InventoryData, error) {
	var parsedAssets []models.AssetData

	allDetails, err := getDetailsForAllItems(settings)
	if err != nil || allDetails == nil {
		return nil, errors.New("Cannot parse prices")
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
				Amount:          amount,
				Name:            description["name"].(string),
				ClassID:         classID,
				MarketHashName:  marketHashName,
				NameColor:       nameColor,
				Tradable:        tradable,
				BackgroundColor: backgroundColor,
				IconURL:         iconURL,
				IconURLLarge:    iconURLLarge,
				Price:           price,
			})
		}
	}
	totalInventoryCount := int(response["total_inventory_count"].(float64))

	return &models.InventoryData{
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

func getDetailsForAllItems(settings config.SteamSettings) (map[string]interface{}, error) {
	client := &http.Client{}
	endpoint := fmt.Sprintf(
		"%s/%d?api_key=%s",
		"https://api.steamapis.com/market/items/",
		settings.GameInventory.AppID,
		settings.SteamAPIs.APIKey,
	)
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting items details")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Cannot close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "Error getting items details")
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "Error decoding response")
	}

	data := response["data"].(map[string]interface{})

	return data, nil
}
