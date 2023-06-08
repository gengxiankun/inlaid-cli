package record

import (
	"bufio"
    "os"
    "io"
    "strings"

    "inlaid-cli/pkg/timetools"
    "inlaid-cli/pkg/stringtools"
)

func Read(path string, startedAt string, endedAt string) (string, error) {
	var finds []string
	// 打开文件流
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 按行读取文件
	fread := bufio.NewReader(file)
	for {
		line, _, err := fread.ReadLine()
		if err == io.EOF {
			break
		}
		if find := stringtools.Contains(string(line), timetools.TraverseTime(startedAt, endedAt)); find {
			finds = append(finds, string(line))
		}
	}

	return strings.Join(finds, "\n"), nil
}
