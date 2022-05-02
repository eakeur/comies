// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package stock

import (
	"context"
	"gomies/app/sdk/types"
	"sync"
)

// Ensure, that ActionsMock does implement Actions.
// If this is not the case, regenerate this file with moq.
var _ Actions = &ActionsMock{}

// ActionsMock is a mock implementation of Actions.
//
// 	func TestSomethingThatUsesActions(t *testing.T) {
//
// 		// make and configure a mocked Actions
// 		mockedActions := &ActionsMock{
// 			ArchiveMovementsFunc: func(ctx context.Context, filter Filter) error {
// 				panic("mock out the ArchiveMovements method")
// 			},
// 			ComputeFunc: func(ctx context.Context, filter Filter) (types.Quantity, error) {
// 				panic("mock out the Compute method")
// 			},
// 			ComputeSomeFunc: func(ctx context.Context, filter Filter, resourceID ...types.ID) ([]types.Quantity, error) {
// 				panic("mock out the ComputeSome method")
// 			},
// 			ListMovementsFunc: func(ctx context.Context, filter Filter) ([]Movement, int, error) {
// 				panic("mock out the ListMovements method")
// 			},
// 			RemoveAllMovementsFunc: func(ctx context.Context, resourceID types.ID) error {
// 				panic("mock out the RemoveAllMovements method")
// 			},
// 			RemoveMovementFunc: func(ctx context.Context, resourceID types.ID, movementID types.ID) error {
// 				panic("mock out the RemoveMovement method")
// 			},
// 			SaveMovementsFunc: func(ctx context.Context, movement ...Movement) ([]Movement, error) {
// 				panic("mock out the SaveMovements method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// ArchiveMovementsFunc mocks the ArchiveMovements method.
	ArchiveMovementsFunc func(ctx context.Context, filter Filter) error

	// ComputeFunc mocks the Compute method.
	ComputeFunc func(ctx context.Context, filter Filter) (types.Quantity, error)

	// ComputeSomeFunc mocks the ComputeSome method.
	ComputeSomeFunc func(ctx context.Context, filter Filter, resourceID ...types.ID) ([]types.Quantity, error)

	// ListMovementsFunc mocks the ListMovements method.
	ListMovementsFunc func(ctx context.Context, filter Filter) ([]Movement, int, error)

	// RemoveAllMovementsFunc mocks the RemoveAllMovements method.
	RemoveAllMovementsFunc func(ctx context.Context, resourceID types.ID) error

	// RemoveMovementFunc mocks the RemoveMovement method.
	RemoveMovementFunc func(ctx context.Context, resourceID types.ID, movementID types.ID) error

	// SaveMovementsFunc mocks the SaveMovements method.
	SaveMovementsFunc func(ctx context.Context, movement ...Movement) ([]Movement, error)

	// calls tracks calls to the methods.
	calls struct {
		// ArchiveMovements holds details about calls to the ArchiveMovements method.
		ArchiveMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// Compute holds details about calls to the Compute method.
		Compute []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// ComputeSome holds details about calls to the ComputeSome method.
		ComputeSome []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
			// ResourceID is the resourceID argument value.
			ResourceID []types.ID
		}
		// ListMovements holds details about calls to the ListMovements method.
		ListMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// RemoveAllMovements holds details about calls to the RemoveAllMovements method.
		RemoveAllMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResourceID is the resourceID argument value.
			ResourceID types.ID
		}
		// RemoveMovement holds details about calls to the RemoveMovement method.
		RemoveMovement []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResourceID is the resourceID argument value.
			ResourceID types.ID
			// MovementID is the movementID argument value.
			MovementID types.ID
		}
		// SaveMovements holds details about calls to the SaveMovements method.
		SaveMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Movement is the movement argument value.
			Movement []Movement
		}
	}
	lockArchiveMovements   sync.RWMutex
	lockCompute            sync.RWMutex
	lockComputeSome        sync.RWMutex
	lockListMovements      sync.RWMutex
	lockRemoveAllMovements sync.RWMutex
	lockRemoveMovement     sync.RWMutex
	lockSaveMovements      sync.RWMutex
}

