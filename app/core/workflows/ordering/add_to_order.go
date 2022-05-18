package ordering

import (
	"context"
	"gomies/app/core/entities/content"
	"gomies/app/core/entities/item"
	"gomies/app/sdk/fault"
	"sync"
)

func (w workflow) AddToOrder(ctx context.Context, i item.Item, contents []content.Content) (ItemAdditionResult, error) {

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, fault.Wrap(fault.ErrMissingID)
	}

	i, err := w.items.Create(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, fault.Wrap(err)
	}

	contentNumber := len(contents)
	reservations := make(chan Reservation, contentNumber)
	failures := make(chan error, contentNumber)
	wg := sync.WaitGroup{}
	wg.Add(contentNumber)

	for _, c := range contents {
		c := c

		go func() {
			defer wg.Done()

			c, err := w.content.Create(ctx, c)
			if err != nil {
				failures <- fault.Wrap(err)
			}

			res, err := w.products.ReserveResources(ctx, i.ID, Reservation{
				ID:        i.ID,
				ProductID: c.ProductID,
				Quantity:  c.Quantity,
				Ignore:    c.Details.IgnoreIngredients,
				Replace:   c.Details.ReplaceIngredients,
			})
			if err != nil {
				failures <- fault.Wrap(err)
			}

			reservations <- res
		}()
	}

	wg.Wait()
	close(failures)
	close(reservations)

	if err := <-failures; err != nil {
		return ItemAdditionResult{}, fault.Wrap(err)
	}

	var result ItemAdditionResult
	for r := range reservations {
		if len(r.Failures) > 0 {
			result.Failed = append(result.Failed, r)
			continue
		}

		result.Succeeded = append(result.Succeeded, r)
	}

	return result, nil
}
