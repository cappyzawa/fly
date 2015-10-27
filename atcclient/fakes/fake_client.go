// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/fly/atcclient"
)

type FakeClient struct {
	AllBuildsStub        func() ([]atc.Build, error)
	allBuildsMutex       sync.RWMutex
	allBuildsArgsForCall []struct{}
	allBuildsReturns     struct {
		result1 []atc.Build
		result2 error
	}
	BuildStub        func(buildID string) (atc.Build, bool, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		buildID string
	}
	buildReturns struct {
		result1 atc.Build
		result2 bool
		result3 error
	}
	BuildEventsStub        func(buildID string) (atcclient.Events, error)
	buildEventsMutex       sync.RWMutex
	buildEventsArgsForCall []struct {
		buildID string
	}
	buildEventsReturns struct {
		result1 atcclient.Events
		result2 error
	}
	AbortBuildStub        func(buildID string) error
	abortBuildMutex       sync.RWMutex
	abortBuildArgsForCall []struct {
		buildID string
	}
	abortBuildReturns struct {
		result1 error
	}
	BuildInputsForJobStub        func(pipelineName string, jobName string) ([]atc.BuildInput, bool, error)
	buildInputsForJobMutex       sync.RWMutex
	buildInputsForJobArgsForCall []struct {
		pipelineName string
		jobName      string
	}
	buildInputsForJobReturns struct {
		result1 []atc.BuildInput
		result2 bool
		result3 error
	}
	CreateBuildStub        func(plan atc.Plan) (atc.Build, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct {
		plan atc.Plan
	}
	createBuildReturns struct {
		result1 atc.Build
		result2 error
	}
	CreateOrUpdatePipelineConfigStub        func(pipelineName string, configVersion string, passedConfig atc.Config, paused *bool) (bool, bool, error)
	createOrUpdatePipelineConfigMutex       sync.RWMutex
	createOrUpdatePipelineConfigArgsForCall []struct {
		pipelineName  string
		configVersion string
		passedConfig  atc.Config
		paused        *bool
	}
	createOrUpdatePipelineConfigReturns struct {
		result1 bool
		result2 bool
		result3 error
	}
	CreatePipeStub        func() (atc.Pipe, error)
	createPipeMutex       sync.RWMutex
	createPipeArgsForCall []struct{}
	createPipeReturns     struct {
		result1 atc.Pipe
		result2 error
	}
	DeletePipelineStub        func(pipelineName string) (bool, error)
	deletePipelineMutex       sync.RWMutex
	deletePipelineArgsForCall []struct {
		pipelineName string
	}
	deletePipelineReturns struct {
		result1 bool
		result2 error
	}
	PausePipelineStub        func(pipelineName string) (bool, error)
	pausePipelineMutex       sync.RWMutex
	pausePipelineArgsForCall []struct {
		pipelineName string
	}
	pausePipelineReturns struct {
		result1 bool
		result2 error
	}
	UnpausePipelineStub        func(pipelineName string) (bool, error)
	unpausePipelineMutex       sync.RWMutex
	unpausePipelineArgsForCall []struct {
		pipelineName string
	}
	unpausePipelineReturns struct {
		result1 bool
		result2 error
	}
	JobStub        func(pipelineName, jobName string) (atc.Job, bool, error)
	jobMutex       sync.RWMutex
	jobArgsForCall []struct {
		pipelineName string
		jobName      string
	}
	jobReturns struct {
		result1 atc.Job
		result2 bool
		result3 error
	}
	JobBuildStub        func(pipelineName, jobName, buildName string) (atc.Build, bool, error)
	jobBuildMutex       sync.RWMutex
	jobBuildArgsForCall []struct {
		pipelineName string
		jobName      string
		buildName    string
	}
	jobBuildReturns struct {
		result1 atc.Build
		result2 bool
		result3 error
	}
	ListContainersStub        func(queryList map[string]string) ([]atc.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct {
		queryList map[string]string
	}
	listContainersReturns struct {
		result1 []atc.Container
		result2 error
	}
	ListPipelinesStub        func() ([]atc.Pipeline, error)
	listPipelinesMutex       sync.RWMutex
	listPipelinesArgsForCall []struct{}
	listPipelinesReturns     struct {
		result1 []atc.Pipeline
		result2 error
	}
	ListVolumesStub        func() ([]atc.Volume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []atc.Volume
		result2 error
	}
	ListWorkersStub        func() ([]atc.Worker, error)
	listWorkersMutex       sync.RWMutex
	listWorkersArgsForCall []struct{}
	listWorkersReturns     struct {
		result1 []atc.Worker
		result2 error
	}
	PipelineConfigStub        func(pipelineName string) (atc.Config, string, bool, error)
	pipelineConfigMutex       sync.RWMutex
	pipelineConfigArgsForCall []struct {
		pipelineName string
	}
	pipelineConfigReturns struct {
		result1 atc.Config
		result2 string
		result3 bool
		result4 error
	}
	GetCLIReaderStub        func(arch, platform string) (io.ReadCloser, error)
	getCLIReaderMutex       sync.RWMutex
	getCLIReaderArgsForCall []struct {
		arch     string
		platform string
	}
	getCLIReaderReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	ListAuthMethodsStub        func() ([]atc.AuthMethod, error)
	listAuthMethodsMutex       sync.RWMutex
	listAuthMethodsArgsForCall []struct{}
	listAuthMethodsReturns     struct {
		result1 []atc.AuthMethod
		result2 error
	}
	AuthTokenStub        func() (atc.AuthToken, error)
	authTokenMutex       sync.RWMutex
	authTokenArgsForCall []struct{}
	authTokenReturns     struct {
		result1 atc.AuthToken
		result2 error
	}
}

func (fake *FakeClient) AllBuilds() ([]atc.Build, error) {
	fake.allBuildsMutex.Lock()
	fake.allBuildsArgsForCall = append(fake.allBuildsArgsForCall, struct{}{})
	fake.allBuildsMutex.Unlock()
	if fake.AllBuildsStub != nil {
		return fake.AllBuildsStub()
	} else {
		return fake.allBuildsReturns.result1, fake.allBuildsReturns.result2
	}
}

func (fake *FakeClient) AllBuildsCallCount() int {
	fake.allBuildsMutex.RLock()
	defer fake.allBuildsMutex.RUnlock()
	return len(fake.allBuildsArgsForCall)
}

func (fake *FakeClient) AllBuildsReturns(result1 []atc.Build, result2 error) {
	fake.AllBuildsStub = nil
	fake.allBuildsReturns = struct {
		result1 []atc.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Build(buildID string) (atc.Build, bool, error) {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(buildID)
	} else {
		return fake.buildReturns.result1, fake.buildReturns.result2, fake.buildReturns.result3
	}
}

func (fake *FakeClient) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeClient) BuildArgsForCall(i int) string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].buildID
}

