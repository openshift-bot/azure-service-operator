// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_BlobContainer_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobContainer_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobContainer_STATUSARM, BlobContainer_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobContainer_STATUSARM runs a test to see if a specific instance of BlobContainer_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobContainer_STATUSARM(subject BlobContainer_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobContainer_STATUSARM
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

// Generator of BlobContainer_STATUSARM instances for property testing - lazily instantiated by
// BlobContainer_STATUSARMGenerator()
var blobContainer_STATUSARMGenerator gopter.Gen

// BlobContainer_STATUSARMGenerator returns a generator of BlobContainer_STATUSARM instances for property testing.
// We first initialize blobContainer_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobContainer_STATUSARMGenerator() gopter.Gen {
	if blobContainer_STATUSARMGenerator != nil {
		return blobContainer_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainer_STATUSARM(generators)
	blobContainer_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainer_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForBlobContainer_STATUSARM(generators)
	blobContainer_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_STATUSARM{}), generators)

	return blobContainer_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobContainer_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobContainer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobContainer_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobContainer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerProperties_STATUSARMGenerator())
}

func Test_ContainerProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerProperties_STATUSARM, ContainerProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerProperties_STATUSARM runs a test to see if a specific instance of ContainerProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerProperties_STATUSARM(subject ContainerProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerProperties_STATUSARM
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

// Generator of ContainerProperties_STATUSARM instances for property testing - lazily instantiated by
// ContainerProperties_STATUSARMGenerator()
var containerProperties_STATUSARMGenerator gopter.Gen

// ContainerProperties_STATUSARMGenerator returns a generator of ContainerProperties_STATUSARM instances for property testing.
// We first initialize containerProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerProperties_STATUSARMGenerator() gopter.Gen {
	if containerProperties_STATUSARMGenerator != nil {
		return containerProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_STATUSARM(generators)
	containerProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForContainerProperties_STATUSARM(generators)
	containerProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUSARM{}), generators)

	return containerProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_STATUS_LeaseDuration_Fixed, ContainerProperties_STATUS_LeaseDuration_Infinite))
	gens["LeaseState"] = gen.PtrOf(gen.OneConstOf(
		ContainerProperties_STATUS_LeaseState_Available,
		ContainerProperties_STATUS_LeaseState_Breaking,
		ContainerProperties_STATUS_LeaseState_Broken,
		ContainerProperties_STATUS_LeaseState_Expired,
		ContainerProperties_STATUS_LeaseState_Leased))
	gens["LeaseStatus"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_STATUS_LeaseStatus_Locked, ContainerProperties_STATUS_LeaseStatus_Unlocked))
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_STATUS_PublicAccess_Blob, ContainerProperties_STATUS_PublicAccess_Container, ContainerProperties_STATUS_PublicAccess_None))
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForContainerProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyProperties_STATUSARMGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_STATUSARMGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldProperties_STATUSARMGenerator())
}

func Test_ImmutabilityPolicyProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperties_STATUSARM, ImmutabilityPolicyProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperties_STATUSARM runs a test to see if a specific instance of ImmutabilityPolicyProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperties_STATUSARM(subject ImmutabilityPolicyProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_STATUSARM
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

// Generator of ImmutabilityPolicyProperties_STATUSARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyProperties_STATUSARMGenerator()
var immutabilityPolicyProperties_STATUSARMGenerator gopter.Gen

// ImmutabilityPolicyProperties_STATUSARMGenerator returns a generator of ImmutabilityPolicyProperties_STATUSARM instances for property testing.
// We first initialize immutabilityPolicyProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyProperties_STATUSARMGenerator() gopter.Gen {
	if immutabilityPolicyProperties_STATUSARMGenerator != nil {
		return immutabilityPolicyProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM(generators)
	immutabilityPolicyProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM(generators)
	immutabilityPolicyProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUSARM{}), generators)

	return immutabilityPolicyProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ImmutabilityPolicyProperty_STATUSARMGenerator())
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryProperty_STATUSARMGenerator())
}

