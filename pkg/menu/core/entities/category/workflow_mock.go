// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package category

import (
	"context"
	"gomies/pkg/sdk/types"
	"sync"
)

// Ensure, that WorkflowMock does implement Workflow.
// If this is not the case, regenerate this file with moq.
var _ Workflow = &WorkflowMock{}

// WorkflowMock is a mock implementation of Workflow.
//
// 	func TestSomethingThatUsesWorkflow(t *testing.T) {
//
// 		// make and configure a mocked Workflow
// 		mockedWorkflow := &WorkflowMock{
// 			GetFunc: func(ctx context.Context, categoryKey Key) (Category, error) {
// 				panic("mock out the Get method")
// 			},
// 			ListFunc: func(ctx context.Context, categoryFilter Filter) ([]Category, error) {
// 				panic("mock out the List method")
// 			},
// 			RemoveFunc: func(ctx context.Context, categoryID types.External) error {
// 				panic("mock out the Remove method")
// 			},
// 			SaveFunc: func(ctx context.Context, cat Category, flag ...types.WritingFlag) (Category, error) {
// 				panic("mock out the Save method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, categoryKey Key) (Category, error)

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, categoryFilter Filter) ([]Category, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(ctx context.Context, categoryID types.External) error

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, cat Category, flag ...types.WritingFlag) (Category, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryKey is the categoryKey argument value.
			CategoryKey Key
		}
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryFilter is the categoryFilter argument value.
			CategoryFilter Filter
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryID is the categoryID argument value.
			CategoryID types.External
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cat is the cat argument value.
			Cat Category
			// Flag is the flag argument value.
			Flag []types.WritingFlag
		}
	}
	lockGet    sync.RWMutex
	lockList   sync.RWMutex
	lockRemove sync.RWMutex
	lockSave   sync.RWMutex
}

// Get calls GetFunc.
func (mock *WorkflowMock) Get(ctx context.Context, categoryKey Key) (Category, error) {
	if mock.GetFunc == nil {
		panic("WorkflowMock.GetFunc: method is nil but Workflow.Get was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		CategoryKey Key
	}{
		Ctx:         ctx,
		CategoryKey: categoryKey,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, categoryKey)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedWorkflow.GetCalls())
func (mock *WorkflowMock) GetCalls() []struct {
	Ctx         context.Context
	CategoryKey Key
} {
	var calls []struct {
		Ctx         context.Context
		CategoryKey Key
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *WorkflowMock) List(ctx context.Context, categoryFilter Filter) ([]Category, error) {
	if mock.ListFunc == nil {
		panic("WorkflowMock.ListFunc: method is nil but Workflow.List was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		CategoryFilter Filter
	}{
		Ctx:            ctx,
		CategoryFilter: categoryFilter,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, categoryFilter)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedWorkflow.ListCalls())
func (mock *WorkflowMock) ListCalls() []struct {
	Ctx            context.Context
	CategoryFilter Filter
} {
	var calls []struct {
		Ctx            context.Context
		CategoryFilter Filter
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *WorkflowMock) Remove(ctx context.Context, categoryID types.External) error {
	if mock.RemoveFunc == nil {
		panic("WorkflowMock.RemoveFunc: method is nil but Workflow.Remove was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CategoryID types.External
	}{
		Ctx:        ctx,
		CategoryID: categoryID,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(ctx, categoryID)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedWorkflow.RemoveCalls())
func (mock *WorkflowMock) RemoveCalls() []struct {
	Ctx        context.Context
	CategoryID types.External
} {
	var calls []struct {
		Ctx        context.Context
		CategoryID types.External
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *WorkflowMock) Save(ctx context.Context, cat Category, flag ...types.WritingFlag) (Category, error) {
	if mock.SaveFunc == nil {
		panic("WorkflowMock.SaveFunc: method is nil but Workflow.Save was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Cat  Category
		Flag []types.WritingFlag
	}{
		Ctx:  ctx,
		Cat:  cat,
		Flag: flag,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, cat, flag...)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedWorkflow.SaveCalls())
func (mock *WorkflowMock) SaveCalls() []struct {
	Ctx  context.Context
	Cat  Category
	Flag []types.WritingFlag
} {
	var calls []struct {
		Ctx  context.Context
		Cat  Category
		Flag []types.WritingFlag
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
