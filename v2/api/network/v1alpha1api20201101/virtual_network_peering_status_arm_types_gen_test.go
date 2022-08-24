// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetworkPeering_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_STATUSARM, VirtualNetworkPeering_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_STATUSARM runs a test to see if a specific instance of VirtualNetworkPeering_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_STATUSARM(subject VirtualNetworkPeering_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_STATUSARM
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

// Generator of VirtualNetworkPeering_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkPeering_STATUSARMGenerator()
var virtualNetworkPeering_STATUSARMGenerator gopter.Gen

// VirtualNetworkPeering_STATUSARMGenerator returns a generator of VirtualNetworkPeering_STATUSARM instances for property testing.
// We first initialize virtualNetworkPeering_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeering_STATUSARMGenerator() gopter.Gen {
	if virtualNetworkPeering_STATUSARMGenerator != nil {
		return virtualNetworkPeering_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUSARM(generators)
	virtualNetworkPeering_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUSARM(generators)
	virtualNetworkPeering_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_STATUSARM{}), generators)

	return virtualNetworkPeering_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeering_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormat_STATUSARMGenerator())
}

func Test_VirtualNetworkPeeringPropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUSARM, VirtualNetworkPeeringPropertiesFormat_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUSARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_STATUSARM(subject VirtualNetworkPeeringPropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormat_STATUSARM
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

// Generator of VirtualNetworkPeeringPropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkPeeringPropertiesFormat_STATUSARMGenerator()
var virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormat_STATUSARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormat_STATUSARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormat_STATUSARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM(generators)
	virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM(generators)
	virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_STATUSARM{}), generators)

	return virtualNetworkPeeringPropertiesFormat_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Initiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSARMGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSARMGenerator())
}

func Test_VirtualNetworkBgpCommunities_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM, VirtualNetworkBgpCommunities_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUSARM(subject VirtualNetworkBgpCommunities_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_STATUSARM
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

// Generator of VirtualNetworkBgpCommunities_STATUSARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_STATUSARMGenerator()
var virtualNetworkBgpCommunities_STATUSARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_STATUSARMGenerator returns a generator of VirtualNetworkBgpCommunities_STATUSARM instances for property testing.
func VirtualNetworkBgpCommunities_STATUSARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_STATUSARMGenerator != nil {
		return virtualNetworkBgpCommunities_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM(generators)
	virtualNetworkBgpCommunities_STATUSARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_STATUSARM{}), generators)

	return virtualNetworkBgpCommunities_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUSARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
