package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// 定义要替换的旧路径和新路径的正则表达式
	oldPathPattern := regexp.MustCompile(`Go语言基础语法\.assets/([^ ]+)`)
	newPathPrefix := "images/"

	// 指定要处理的目录（相对路径）
	targetDir := "./kuangshen"

	// 获取指定目录下所有 Markdown 文件
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理 .md 文件
		if filepath.Ext(path) == ".md" {
			// 读取文件内容
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// 替换路径
			newContent := oldPathPattern.ReplaceAllString(string(content), newPathPrefix+"$1")

			// 写回文件
			err = ioutil.WriteFile(path, []byte(newContent), info.Mode())
			if err != nil {
				return err
			}

			fmt.Printf("已更新文件: %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历文件时出错:", err)
	}
}
