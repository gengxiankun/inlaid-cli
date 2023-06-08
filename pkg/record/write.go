package record

import (
	"bufio"
    "fmt"
    "os"
    "time"

    "github.com/spf13/viper"
)

func Write(project string, context string) {
	if len(context) == 0 {
		return 
	}
	// 内容预处理
	context = getTime() + " " + project + " " + context + "\r\n"
	path := viper.GetString("path")
	// 打开文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	// 写入文件
	write := bufio.NewWriter(file)
	write.WriteString(context)
	write.Flush()
}

func getTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
