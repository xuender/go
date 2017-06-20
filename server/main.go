package main

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/xuender/goutils"

	"./example"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/msg", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<html><body><script>
		var s = new WebSocket('ws://localhost:8888/ws');
		s.onopen = function(event) {
			// s.send('I am the client and I\'m listening!');
			s.onmessage = function(event) {
				console.log('Client received a message',event);
			};
			// 监听Socket的关闭
			s.onclose = function(event) {
				console.log('Client notified socket has closed',event);
			};
		};
		</script></body></html>`)
	})
	e.GET("/ws", func(c echo.Context) error {
		return wsHandler(c.Response().Writer, c.Request())
	})
	e.Logger.Fatal(e.Start(":8888"))
}

func wsHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	goutils.CheckError(err)
	go func() {
		for {
			msg := &example.Message{
				Success: true,
				Data:    "测试数据",
			}
			data, _ := proto.Marshal(msg)
			if conn.WriteMessage(websocket.BinaryMessage, data) != nil {
				break
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		// err := conn.ReadJSON(&t)
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
	return nil
}
