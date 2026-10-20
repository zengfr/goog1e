package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonEncode(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonByte)
}
func JsonDecode(jsonString string) interface{} {
	defer func() interface{} {
		return make(map[string]interface{})
	}()
	jsonResult := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(jsonString), &jsonResult)
	if err1 != nil {
		panic(err1.Error() + "jsonString:" + jsonString)
	}
	return jsonResult
}
func JsonDecodeWithType(jsonString string, jsonResult interface{}) error {
	defer func() interface{} {
		return make(map[string]interface{})
	}()
	err1 := json.Unmarshal([]byte(jsonString), &jsonResult)
	if err1 != nil {
		return err1
	}
	return nil
}

func JsonDecodeWithErr(jsonString string) (interface{}, error) {
	defer func() interface{} {
		return make(map[string]interface{})
	}()
	jsonResult := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(jsonString), &jsonResult)
	if err1 != nil {
		return nil, err1
	}
	return jsonResult, nil
}

func FormatJsonStringfy(jsonParams interface{}) string {
	var out bytes.Buffer
	b, _ := json.Marshal(jsonParams)
	json.Indent(&out, b, "", "\t")
	return out.String()
}
func PrintJson(params interface{}) {
	fmt.Println(FormatJsonStringfy(params))
}
