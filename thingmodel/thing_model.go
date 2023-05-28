package thingmodel

import "encoding/json"

type ThingSpecificationLanguage struct {
	Schema     string       `json:"schema"`
	Profile    Profile      `json:"profile"`
	Properties []Properties `json:"properties"`
	Events     []Events     `json:"events"`
	Services   []Services   `json:"services"`
}

type Profile struct {
	Version    string `json:"version"`
	ProductKey string `json:"productKey"`
}

type Specs struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	DataType   DataType `json:"dataType"`
}

type DataType struct {
	Type  string      `json:"type"`
	Specs interface{} `json:"specs"`
}

type FunctionParam struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	DataType   DataType `json:"dataType"`
}

// Function 基础功能
type Function struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Required   bool   `json:"required"`
	Desc       string `json:"desc"`
}

// Properties 属性
type Properties struct {
	Function
	AccessMode string   `json:"accessMode"`
	DataType   DataType `json:"dataType"`
}

// Events 事件
type Events struct {
	Function
	Type       string        `json:"type"`
	Method     string        `json:"method"`
	OutputData []interface{} `json:"outputData"`
}

// Services 服务
type Services struct {
	CallType   string        `json:"callType"`
	Method     string        `json:"method"`
	InputData  []interface{} `json:"inputData"`
	OutputData []interface{} `json:"outputData"`
}

func NewThingSpecificationLanguage(data []byte) (*ThingSpecificationLanguage, error) {
	t := &ThingSpecificationLanguage{}
	err := json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
