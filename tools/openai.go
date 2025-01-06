package tools

import (
	"github.com/tidwall/gjson"
	"time"
)

type OpenAIResult struct {
	Err     error
	Result  string
	AIReply string
}

type OpenAI struct {
	APIAddr     string
	APIKey      string
	Model       string
	HttpProxy   string
	Prompt      string
	Timeout     int
	Temperature float32
}

func NewOpenAI(conf OpenAI) *OpenAI {
	return &OpenAI{
		APIAddr:     NewString().DefaultIfEmpty(conf.APIAddr, "https://api.openai.com"),
		APIKey:      conf.APIKey,
		Model:       conf.Model,
		HttpProxy:   conf.HttpProxy,
		Prompt:      conf.Prompt,
		Timeout:     int(NewNumber().DefaultInt64IfZero(conf.Timeout, 10)),
		Temperature: float32(NewNumber().DefaultFloat64IfZero(conf.Temperature, 0.7)),
	}
}

func (o *OpenAI) OnceTalk(yourMessage string) OpenAIResult {
	cli := NewHttpClient()
	if o.HttpProxy != "" {
		cli.Proxy(o.HttpProxy)
	}
	messageArr := make([]map[string]string, 0)
	messageArr = append(messageArr, map[string]string{"role": "system", "content": o.Prompt})
	messageArr = append(messageArr, map[string]string{"role": "user", "content": yourMessage})
	json, err := cli.Timeout(time.Second*time.Duration(o.Timeout)).Headers(map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.APIKey,
	}).PostString(o.APIAddr+"/v1/chat/completions", map[string]any{
		"model":       o.Model,
		"messages":    messageArr,
		"temperature": 0.7,
	})
	if err != nil {
		return OpenAIResult{
			Err: err,
		}
	}
	content := gjson.Get(json, "choices.0.message.content").String()

	return OpenAIResult{
		Result:  json,
		AIReply: content,
	}
}
