package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

type Errors struct {
}

type MoneyConvert struct {
}

type Environment struct {
}

func (e Errors) HandleError(c *gin.Context, httpStatus int, errMsg string, err error) {
	if err != nil {
		c.JSON(httpStatus, gin.H{
			"error": errMsg,
		})
		log.Println(err)
	}
}

func (mc MoneyConvert) FromCentsToVault(value int) string {
	if value <= 0 {
		return "0"
	}

	return fmt.Sprintf("%.2f", float64(value)/100)
}

func (mc MoneyConvert) FromVaultToCents(value string) (int, error) {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return int(parsedValue * 100), nil
}

func (e Environment) GetEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
