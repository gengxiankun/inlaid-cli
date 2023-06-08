package openai

type ChatGPTRequestBody struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatGPTResponseBody struct {
    ID      string                 `json:"id"`
    Object  string                 `json:"object"`
    Created int                    `json:"created"`
    Model   string                 `json:"model"`
    Choices []ChoiceItem           `json:"choices"`
    Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
    Message      Message `json:"message"`
    Index        int    `json:"index"`
    FinishReason string `json:"finish_reason"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}
