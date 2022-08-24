// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerUserDefinedFunction
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to SqlDatabaseContainerUserDefinedFunction via AssignProperties_To_SqlDatabaseContainerUserDefinedFunction & AssignProperties_From_SqlDatabaseContainerUserDefinedFunction returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.AssignProperties_To_SqlDatabaseContainerUserDefinedFunction(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerUserDefinedFunction
	err = actual.AssignProperties_From_SqlDatabaseContainerUserDefinedFunction(&other)
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerUserDefinedFunction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction runs a test to see if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerUserDefinedFunction
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

// Generator of SqlDatabaseContainerUserDefinedFunction instances for property testing - lazily instantiated by
// SqlDatabaseContainerUserDefinedFunctionGenerator()
var sqlDatabaseContainerUserDefinedFunctionGenerator gopter.Gen

// SqlDatabaseContainerUserDefinedFunctionGenerator returns a generator of SqlDatabaseContainerUserDefinedFunction instances for property testing.
func SqlDatabaseContainerUserDefinedFunctionGenerator() gopter.Gen {
	if sqlDatabaseContainerUserDefinedFunctionGenerator != nil {
		return sqlDatabaseContainerUserDefinedFunctionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(generators)
	sqlDatabaseContainerUserDefinedFunctionGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerUserDefinedFunction{}), generators)

	return sqlDatabaseContainerUserDefinedFunctionGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator()
	gens["Status"] = SqlUserDefinedFunctionGetResults_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec to DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec via AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec & AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec tests if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(&other)
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

func Test_DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec, DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(subject DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec instances for property testing -
// lazily instantiated by DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator()
var databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec{}), generators)

	return databaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_UserDefinedFunctions_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResourceGenerator())
}

func Test_SqlUserDefinedFunctionGetResults_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionGetResults_STATUS to SqlUserDefinedFunctionGetResults_STATUS via AssignProperties_To_SqlUserDefinedFunctionGetResults_STATUS & AssignProperties_From_SqlUserDefinedFunctionGetResults_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResults_STATUS, SqlUserDefinedFunctionGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResults_STATUS tests if a specific instance of SqlUserDefinedFunctionGetResults_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResults_STATUS(subject SqlUserDefinedFunctionGetResults_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlUserDefinedFunctionGetResults_STATUS
	err := copied.AssignProperties_To_SqlUserDefinedFunctionGetResults_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionGetResults_STATUS
	err = actual.AssignProperties_From_SqlUserDefinedFunctionGetResults_STATUS(&other)
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

func Test_SqlUserDefinedFunctionGetResults_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetResults_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetResults_STATUS, SqlUserDefinedFunctionGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetResults_STATUS runs a test to see if a specific instance of SqlUserDefinedFunctionGetResults_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetResults_STATUS(subject SqlUserDefinedFunctionGetResults_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetResults_STATUS
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

// Generator of SqlUserDefinedFunctionGetResults_STATUS instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionGetResults_STATUSGenerator()
var sqlUserDefinedFunctionGetResults_STATUSGenerator gopter.Gen

// SqlUserDefinedFunctionGetResults_STATUSGenerator returns a generator of SqlUserDefinedFunctionGetResults_STATUS instances for property testing.
// We first initialize sqlUserDefinedFunctionGetResults_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlUserDefinedFunctionGetResults_STATUSGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetResults_STATUSGenerator != nil {
		return sqlUserDefinedFunctionGetResults_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS(generators)
	sqlUserDefinedFunctionGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS(generators)
	sqlUserDefinedFunctionGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_STATUS{}), generators)

	return sqlUserDefinedFunctionGetResults_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_STATUS_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionGetProperties_STATUS_Resource to SqlUserDefinedFunctionGetProperties_STATUS_Resource via AssignProperties_To_SqlUserDefinedFunctionGetProperties_STATUS_Resource & AssignProperties_From_SqlUserDefinedFunctionGetProperties_STATUS_Resource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource, SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource tests if a specific instance of SqlUserDefinedFunctionGetProperties_STATUS_Resource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource(subject SqlUserDefinedFunctionGetProperties_STATUS_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlUserDefinedFunctionGetProperties_STATUS_Resource
	err := copied.AssignProperties_To_SqlUserDefinedFunctionGetProperties_STATUS_Resource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionGetProperties_STATUS_Resource
	err = actual.AssignProperties_From_SqlUserDefinedFunctionGetProperties_STATUS_Resource(&other)
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

func Test_SqlUserDefinedFunctionGetProperties_STATUS_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_STATUS_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource, SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_STATUS_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetProperties_STATUS_Resource(subject SqlUserDefinedFunctionGetProperties_STATUS_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_STATUS_Resource
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

// Generator of SqlUserDefinedFunctionGetProperties_STATUS_Resource instances for property testing - lazily instantiated
// by SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator()
var sqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator gopter.Gen

// SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator returns a generator of SqlUserDefinedFunctionGetProperties_STATUS_Resource instances for property testing.
func SqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator != nil {
		return sqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_STATUS_Resource(generators)
	sqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_STATUS_Resource{}), generators)

	return sqlUserDefinedFunctionGetProperties_STATUS_ResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_STATUS_Resource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetProperties_STATUS_Resource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlUserDefinedFunctionResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionResource to SqlUserDefinedFunctionResource via AssignProperties_To_SqlUserDefinedFunctionResource & AssignProperties_From_SqlUserDefinedFunctionResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionResource tests if a specific instance of SqlUserDefinedFunctionResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlUserDefinedFunctionResource
	err := copied.AssignProperties_To_SqlUserDefinedFunctionResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionResource
	err = actual.AssignProperties_From_SqlUserDefinedFunctionResource(&other)
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

func Test_SqlUserDefinedFunctionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResource runs a test to see if a specific instance of SqlUserDefinedFunctionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource
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

// Generator of SqlUserDefinedFunctionResource instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResourceGenerator()
var sqlUserDefinedFunctionResourceGenerator gopter.Gen

// SqlUserDefinedFunctionResourceGenerator returns a generator of SqlUserDefinedFunctionResource instances for property testing.
func SqlUserDefinedFunctionResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceGenerator != nil {
		return sqlUserDefinedFunctionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(generators)
	sqlUserDefinedFunctionResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource{}), generators)

	return sqlUserDefinedFunctionResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
