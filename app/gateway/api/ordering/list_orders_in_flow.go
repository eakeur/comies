package ordering

import (
	"comies/app/gateway/api/handler"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var ws = websocket.Upgrader{}

func (s Service) ListOrdersInFlow(ctx context.Context, w http.ResponseWriter, r *http.Request) handler.Response {
	server, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return handler.Fail(err)
	}

	defer func(server *websocket.Conn) {
		_ = server.Close()
	}(server)

	channel, err := s.ordering.Channel(ctx)
	if err != nil {
		return handler.Fail(err)
	}

	for {
		select {
		case not := <-channel:
			ord := NewOrder(not.Order)
			ord.Items = NewItemList(not.Items)

			_ = server.WriteJSON(ord)
		}

		time.Sleep(time.Second * 3)
	}
}
