package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

//func(ResponseWriter, *Request)
func Index(w http.ResponseWriter, r *http.Request) {
	//设置返回值类型Content-Type为json
	w.Header().Set("Content-Type", "application/json")
	indexData := IndexData{
		Title: "go博客",
		Desc:  "欢迎来到新手教程",
	}
	//序列化将返回值转为json格式
	jsonBytes, err := json.Marshal(indexData)
	if err != nil {
		log.Println(err)
	}
	//打印到前台
	w.Write(jsonBytes)
}

func Index2(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "Title:title"
	indexData.Desc = "Desc:desc"
	t := template.New("index.html")
	//获取当前路径
	dir, _ := os.Getwd()
	t, _ = t.ParseFiles(dir + "/template/index.html")
	t.Execute(w, indexData)

}

func main() {
	//开启服务
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//处理要写在监听端口前，否则报错
	http.HandleFunc("/", Index)
	http.HandleFunc("/index", Index2)

	//监听端口
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
