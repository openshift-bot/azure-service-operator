// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ManagedClustersAgentPool_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClustersAgentPool via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClustersAgentPool, ManagedClustersAgentPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClustersAgentPool runs a test to see if a specific instance of ManagedClustersAgentPool round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClustersAgentPool(subject ManagedClustersAgentPool) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClustersAgentPool
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClustersAgentPool instances for property testing - lazily instantiated by
// ManagedClustersAgentPoolGenerator()
var managedClustersAgentPoolGenerator gopter.Gen

// ManagedClustersAgentPoolGenerator returns a generator of ManagedClustersAgentPool instances for property testing.
func ManagedClustersAgentPoolGenerator() gopter.Gen {
	if managedClustersAgentPoolGenerator != nil {
		return managedClustersAgentPoolGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagedClustersAgentPool(generators)
	managedClustersAgentPoolGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPool{}), generators)

	return managedClustersAgentPoolGenerator
}

// AddRelatedPropertyGeneratorsForManagedClustersAgentPool is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClustersAgentPool(gens map[string]gopter.Gen) {
	gens["Spec"] = ManagedClusters_AgentPools_SpecGenerator()
	gens["Status"] = AgentPool_STATUSGenerator()
}

func Test_AgentPool_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPool_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPool_STATUS, AgentPool_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPool_STATUS runs a test to see if a specific instance of AgentPool_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPool_STATUS(subject AgentPool_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPool_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPool_STATUS instances for property testing - lazily instantiated by AgentPool_STATUSGenerator()
var agentPool_STATUSGenerator gopter.Gen

// AgentPool_STATUSGenerator returns a generator of AgentPool_STATUS instances for property testing.
// We first initialize agentPool_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPool_STATUSGenerator() gopter.Gen {
	if agentPool_STATUSGenerator != nil {
		return agentPool_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_STATUS(generators)
	agentPool_STATUSGenerator = gen.Struct(reflect.TypeOf(AgentPool_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_STATUS(generators)
	AddRelatedPropertyGeneratorsForAgentPool_STATUS(generators)
	agentPool_STATUSGenerator = gen.Struct(reflect.TypeOf(AgentPool_STATUS{}), generators)

	return agentPool_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAgentPool_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPool_STATUS(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["OsSKU"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetPriority"] = gen.PtrOf(gen.AlphaString())
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPool_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPool_STATUS(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfig_STATUSGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfig_STATUSGenerator())
	gens["PowerState"] = gen.PtrOf(PowerState_STATUSGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettings_STATUSGenerator())
}

func Test_ManagedClusters_AgentPools_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_AgentPools_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_AgentPools_Spec, ManagedClusters_AgentPools_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_AgentPools_Spec runs a test to see if a specific instance of ManagedClusters_AgentPools_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_AgentPools_Spec(subject ManagedClusters_AgentPools_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_AgentPools_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClusters_AgentPools_Spec instances for property testing - lazily instantiated by
// ManagedClusters_AgentPools_SpecGenerator()
var managedClusters_AgentPools_SpecGenerator gopter.Gen

// ManagedClusters_AgentPools_SpecGenerator returns a generator of ManagedClusters_AgentPools_Spec instances for property testing.
// We first initialize managedClusters_AgentPools_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusters_AgentPools_SpecGenerator() gopter.Gen {
	if managedClusters_AgentPools_SpecGenerator != nil {
		return managedClusters_AgentPools_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPools_Spec(generators)
	managedClusters_AgentPools_SpecGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPools_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPools_Spec(generators)
	AddRelatedPropertyGeneratorsForManagedClusters_AgentPools_Spec(generators)
	managedClusters_AgentPools_SpecGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPools_Spec{}), generators)

	return managedClusters_AgentPools_SpecGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_AgentPools_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_AgentPools_Spec(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.AlphaString())
	gens["OsSKU"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetPriority"] = gen.PtrOf(gen.AlphaString())
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusters_AgentPools_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_AgentPools_Spec(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfigGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfigGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettingsGenerator())
}

func Test_AgentPoolUpgradeSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings, AgentPoolUpgradeSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings runs a test to see if a specific instance of AgentPoolUpgradeSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings(subject AgentPoolUpgradeSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPoolUpgradeSettings instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettingsGenerator()
var agentPoolUpgradeSettingsGenerator gopter.Gen

// AgentPoolUpgradeSettingsGenerator returns a generator of AgentPoolUpgradeSettings instances for property testing.
func AgentPoolUpgradeSettingsGenerator() gopter.Gen {
	if agentPoolUpgradeSettingsGenerator != nil {
		return agentPoolUpgradeSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings(generators)
	agentPoolUpgradeSettingsGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings{}), generators)

	return agentPoolUpgradeSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_AgentPoolUpgradeSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS, AgentPoolUpgradeSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS runs a test to see if a specific instance of AgentPoolUpgradeSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS(subject AgentPoolUpgradeSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPoolUpgradeSettings_STATUS instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettings_STATUSGenerator()
var agentPoolUpgradeSettings_STATUSGenerator gopter.Gen

// AgentPoolUpgradeSettings_STATUSGenerator returns a generator of AgentPoolUpgradeSettings_STATUS instances for property testing.
func AgentPoolUpgradeSettings_STATUSGenerator() gopter.Gen {
	if agentPoolUpgradeSettings_STATUSGenerator != nil {
		return agentPoolUpgradeSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS(generators)
	agentPoolUpgradeSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_STATUS{}), generators)

	return agentPoolUpgradeSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig, KubeletConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig runs a test to see if a specific instance of KubeletConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig(subject KubeletConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KubeletConfig instances for property testing - lazily instantiated by KubeletConfigGenerator()
var kubeletConfigGenerator gopter.Gen

// KubeletConfigGenerator returns a generator of KubeletConfig instances for property testing.
func KubeletConfigGenerator() gopter.Gen {
	if kubeletConfigGenerator != nil {
		return kubeletConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig(generators)
	kubeletConfigGenerator = gen.Struct(reflect.TypeOf(KubeletConfig{}), generators)

	return kubeletConfigGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig_STATUS, KubeletConfig_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig_STATUS runs a test to see if a specific instance of KubeletConfig_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig_STATUS(subject KubeletConfig_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KubeletConfig_STATUS instances for property testing - lazily instantiated by
// KubeletConfig_STATUSGenerator()
var kubeletConfig_STATUSGenerator gopter.Gen

// KubeletConfig_STATUSGenerator returns a generator of KubeletConfig_STATUS instances for property testing.
func KubeletConfig_STATUSGenerator() gopter.Gen {
	if kubeletConfig_STATUSGenerator != nil {
		return kubeletConfig_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig_STATUS(generators)
	kubeletConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_STATUS{}), generators)

	return kubeletConfig_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig, LinuxOSConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig runs a test to see if a specific instance of LinuxOSConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig(subject LinuxOSConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LinuxOSConfig instances for property testing - lazily instantiated by LinuxOSConfigGenerator()
var linuxOSConfigGenerator gopter.Gen

// LinuxOSConfigGenerator returns a generator of LinuxOSConfig instances for property testing.
// We first initialize linuxOSConfigGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfigGenerator() gopter.Gen {
	if linuxOSConfigGenerator != nil {
		return linuxOSConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig(generators)
	linuxOSConfigGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig(generators)
	linuxOSConfigGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig{}), generators)

	return linuxOSConfigGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfigGenerator())
}

func Test_LinuxOSConfig_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig_STATUS, LinuxOSConfig_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig_STATUS runs a test to see if a specific instance of LinuxOSConfig_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig_STATUS(subject LinuxOSConfig_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LinuxOSConfig_STATUS instances for property testing - lazily instantiated by
// LinuxOSConfig_STATUSGenerator()
var linuxOSConfig_STATUSGenerator gopter.Gen

// LinuxOSConfig_STATUSGenerator returns a generator of LinuxOSConfig_STATUS instances for property testing.
// We first initialize linuxOSConfig_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfig_STATUSGenerator() gopter.Gen {
	if linuxOSConfig_STATUSGenerator != nil {
		return linuxOSConfig_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS(generators)
	linuxOSConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS(generators)
	linuxOSConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUS{}), generators)

	return linuxOSConfig_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfig_STATUSGenerator())
}

func Test_SysctlConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig, SysctlConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig runs a test to see if a specific instance of SysctlConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig(subject SysctlConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SysctlConfig instances for property testing - lazily instantiated by SysctlConfigGenerator()
var sysctlConfigGenerator gopter.Gen

// SysctlConfigGenerator returns a generator of SysctlConfig instances for property testing.
func SysctlConfigGenerator() gopter.Gen {
	if sysctlConfigGenerator != nil {
		return sysctlConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig(generators)
	sysctlConfigGenerator = gen.Struct(reflect.TypeOf(SysctlConfig{}), generators)

	return sysctlConfigGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}

func Test_SysctlConfig_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig_STATUS, SysctlConfig_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig_STATUS runs a test to see if a specific instance of SysctlConfig_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig_STATUS(subject SysctlConfig_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SysctlConfig_STATUS instances for property testing - lazily instantiated by
// SysctlConfig_STATUSGenerator()
var sysctlConfig_STATUSGenerator gopter.Gen

// SysctlConfig_STATUSGenerator returns a generator of SysctlConfig_STATUS instances for property testing.
func SysctlConfig_STATUSGenerator() gopter.Gen {
	if sysctlConfig_STATUSGenerator != nil {
		return sysctlConfig_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig_STATUS(generators)
	sysctlConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_STATUS{}), generators)

	return sysctlConfig_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig_STATUS(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}
