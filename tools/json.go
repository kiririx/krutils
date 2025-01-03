package tools

import "encoding/json"

type JSON struct{}

func NewJSON() *JSON {
	return &JSON{}
}

func (*JSON) JSON2Map(jsonStr string) (map[string]any, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (*JSON) Map2JSON(m any) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