func Test_ImmutableStorageWithVersioning_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_STATUSARM, ImmutableStorageWithVersioning_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_STATUSARM runs a test to see if a specific instance of ImmutableStorageWithVersioning_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_STATUSARM(subject ImmutableStorageWithVersioning_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_STATUSARM
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

// Generator of ImmutableStorageWithVersioning_STATUSARM instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioning_STATUSARMGenerator()
var immutableStorageWithVersioning_STATUSARMGenerator gopter.Gen

// ImmutableStorageWithVersioning_STATUSARMGenerator returns a generator of ImmutableStorageWithVersioning_STATUSARM instances for property testing.
func ImmutableStorageWithVersioning_STATUSARMGenerator() gopter.Gen {
	if immutableStorageWithVersioning_STATUSARMGenerator != nil {
		return immutableStorageWithVersioning_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUSARM(generators)
	immutableStorageWithVersioning_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_STATUSARM{}), generators)

	return immutableStorageWithVersioning_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUSARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.OneConstOf(ImmutableStorageWithVersioning_STATUS_MigrationState_Completed, ImmutableStorageWithVersioning_STATUS_MigrationState_InProgress))
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldProperties_STATUSARM, LegalHoldProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldProperties_STATUSARM runs a test to see if a specific instance of LegalHoldProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldProperties_STATUSARM(subject LegalHoldProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_STATUSARM
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

// Generator of LegalHoldProperties_STATUSARM instances for property testing - lazily instantiated by
// LegalHoldProperties_STATUSARMGenerator()
var legalHoldProperties_STATUSARMGenerator gopter.Gen

// LegalHoldProperties_STATUSARMGenerator returns a generator of LegalHoldProperties_STATUSARM instances for property testing.
// We first initialize legalHoldProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldProperties_STATUSARMGenerator() gopter.Gen {
	if legalHoldProperties_STATUSARMGenerator != nil {
		return legalHoldProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUSARM(generators)
	legalHoldProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUSARM(generators)
	legalHoldProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUSARM{}), generators)

	return legalHoldProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Tags"] = gen.SliceOf(TagProperty_STATUSARMGenerator())
}

func Test_ImmutabilityPolicyProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperty_STATUSARM, ImmutabilityPolicyProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperty_STATUSARM runs a test to see if a specific instance of ImmutabilityPolicyProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperty_STATUSARM(subject ImmutabilityPolicyProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperty_STATUSARM
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

// Generator of ImmutabilityPolicyProperty_STATUSARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyProperty_STATUSARMGenerator()
var immutabilityPolicyProperty_STATUSARMGenerator gopter.Gen

// ImmutabilityPolicyProperty_STATUSARMGenerator returns a generator of ImmutabilityPolicyProperty_STATUSARM instances for property testing.
func ImmutabilityPolicyProperty_STATUSARMGenerator() gopter.Gen {
	if immutabilityPolicyProperty_STATUSARMGenerator != nil {
		return immutabilityPolicyProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUSARM(generators)
	immutabilityPolicyProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperty_STATUSARM{}), generators)

	return immutabilityPolicyProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilityPolicyProperty_STATUS_State_Locked, ImmutabilityPolicyProperty_STATUS_State_Unlocked))
}

func Test_TagProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagProperty_STATUSARM, TagProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagProperty_STATUSARM runs a test to see if a specific instance of TagProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagProperty_STATUSARM(subject TagProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_STATUSARM
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

// Generator of TagProperty_STATUSARM instances for property testing - lazily instantiated by
// TagProperty_STATUSARMGenerator()
var tagProperty_STATUSARMGenerator gopter.Gen

// TagProperty_STATUSARMGenerator returns a generator of TagProperty_STATUSARM instances for property testing.
func TagProperty_STATUSARMGenerator() gopter.Gen {
	if tagProperty_STATUSARMGenerator != nil {
		return tagProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagProperty_STATUSARM(generators)
	tagProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(TagProperty_STATUSARM{}), generators)

	return tagProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForTagProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryProperty_STATUSARM, UpdateHistoryProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryProperty_STATUSARM runs a test to see if a specific instance of UpdateHistoryProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryProperty_STATUSARM(subject UpdateHistoryProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_STATUSARM
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

// Generator of UpdateHistoryProperty_STATUSARM instances for property testing - lazily instantiated by
// UpdateHistoryProperty_STATUSARMGenerator()
var updateHistoryProperty_STATUSARMGenerator gopter.Gen

// UpdateHistoryProperty_STATUSARMGenerator returns a generator of UpdateHistoryProperty_STATUSARM instances for property testing.
func UpdateHistoryProperty_STATUSARMGenerator() gopter.Gen {
	if updateHistoryProperty_STATUSARMGenerator != nil {
		return updateHistoryProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUSARM(generators)
	updateHistoryProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_STATUSARM{}), generators)

	return updateHistoryProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.OneConstOf(UpdateHistoryProperty_STATUS_Update_Extend, UpdateHistoryProperty_STATUS_Update_Lock, UpdateHistoryProperty_STATUS_Update_Put))
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}