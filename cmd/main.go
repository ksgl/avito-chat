package main

import (
	"avito-chat/api/chats"
	"avito-chat/api/messages"
	"avito-chat/api/users"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	r := fasthttprouter.New()
	r.POST("/users/add", users.Add)
	r.POST("/chats/add", chats.Add)
	r.POST("/messages/add", messages.Add)
	r.POST("/chats/get", chats.Get)
	r.POST("/messages/get", messages.Get)

	log.Println("Listening on localhost:9000...")
	fasthttp.ListenAndServe(":9000", r.Handler)
}
