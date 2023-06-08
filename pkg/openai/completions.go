package openai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
	"log"

	"github.com/spf13/viper"
)

func Completions(messages []Message) (string, error) {
	// 构建请求体
	requestBody, err := json.Marshal(ChatGPTRequestBody {
		Model: "gpt-3.5-turbo",
		Messages: messages,
	})

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string: %v", string(requestBody))

	// 构建请求
	req, err := http.NewRequest("POST", "https://openai.api2d.net/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	apiKey := viper.GetString(`openai.key`)
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + apiKey)

    // 发送请求
    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
        return "", errors.New(fmt.Sprintf("status code != 200, code is %d", response.StatusCode))
    }

    // 解析响应内容
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", err
    }
    responseBody := &ChatGPTResponseBody{}
    err = json.Unmarshal(body, responseBody)
    if err != nil {
        return "", err
    }
	var reply string
    if len(responseBody.Choices) > 0 {
        reply = responseBody.Choices[0].Message.Content
    }
    // log.Printf("response content: %s\n", reply)

	return reply, nil
}
