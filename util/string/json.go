package _string

import "encoding/json"

func MarshalToJSON(input *string) error {
	inputByte, err := json.Marshal(*input)
	if err != nil {
		return err
	}
	*input = string(inputByte)
	return nil
}
