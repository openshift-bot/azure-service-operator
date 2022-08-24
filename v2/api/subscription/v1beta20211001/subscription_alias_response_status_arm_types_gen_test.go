// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

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

func Test_SubscriptionAliasResponse_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponse_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponse_STATUSARM, SubscriptionAliasResponse_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponse_STATUSARM runs a test to see if a specific instance of SubscriptionAliasResponse_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponse_STATUSARM(subject SubscriptionAliasResponse_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponse_STATUSARM
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

// Generator of SubscriptionAliasResponse_STATUSARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponse_STATUSARMGenerator()
var subscriptionAliasResponse_STATUSARMGenerator gopter.Gen

// SubscriptionAliasResponse_STATUSARMGenerator returns a generator of SubscriptionAliasResponse_STATUSARM instances for property testing.
// We first initialize subscriptionAliasResponse_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubscriptionAliasResponse_STATUSARMGenerator() gopter.Gen {
	if subscriptionAliasResponse_STATUSARMGenerator != nil {
		return subscriptionAliasResponse_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM(generators)
	subscriptionAliasResponse_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM(generators)
	subscriptionAliasResponse_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_STATUSARM{}), generators)

	return subscriptionAliasResponse_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubscriptionAliasResponse_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionAliasResponseProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_SubscriptionAliasResponseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUSARM, SubscriptionAliasResponseProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUSARM runs a test to see if a specific instance of SubscriptionAliasResponseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUSARM(subject SubscriptionAliasResponseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponseProperties_STATUSARM
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

// Generator of SubscriptionAliasResponseProperties_STATUSARM instances for property testing - lazily instantiated by
// SubscriptionAliasResponseProperties_STATUSARMGenerator()
var subscriptionAliasResponseProperties_STATUSARMGenerator gopter.Gen

// SubscriptionAliasResponseProperties_STATUSARMGenerator returns a generator of SubscriptionAliasResponseProperties_STATUSARM instances for property testing.
func SubscriptionAliasResponseProperties_STATUSARMGenerator() gopter.Gen {
	if subscriptionAliasResponseProperties_STATUSARMGenerator != nil {
		return subscriptionAliasResponseProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUSARM(generators)
	subscriptionAliasResponseProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponseProperties_STATUSARM{}), generators)

	return subscriptionAliasResponseProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AcceptOwnershipState"] = gen.PtrOf(gen.OneConstOf(AcceptOwnershipState_STATUS_Completed, AcceptOwnershipState_STATUS_Expired, AcceptOwnershipState_STATUS_Pending))
	gens["AcceptOwnershipUrl"] = gen.PtrOf(gen.AlphaString())
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SubscriptionAliasResponseProperties_STATUS_ProvisioningState_Accepted, SubscriptionAliasResponseProperties_STATUS_ProvisioningState_Failed, SubscriptionAliasResponseProperties_STATUS_ProvisioningState_Succeeded))
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_STATUS_DevTest, Workload_STATUS_Production))
}

func Test_SystemData_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUSARM, SystemData_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUSARM runs a test to see if a specific instance of SystemData_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUSARM(subject SystemData_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUSARM
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

// Generator of SystemData_STATUSARM instances for property testing - lazily instantiated by
// SystemData_STATUSARMGenerator()
var systemData_STATUSARMGenerator gopter.Gen

// SystemData_STATUSARMGenerator returns a generator of SystemData_STATUSARM instances for property testing.
func SystemData_STATUSARMGenerator() gopter.Gen {
	if systemData_STATUSARMGenerator != nil {
		return systemData_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUSARM(generators)
	systemData_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUSARM{}), generators)

	return systemData_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUSARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_STATUS_CreatedByType_Application,
		SystemData_STATUS_CreatedByType_Key,
		SystemData_STATUS_CreatedByType_ManagedIdentity,
		SystemData_STATUS_CreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_STATUS_LastModifiedByType_Application,
		SystemData_STATUS_LastModifiedByType_Key,
		SystemData_STATUS_LastModifiedByType_ManagedIdentity,
		SystemData_STATUS_LastModifiedByType_User))
}
