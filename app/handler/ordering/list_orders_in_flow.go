package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/handler/rest"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var ws = websocket.Upgrader{}

func ListOrdersInFlow(ctx context.Context, w http.ResponseWriter, r *http.Request) rest.Response {
	server, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return rest.Fail(err)
	}

	defer func(server *websocket.Conn) {
		_ = server.Close()
	}(server)

	channel, err := ordering.Channel(ctx)
	if err != nil {
		return rest.Fail(err)
	}

	for {
		not := <-channel
		ord := NewOrder(not.Order)
		ord.Items = NewItemList(not.Items)

		_ = server.WriteJSON(ord)

		time.Sleep(time.Second * 3)
	}
}
