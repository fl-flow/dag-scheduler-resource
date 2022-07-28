package resource

import (
  "net/http"
  "github.com/gorilla/websocket"
)


var upgrader *websocket.Upgrader


func init() {

  upgrader = &websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
      return true
    },
  }

}
