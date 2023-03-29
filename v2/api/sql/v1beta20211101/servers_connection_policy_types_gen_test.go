// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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

func Test_ServersConnectionPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersConnectionPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersConnectionPolicy, ServersConnectionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersConnectionPolicy tests if a specific instance of ServersConnectionPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersConnectionPolicy(subject ServersConnectionPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersConnectionPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersConnectionPolicy
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersConnectionPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersConnectionPolicy to ServersConnectionPolicy via AssignProperties_To_ServersConnectionPolicy & AssignProperties_From_ServersConnectionPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersConnectionPolicy, ServersConnectionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersConnectionPolicy tests if a specific instance of ServersConnectionPolicy can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServersConnectionPolicy(subject ServersConnectionPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersConnectionPolicy
	err := copied.AssignProperties_To_ServersConnectionPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersConnectionPolicy
	err = actual.AssignProperties_From_ServersConnectionPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersConnectionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConnectionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConnectionPolicy, ServersConnectionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConnectionPolicy runs a test to see if a specific instance of ServersConnectionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConnectionPolicy(subject ServersConnectionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConnectionPolicy
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

// Generator of ServersConnectionPolicy instances for property testing - lazily instantiated by
// ServersConnectionPolicyGenerator()
var serversConnectionPolicyGenerator gopter.Gen

// ServersConnectionPolicyGenerator returns a generator of ServersConnectionPolicy instances for property testing.
func ServersConnectionPolicyGenerator() gopter.Gen {
	if serversConnectionPolicyGenerator != nil {
		return serversConnectionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersConnectionPolicy(generators)
	serversConnectionPolicyGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicy{}), generators)

	return serversConnectionPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersConnectionPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersConnectionPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_ConnectionPolicy_SpecGenerator()
	gens["Status"] = Servers_ConnectionPolicy_STATUSGenerator()
}

func Test_Servers_ConnectionPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_ConnectionPolicy_Spec to Servers_ConnectionPolicy_Spec via AssignProperties_To_Servers_ConnectionPolicy_Spec & AssignProperties_From_Servers_ConnectionPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_ConnectionPolicy_Spec, Servers_ConnectionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_ConnectionPolicy_Spec tests if a specific instance of Servers_ConnectionPolicy_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_ConnectionPolicy_Spec(subject Servers_ConnectionPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_ConnectionPolicy_Spec
	err := copied.AssignProperties_To_Servers_ConnectionPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_ConnectionPolicy_Spec
	err = actual.AssignProperties_From_Servers_ConnectionPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_ConnectionPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ConnectionPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ConnectionPolicy_Spec, Servers_ConnectionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ConnectionPolicy_Spec runs a test to see if a specific instance of Servers_ConnectionPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ConnectionPolicy_Spec(subject Servers_ConnectionPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ConnectionPolicy_Spec
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

// Generator of Servers_ConnectionPolicy_Spec instances for property testing - lazily instantiated by
// Servers_ConnectionPolicy_SpecGenerator()
var servers_ConnectionPolicy_SpecGenerator gopter.Gen

// Servers_ConnectionPolicy_SpecGenerator returns a generator of Servers_ConnectionPolicy_Spec instances for property testing.
func Servers_ConnectionPolicy_SpecGenerator() gopter.Gen {
	if servers_ConnectionPolicy_SpecGenerator != nil {
		return servers_ConnectionPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_Spec(generators)
	servers_ConnectionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_ConnectionPolicy_Spec{}), generators)

	return servers_ConnectionPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["ConnectionType"] = gen.PtrOf(gen.OneConstOf(ServerConnectionPolicyProperties_ConnectionType_Default, ServerConnectionPolicyProperties_ConnectionType_Proxy, ServerConnectionPolicyProperties_ConnectionType_Redirect))
}

func Test_Servers_ConnectionPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_ConnectionPolicy_STATUS to Servers_ConnectionPolicy_STATUS via AssignProperties_To_Servers_ConnectionPolicy_STATUS & AssignProperties_From_Servers_ConnectionPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_ConnectionPolicy_STATUS, Servers_ConnectionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_ConnectionPolicy_STATUS tests if a specific instance of Servers_ConnectionPolicy_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_ConnectionPolicy_STATUS(subject Servers_ConnectionPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_ConnectionPolicy_STATUS
	err := copied.AssignProperties_To_Servers_ConnectionPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_ConnectionPolicy_STATUS
	err = actual.AssignProperties_From_Servers_ConnectionPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_ConnectionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ConnectionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ConnectionPolicy_STATUS, Servers_ConnectionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ConnectionPolicy_STATUS runs a test to see if a specific instance of Servers_ConnectionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ConnectionPolicy_STATUS(subject Servers_ConnectionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ConnectionPolicy_STATUS
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

// Generator of Servers_ConnectionPolicy_STATUS instances for property testing - lazily instantiated by
// Servers_ConnectionPolicy_STATUSGenerator()
var servers_ConnectionPolicy_STATUSGenerator gopter.Gen

// Servers_ConnectionPolicy_STATUSGenerator returns a generator of Servers_ConnectionPolicy_STATUS instances for property testing.
func Servers_ConnectionPolicy_STATUSGenerator() gopter.Gen {
	if servers_ConnectionPolicy_STATUSGenerator != nil {
		return servers_ConnectionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_STATUS(generators)
	servers_ConnectionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_ConnectionPolicy_STATUS{}), generators)

	return servers_ConnectionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ConnectionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["ConnectionType"] = gen.PtrOf(gen.OneConstOf(ServerConnectionPolicyProperties_ConnectionType_STATUS_Default, ServerConnectionPolicyProperties_ConnectionType_STATUS_Proxy, ServerConnectionPolicyProperties_ConnectionType_STATUS_Redirect))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}