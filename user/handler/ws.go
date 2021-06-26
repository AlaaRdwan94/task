package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func (h UserHandler) Send(c *gin.Context) {
	Upgrader.CheckOrigin = func(r *http.Request) bool{
		return true
	}
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	msg := []byte("we start our ws chat .")
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println(err)
	}
	//mt, message, err := conn.ReadMessage()
	//if err != nil {
	//	log.Println("error read message")
	//	log.Fatal(err)
	//}
}

