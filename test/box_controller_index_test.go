package test

/*
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/resources"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBoxController_Index(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoxRepo := mocks.NewMockBoxRepository(mockCtrl)

	// Generate sample data for testing
	boxItems := []models.BoxItem{
		{BoxID: 1, ItemID: 1, Rarity: "common"},
	}

	boxes := models.Boxes{
		{
			ID:       1,
			UUID:     uuid.New(),
			Title:    "Test Box",
			Image:    "https://example.com/image.jpg",
			AltImage: "https://example.com/alt-image.jpg",
			Price:    100,
			Active:   1,
			BoxItems: boxItems,
		},
	}

	// Mock the box repository
	mockBoxRepo.EXPECT().FindAllWithItems().Return(boxes)

	// Create a BoxController with the mock repository
	boxController := controllers.BoxController{
		BoxManager: controllers.BoxManager{
			BoxRepository: mockBoxRepo,
		},
	}

	router := gin.Default()
	router.GET("/boxes", boxController.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/boxes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var jsonResponse []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &jsonResponse)

	resource := resources.BoxesResource{Boxes: boxes}
	expectedResponse := resource.ToJSON()

	assert.Equal(t, expectedResponse, jsonResponse)
}

func TestBoxManager_FindAllWithItems(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoxRepo := mocks.NewMockBoxRepository(mockCtrl)

	// Generate sample data for testing
	boxItems := []models.BoxItem{
		{BoxID: 1, ItemID: 1, Rarity: "common"},
	}

	boxes := models.Boxes{
		{
			ID:       1,
			UUID:     uuid.New(),
			Title:    "Test Box",
			Image:    "https://example.com/image.jpg",
			AltImage: "https://example.com/alt-image.jpg",
			Price:    100,
			Active:   1,
			BoxItems: boxItems,
		},
	}

	// Mock the box repository
	mockBoxRepo.EXPECT().FindAllWithItems().Return(boxes)

	// Create a BoxManager with the mock repository
	boxManager := controllers.BoxManager{
		BoxRepository: mockBoxRepo,
	}
	// Call the method and compare the result
	result := boxManager.FindAllWithItems()
	assert.Equal(t, boxes, result)
}
*/
