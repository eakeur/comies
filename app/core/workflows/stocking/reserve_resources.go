package stocking

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"sync"
	"time"
)

func (w workflow) ReserveResources(ctx context.Context, reservationID types.ID, reservations []Reservation) ([]ReservationResult, error) {

	reservationsNumber := len(reservations)
	responses := make([]ReservationResult, reservationsNumber)
	errChan := make(chan error, reservationsNumber)

	wg := sync.WaitGroup{}
	for i, reservation := range reservations {
		i := i
		reservation := reservation
		wg.Add(1)

		go func() {

			defer wg.Done()
			res, err := w.checkResource(ctx, reservationID, reservation)
			if err != nil {
				errChan <- fault.Wrap(err).Params(map[string]interface{}{
					"reservation_id": reservationID.String(),
					"resource_id":    reservation.ResourceID.String(),
				})
			}
			responses[i] = res
		}()
	}
	wg.Wait()

	if len(errChan) > 0 {
		return nil, fault.Wrap(<-errChan)
	}

	return responses, nil

}

func (w workflow) checkResource(ctx context.Context, reservationID types.ID, reservation Reservation) (ReservationResult, error) {

	mv := movement.Movement{
		Type:     movement.ReservedMovement,
		Date:     time.Now(),
		Quantity: reservation.Quantity,
		AgentID:  reservationID,
	}

	res := ReservationResult{
		ResourceID: reservation.ResourceID,
		Want:       reservation.Quantity,
	}

	actual, err := w.movements.GetBalanceByResourceID(ctx, reservation.ResourceID, movement.Filter{})
	if err != nil {
		return ReservationResult{}, fault.Wrap(err)
	}
	res.Got = actual

	stk, err := w.stocks.GetByID(ctx, reservation.ResourceID)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err)
	}

	if err := mv.Validate(); err != nil {
		return ReservationResult{}, fault.Wrap(err)
	}

	actual += mv.Value()

	if actual < stk.MinimumQuantity {
		res.Error = stock.ErrStockEmpty
		return res, nil
	}

	_, err = w.movements.Create(ctx, mv)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err)
	}

	return res, nil
}