// ArchiveMovements calls ArchiveMovementsFunc.
func (mock *ActionsMock) ArchiveMovements(ctx context.Context, filter Filter) error {
	if mock.ArchiveMovementsFunc == nil {
		panic("ActionsMock.ArchiveMovementsFunc: method is nil but Actions.ArchiveMovements was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockArchiveMovements.Lock()
	mock.calls.ArchiveMovements = append(mock.calls.ArchiveMovements, callInfo)
	mock.lockArchiveMovements.Unlock()
	return mock.ArchiveMovementsFunc(ctx, filter)
}

// ArchiveMovementsCalls gets all the calls that were made to ArchiveMovements.
// Check the length with:
//     len(mockedActions.ArchiveMovementsCalls())
func (mock *ActionsMock) ArchiveMovementsCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockArchiveMovements.RLock()
	calls = mock.calls.ArchiveMovements
	mock.lockArchiveMovements.RUnlock()
	return calls
}

// Compute calls ComputeFunc.
func (mock *ActionsMock) Compute(ctx context.Context, filter Filter) (types.Quantity, error) {
	if mock.ComputeFunc == nil {
		panic("ActionsMock.ComputeFunc: method is nil but Actions.Compute was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockCompute.Lock()
	mock.calls.Compute = append(mock.calls.Compute, callInfo)
	mock.lockCompute.Unlock()
	return mock.ComputeFunc(ctx, filter)
}

// ComputeCalls gets all the calls that were made to Compute.
// Check the length with:
//     len(mockedActions.ComputeCalls())
func (mock *ActionsMock) ComputeCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockCompute.RLock()
	calls = mock.calls.Compute
	mock.lockCompute.RUnlock()
	return calls
}

// ComputeSome calls ComputeSomeFunc.
func (mock *ActionsMock) ComputeSome(ctx context.Context, filter Filter, resourceID ...types.ID) ([]types.Quantity, error) {
	if mock.ComputeSomeFunc == nil {
		panic("ActionsMock.ComputeSomeFunc: method is nil but Actions.ComputeSome was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Filter     Filter
		ResourceID []types.ID
	}{
		Ctx:        ctx,
		Filter:     filter,
		ResourceID: resourceID,
	}
	mock.lockComputeSome.Lock()
	mock.calls.ComputeSome = append(mock.calls.ComputeSome, callInfo)
	mock.lockComputeSome.Unlock()
	return mock.ComputeSomeFunc(ctx, filter, resourceID...)
}

// ComputeSomeCalls gets all the calls that were made to ComputeSome.
// Check the length with:
//     len(mockedActions.ComputeSomeCalls())
func (mock *ActionsMock) ComputeSomeCalls() []struct {
	Ctx        context.Context
	Filter     Filter
	ResourceID []types.ID
} {
	var calls []struct {
		Ctx        context.Context
		Filter     Filter
		ResourceID []types.ID
	}
	mock.lockComputeSome.RLock()
	calls = mock.calls.ComputeSome
	mock.lockComputeSome.RUnlock()
	return calls
}

// ListMovements calls ListMovementsFunc.
func (mock *ActionsMock) ListMovements(ctx context.Context, filter Filter) ([]Movement, int, error) {
	if mock.ListMovementsFunc == nil {
		panic("ActionsMock.ListMovementsFunc: method is nil but Actions.ListMovements was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockListMovements.Lock()
	mock.calls.ListMovements = append(mock.calls.ListMovements, callInfo)
	mock.lockListMovements.Unlock()
	return mock.ListMovementsFunc(ctx, filter)
}

// ListMovementsCalls gets all the calls that were made to ListMovements.
// Check the length with:
//     len(mockedActions.ListMovementsCalls())
func (mock *ActionsMock) ListMovementsCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockListMovements.RLock()
	calls = mock.calls.ListMovements
	mock.lockListMovements.RUnlock()
	return calls
}

// RemoveAllMovements calls RemoveAllMovementsFunc.
func (mock *ActionsMock) RemoveAllMovements(ctx context.Context, resourceID types.ID) error {
	if mock.RemoveAllMovementsFunc == nil {
		panic("ActionsMock.RemoveAllMovementsFunc: method is nil but Actions.RemoveAllMovements was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ResourceID types.ID
	}{
		Ctx:        ctx,
		ResourceID: resourceID,
	}
	mock.lockRemoveAllMovements.Lock()
	mock.calls.RemoveAllMovements = append(mock.calls.RemoveAllMovements, callInfo)
	mock.lockRemoveAllMovements.Unlock()
	return mock.RemoveAllMovementsFunc(ctx, resourceID)
}

// RemoveAllMovementsCalls gets all the calls that were made to RemoveAllMovements.
// Check the length with:
//     len(mockedActions.RemoveAllMovementsCalls())
func (mock *ActionsMock) RemoveAllMovementsCalls() []struct {
	Ctx        context.Context
	ResourceID types.ID
} {
	var calls []struct {
		Ctx        context.Context
		ResourceID types.ID
	}
	mock.lockRemoveAllMovements.RLock()
	calls = mock.calls.RemoveAllMovements
	mock.lockRemoveAllMovements.RUnlock()
	return calls
}

// RemoveMovement calls RemoveMovementFunc.
func (mock *ActionsMock) RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error {
	if mock.RemoveMovementFunc == nil {
		panic("ActionsMock.RemoveMovementFunc: method is nil but Actions.RemoveMovement was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ResourceID types.ID
		MovementID types.ID
	}{
		Ctx:        ctx,
		ResourceID: resourceID,
		MovementID: movementID,
	}
	mock.lockRemoveMovement.Lock()
	mock.calls.RemoveMovement = append(mock.calls.RemoveMovement, callInfo)
	mock.lockRemoveMovement.Unlock()
	return mock.RemoveMovementFunc(ctx, resourceID, movementID)
}

// RemoveMovementCalls gets all the calls that were made to RemoveMovement.
// Check the length with:
//     len(mockedActions.RemoveMovementCalls())
func (mock *ActionsMock) RemoveMovementCalls() []struct {
	Ctx        context.Context
	ResourceID types.ID
	MovementID types.ID
} {
	var calls []struct {
		Ctx        context.Context
		ResourceID types.ID
		MovementID types.ID
	}
	mock.lockRemoveMovement.RLock()
	calls = mock.calls.RemoveMovement
	mock.lockRemoveMovement.RUnlock()
	return calls
}

// SaveMovements calls SaveMovementsFunc.
func (mock *ActionsMock) SaveMovements(ctx context.Context, movement ...Movement) ([]Movement, error) {
	if mock.SaveMovementsFunc == nil {
		panic("ActionsMock.SaveMovementsFunc: method is nil but Actions.SaveMovements was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Movement []Movement
	}{
		Ctx:      ctx,
		Movement: movement,
	}
	mock.lockSaveMovements.Lock()
	mock.calls.SaveMovements = append(mock.calls.SaveMovements, callInfo)
	mock.lockSaveMovements.Unlock()
	return mock.SaveMovementsFunc(ctx, movement...)
}

// SaveMovementsCalls gets all the calls that were made to SaveMovements.
// Check the length with:
//     len(mockedActions.SaveMovementsCalls())
func (mock *ActionsMock) SaveMovementsCalls() []struct {
	Ctx      context.Context
	Movement []Movement
} {
	var calls []struct {
		Ctx      context.Context
		Movement []Movement
	}
	mock.lockSaveMovements.RLock()
	calls = mock.calls.SaveMovements
	mock.lockSaveMovements.RUnlock()
	return calls
}
