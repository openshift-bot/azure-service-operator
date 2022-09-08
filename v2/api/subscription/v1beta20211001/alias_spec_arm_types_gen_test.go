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

func Test_Alias_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias_SpecARM, Alias_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias_SpecARM runs a test to see if a specific instance of Alias_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias_SpecARM(subject Alias_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias_SpecARM
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

// Generator of Alias_SpecARM instances for property testing - lazily instantiated by Alias_SpecARMGenerator()
var alias_SpecARMGenerator gopter.Gen

// Alias_SpecARMGenerator returns a generator of Alias_SpecARM instances for property testing.
// We first initialize alias_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Alias_SpecARMGenerator() gopter.Gen {
	if alias_SpecARMGenerator != nil {
		return alias_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_SpecARM(generators)
	alias_SpecARMGenerator = gen.Struct(reflect.TypeOf(Alias_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_SpecARM(generators)
	AddRelatedPropertyGeneratorsForAlias_SpecARM(generators)
	alias_SpecARMGenerator = gen.Struct(reflect.TypeOf(Alias_SpecARM{}), generators)

	return alias_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForAlias_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAlias_SpecARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAlias_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PutAliasRequestPropertiesARMGenerator())
}

func Test_PutAliasRequestPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestPropertiesARM, PutAliasRequestPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestPropertiesARM runs a test to see if a specific instance of PutAliasRequestPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestPropertiesARM(subject PutAliasRequestPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestPropertiesARM
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

// Generator of PutAliasRequestPropertiesARM instances for property testing - lazily instantiated by
// PutAliasRequestPropertiesARMGenerator()
var putAliasRequestPropertiesARMGenerator gopter.Gen

// PutAliasRequestPropertiesARMGenerator returns a generator of PutAliasRequestPropertiesARM instances for property testing.
// We first initialize putAliasRequestPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PutAliasRequestPropertiesARMGenerator() gopter.Gen {
	if putAliasRequestPropertiesARMGenerator != nil {
		return putAliasRequestPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestPropertiesARM(generators)
	putAliasRequestPropertiesARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForPutAliasRequestPropertiesARM(generators)
	putAliasRequestPropertiesARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestPropertiesARM{}), generators)

	return putAliasRequestPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestPropertiesARM(gens map[string]gopter.Gen) {
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(PutAliasRequestProperties_Workload_DevTest, PutAliasRequestProperties_Workload_Production))
}

// AddRelatedPropertyGeneratorsForPutAliasRequestPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPutAliasRequestPropertiesARM(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.PtrOf(PutAliasRequestAdditionalPropertiesARMGenerator())
}

func Test_PutAliasRequestAdditionalPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestAdditionalPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestAdditionalPropertiesARM, PutAliasRequestAdditionalPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestAdditionalPropertiesARM runs a test to see if a specific instance of PutAliasRequestAdditionalPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestAdditionalPropertiesARM(subject PutAliasRequestAdditionalPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestAdditionalPropertiesARM
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

// Generator of PutAliasRequestAdditionalPropertiesARM instances for property testing - lazily instantiated by
// PutAliasRequestAdditionalPropertiesARMGenerator()
var putAliasRequestAdditionalPropertiesARMGenerator gopter.Gen

// PutAliasRequestAdditionalPropertiesARMGenerator returns a generator of PutAliasRequestAdditionalPropertiesARM instances for property testing.
func PutAliasRequestAdditionalPropertiesARMGenerator() gopter.Gen {
	if putAliasRequestAdditionalPropertiesARMGenerator != nil {
		return putAliasRequestAdditionalPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalPropertiesARM(generators)
	putAliasRequestAdditionalPropertiesARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestAdditionalPropertiesARM{}), generators)

	return putAliasRequestAdditionalPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalPropertiesARM(gens map[string]gopter.Gen) {
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionTenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}