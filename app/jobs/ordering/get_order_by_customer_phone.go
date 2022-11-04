package ordering

import (
	"comies/app/core/ordering/order"
	"context"
)

func (w jobs) GetOrderByCustomerPhone(ctx context.Context, phone string) (order.Order, error) {
	return w.orders.GetByCustomerPhone(ctx, phone)
}
