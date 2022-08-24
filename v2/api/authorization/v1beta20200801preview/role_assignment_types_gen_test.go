// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200801preview

import (
	"encoding/json"
	v20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801previewstorage"
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

func Test_RoleAssignment_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment to hub returns original",
		prop.ForAll(RunResourceConversionTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRoleAssignment tests if a specific instance of RoleAssignment round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRoleAssignment(subject RoleAssignment) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200801ps.RoleAssignment
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RoleAssignment
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

func Test_RoleAssignment_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment to RoleAssignment via AssignProperties_To_RoleAssignment & AssignProperties_From_RoleAssignment returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignment tests if a specific instance of RoleAssignment can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignment(subject RoleAssignment) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignment
	err := copied.AssignProperties_To_RoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignment
	err = actual.AssignProperties_From_RoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment runs a test to see if a specific instance of RoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment(subject RoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment
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

// Generator of RoleAssignment instances for property testing - lazily instantiated by RoleAssignmentGenerator()
var roleAssignmentGenerator gopter.Gen

// RoleAssignmentGenerator returns a generator of RoleAssignment instances for property testing.
func RoleAssignmentGenerator() gopter.Gen {
	if roleAssignmentGenerator != nil {
		return roleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoleAssignment(generators)
	roleAssignmentGenerator = gen.Struct(reflect.TypeOf(RoleAssignment{}), generators)

	return roleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = RoleAssignments_SpecGenerator()
	gens["Status"] = RoleAssignment_STATUSGenerator()
}

func Test_RoleAssignment_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment_STATUS to RoleAssignment_STATUS via AssignProperties_To_RoleAssignment_STATUS & AssignProperties_From_RoleAssignment_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignment_STATUS, RoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignment_STATUS tests if a specific instance of RoleAssignment_STATUS can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignment_STATUS(subject RoleAssignment_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignment_STATUS
	err := copied.AssignProperties_To_RoleAssignment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignment_STATUS
	err = actual.AssignProperties_From_RoleAssignment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_STATUS, RoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_STATUS runs a test to see if a specific instance of RoleAssignment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_STATUS(subject RoleAssignment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_STATUS
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

// Generator of RoleAssignment_STATUS instances for property testing - lazily instantiated by
// RoleAssignment_STATUSGenerator()
var roleAssignment_STATUSGenerator gopter.Gen

// RoleAssignment_STATUSGenerator returns a generator of RoleAssignment_STATUS instances for property testing.
func RoleAssignment_STATUSGenerator() gopter.Gen {
	if roleAssignment_STATUSGenerator != nil {
		return roleAssignment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(generators)
	roleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUS{}), generators)

	return roleAssignment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentProperties_STATUS_PrincipalType_ForeignGroup,
		RoleAssignmentProperties_STATUS_PrincipalType_Group,
		RoleAssignmentProperties_STATUS_PrincipalType_ServicePrincipal,
		RoleAssignmentProperties_STATUS_PrincipalType_User))
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleAssignments_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignments_Spec to RoleAssignments_Spec via AssignProperties_To_RoleAssignments_Spec & AssignProperties_From_RoleAssignments_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignments_Spec, RoleAssignments_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignments_Spec tests if a specific instance of RoleAssignments_Spec can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignments_Spec(subject RoleAssignments_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignments_Spec
	err := copied.AssignProperties_To_RoleAssignments_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignments_Spec
	err = actual.AssignProperties_From_RoleAssignments_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignments_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignments_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignments_Spec, RoleAssignments_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignments_Spec runs a test to see if a specific instance of RoleAssignments_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignments_Spec(subject RoleAssignments_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignments_Spec
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

// Generator of RoleAssignments_Spec instances for property testing - lazily instantiated by
// RoleAssignments_SpecGenerator()
var roleAssignments_SpecGenerator gopter.Gen

// RoleAssignments_SpecGenerator returns a generator of RoleAssignments_Spec instances for property testing.
func RoleAssignments_SpecGenerator() gopter.Gen {
	if roleAssignments_SpecGenerator != nil {
		return roleAssignments_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignments_Spec(generators)
	roleAssignments_SpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignments_Spec{}), generators)

	return roleAssignments_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignments_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignments_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentProperties_PrincipalType_ForeignGroup,
		RoleAssignmentProperties_PrincipalType_Group,
		RoleAssignmentProperties_PrincipalType_ServicePrincipal,
		RoleAssignmentProperties_PrincipalType_User))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
