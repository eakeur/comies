package contacting

import (
	"context"
	"gomies/app/core/entities/phone"
	"gomies/app/sdk/fault"
	"sync"
)

func (w workflow) CreatePhones(ctx context.Context, phones []phone.Phone) ([]phone.Phone, error) {

	wg := sync.WaitGroup{}
	failures := make(chan error, len(phones))

	for i, p := range phones {
		i := i
		p := p
		wg.Add(1)

		go func() {
			defer wg.Done()
			ph, err := w.phones.Save(ctx, p)
			if err != nil {
				failures <- err
			}

			phones[i] = ph
		}()
	}

	wg.Wait()
	close(failures)

	if err := <-failures; err != nil {
		return nil, fault.Wrap(err)
	}

	return phones, nil
}
