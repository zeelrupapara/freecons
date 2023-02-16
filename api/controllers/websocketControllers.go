package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (db *DatabaseControllers) GetCountOfIcons(c echo.Context) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
	}
	if ws == nil {
		c.Logger().Error("Waiting for websocket connection...")
		time.Sleep(2 * time.Second)
	}
	defer ws.Close()
	for {
		diffCounts, err := db.GetCountIconsData()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Println(diffCounts)
		bytesDiffConts, err := json.Marshal(&diffCounts)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		err = ws.WriteMessage(websocket.TextMessage, bytesDiffConts)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
