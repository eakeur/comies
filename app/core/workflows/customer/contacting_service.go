package customer

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out contacting_service_mock.go . ContactingService:ContactingServiceMock

type (
	ContactingService interface {
		SaveContact(ctx context.Context, addresses []Address, phones []Phone) error
		RemoveAllCustomerContacts(ctx context.Context, customerUID types.ID) error
	}

	Address struct {
		ID         types.ID
		TargetID   types.ID
		Code       string
		Street     string
		Number     string
		Complement string
		District   string
		City       string
		State      string
		Country    string
	}

	Phone struct {
		ID          types.ID
		TargetID    types.ID
		CountryCode string
		AreaCode    string
		Number      string
	}
)
