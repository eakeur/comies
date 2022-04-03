package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"sync"
)

func (w workflow) SaveContact(ctx context.Context, contact contacting.Contact) (contacting.Contact, error) {
	const operation = "Workflows.Contacting.SaveContact"

	wg := sync.WaitGroup{}
	var err error
	if contact.Phones != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			phones, perr := w.contacts.SavePhones(ctx, contact.TargetID, contact.Phones...)
			if perr != nil {
				perr = err
			}
			contact.Phones = phones
		}()
	}
	if contact.Addresses != nil {
		wg.Add(1)
		go func() {

			defer wg.Done()
			addresses, aerr := w.contacts.SaveAddresses(ctx, contact.TargetID, contact.Addresses...)
			if aerr != nil {
				err = aerr
			}
			contact.Addresses = addresses

		}()
	}
	wg.Wait()

	if err != nil {
		return contacting.Contact{}, fault.Wrap(err, operation)
	}

	return contact, nil
}
