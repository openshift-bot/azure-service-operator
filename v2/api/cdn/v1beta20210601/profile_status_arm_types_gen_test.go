// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_Profile_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfile_STATUSARM, Profile_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfile_STATUSARM runs a test to see if a specific instance of Profile_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfile_STATUSARM(subject Profile_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_STATUSARM
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

// Generator of Profile_STATUSARM instances for property testing - lazily instantiated by Profile_STATUSARMGenerator()
var profile_STATUSARMGenerator gopter.Gen

// Profile_STATUSARMGenerator returns a generator of Profile_STATUSARM instances for property testing.
// We first initialize profile_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profile_STATUSARMGenerator() gopter.Gen {
	if profile_STATUSARMGenerator != nil {
		return profile_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_STATUSARM(generators)
	profile_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Profile_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForProfile_STATUSARM(generators)
	profile_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Profile_STATUSARM{}), generators)

	return profile_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForProfile_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfile_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfile_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfileProperties_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_ProfileProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileProperties_STATUSARM, ProfileProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileProperties_STATUSARM runs a test to see if a specific instance of ProfileProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileProperties_STATUSARM(subject ProfileProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties_STATUSARM
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

// Generator of ProfileProperties_STATUSARM instances for property testing - lazily instantiated by
// ProfileProperties_STATUSARMGenerator()
var profileProperties_STATUSARMGenerator gopter.Gen

// ProfileProperties_STATUSARMGenerator returns a generator of ProfileProperties_STATUSARM instances for property testing.
func ProfileProperties_STATUSARMGenerator() gopter.Gen {
	if profileProperties_STATUSARMGenerator != nil {
		return profileProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileProperties_STATUSARM(generators)
	profileProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_STATUSARM{}), generators)

	return profileProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForProfileProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["FrontDoorId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProfileProperties_STATUS_ProvisioningState_Creating,
		ProfileProperties_STATUS_ProvisioningState_Deleting,
		ProfileProperties_STATUS_ProvisioningState_Failed,
		ProfileProperties_STATUS_ProvisioningState_Succeeded,
		ProfileProperties_STATUS_ProvisioningState_Updating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ProfileProperties_STATUS_ResourceState_Active,
		ProfileProperties_STATUS_ResourceState_Creating,
		ProfileProperties_STATUS_ResourceState_Deleting,
		ProfileProperties_STATUS_ResourceState_Disabled))
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_STATUS_Name_Custom_Verizon,
		Sku_STATUS_Name_Premium_AzureFrontDoor,
		Sku_STATUS_Name_Premium_Verizon,
		Sku_STATUS_Name_StandardPlus_955BandWidth_ChinaCdn,
		Sku_STATUS_Name_StandardPlus_AvgBandWidth_ChinaCdn,
		Sku_STATUS_Name_StandardPlus_ChinaCdn,
		Sku_STATUS_Name_Standard_955BandWidth_ChinaCdn,
		Sku_STATUS_Name_Standard_Akamai,
		Sku_STATUS_Name_Standard_AvgBandWidth_ChinaCdn,
		Sku_STATUS_Name_Standard_AzureFrontDoor,
		Sku_STATUS_Name_Standard_ChinaCdn,
		Sku_STATUS_Name_Standard_Microsoft,
		Sku_STATUS_Name_Standard_Verizon))
}
