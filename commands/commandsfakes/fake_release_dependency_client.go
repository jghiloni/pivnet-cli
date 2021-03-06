// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeReleaseDependencyClient struct {
	ListStub        func(productSlug string, releaseVersion string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	listReturns struct {
		result1 error
	}
	AddStub        func(productSlug string, releaseVersion string, dependentProductSlug string, dependentReleaseVersion string) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		productSlug             string
		releaseVersion          string
		dependentProductSlug    string
		dependentReleaseVersion string
	}
	addReturns struct {
		result1 error
	}
	RemoveStub        func(productSlug string, releaseVersion string, dependentProductSlug string, dependentReleaseVersion string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		productSlug             string
		releaseVersion          string
		dependentProductSlug    string
		dependentReleaseVersion string
	}
	removeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseDependencyClient) List(productSlug string, releaseVersion string) error {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("List", []interface{}{productSlug, releaseVersion})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug, releaseVersion)
	} else {
		return fake.listReturns.result1
	}
}

func (fake *FakeReleaseDependencyClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeReleaseDependencyClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeReleaseDependencyClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) Add(productSlug string, releaseVersion string, dependentProductSlug string, dependentReleaseVersion string) error {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		productSlug             string
		releaseVersion          string
		dependentProductSlug    string
		dependentReleaseVersion string
	}{productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion})
	fake.recordInvocation("Add", []interface{}{productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion)
	} else {
		return fake.addReturns.result1
	}
}

func (fake *FakeReleaseDependencyClient) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeReleaseDependencyClient) AddArgsForCall(i int) (string, string, string, string) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].productSlug, fake.addArgsForCall[i].releaseVersion, fake.addArgsForCall[i].dependentProductSlug, fake.addArgsForCall[i].dependentReleaseVersion
}

func (fake *FakeReleaseDependencyClient) AddReturns(result1 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) Remove(productSlug string, releaseVersion string, dependentProductSlug string, dependentReleaseVersion string) error {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		productSlug             string
		releaseVersion          string
		dependentProductSlug    string
		dependentReleaseVersion string
	}{productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion})
	fake.recordInvocation("Remove", []interface{}{productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(productSlug, releaseVersion, dependentProductSlug, dependentReleaseVersion)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeReleaseDependencyClient) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeReleaseDependencyClient) RemoveArgsForCall(i int) (string, string, string, string) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].productSlug, fake.removeArgsForCall[i].releaseVersion, fake.removeArgsForCall[i].dependentProductSlug, fake.removeArgsForCall[i].dependentReleaseVersion
}

func (fake *FakeReleaseDependencyClient) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseDependencyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReleaseDependencyClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ReleaseDependencyClient = new(FakeReleaseDependencyClient)