func (fake *FakeClient) BuildReturns(result1 atc.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 atc.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildEvents(buildID string) (atcclient.Events, error) {
	fake.buildEventsMutex.Lock()
	fake.buildEventsArgsForCall = append(fake.buildEventsArgsForCall, struct {
		buildID string
	}{buildID})
	fake.buildEventsMutex.Unlock()
	if fake.BuildEventsStub != nil {
		return fake.BuildEventsStub(buildID)
	} else {
		return fake.buildEventsReturns.result1, fake.buildEventsReturns.result2
	}
}

func (fake *FakeClient) BuildEventsCallCount() int {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return len(fake.buildEventsArgsForCall)
}

func (fake *FakeClient) BuildEventsArgsForCall(i int) string {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return fake.buildEventsArgsForCall[i].buildID
}

func (fake *FakeClient) BuildEventsReturns(result1 atcclient.Events, result2 error) {
	fake.BuildEventsStub = nil
	fake.buildEventsReturns = struct {
		result1 atcclient.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AbortBuild(buildID string) error {
	fake.abortBuildMutex.Lock()
	fake.abortBuildArgsForCall = append(fake.abortBuildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.abortBuildMutex.Unlock()
	if fake.AbortBuildStub != nil {
		return fake.AbortBuildStub(buildID)
	} else {
		return fake.abortBuildReturns.result1
	}
}

func (fake *FakeClient) AbortBuildCallCount() int {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return len(fake.abortBuildArgsForCall)
}

func (fake *FakeClient) AbortBuildArgsForCall(i int) string {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return fake.abortBuildArgsForCall[i].buildID
}

func (fake *FakeClient) AbortBuildReturns(result1 error) {
	fake.AbortBuildStub = nil
	fake.abortBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) BuildInputsForJob(pipelineName string, jobName string) ([]atc.BuildInput, bool, error) {
	fake.buildInputsForJobMutex.Lock()
	fake.buildInputsForJobArgsForCall = append(fake.buildInputsForJobArgsForCall, struct {
		pipelineName string
		jobName      string
	}{pipelineName, jobName})
	fake.buildInputsForJobMutex.Unlock()
	if fake.BuildInputsForJobStub != nil {
		return fake.BuildInputsForJobStub(pipelineName, jobName)
	} else {
		return fake.buildInputsForJobReturns.result1, fake.buildInputsForJobReturns.result2, fake.buildInputsForJobReturns.result3
	}
}

func (fake *FakeClient) BuildInputsForJobCallCount() int {
	fake.buildInputsForJobMutex.RLock()
	defer fake.buildInputsForJobMutex.RUnlock()
	return len(fake.buildInputsForJobArgsForCall)
}

func (fake *FakeClient) BuildInputsForJobArgsForCall(i int) (string, string) {
	fake.buildInputsForJobMutex.RLock()
	defer fake.buildInputsForJobMutex.RUnlock()
	return fake.buildInputsForJobArgsForCall[i].pipelineName, fake.buildInputsForJobArgsForCall[i].jobName
}

func (fake *FakeClient) BuildInputsForJobReturns(result1 []atc.BuildInput, result2 bool, result3 error) {
	fake.BuildInputsForJobStub = nil
	fake.buildInputsForJobReturns = struct {
		result1 []atc.BuildInput
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) CreateBuild(plan atc.Plan) (atc.Build, error) {
	fake.createBuildMutex.Lock()
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct {
		plan atc.Plan
	}{plan})
	fake.createBuildMutex.Unlock()
	if fake.CreateBuildStub != nil {
		return fake.CreateBuildStub(plan)
	} else {
		return fake.createBuildReturns.result1, fake.createBuildReturns.result2
	}
}

func (fake *FakeClient) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *FakeClient) CreateBuildArgsForCall(i int) atc.Plan {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return fake.createBuildArgsForCall[i].plan
}

func (fake *FakeClient) CreateBuildReturns(result1 atc.Build, result2 error) {
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 atc.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateOrUpdatePipelineConfig(pipelineName string, configVersion string, passedConfig atc.Config, paused *bool) (bool, bool, error) {
	fake.createOrUpdatePipelineConfigMutex.Lock()
	fake.createOrUpdatePipelineConfigArgsForCall = append(fake.createOrUpdatePipelineConfigArgsForCall, struct {
		pipelineName  string
		configVersion string
		passedConfig  atc.Config
		paused        *bool
	}{pipelineName, configVersion, passedConfig, paused})
	fake.createOrUpdatePipelineConfigMutex.Unlock()
	if fake.CreateOrUpdatePipelineConfigStub != nil {
		return fake.CreateOrUpdatePipelineConfigStub(pipelineName, configVersion, passedConfig, paused)
	} else {
		return fake.createOrUpdatePipelineConfigReturns.result1, fake.createOrUpdatePipelineConfigReturns.result2, fake.createOrUpdatePipelineConfigReturns.result3
	}
}

func (fake *FakeClient) CreateOrUpdatePipelineConfigCallCount() int {
	fake.createOrUpdatePipelineConfigMutex.RLock()
	defer fake.createOrUpdatePipelineConfigMutex.RUnlock()
	return len(fake.createOrUpdatePipelineConfigArgsForCall)
}

func (fake *FakeClient) CreateOrUpdatePipelineConfigArgsForCall(i int) (string, string, atc.Config, *bool) {
	fake.createOrUpdatePipelineConfigMutex.RLock()
	defer fake.createOrUpdatePipelineConfigMutex.RUnlock()
	return fake.createOrUpdatePipelineConfigArgsForCall[i].pipelineName, fake.createOrUpdatePipelineConfigArgsForCall[i].configVersion, fake.createOrUpdatePipelineConfigArgsForCall[i].passedConfig, fake.createOrUpdatePipelineConfigArgsForCall[i].paused
}

func (fake *FakeClient) CreateOrUpdatePipelineConfigReturns(result1 bool, result2 bool, result3 error) {
	fake.CreateOrUpdatePipelineConfigStub = nil
	fake.createOrUpdatePipelineConfigReturns = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) CreatePipe() (atc.Pipe, error) {
	fake.createPipeMutex.Lock()
	fake.createPipeArgsForCall = append(fake.createPipeArgsForCall, struct{}{})
	fake.createPipeMutex.Unlock()
	if fake.CreatePipeStub != nil {
		return fake.CreatePipeStub()
	} else {
		return fake.createPipeReturns.result1, fake.createPipeReturns.result2
	}
}

func (fake *FakeClient) CreatePipeCallCount() int {
	fake.createPipeMutex.RLock()
	defer fake.createPipeMutex.RUnlock()
	return len(fake.createPipeArgsForCall)
}

func (fake *FakeClient) CreatePipeReturns(result1 atc.Pipe, result2 error) {
	fake.CreatePipeStub = nil
	fake.createPipeReturns = struct {
		result1 atc.Pipe
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeletePipeline(pipelineName string) (bool, error) {
	fake.deletePipelineMutex.Lock()
	fake.deletePipelineArgsForCall = append(fake.deletePipelineArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.deletePipelineMutex.Unlock()
	if fake.DeletePipelineStub != nil {
		return fake.DeletePipelineStub(pipelineName)
	} else {
		return fake.deletePipelineReturns.result1, fake.deletePipelineReturns.result2
	}
}

func (fake *FakeClient) DeletePipelineCallCount() int {
	fake.deletePipelineMutex.RLock()
	defer fake.deletePipelineMutex.RUnlock()
	return len(fake.deletePipelineArgsForCall)
}

func (fake *FakeClient) DeletePipelineArgsForCall(i int) string {
	fake.deletePipelineMutex.RLock()
	defer fake.deletePipelineMutex.RUnlock()
	return fake.deletePipelineArgsForCall[i].pipelineName
}

func (fake *FakeClient) DeletePipelineReturns(result1 bool, result2 error) {
	fake.DeletePipelineStub = nil
	fake.deletePipelineReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PausePipeline(pipelineName string) (bool, error) {
	fake.pausePipelineMutex.Lock()
	fake.pausePipelineArgsForCall = append(fake.pausePipelineArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.pausePipelineMutex.Unlock()
	if fake.PausePipelineStub != nil {
		return fake.PausePipelineStub(pipelineName)
	} else {
		return fake.pausePipelineReturns.result1, fake.pausePipelineReturns.result2
	}
}

func (fake *FakeClient) PausePipelineCallCount() int {
	fake.pausePipelineMutex.RLock()
	defer fake.pausePipelineMutex.RUnlock()
	return len(fake.pausePipelineArgsForCall)
}

func (fake *FakeClient) PausePipelineArgsForCall(i int) string {
	fake.pausePipelineMutex.RLock()
	defer fake.pausePipelineMutex.RUnlock()
	return fake.pausePipelineArgsForCall[i].pipelineName
}

func (fake *FakeClient) PausePipelineReturns(result1 bool, result2 error) {
	fake.PausePipelineStub = nil
	fake.pausePipelineReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UnpausePipeline(pipelineName string) (bool, error) {
	fake.unpausePipelineMutex.Lock()
	fake.unpausePipelineArgsForCall = append(fake.unpausePipelineArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.unpausePipelineMutex.Unlock()
	if fake.UnpausePipelineStub != nil {
		return fake.UnpausePipelineStub(pipelineName)
	} else {
		return fake.unpausePipelineReturns.result1, fake.unpausePipelineReturns.result2
	}
}

func (fake *FakeClient) UnpausePipelineCallCount() int {
	fake.unpausePipelineMutex.RLock()
	defer fake.unpausePipelineMutex.RUnlock()
	return len(fake.unpausePipelineArgsForCall)
}

func (fake *FakeClient) UnpausePipelineArgsForCall(i int) string {
	fake.unpausePipelineMutex.RLock()
	defer fake.unpausePipelineMutex.RUnlock()
	return fake.unpausePipelineArgsForCall[i].pipelineName
}

func (fake *FakeClient) UnpausePipelineReturns(result1 bool, result2 error) {
	fake.UnpausePipelineStub = nil
	fake.unpausePipelineReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Job(pipelineName string, jobName string) (atc.Job, bool, error) {
	fake.jobMutex.Lock()
	fake.jobArgsForCall = append(fake.jobArgsForCall, struct {
		pipelineName string
		jobName      string
	}{pipelineName, jobName})
	fake.jobMutex.Unlock()
	if fake.JobStub != nil {
		return fake.JobStub(pipelineName, jobName)
	} else {
		return fake.jobReturns.result1, fake.jobReturns.result2, fake.jobReturns.result3
	}
}

func (fake *FakeClient) JobCallCount() int {
	fake.jobMutex.RLock()
	defer fake.jobMutex.RUnlock()
	return len(fake.jobArgsForCall)
}

func (fake *FakeClient) JobArgsForCall(i int) (string, string) {
	fake.jobMutex.RLock()
	defer fake.jobMutex.RUnlock()
	return fake.jobArgsForCall[i].pipelineName, fake.jobArgsForCall[i].jobName
}

func (fake *FakeClient) JobReturns(result1 atc.Job, result2 bool, result3 error) {
	fake.JobStub = nil
	fake.jobReturns = struct {
		result1 atc.Job
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) JobBuild(pipelineName string, jobName string, buildName string) (atc.Build, bool, error) {
	fake.jobBuildMutex.Lock()
	fake.jobBuildArgsForCall = append(fake.jobBuildArgsForCall, struct {
		pipelineName string
		jobName      string
		buildName    string
	}{pipelineName, jobName, buildName})
	fake.jobBuildMutex.Unlock()
	if fake.JobBuildStub != nil {
		return fake.JobBuildStub(pipelineName, jobName, buildName)
	} else {
		return fake.jobBuildReturns.result1, fake.jobBuildReturns.result2, fake.jobBuildReturns.result3
	}
}

func (fake *FakeClient) JobBuildCallCount() int {
	fake.jobBuildMutex.RLock()
	defer fake.jobBuildMutex.RUnlock()
	return len(fake.jobBuildArgsForCall)
}

func (fake *FakeClient) JobBuildArgsForCall(i int) (string, string, string) {
	fake.jobBuildMutex.RLock()
	defer fake.jobBuildMutex.RUnlock()
	return fake.jobBuildArgsForCall[i].pipelineName, fake.jobBuildArgsForCall[i].jobName, fake.jobBuildArgsForCall[i].buildName
}

func (fake *FakeClient) JobBuildReturns(result1 atc.Build, result2 bool, result3 error) {
	fake.JobBuildStub = nil
	fake.jobBuildReturns = struct {
		result1 atc.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListContainers(queryList map[string]string) ([]atc.Container, error) {
	fake.listContainersMutex.Lock()
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct {
		queryList map[string]string
	}{queryList})
	fake.listContainersMutex.Unlock()
	if fake.ListContainersStub != nil {
		return fake.ListContainersStub(queryList)
	} else {
		return fake.listContainersReturns.result1, fake.listContainersReturns.result2
	}
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersArgsForCall(i int) map[string]string {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return fake.listContainersArgsForCall[i].queryList
}

func (fake *FakeClient) ListContainersReturns(result1 []atc.Container, result2 error) {
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []atc.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPipelines() ([]atc.Pipeline, error) {
	fake.listPipelinesMutex.Lock()
	fake.listPipelinesArgsForCall = append(fake.listPipelinesArgsForCall, struct{}{})
	fake.listPipelinesMutex.Unlock()
	if fake.ListPipelinesStub != nil {
		return fake.ListPipelinesStub()
	} else {
		return fake.listPipelinesReturns.result1, fake.listPipelinesReturns.result2
	}
}

func (fake *FakeClient) ListPipelinesCallCount() int {
	fake.listPipelinesMutex.RLock()
	defer fake.listPipelinesMutex.RUnlock()
	return len(fake.listPipelinesArgsForCall)
}

func (fake *FakeClient) ListPipelinesReturns(result1 []atc.Pipeline, result2 error) {
	fake.ListPipelinesStub = nil
	fake.listPipelinesReturns = struct {
		result1 []atc.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListVolumes() ([]atc.Volume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeClient) ListVolumesReturns(result1 []atc.Volume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []atc.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListWorkers() ([]atc.Worker, error) {
	fake.listWorkersMutex.Lock()
	fake.listWorkersArgsForCall = append(fake.listWorkersArgsForCall, struct{}{})
	fake.listWorkersMutex.Unlock()
	if fake.ListWorkersStub != nil {
		return fake.ListWorkersStub()
	} else {
		return fake.listWorkersReturns.result1, fake.listWorkersReturns.result2
	}
}

func (fake *FakeClient) ListWorkersCallCount() int {
	fake.listWorkersMutex.RLock()
	defer fake.listWorkersMutex.RUnlock()
	return len(fake.listWorkersArgsForCall)
}

func (fake *FakeClient) ListWorkersReturns(result1 []atc.Worker, result2 error) {
	fake.ListWorkersStub = nil
	fake.listWorkersReturns = struct {
		result1 []atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PipelineConfig(pipelineName string) (atc.Config, string, bool, error) {
	fake.pipelineConfigMutex.Lock()
	fake.pipelineConfigArgsForCall = append(fake.pipelineConfigArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.pipelineConfigMutex.Unlock()
	if fake.PipelineConfigStub != nil {
		return fake.PipelineConfigStub(pipelineName)
	} else {
		return fake.pipelineConfigReturns.result1, fake.pipelineConfigReturns.result2, fake.pipelineConfigReturns.result3, fake.pipelineConfigReturns.result4
	}
}

func (fake *FakeClient) PipelineConfigCallCount() int {
	fake.pipelineConfigMutex.RLock()
	defer fake.pipelineConfigMutex.RUnlock()
	return len(fake.pipelineConfigArgsForCall)
}

func (fake *FakeClient) PipelineConfigArgsForCall(i int) string {
	fake.pipelineConfigMutex.RLock()
	defer fake.pipelineConfigMutex.RUnlock()
	return fake.pipelineConfigArgsForCall[i].pipelineName
}

func (fake *FakeClient) PipelineConfigReturns(result1 atc.Config, result2 string, result3 bool, result4 error) {
	fake.PipelineConfigStub = nil
	fake.pipelineConfigReturns = struct {
		result1 atc.Config
		result2 string
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeClient) GetCLIReader(arch string, platform string) (io.ReadCloser, error) {
	fake.getCLIReaderMutex.Lock()
	fake.getCLIReaderArgsForCall = append(fake.getCLIReaderArgsForCall, struct {
		arch     string
		platform string
	}{arch, platform})
	fake.getCLIReaderMutex.Unlock()
	if fake.GetCLIReaderStub != nil {
		return fake.GetCLIReaderStub(arch, platform)
	} else {
		return fake.getCLIReaderReturns.result1, fake.getCLIReaderReturns.result2
	}
}

func (fake *FakeClient) GetCLIReaderCallCount() int {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return len(fake.getCLIReaderArgsForCall)
}

func (fake *FakeClient) GetCLIReaderArgsForCall(i int) (string, string) {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return fake.getCLIReaderArgsForCall[i].arch, fake.getCLIReaderArgsForCall[i].platform
}

func (fake *FakeClient) GetCLIReaderReturns(result1 io.ReadCloser, result2 error) {
	fake.GetCLIReaderStub = nil
	fake.getCLIReaderReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListAuthMethods() ([]atc.AuthMethod, error) {
	fake.listAuthMethodsMutex.Lock()
	fake.listAuthMethodsArgsForCall = append(fake.listAuthMethodsArgsForCall, struct{}{})
	fake.listAuthMethodsMutex.Unlock()
	if fake.ListAuthMethodsStub != nil {
		return fake.ListAuthMethodsStub()
	} else {
		return fake.listAuthMethodsReturns.result1, fake.listAuthMethodsReturns.result2
	}
}

func (fake *FakeClient) ListAuthMethodsCallCount() int {
	fake.listAuthMethodsMutex.RLock()
	defer fake.listAuthMethodsMutex.RUnlock()
	return len(fake.listAuthMethodsArgsForCall)
}

func (fake *FakeClient) ListAuthMethodsReturns(result1 []atc.AuthMethod, result2 error) {
	fake.ListAuthMethodsStub = nil
	fake.listAuthMethodsReturns = struct {
		result1 []atc.AuthMethod
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AuthToken() (atc.AuthToken, error) {
	fake.authTokenMutex.Lock()
	fake.authTokenArgsForCall = append(fake.authTokenArgsForCall, struct{}{})
	fake.authTokenMutex.Unlock()
	if fake.AuthTokenStub != nil {
		return fake.AuthTokenStub()
	} else {
		return fake.authTokenReturns.result1, fake.authTokenReturns.result2
	}
}

func (fake *FakeClient) AuthTokenCallCount() int {
	fake.authTokenMutex.RLock()
	defer fake.authTokenMutex.RUnlock()
	return len(fake.authTokenArgsForCall)
}

func (fake *FakeClient) AuthTokenReturns(result1 atc.AuthToken, result2 error) {
	fake.AuthTokenStub = nil
	fake.authTokenReturns = struct {
		result1 atc.AuthToken
		result2 error
	}{result1, result2}
}

var _ atcclient.Client = new(FakeClient)
