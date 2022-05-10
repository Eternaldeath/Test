package main

import (
	"fmt"
	"io"
	"net/http"

	"os"
)

func SpiderPage() {
	url := "https://www.zhihu.com/hot"
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	fmt.Println("结果是：", result)
	// 将读取的整个网页数据保存成一个文件
	f, err := os.Create("zhihu.html")
	if err != nil {
		fmt.Println("create err：", err)
		return
	}
	// 这里不能用 defer ，因为 defer 在函数调用完执行
	// 而这里要求每一页请求都打开关闭一次
	f.WriteString(result)
	f.Close()
}

// 循环读取网页
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("网页读取完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
