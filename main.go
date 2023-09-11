package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// home 页
func homePage(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{
		"code":     "0",
		"document": "how to use documentation is in confluence: /lab/test/API使用-RTC-Tokyo, or 微信 contact 李璐璐.",
		"msg":      "testing rtc between shanghai and tokyo...",
	}
	json.NewEncoder(w).Encode(msg)
}

// join channel 页
func joinChannel(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{
		"code": "101",
		"msg":  "do not support Get Request, please use Post request.",
	}
	msg1 := map[string]string{
		"code": "201",
		"msg":  "join_channel failed, please check the API documentation, or contact LuLu Lee at wechat.",
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(msg)
	} else {
		time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
		json.NewEncoder(w).Encode(msg1)
	}
}

// 接口的定义
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rtc/join_channel", joinChannel)
	http.ListenAndServe(":39527", nil)
}

// 主程序
func main() {
	handleRequests()
}
