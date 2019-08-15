package main

import (
	"github.com/micro/go-micro/util/log"
	"net/http"

	"./aisitar-micro/user-web/handler"
	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/call", handler.UserCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
