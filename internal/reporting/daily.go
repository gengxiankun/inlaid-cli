package reporting

import (
	"time"

	"inlaid-cli/pkg/openai"
	"inlaid-cli/pkg/record"
)

func Daily(path string) (string, error) {
	today := time.Now().Format("2006-01-02")
	content, _ := record.Read(path, today, today)
	messages := []openai.Message{
		openai.Message {
			"system",
			"请针对代码提交的commit信息转化为简短的工作汇报内容，需要以日报的形式返回，包含今日完成、明日计划",
		},
		openai.Message {
			"user",
			content,
		},
	}
	reply, err := openai.Completions(messages)
	if err != nil {
		return "", err
	}
	return reply, nil
}
