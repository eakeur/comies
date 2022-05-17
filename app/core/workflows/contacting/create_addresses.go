package contacting

import (
	"context"
	"gomies/app/core/entities/address"
	"gomies/app/sdk/fault"
	"sync"
)

func (w workflow) CreateAddresses(ctx context.Context, addresses []address.Address) ([]address.Address, error) {

	wg := sync.WaitGroup{}
	failures := make(chan error, len(addresses))

	for i, a := range addresses {
		i := i
		a := a
		wg.Add(1)

		go func() {
			defer wg.Done()
			addr, err := w.addresses.Save(ctx, a)
			if err != nil {
				failures <- err
			}

			addresses[i] = addr
		}()
	}

	wg.Wait()
	close(failures)

	if err := <-failures; err != nil {
		return nil, fault.Wrap(err)
	}

	return addresses, nil
}
