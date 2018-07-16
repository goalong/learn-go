package main

import (
	"fmt"
	"net/http"


	"github.com/goalong/learn-go/gin_blog/pkg/setting"
	"github.com/goalong/learn-go/gin_blog/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Port),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}