package main

import (
	"errors"
	"fmt"
	"gomies/app/core/entities/content"
	"gomies/app/sdk/fault"
)

func main() {
	err := createContent()
	if errors.Is(err, fault.ErrMissingID) {
		fmt.Println(fault.Wrap(err))
	}

}
func createContent() error {
	return fault.Wrap(content.Content{Quantity: 5}.Validate())
}
