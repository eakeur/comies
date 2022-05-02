// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package transaction

import (
	"context"
	"sync"
)

// Ensure, that ManagerMock does implement Manager.
// If this is not the case, regenerate this file with moq.
var _ Manager = &ManagerMock{}

// ManagerMock is a mock implementation of Manager.
//
// 	func TestSomethingThatUsesManager(t *testing.T) {
//
// 		// make and configure a mocked Manager
// 		mockedManager := &ManagerMock{
// 			BeginFunc: func(contextMoqParam context.Context) context.Context {
// 				panic("mock out the Begin method")
// 			},
// 			CommitFunc: func(contextMoqParam context.Context)  {
// 				panic("mock out the Commit method")
// 			},
// 			EndFunc: func(ctx context.Context)  {
// 				panic("mock out the End method")
// 			},
// 			RollbackFunc: func(contextMoqParam context.Context)  {
// 				panic("mock out the Rollback method")
// 			},
// 		}
//
// 		// use mockedManager in code that requires Manager
// 		// and then make assertions.
//
// 	}
type ManagerMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func(contextMoqParam context.Context) context.Context

	// CommitFunc mocks the Commit method.
	CommitFunc func(contextMoqParam context.Context)

	// EndFunc mocks the End method.
	EndFunc func(ctx context.Context)

	// RollbackFunc mocks the Rollback method.
	RollbackFunc func(contextMoqParam context.Context)

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// Commit holds details about calls to the Commit method.
		Commit []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// End holds details about calls to the End method.
		End []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Rollback holds details about calls to the Rollback method.
		Rollback []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockBegin    sync.RWMutex
	lockCommit   sync.RWMutex
	lockEnd      sync.RWMutex
	lockRollback sync.RWMutex
}

// Begin calls BeginFunc.
func (mock *ManagerMock) Begin(contextMoqParam context.Context) context.Context {
	if mock.BeginFunc == nil {
		panic("ManagerMock.BeginFunc: method is nil but Manager.Begin was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	mock.lockBegin.Unlock()
	return mock.BeginFunc(contextMoqParam)
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedManager.BeginCalls())
func (mock *ManagerMock) BeginCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockBegin.RLock()
	calls = mock.calls.Begin
	mock.lockBegin.RUnlock()
	return calls
}

// Commit calls CommitFunc.
func (mock *ManagerMock) Commit(contextMoqParam context.Context) {
	if mock.CommitFunc == nil {
		panic("ManagerMock.CommitFunc: method is nil but Manager.Commit was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	mock.lockCommit.Unlock()
	mock.CommitFunc(contextMoqParam)
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//     len(mockedManager.CommitCalls())
func (mock *ManagerMock) CommitCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockCommit.RLock()
	calls = mock.calls.Commit
	mock.lockCommit.RUnlock()
	return calls
}

// End calls EndFunc.
func (mock *ManagerMock) End(ctx context.Context) {
	if mock.EndFunc == nil {
		panic("ManagerMock.EndFunc: method is nil but Manager.End was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockEnd.Lock()
	mock.calls.End = append(mock.calls.End, callInfo)
	mock.lockEnd.Unlock()
	mock.EndFunc(ctx)
}

// EndCalls gets all the calls that were made to End.
// Check the length with:
//     len(mockedManager.EndCalls())
func (mock *ManagerMock) EndCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockEnd.RLock()
	calls = mock.calls.End
	mock.lockEnd.RUnlock()
	return calls
}

// Rollback calls RollbackFunc.
func (mock *ManagerMock) Rollback(contextMoqParam context.Context) {
	if mock.RollbackFunc == nil {
		panic("ManagerMock.RollbackFunc: method is nil but Manager.Rollback was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockRollback.Lock()
	mock.calls.Rollback = append(mock.calls.Rollback, callInfo)
	mock.lockRollback.Unlock()
	mock.RollbackFunc(contextMoqParam)
}

// RollbackCalls gets all the calls that were made to Rollback.
// Check the length with:
//     len(mockedManager.RollbackCalls())
func (mock *ManagerMock) RollbackCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockRollback.RLock()
	calls = mock.calls.Rollback
	mock.lockRollback.RUnlock()
	return calls
}
