package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// 获取当前目录
	root := "./img"

	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		// 如果当前路径是目录，跳过
		if info.IsDir() {
			return nil
		}
		// if _, ok := mp[path]; !ok {
		// 	os.Remove(path)
		// }

		// 打印文件路径
		fmt.Println(path)

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
