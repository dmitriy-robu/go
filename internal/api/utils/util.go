package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

type Errors struct {
	Code    int
	Message string
	Err     error
}

func NewErrors(code int, msg string, err error) *Errors {
	return &Errors{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func (e *Errors) HandleError(c *gin.Context) {
	if e.Err != nil {
		c.JSON(e.Code, gin.H{
			"error": e.Message,
		})
		log.Println(e.Err)
	}
}

type MoneyConvert struct {
}

func (mc MoneyConvert) FromCentsToVault(value uint) string {
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

type Environment struct {
}

func (e Environment) GetEnvOrDefault(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	switch defaultValue.(type) {
	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	case string:
		return value
	default:
		return defaultValue
	}
}

func GetTimeNow() time.Time {
	return time.Now()
}
