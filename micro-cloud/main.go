package main

import (
	"net/http"
	"time"
	"micro-cloud/controller"
	"fmt"
)

func main() {

	microcloud.InitDB()
	microcloud.CreateTable()

	server := &http.Server{
		Addr:        ":8080",
		Handler:     microcloud.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(microcloud.Router)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("start server error")
	}
	fmt.Println("start server success")
}

func RegiterRouter(handler *microcloud.RouterHandler) {
	new(controller.UserConterller).Router(handler)
}
