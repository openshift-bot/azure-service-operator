// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

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

func Test_ComputeResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ComputeResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComputeResource_STATUSARM, ComputeResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComputeResource_STATUSARM runs a test to see if a specific instance of ComputeResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComputeResource_STATUSARM(subject ComputeResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ComputeResource_STATUSARM
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

// Generator of ComputeResource_STATUSARM instances for property testing - lazily instantiated by
// ComputeResource_STATUSARMGenerator()
var computeResource_STATUSARMGenerator gopter.Gen

// ComputeResource_STATUSARMGenerator returns a generator of ComputeResource_STATUSARM instances for property testing.
// We first initialize computeResource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ComputeResource_STATUSARMGenerator() gopter.Gen {
	if computeResource_STATUSARMGenerator != nil {
		return computeResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeResource_STATUSARM(generators)
	computeResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ComputeResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComputeResource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForComputeResource_STATUSARM(generators)
	computeResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ComputeResource_STATUSARM{}), generators)

	return computeResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForComputeResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComputeResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComputeResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComputeResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(Compute_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_Compute_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Compute_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompute_STATUSARM, Compute_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompute_STATUSARM runs a test to see if a specific instance of Compute_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompute_STATUSARM(subject Compute_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Compute_STATUSARM
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

// Generator of Compute_STATUSARM instances for property testing - lazily instantiated by Compute_STATUSARMGenerator()
var compute_STATUSARMGenerator gopter.Gen

// Compute_STATUSARMGenerator returns a generator of Compute_STATUSARM instances for property testing.
// We first initialize compute_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Compute_STATUSARMGenerator() gopter.Gen {
	if compute_STATUSARMGenerator != nil {
		return compute_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompute_STATUSARM(generators)
	compute_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Compute_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompute_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForCompute_STATUSARM(generators)
	compute_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Compute_STATUSARM{}), generators)

	return compute_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForCompute_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompute_STATUSARM(gens map[string]gopter.Gen) {
	gens["ComputeLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ComputeType"] = gen.PtrOf(gen.OneConstOf(
		ComputeType_STATUS_AKS,
		ComputeType_STATUS_AmlCompute,
		ComputeType_STATUS_ComputeInstance,
		ComputeType_STATUS_DataFactory,
		ComputeType_STATUS_DataLakeAnalytics,
		ComputeType_STATUS_Databricks,
		ComputeType_STATUS_HDInsight,
		ComputeType_STATUS_Kubernetes,
		ComputeType_STATUS_SynapseSpark,
		ComputeType_STATUS_VirtualMachine))
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAttachedCompute"] = gen.PtrOf(gen.Bool())
	gens["ModifiedOn"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		Compute_STATUS_ProvisioningState_Canceled,
		Compute_STATUS_ProvisioningState_Creating,
		Compute_STATUS_ProvisioningState_Deleting,
		Compute_STATUS_ProvisioningState_Failed,
		Compute_STATUS_ProvisioningState_Succeeded,
		Compute_STATUS_ProvisioningState_Unknown,
		Compute_STATUS_ProvisioningState_Updating))
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCompute_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCompute_STATUSARM(gens map[string]gopter.Gen) {
	gens["ProvisioningErrors"] = gen.SliceOf(ErrorResponse_STATUSARMGenerator())
}

func Test_Identity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUSARM, Identity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUSARM runs a test to see if a specific instance of Identity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUSARM(subject Identity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUSARM
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

// Generator of Identity_STATUSARM instances for property testing - lazily instantiated by Identity_STATUSARMGenerator()
var identity_STATUSARMGenerator gopter.Gen

// Identity_STATUSARMGenerator returns a generator of Identity_STATUSARM instances for property testing.
// We first initialize identity_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_STATUSARMGenerator() gopter.Gen {
	if identity_STATUSARMGenerator != nil {
		return identity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUSARM(generators)
	identity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_STATUSARM(generators)
	identity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	return identity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_STATUS_Type_None,
		Identity_STATUS_Type_SystemAssigned,
		Identity_STATUS_Type_SystemAssignedUserAssigned,
		Identity_STATUS_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserAssignedIdentity_STATUSARMGenerator())
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
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
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

func Test_ErrorResponse_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorResponse_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorResponse_STATUSARM, ErrorResponse_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorResponse_STATUSARM runs a test to see if a specific instance of ErrorResponse_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorResponse_STATUSARM(subject ErrorResponse_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorResponse_STATUSARM
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

// Generator of ErrorResponse_STATUSARM instances for property testing - lazily instantiated by
// ErrorResponse_STATUSARMGenerator()
var errorResponse_STATUSARMGenerator gopter.Gen

// ErrorResponse_STATUSARMGenerator returns a generator of ErrorResponse_STATUSARM instances for property testing.
func ErrorResponse_STATUSARMGenerator() gopter.Gen {
	if errorResponse_STATUSARMGenerator != nil {
		return errorResponse_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForErrorResponse_STATUSARM(generators)
	errorResponse_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ErrorResponse_STATUSARM{}), generators)

	return errorResponse_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForErrorResponse_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorResponse_STATUSARM(gens map[string]gopter.Gen) {
	gens["Error"] = gen.PtrOf(ErrorDetail_STATUSARMGenerator())
}

func Test_UserAssignedIdentity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUSARM, UserAssignedIdentity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUSARM runs a test to see if a specific instance of UserAssignedIdentity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUSARM(subject UserAssignedIdentity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUSARM
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

// Generator of UserAssignedIdentity_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSARMGenerator()
var userAssignedIdentity_STATUSARMGenerator gopter.Gen

// UserAssignedIdentity_STATUSARMGenerator returns a generator of UserAssignedIdentity_STATUSARM instances for property testing.
func UserAssignedIdentity_STATUSARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSARMGenerator != nil {
		return userAssignedIdentity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(generators)
	userAssignedIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUSARM{}), generators)

	return userAssignedIdentity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ErrorDetail_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetail_STATUSARM, ErrorDetail_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetail_STATUSARM runs a test to see if a specific instance of ErrorDetail_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetail_STATUSARM(subject ErrorDetail_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_STATUSARM
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

// Generator of ErrorDetail_STATUSARM instances for property testing - lazily instantiated by
// ErrorDetail_STATUSARMGenerator()
var errorDetail_STATUSARMGenerator gopter.Gen

// ErrorDetail_STATUSARMGenerator returns a generator of ErrorDetail_STATUSARM instances for property testing.
// We first initialize errorDetail_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetail_STATUSARMGenerator() gopter.Gen {
	if errorDetail_STATUSARMGenerator != nil {
		return errorDetail_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUSARM(generators)
	errorDetail_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetail_STATUSARM(generators)
	errorDetail_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUSARM{}), generators)

	return errorDetail_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetail_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetail_STATUSARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetail_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetail_STATUSARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfo_STATUSARMGenerator())
	gens["Details"] = gen.SliceOf(ErrorDetail_STATUS_UnrolledARMGenerator())
}

func Test_ErrorAdditionalInfo_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorAdditionalInfo_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorAdditionalInfo_STATUSARM, ErrorAdditionalInfo_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorAdditionalInfo_STATUSARM runs a test to see if a specific instance of ErrorAdditionalInfo_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorAdditionalInfo_STATUSARM(subject ErrorAdditionalInfo_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorAdditionalInfo_STATUSARM
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

// Generator of ErrorAdditionalInfo_STATUSARM instances for property testing - lazily instantiated by
// ErrorAdditionalInfo_STATUSARMGenerator()
var errorAdditionalInfo_STATUSARMGenerator gopter.Gen

// ErrorAdditionalInfo_STATUSARMGenerator returns a generator of ErrorAdditionalInfo_STATUSARM instances for property testing.
func ErrorAdditionalInfo_STATUSARMGenerator() gopter.Gen {
	if errorAdditionalInfo_STATUSARMGenerator != nil {
		return errorAdditionalInfo_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUSARM(generators)
	errorAdditionalInfo_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ErrorAdditionalInfo_STATUSARM{}), generators)

	return errorAdditionalInfo_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorAdditionalInfo_STATUSARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ErrorDetail_STATUS_UnrolledARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ErrorDetail_STATUS_UnrolledARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForErrorDetail_STATUS_UnrolledARM, ErrorDetail_STATUS_UnrolledARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForErrorDetail_STATUS_UnrolledARM runs a test to see if a specific instance of ErrorDetail_STATUS_UnrolledARM round trips to JSON and back losslessly
func RunJSONSerializationTestForErrorDetail_STATUS_UnrolledARM(subject ErrorDetail_STATUS_UnrolledARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ErrorDetail_STATUS_UnrolledARM
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

// Generator of ErrorDetail_STATUS_UnrolledARM instances for property testing - lazily instantiated by
// ErrorDetail_STATUS_UnrolledARMGenerator()
var errorDetail_STATUS_UnrolledARMGenerator gopter.Gen

// ErrorDetail_STATUS_UnrolledARMGenerator returns a generator of ErrorDetail_STATUS_UnrolledARM instances for property testing.
// We first initialize errorDetail_STATUS_UnrolledARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ErrorDetail_STATUS_UnrolledARMGenerator() gopter.Gen {
	if errorDetail_STATUS_UnrolledARMGenerator != nil {
		return errorDetail_STATUS_UnrolledARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM(generators)
	errorDetail_STATUS_UnrolledARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_UnrolledARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM(generators)
	AddRelatedPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM(generators)
	errorDetail_STATUS_UnrolledARMGenerator = gen.Struct(reflect.TypeOf(ErrorDetail_STATUS_UnrolledARM{}), generators)

	return errorDetail_STATUS_UnrolledARMGenerator
}

// AddIndependentPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForErrorDetail_STATUS_UnrolledARM(gens map[string]gopter.Gen) {
	gens["AdditionalInfo"] = gen.SliceOf(ErrorAdditionalInfo_STATUSARMGenerator())
}
