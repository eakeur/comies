package ordering

import (
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var ws = websocket.Upgrader{}

func (s Service) ListOrdersInFlow(ctx context.Context, w http.ResponseWriter, r *http.Request) response.Response {
	server, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	defer func(server *websocket.Conn) {
		_ = server.Close()
	}(server)

	channel, err := s.ordering.Channel(ctx)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	for {
		select {
		case not := <-channel:
			items := make([]Item, len(not.Items))
			for ind, it := range not.Items {
				items[ind] = Item{
					ID:           it.ID,
					OrderID:      it.OrderID,
					ProductID:    it.ProductID,
					Price:        it.Price,
					Status:       it.Status,
					Quantity:     it.Quantity,
					Observations: it.Observations,
				}
			}
			fail := server.WriteJSON(Order{
				ID:             not.ID,
				Identification: not.Identification,
				PlacedAt:       not.PlacedAt,
				Observations:   not.Observations,
				FinalPrice:     not.FinalPrice,
				Address:        not.Address,
				Phone:          not.Phone,
				Items:          items,
			})
			if fail != nil {
				return failures.Handle(throw.Error(fail))
			}
		}

		time.Sleep(time.Second * 3)
	}
}
