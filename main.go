package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr = flag.String("addr", ":9876", "-addr=:8080")

func main() {

	flag.Parse()


	fmt.Println("启动...", *addr)

	// 只处理一个hello
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Sprintln("启动失败", err)
	}

	fmt.Println("退出程序")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}


