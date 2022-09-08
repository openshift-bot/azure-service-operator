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

func Test_VirtualNetwork_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_SpecARM, VirtualNetwork_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_SpecARM runs a test to see if a specific instance of VirtualNetwork_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_SpecARM(subject VirtualNetwork_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_SpecARM
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

// Generator of VirtualNetwork_SpecARM instances for property testing - lazily instantiated by
// VirtualNetwork_SpecARMGenerator()
var virtualNetwork_SpecARMGenerator gopter.Gen

// VirtualNetwork_SpecARMGenerator returns a generator of VirtualNetwork_SpecARM instances for property testing.
// We first initialize virtualNetwork_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_SpecARMGenerator() gopter.Gen {
	if virtualNetwork_SpecARMGenerator != nil {
		return virtualNetwork_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_SpecARM(generators)
	virtualNetwork_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_SpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_SpecARM(generators)
	virtualNetwork_SpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_SpecARM{}), generators)

	return virtualNetwork_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_SpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetwork_Spec_PropertiesARMGenerator())
}

func Test_VirtualNetwork_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_PropertiesARM, VirtualNetwork_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_PropertiesARM runs a test to see if a specific instance of VirtualNetwork_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_PropertiesARM(subject VirtualNetwork_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_PropertiesARM
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

// Generator of VirtualNetwork_Spec_PropertiesARM instances for property testing - lazily instantiated by
// VirtualNetwork_Spec_PropertiesARMGenerator()
var virtualNetwork_Spec_PropertiesARMGenerator gopter.Gen

// VirtualNetwork_Spec_PropertiesARMGenerator returns a generator of VirtualNetwork_Spec_PropertiesARM instances for property testing.
// We first initialize virtualNetwork_Spec_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_PropertiesARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_PropertiesARMGenerator != nil {
		return virtualNetwork_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM(generators)
	virtualNetwork_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM(generators)
	virtualNetwork_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_PropertiesARM{}), generators)

	return virtualNetwork_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceARMGenerator())
	gens["Subnets"] = gen.SliceOf(VirtualNetwork_Spec_Properties_SubnetsARMGenerator())
	gens["VirtualNetworkPeerings"] = gen.SliceOf(VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator())
}

func Test_DhcpOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptionsARM, DhcpOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptionsARM runs a test to see if a specific instance of DhcpOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptionsARM(subject DhcpOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptionsARM
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

// Generator of DhcpOptionsARM instances for property testing - lazily instantiated by DhcpOptionsARMGenerator()
var dhcpOptionsARMGenerator gopter.Gen

// DhcpOptionsARMGenerator returns a generator of DhcpOptionsARM instances for property testing.
func DhcpOptionsARMGenerator() gopter.Gen {
	if dhcpOptionsARMGenerator != nil {
		return dhcpOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptionsARM(generators)
	dhcpOptionsARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptionsARM{}), generators)

	return dhcpOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptionsARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetwork_Spec_Properties_SubnetsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_Properties_SubnetsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_Properties_SubnetsARM, VirtualNetwork_Spec_Properties_SubnetsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_Properties_SubnetsARM runs a test to see if a specific instance of VirtualNetwork_Spec_Properties_SubnetsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_Properties_SubnetsARM(subject VirtualNetwork_Spec_Properties_SubnetsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_Properties_SubnetsARM
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

// Generator of VirtualNetwork_Spec_Properties_SubnetsARM instances for property testing - lazily instantiated by
// VirtualNetwork_Spec_Properties_SubnetsARMGenerator()
var virtualNetwork_Spec_Properties_SubnetsARMGenerator gopter.Gen

// VirtualNetwork_Spec_Properties_SubnetsARMGenerator returns a generator of VirtualNetwork_Spec_Properties_SubnetsARM instances for property testing.
// We first initialize virtualNetwork_Spec_Properties_SubnetsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_Properties_SubnetsARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_Properties_SubnetsARMGenerator != nil {
		return virtualNetwork_Spec_Properties_SubnetsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM(generators)
	virtualNetwork_Spec_Properties_SubnetsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_SubnetsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM(generators)
	virtualNetwork_Spec_Properties_SubnetsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_SubnetsARM{}), generators)

	return virtualNetwork_Spec_Properties_SubnetsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_SubnetsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator())
}

func Test_VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM, VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM runs a test to see if a specific instance of VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(subject VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM
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

// Generator of VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM instances for property testing - lazily
// instantiated by VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator()
var virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator gopter.Gen

// VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator returns a generator of VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM instances for property testing.
// We first initialize virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator != nil {
		return virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(generators)
	virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(generators)
	virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM{}), generators)

	return virtualNetwork_Spec_Properties_VirtualNetworkPeeringsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_VirtualNetworkPeeringsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormatARMGenerator())
}

func Test_VirtualNetworkBgpCommunitiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunitiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM, VirtualNetworkBgpCommunitiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM runs a test to see if a specific instance of VirtualNetworkBgpCommunitiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM(subject VirtualNetworkBgpCommunitiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunitiesARM
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

// Generator of VirtualNetworkBgpCommunitiesARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunitiesARMGenerator()
var virtualNetworkBgpCommunitiesARMGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesARMGenerator returns a generator of VirtualNetworkBgpCommunitiesARM instances for property testing.
func VirtualNetworkBgpCommunitiesARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesARMGenerator != nil {
		return virtualNetworkBgpCommunitiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM(generators)
	virtualNetworkBgpCommunitiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunitiesARM{}), generators)

	return virtualNetworkBgpCommunitiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM(gens map[string]gopter.Gen) {
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetwork_Spec_Properties_Subnets_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_Properties_Subnets_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM, VirtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM runs a test to see if a specific instance of VirtualNetwork_Spec_Properties_Subnets_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(subject VirtualNetwork_Spec_Properties_Subnets_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_Properties_Subnets_PropertiesARM
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

// Generator of VirtualNetwork_Spec_Properties_Subnets_PropertiesARM instances for property testing - lazily
// instantiated by VirtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator()
var virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator gopter.Gen

// VirtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator returns a generator of VirtualNetwork_Spec_Properties_Subnets_PropertiesARM instances for property testing.
// We first initialize virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator != nil {
		return virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(generators)
	virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_Subnets_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(generators)
	virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_Subnets_PropertiesARM{}), generators)

	return virtualNetwork_Spec_Properties_Subnets_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["PrivateEndpointNetworkPolicies"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkServiceNetworkPolicies"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_PropertiesARM(gens map[string]gopter.Gen) {
	gens["Delegations"] = gen.SliceOf(VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceARMGenerator())
	gens["NatGateway"] = gen.PtrOf(SubResourceARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(SubResourceARMGenerator())
	gens["RouteTable"] = gen.PtrOf(SubResourceARMGenerator())
	gens["ServiceEndpointPolicies"] = gen.SliceOf(SubResourceARMGenerator())
	gens["ServiceEndpoints"] = gen.SliceOf(ServiceEndpointPropertiesFormatARMGenerator())
}

func Test_VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM, VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM runs a test to see if a specific instance of VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(subject VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM
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

// Generator of VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM instances for property testing - lazily
// instantiated by VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator()
var virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator gopter.Gen

// VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator returns a generator of VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM instances for property testing.
// We first initialize virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator != nil {
		return virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(generators)
	virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(generators)
	virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM{}), generators)

	return virtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_Properties_Subnets_Properties_DelegationsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServiceDelegationPropertiesFormatARMGenerator())
}