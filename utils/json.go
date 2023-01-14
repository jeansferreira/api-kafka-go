package utils

import "encoding/json"

func BytesToJson[T any](str []byte) T {
	var data T
	json.Unmarshal(str, &data)
	return data
}

func JsonToByes[T any](data T) []byte {
	str, _ := json.Marshal(data)
	return str
}
