package jsonx

import "encoding/json"

func JSON2Map(jsonStr string) (map[string]any, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Map2JSON[T any](m map[string]T) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
