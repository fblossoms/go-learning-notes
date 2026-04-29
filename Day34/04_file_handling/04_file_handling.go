package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05 -0700"
}

func main() {
	f, err := os.Open("logs/04file.log") // 有Read能力
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f) // 包装f提供buffer 缓冲能力
	for scanner.Scan() {
		line := scanner.Text() // 读取当前行
		data := make(map[string]any)
		err := json.Unmarshal([]byte(line), &data)
		if err != nil {
			continue
		}
		fmt.Printf("%T %[1]v", data)

	}
}
