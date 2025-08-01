// Code generated by counterfeiter. DO NOT EDIT.
package modelsfakes

import (
	"context"
	"sync"

	"github.com/slsa-framework/slsa-source-poc/pkg/slsa"
	"github.com/slsa-framework/slsa-source-poc/pkg/sourcetool/models"
)

type FakeVcsBackend struct {
	ConfigureControlsStub        func(*models.Repository, []*models.Branch, []models.ControlConfiguration) error
	configureControlsMutex       sync.RWMutex
	configureControlsArgsForCall []struct {
		arg1 *models.Repository
		arg2 []*models.Branch
		arg3 []models.ControlConfiguration
	}
	configureControlsReturns struct {
		result1 error
	}
	configureControlsReturnsOnCall map[int]struct {
		result1 error
	}
	ControlConfigurationDescrStub        func(*models.Branch, models.ControlConfiguration) string
	controlConfigurationDescrMutex       sync.RWMutex
	controlConfigurationDescrArgsForCall []struct {
		arg1 *models.Branch
		arg2 models.ControlConfiguration
	}
	controlConfigurationDescrReturns struct {
		result1 string
	}
	controlConfigurationDescrReturnsOnCall map[int]struct {
		result1 string
	}
	ControlPrecheckStub        func(*models.Repository, []*models.Branch, models.ControlConfiguration) (bool, string, models.ControlPreRemediationFn, error)
	controlPrecheckMutex       sync.RWMutex
	controlPrecheckArgsForCall []struct {
		arg1 *models.Repository
		arg2 []*models.Branch
		arg3 models.ControlConfiguration
	}
	controlPrecheckReturns struct {
		result1 bool
		result2 string
		result3 models.ControlPreRemediationFn
		result4 error
	}
	controlPrecheckReturnsOnCall map[int]struct {
		result1 bool
		result2 string
		result3 models.ControlPreRemediationFn
		result4 error
	}
	GetBranchControlsStub        func(context.Context, *models.Repository, *models.Branch) (*slsa.ControlSetStatus, error)
	getBranchControlsMutex       sync.RWMutex
	getBranchControlsArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
	}
	getBranchControlsReturns struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}
	getBranchControlsReturnsOnCall map[int]struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}
	GetBranchControlsAtCommitStub        func(context.Context, *models.Repository, *models.Branch, *models.Commit) (*slsa.ControlSetStatus, error)
	getBranchControlsAtCommitMutex       sync.RWMutex
	getBranchControlsAtCommitArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
		arg4 *models.Commit
	}
	getBranchControlsAtCommitReturns struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}
	getBranchControlsAtCommitReturnsOnCall map[int]struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}
	GetLatestCommitStub        func(context.Context, *models.Repository, *models.Branch) (*models.Commit, error)
	getLatestCommitMutex       sync.RWMutex
	getLatestCommitArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
	}
	getLatestCommitReturns struct {
		result1 *models.Commit
		result2 error
	}
	getLatestCommitReturnsOnCall map[int]struct {
		result1 *models.Commit
		result2 error
	}
	GetTagControlsStub        func(context.Context, *models.Tag) (*slsa.Controls, error)
	getTagControlsMutex       sync.RWMutex
	getTagControlsArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Tag
	}
	getTagControlsReturns struct {
		result1 *slsa.Controls
		result2 error
	}
	getTagControlsReturnsOnCall map[int]struct {
		result1 *slsa.Controls
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVcsBackend) ConfigureControls(arg1 *models.Repository, arg2 []*models.Branch, arg3 []models.ControlConfiguration) error {
	var arg2Copy []*models.Branch
	if arg2 != nil {
		arg2Copy = make([]*models.Branch, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []models.ControlConfiguration
	if arg3 != nil {
		arg3Copy = make([]models.ControlConfiguration, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.configureControlsMutex.Lock()
	ret, specificReturn := fake.configureControlsReturnsOnCall[len(fake.configureControlsArgsForCall)]
	fake.configureControlsArgsForCall = append(fake.configureControlsArgsForCall, struct {
		arg1 *models.Repository
		arg2 []*models.Branch
		arg3 []models.ControlConfiguration
	}{arg1, arg2Copy, arg3Copy})
	stub := fake.ConfigureControlsStub
	fakeReturns := fake.configureControlsReturns
	fake.recordInvocation("ConfigureControls", []interface{}{arg1, arg2Copy, arg3Copy})
	fake.configureControlsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVcsBackend) ConfigureControlsCallCount() int {
	fake.configureControlsMutex.RLock()
	defer fake.configureControlsMutex.RUnlock()
	return len(fake.configureControlsArgsForCall)
}

func (fake *FakeVcsBackend) ConfigureControlsCalls(stub func(*models.Repository, []*models.Branch, []models.ControlConfiguration) error) {
	fake.configureControlsMutex.Lock()
	defer fake.configureControlsMutex.Unlock()
	fake.ConfigureControlsStub = stub
}

func (fake *FakeVcsBackend) ConfigureControlsArgsForCall(i int) (*models.Repository, []*models.Branch, []models.ControlConfiguration) {
	fake.configureControlsMutex.RLock()
	defer fake.configureControlsMutex.RUnlock()
	argsForCall := fake.configureControlsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVcsBackend) ConfigureControlsReturns(result1 error) {
	fake.configureControlsMutex.Lock()
	defer fake.configureControlsMutex.Unlock()
	fake.ConfigureControlsStub = nil
	fake.configureControlsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVcsBackend) ConfigureControlsReturnsOnCall(i int, result1 error) {
	fake.configureControlsMutex.Lock()
	defer fake.configureControlsMutex.Unlock()
	fake.ConfigureControlsStub = nil
	if fake.configureControlsReturnsOnCall == nil {
		fake.configureControlsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureControlsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVcsBackend) ControlConfigurationDescr(arg1 *models.Branch, arg2 models.ControlConfiguration) string {
	fake.controlConfigurationDescrMutex.Lock()
	ret, specificReturn := fake.controlConfigurationDescrReturnsOnCall[len(fake.controlConfigurationDescrArgsForCall)]
	fake.controlConfigurationDescrArgsForCall = append(fake.controlConfigurationDescrArgsForCall, struct {
		arg1 *models.Branch
		arg2 models.ControlConfiguration
	}{arg1, arg2})
	stub := fake.ControlConfigurationDescrStub
	fakeReturns := fake.controlConfigurationDescrReturns
	fake.recordInvocation("ControlConfigurationDescr", []interface{}{arg1, arg2})
	fake.controlConfigurationDescrMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVcsBackend) ControlConfigurationDescrCallCount() int {
	fake.controlConfigurationDescrMutex.RLock()
	defer fake.controlConfigurationDescrMutex.RUnlock()
	return len(fake.controlConfigurationDescrArgsForCall)
}

func (fake *FakeVcsBackend) ControlConfigurationDescrCalls(stub func(*models.Branch, models.ControlConfiguration) string) {
	fake.controlConfigurationDescrMutex.Lock()
	defer fake.controlConfigurationDescrMutex.Unlock()
	fake.ControlConfigurationDescrStub = stub
}

func (fake *FakeVcsBackend) ControlConfigurationDescrArgsForCall(i int) (*models.Branch, models.ControlConfiguration) {
	fake.controlConfigurationDescrMutex.RLock()
	defer fake.controlConfigurationDescrMutex.RUnlock()
	argsForCall := fake.controlConfigurationDescrArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVcsBackend) ControlConfigurationDescrReturns(result1 string) {
	fake.controlConfigurationDescrMutex.Lock()
	defer fake.controlConfigurationDescrMutex.Unlock()
	fake.ControlConfigurationDescrStub = nil
	fake.controlConfigurationDescrReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVcsBackend) ControlConfigurationDescrReturnsOnCall(i int, result1 string) {
	fake.controlConfigurationDescrMutex.Lock()
	defer fake.controlConfigurationDescrMutex.Unlock()
	fake.ControlConfigurationDescrStub = nil
	if fake.controlConfigurationDescrReturnsOnCall == nil {
		fake.controlConfigurationDescrReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.controlConfigurationDescrReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVcsBackend) ControlPrecheck(arg1 *models.Repository, arg2 []*models.Branch, arg3 models.ControlConfiguration) (bool, string, models.ControlPreRemediationFn, error) {
	var arg2Copy []*models.Branch
	if arg2 != nil {
		arg2Copy = make([]*models.Branch, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.controlPrecheckMutex.Lock()
	ret, specificReturn := fake.controlPrecheckReturnsOnCall[len(fake.controlPrecheckArgsForCall)]
	fake.controlPrecheckArgsForCall = append(fake.controlPrecheckArgsForCall, struct {
		arg1 *models.Repository
		arg2 []*models.Branch
		arg3 models.ControlConfiguration
	}{arg1, arg2Copy, arg3})
	stub := fake.ControlPrecheckStub
	fakeReturns := fake.controlPrecheckReturns
	fake.recordInvocation("ControlPrecheck", []interface{}{arg1, arg2Copy, arg3})
	fake.controlPrecheckMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeVcsBackend) ControlPrecheckCallCount() int {
	fake.controlPrecheckMutex.RLock()
	defer fake.controlPrecheckMutex.RUnlock()
	return len(fake.controlPrecheckArgsForCall)
}

func (fake *FakeVcsBackend) ControlPrecheckCalls(stub func(*models.Repository, []*models.Branch, models.ControlConfiguration) (bool, string, models.ControlPreRemediationFn, error)) {
	fake.controlPrecheckMutex.Lock()
	defer fake.controlPrecheckMutex.Unlock()
	fake.ControlPrecheckStub = stub
}

func (fake *FakeVcsBackend) ControlPrecheckArgsForCall(i int) (*models.Repository, []*models.Branch, models.ControlConfiguration) {
	fake.controlPrecheckMutex.RLock()
	defer fake.controlPrecheckMutex.RUnlock()
	argsForCall := fake.controlPrecheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVcsBackend) ControlPrecheckReturns(result1 bool, result2 string, result3 models.ControlPreRemediationFn, result4 error) {
	fake.controlPrecheckMutex.Lock()
	defer fake.controlPrecheckMutex.Unlock()
	fake.ControlPrecheckStub = nil
	fake.controlPrecheckReturns = struct {
		result1 bool
		result2 string
		result3 models.ControlPreRemediationFn
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeVcsBackend) ControlPrecheckReturnsOnCall(i int, result1 bool, result2 string, result3 models.ControlPreRemediationFn, result4 error) {
	fake.controlPrecheckMutex.Lock()
	defer fake.controlPrecheckMutex.Unlock()
	fake.ControlPrecheckStub = nil
	if fake.controlPrecheckReturnsOnCall == nil {
		fake.controlPrecheckReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 string
			result3 models.ControlPreRemediationFn
			result4 error
		})
	}
	fake.controlPrecheckReturnsOnCall[i] = struct {
		result1 bool
		result2 string
		result3 models.ControlPreRemediationFn
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeVcsBackend) GetBranchControls(arg1 context.Context, arg2 *models.Repository, arg3 *models.Branch) (*slsa.ControlSetStatus, error) {
	fake.getBranchControlsMutex.Lock()
	ret, specificReturn := fake.getBranchControlsReturnsOnCall[len(fake.getBranchControlsArgsForCall)]
	fake.getBranchControlsArgsForCall = append(fake.getBranchControlsArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
	}{arg1, arg2, arg3})
	stub := fake.GetBranchControlsStub
	fakeReturns := fake.getBranchControlsReturns
	fake.recordInvocation("GetBranchControls", []interface{}{arg1, arg2, arg3})
	fake.getBranchControlsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVcsBackend) GetBranchControlsCallCount() int {
	fake.getBranchControlsMutex.RLock()
	defer fake.getBranchControlsMutex.RUnlock()
	return len(fake.getBranchControlsArgsForCall)
}

func (fake *FakeVcsBackend) GetBranchControlsCalls(stub func(context.Context, *models.Repository, *models.Branch) (*slsa.ControlSetStatus, error)) {
	fake.getBranchControlsMutex.Lock()
	defer fake.getBranchControlsMutex.Unlock()
	fake.GetBranchControlsStub = stub
}

func (fake *FakeVcsBackend) GetBranchControlsArgsForCall(i int) (context.Context, *models.Repository, *models.Branch) {
	fake.getBranchControlsMutex.RLock()
	defer fake.getBranchControlsMutex.RUnlock()
	argsForCall := fake.getBranchControlsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVcsBackend) GetBranchControlsReturns(result1 *slsa.ControlSetStatus, result2 error) {
	fake.getBranchControlsMutex.Lock()
	defer fake.getBranchControlsMutex.Unlock()
	fake.GetBranchControlsStub = nil
	fake.getBranchControlsReturns = struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetBranchControlsReturnsOnCall(i int, result1 *slsa.ControlSetStatus, result2 error) {
	fake.getBranchControlsMutex.Lock()
	defer fake.getBranchControlsMutex.Unlock()
	fake.GetBranchControlsStub = nil
	if fake.getBranchControlsReturnsOnCall == nil {
		fake.getBranchControlsReturnsOnCall = make(map[int]struct {
			result1 *slsa.ControlSetStatus
			result2 error
		})
	}
	fake.getBranchControlsReturnsOnCall[i] = struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommit(arg1 context.Context, arg2 *models.Repository, arg3 *models.Branch, arg4 *models.Commit) (*slsa.ControlSetStatus, error) {
	fake.getBranchControlsAtCommitMutex.Lock()
	ret, specificReturn := fake.getBranchControlsAtCommitReturnsOnCall[len(fake.getBranchControlsAtCommitArgsForCall)]
	fake.getBranchControlsAtCommitArgsForCall = append(fake.getBranchControlsAtCommitArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
		arg4 *models.Commit
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetBranchControlsAtCommitStub
	fakeReturns := fake.getBranchControlsAtCommitReturns
	fake.recordInvocation("GetBranchControlsAtCommit", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBranchControlsAtCommitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommitCallCount() int {
	fake.getBranchControlsAtCommitMutex.RLock()
	defer fake.getBranchControlsAtCommitMutex.RUnlock()
	return len(fake.getBranchControlsAtCommitArgsForCall)
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommitCalls(stub func(context.Context, *models.Repository, *models.Branch, *models.Commit) (*slsa.ControlSetStatus, error)) {
	fake.getBranchControlsAtCommitMutex.Lock()
	defer fake.getBranchControlsAtCommitMutex.Unlock()
	fake.GetBranchControlsAtCommitStub = stub
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommitArgsForCall(i int) (context.Context, *models.Repository, *models.Branch, *models.Commit) {
	fake.getBranchControlsAtCommitMutex.RLock()
	defer fake.getBranchControlsAtCommitMutex.RUnlock()
	argsForCall := fake.getBranchControlsAtCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommitReturns(result1 *slsa.ControlSetStatus, result2 error) {
	fake.getBranchControlsAtCommitMutex.Lock()
	defer fake.getBranchControlsAtCommitMutex.Unlock()
	fake.GetBranchControlsAtCommitStub = nil
	fake.getBranchControlsAtCommitReturns = struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetBranchControlsAtCommitReturnsOnCall(i int, result1 *slsa.ControlSetStatus, result2 error) {
	fake.getBranchControlsAtCommitMutex.Lock()
	defer fake.getBranchControlsAtCommitMutex.Unlock()
	fake.GetBranchControlsAtCommitStub = nil
	if fake.getBranchControlsAtCommitReturnsOnCall == nil {
		fake.getBranchControlsAtCommitReturnsOnCall = make(map[int]struct {
			result1 *slsa.ControlSetStatus
			result2 error
		})
	}
	fake.getBranchControlsAtCommitReturnsOnCall[i] = struct {
		result1 *slsa.ControlSetStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetLatestCommit(arg1 context.Context, arg2 *models.Repository, arg3 *models.Branch) (*models.Commit, error) {
	fake.getLatestCommitMutex.Lock()
	ret, specificReturn := fake.getLatestCommitReturnsOnCall[len(fake.getLatestCommitArgsForCall)]
	fake.getLatestCommitArgsForCall = append(fake.getLatestCommitArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Repository
		arg3 *models.Branch
	}{arg1, arg2, arg3})
	stub := fake.GetLatestCommitStub
	fakeReturns := fake.getLatestCommitReturns
	fake.recordInvocation("GetLatestCommit", []interface{}{arg1, arg2, arg3})
	fake.getLatestCommitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVcsBackend) GetLatestCommitCallCount() int {
	fake.getLatestCommitMutex.RLock()
	defer fake.getLatestCommitMutex.RUnlock()
	return len(fake.getLatestCommitArgsForCall)
}

func (fake *FakeVcsBackend) GetLatestCommitCalls(stub func(context.Context, *models.Repository, *models.Branch) (*models.Commit, error)) {
	fake.getLatestCommitMutex.Lock()
	defer fake.getLatestCommitMutex.Unlock()
	fake.GetLatestCommitStub = stub
}

func (fake *FakeVcsBackend) GetLatestCommitArgsForCall(i int) (context.Context, *models.Repository, *models.Branch) {
	fake.getLatestCommitMutex.RLock()
	defer fake.getLatestCommitMutex.RUnlock()
	argsForCall := fake.getLatestCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVcsBackend) GetLatestCommitReturns(result1 *models.Commit, result2 error) {
	fake.getLatestCommitMutex.Lock()
	defer fake.getLatestCommitMutex.Unlock()
	fake.GetLatestCommitStub = nil
	fake.getLatestCommitReturns = struct {
		result1 *models.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetLatestCommitReturnsOnCall(i int, result1 *models.Commit, result2 error) {
	fake.getLatestCommitMutex.Lock()
	defer fake.getLatestCommitMutex.Unlock()
	fake.GetLatestCommitStub = nil
	if fake.getLatestCommitReturnsOnCall == nil {
		fake.getLatestCommitReturnsOnCall = make(map[int]struct {
			result1 *models.Commit
			result2 error
		})
	}
	fake.getLatestCommitReturnsOnCall[i] = struct {
		result1 *models.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetTagControls(arg1 context.Context, arg2 *models.Tag) (*slsa.Controls, error) {
	fake.getTagControlsMutex.Lock()
	ret, specificReturn := fake.getTagControlsReturnsOnCall[len(fake.getTagControlsArgsForCall)]
	fake.getTagControlsArgsForCall = append(fake.getTagControlsArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Tag
	}{arg1, arg2})
	stub := fake.GetTagControlsStub
	fakeReturns := fake.getTagControlsReturns
	fake.recordInvocation("GetTagControls", []interface{}{arg1, arg2})
	fake.getTagControlsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVcsBackend) GetTagControlsCallCount() int {
	fake.getTagControlsMutex.RLock()
	defer fake.getTagControlsMutex.RUnlock()
	return len(fake.getTagControlsArgsForCall)
}

func (fake *FakeVcsBackend) GetTagControlsCalls(stub func(context.Context, *models.Tag) (*slsa.Controls, error)) {
	fake.getTagControlsMutex.Lock()
	defer fake.getTagControlsMutex.Unlock()
	fake.GetTagControlsStub = stub
}

func (fake *FakeVcsBackend) GetTagControlsArgsForCall(i int) (context.Context, *models.Tag) {
	fake.getTagControlsMutex.RLock()
	defer fake.getTagControlsMutex.RUnlock()
	argsForCall := fake.getTagControlsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVcsBackend) GetTagControlsReturns(result1 *slsa.Controls, result2 error) {
	fake.getTagControlsMutex.Lock()
	defer fake.getTagControlsMutex.Unlock()
	fake.GetTagControlsStub = nil
	fake.getTagControlsReturns = struct {
		result1 *slsa.Controls
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) GetTagControlsReturnsOnCall(i int, result1 *slsa.Controls, result2 error) {
	fake.getTagControlsMutex.Lock()
	defer fake.getTagControlsMutex.Unlock()
	fake.GetTagControlsStub = nil
	if fake.getTagControlsReturnsOnCall == nil {
		fake.getTagControlsReturnsOnCall = make(map[int]struct {
			result1 *slsa.Controls
			result2 error
		})
	}
	fake.getTagControlsReturnsOnCall[i] = struct {
		result1 *slsa.Controls
		result2 error
	}{result1, result2}
}

func (fake *FakeVcsBackend) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVcsBackend) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ models.VcsBackend = new(FakeVcsBackend)
