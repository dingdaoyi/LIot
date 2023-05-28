package thingmodel

import (
	"github.com/dingdaoyi/LIot/config"
	"os"
	"testing"
)

func TestNewThingSpecificationLanguage(t *testing.T) {
	logger := config.GetLogger()
	file, err := os.ReadFile("tmplate/aliyun/gateway.json")
	if err != nil {
		logger.Err(err)
		panic(err)
	}
	result, err := NewThingSpecificationLanguage(file)
	if err != nil {
		logger.Err(err)
		panic(err)
	}
	properties := result.Properties
	if len(properties) > 0 {
		for _, eee := range properties {
			logger.Info().Msg(eee.Name)
		}
	}
}
