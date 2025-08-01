// Code generated by counterfeiter. DO NOT EDIT.
package modelsfakes

import (
	"context"
	"sync"

	v1a "github.com/in-toto/attestation/go/predicates/vsa/v1"
	v1 "github.com/in-toto/attestation/go/v1"
	"github.com/slsa-framework/slsa-source-poc/pkg/provenance"
	"github.com/slsa-framework/slsa-source-poc/pkg/sourcetool/models"
)

type FakeAttestationStorageReader struct {
	GetCommitProvenanceStub        func(context.Context, *models.Branch, *models.Commit) (*v1.Statement, *provenance.SourceProvenancePred, error)
	getCommitProvenanceMutex       sync.RWMutex
	getCommitProvenanceArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Branch
		arg3 *models.Commit
	}
	getCommitProvenanceReturns struct {
		result1 *v1.Statement
		result2 *provenance.SourceProvenancePred
		result3 error
	}
	getCommitProvenanceReturnsOnCall map[int]struct {
		result1 *v1.Statement
		result2 *provenance.SourceProvenancePred
		result3 error
	}
	GetCommitVsaStub        func(context.Context, *models.Branch, *models.Commit) (*v1.Statement, *v1a.VerificationSummary, error)
	getCommitVsaMutex       sync.RWMutex
	getCommitVsaArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Branch
		arg3 *models.Commit
	}
	getCommitVsaReturns struct {
		result1 *v1.Statement
		result2 *v1a.VerificationSummary
		result3 error
	}
	getCommitVsaReturnsOnCall map[int]struct {
		result1 *v1.Statement
		result2 *v1a.VerificationSummary
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAttestationStorageReader) GetCommitProvenance(arg1 context.Context, arg2 *models.Branch, arg3 *models.Commit) (*v1.Statement, *provenance.SourceProvenancePred, error) {
	fake.getCommitProvenanceMutex.Lock()
	ret, specificReturn := fake.getCommitProvenanceReturnsOnCall[len(fake.getCommitProvenanceArgsForCall)]
	fake.getCommitProvenanceArgsForCall = append(fake.getCommitProvenanceArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Branch
		arg3 *models.Commit
	}{arg1, arg2, arg3})
	stub := fake.GetCommitProvenanceStub
	fakeReturns := fake.getCommitProvenanceReturns
	fake.recordInvocation("GetCommitProvenance", []interface{}{arg1, arg2, arg3})
	fake.getCommitProvenanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAttestationStorageReader) GetCommitProvenanceCallCount() int {
	fake.getCommitProvenanceMutex.RLock()
	defer fake.getCommitProvenanceMutex.RUnlock()
	return len(fake.getCommitProvenanceArgsForCall)
}

func (fake *FakeAttestationStorageReader) GetCommitProvenanceCalls(stub func(context.Context, *models.Branch, *models.Commit) (*v1.Statement, *provenance.SourceProvenancePred, error)) {
	fake.getCommitProvenanceMutex.Lock()
	defer fake.getCommitProvenanceMutex.Unlock()
	fake.GetCommitProvenanceStub = stub
}

func (fake *FakeAttestationStorageReader) GetCommitProvenanceArgsForCall(i int) (context.Context, *models.Branch, *models.Commit) {
	fake.getCommitProvenanceMutex.RLock()
	defer fake.getCommitProvenanceMutex.RUnlock()
	argsForCall := fake.getCommitProvenanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAttestationStorageReader) GetCommitProvenanceReturns(result1 *v1.Statement, result2 *provenance.SourceProvenancePred, result3 error) {
	fake.getCommitProvenanceMutex.Lock()
	defer fake.getCommitProvenanceMutex.Unlock()
	fake.GetCommitProvenanceStub = nil
	fake.getCommitProvenanceReturns = struct {
		result1 *v1.Statement
		result2 *provenance.SourceProvenancePred
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAttestationStorageReader) GetCommitProvenanceReturnsOnCall(i int, result1 *v1.Statement, result2 *provenance.SourceProvenancePred, result3 error) {
	fake.getCommitProvenanceMutex.Lock()
	defer fake.getCommitProvenanceMutex.Unlock()
	fake.GetCommitProvenanceStub = nil
	if fake.getCommitProvenanceReturnsOnCall == nil {
		fake.getCommitProvenanceReturnsOnCall = make(map[int]struct {
			result1 *v1.Statement
			result2 *provenance.SourceProvenancePred
			result3 error
		})
	}
	fake.getCommitProvenanceReturnsOnCall[i] = struct {
		result1 *v1.Statement
		result2 *provenance.SourceProvenancePred
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAttestationStorageReader) GetCommitVsa(arg1 context.Context, arg2 *models.Branch, arg3 *models.Commit) (*v1.Statement, *v1a.VerificationSummary, error) {
	fake.getCommitVsaMutex.Lock()
	ret, specificReturn := fake.getCommitVsaReturnsOnCall[len(fake.getCommitVsaArgsForCall)]
	fake.getCommitVsaArgsForCall = append(fake.getCommitVsaArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Branch
		arg3 *models.Commit
	}{arg1, arg2, arg3})
	stub := fake.GetCommitVsaStub
	fakeReturns := fake.getCommitVsaReturns
	fake.recordInvocation("GetCommitVsa", []interface{}{arg1, arg2, arg3})
	fake.getCommitVsaMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAttestationStorageReader) GetCommitVsaCallCount() int {
	fake.getCommitVsaMutex.RLock()
	defer fake.getCommitVsaMutex.RUnlock()
	return len(fake.getCommitVsaArgsForCall)
}

func (fake *FakeAttestationStorageReader) GetCommitVsaCalls(stub func(context.Context, *models.Branch, *models.Commit) (*v1.Statement, *v1a.VerificationSummary, error)) {
	fake.getCommitVsaMutex.Lock()
	defer fake.getCommitVsaMutex.Unlock()
	fake.GetCommitVsaStub = stub
}

func (fake *FakeAttestationStorageReader) GetCommitVsaArgsForCall(i int) (context.Context, *models.Branch, *models.Commit) {
	fake.getCommitVsaMutex.RLock()
	defer fake.getCommitVsaMutex.RUnlock()
	argsForCall := fake.getCommitVsaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAttestationStorageReader) GetCommitVsaReturns(result1 *v1.Statement, result2 *v1a.VerificationSummary, result3 error) {
	fake.getCommitVsaMutex.Lock()
	defer fake.getCommitVsaMutex.Unlock()
	fake.GetCommitVsaStub = nil
	fake.getCommitVsaReturns = struct {
		result1 *v1.Statement
		result2 *v1a.VerificationSummary
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAttestationStorageReader) GetCommitVsaReturnsOnCall(i int, result1 *v1.Statement, result2 *v1a.VerificationSummary, result3 error) {
	fake.getCommitVsaMutex.Lock()
	defer fake.getCommitVsaMutex.Unlock()
	fake.GetCommitVsaStub = nil
	if fake.getCommitVsaReturnsOnCall == nil {
		fake.getCommitVsaReturnsOnCall = make(map[int]struct {
			result1 *v1.Statement
			result2 *v1a.VerificationSummary
			result3 error
		})
	}
	fake.getCommitVsaReturnsOnCall[i] = struct {
		result1 *v1.Statement
		result2 *v1a.VerificationSummary
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAttestationStorageReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAttestationStorageReader) recordInvocation(key string, args []interface{}) {
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

var _ models.AttestationStorageReader = new(FakeAttestationStorageReader)
