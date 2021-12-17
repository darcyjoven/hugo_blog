package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 循环content中post，新建一笔文章
func main() {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	fileName := fmt.Sprintf("%d-%d-%d-", time.Now().Year(), time.Now().Month(), time.Now().Day())
	var maxIndex = 1
	fmt.Println(filepath.Join(filepath.Dir(path), "content", "post"))
	filepath.Walk(filepath.Join(filepath.Dir(path), "content", "post"), func(path string, info fs.FileInfo, err error) error {
		//遍历post
		// fmt.Println(info.Name())
		if strings.Index(info.Name(), fileName) > -1 {
			a, _ := strconv.Atoi(strings.Replace(strings.Replace(info.Name(), fileName, "", 1), ".md", "", 1))
			if a >= maxIndex {
				maxIndex = a + 1
			}
		}
		return nil
	})
	fileName = fmt.Sprintf("%s%d.md", fileName, maxIndex)

	text := fmt.Sprintf(`---
title: "%s"
date:  %s
lastmod:  %s
draft: false
tags: [ ]
categories: [ ]
author: "darcy"

contentCopyright: '<a rel="license noopener" href="https://en.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License" target="_blank">Creative Commons Attribution-ShareAlike License</a>'

---`,
		strings.Replace(fileName, ".md", "", 1),
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
	)

	file := filepath.Join(filepath.Dir(path), "content", "post", fileName)
	fmt.Println(file)
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.Write([]byte(text))
	if err != nil {
		fmt.Println(err)
	}

}
