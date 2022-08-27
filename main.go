// Package main
// @Description: 主程序入口
package main

import (
	"github.com/sivanWu0222/GoAppTemplate/pkg/logging"
)

func init() {
	logging.InitLogger()
	logging.SetLevel(logging.InfoLevel)
}

func main() {
	logging.Error("123")
	logging.InfoWithFields(map[string]interface{}{
		"err": "wrong answer",
		"url": "http://localhost:8080",
	}, "请求成功")

}
