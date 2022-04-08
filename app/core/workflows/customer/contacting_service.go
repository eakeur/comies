package customer

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out contacting_service_mock.go . ContactingService:ContactingServiceMock

type (
	ContactingService interface {
		SaveContact(ctx context.Context, addresses []Address, phones []Phone) error
		RemoveAllCustomerContacts(ctx context.Context, customerUID types.UID) error
	}

	Address struct {
		ID         types.UID
		TargetID   types.UID
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
		ID          types.UID
		TargetID    types.UID
		CountryCode string
		AreaCode    string
		Number      string
	}
)
