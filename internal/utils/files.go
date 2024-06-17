package utils

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
)

func ReadFileBytes(path string) (result []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	result, err = io.ReadAll(file)

	return
}

func ParseXmlResponseFromFile(path string, targetPtr interface{}) error {
	byteValue, err := ReadFileBytes(path)
	if err != nil {
		return err
	}

	return xml.Unmarshal(byteValue, targetPtr)
}

func ParseJsonResponseFromFile(path string, targetPtr interface{}) error {
	byteValue, err := ReadFileBytes(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, targetPtr)
}
