// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

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

func Test_RoleAssignment_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_STATUSARM, RoleAssignment_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_STATUSARM runs a test to see if a specific instance of RoleAssignment_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_STATUSARM(subject RoleAssignment_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_STATUSARM
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

// Generator of RoleAssignment_STATUSARM instances for property testing - lazily instantiated by
// RoleAssignment_STATUSARMGenerator()
var roleAssignment_STATUSARMGenerator gopter.Gen

// RoleAssignment_STATUSARMGenerator returns a generator of RoleAssignment_STATUSARM instances for property testing.
// We first initialize roleAssignment_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignment_STATUSARMGenerator() gopter.Gen {
	if roleAssignment_STATUSARMGenerator != nil {
		return roleAssignment_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUSARM(generators)
	roleAssignment_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRoleAssignment_STATUSARM(generators)
	roleAssignment_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUSARM{}), generators)

	return roleAssignment_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleAssignment_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleAssignmentProperties_STATUSARMGenerator())
}

func Test_RoleAssignmentProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentProperties_STATUSARM, RoleAssignmentProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentProperties_STATUSARM runs a test to see if a specific instance of RoleAssignmentProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentProperties_STATUSARM(subject RoleAssignmentProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentProperties_STATUSARM
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

// Generator of RoleAssignmentProperties_STATUSARM instances for property testing - lazily instantiated by
// RoleAssignmentProperties_STATUSARMGenerator()
var roleAssignmentProperties_STATUSARMGenerator gopter.Gen

// RoleAssignmentProperties_STATUSARMGenerator returns a generator of RoleAssignmentProperties_STATUSARM instances for property testing.
func RoleAssignmentProperties_STATUSARMGenerator() gopter.Gen {
	if roleAssignmentProperties_STATUSARMGenerator != nil {
		return roleAssignmentProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUSARM(generators)
	roleAssignmentProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentProperties_STATUSARM{}), generators)

	return roleAssignmentProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentProperties_STATUS_PrincipalType_ForeignGroup,
		RoleAssignmentProperties_STATUS_PrincipalType_Group,
		RoleAssignmentProperties_STATUS_PrincipalType_ServicePrincipal,
		RoleAssignmentProperties_STATUS_PrincipalType_User))
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}
