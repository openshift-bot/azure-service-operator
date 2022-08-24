// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_SqlDatabaseContainerTrigger_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerTrigger to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerTrigger, SqlDatabaseContainerTriggerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerTrigger tests if a specific instance of SqlDatabaseContainerTrigger round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerTrigger(subject SqlDatabaseContainerTrigger) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerTrigger
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerTrigger
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

func Test_SqlDatabaseContainerTrigger_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerTrigger to SqlDatabaseContainerTrigger via AssignProperties_To_SqlDatabaseContainerTrigger & AssignProperties_From_SqlDatabaseContainerTrigger returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerTrigger, SqlDatabaseContainerTriggerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerTrigger tests if a specific instance of SqlDatabaseContainerTrigger can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerTrigger(subject SqlDatabaseContainerTrigger) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerTrigger
	err := copied.AssignProperties_To_SqlDatabaseContainerTrigger(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerTrigger
	err = actual.AssignProperties_From_SqlDatabaseContainerTrigger(&other)
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

func Test_SqlDatabaseContainerTrigger_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerTrigger via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerTrigger, SqlDatabaseContainerTriggerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerTrigger runs a test to see if a specific instance of SqlDatabaseContainerTrigger round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerTrigger(subject SqlDatabaseContainerTrigger) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerTrigger
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

// Generator of SqlDatabaseContainerTrigger instances for property testing - lazily instantiated by
// SqlDatabaseContainerTriggerGenerator()
var sqlDatabaseContainerTriggerGenerator gopter.Gen

// SqlDatabaseContainerTriggerGenerator returns a generator of SqlDatabaseContainerTrigger instances for property testing.
func SqlDatabaseContainerTriggerGenerator() gopter.Gen {
	if sqlDatabaseContainerTriggerGenerator != nil {
		return sqlDatabaseContainerTriggerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger(generators)
	sqlDatabaseContainerTriggerGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerTrigger{}), generators)

	return sqlDatabaseContainerTriggerGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerTrigger(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator()
	gens["Status"] = SqlTriggerGetResults_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec to DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec via AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec & AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec, DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec tests if a specific instance of DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(subject DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(&other)
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

func Test_DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec, DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(subject DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator()
var databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_Triggers_Spec{}), generators)

	return databaseAccounts_SqlDatabases_Containers_Triggers_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_Triggers_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlTriggerResourceGenerator())
}

func Test_SqlTriggerGetResults_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlTriggerGetResults_STATUS to SqlTriggerGetResults_STATUS via AssignProperties_To_SqlTriggerGetResults_STATUS & AssignProperties_From_SqlTriggerGetResults_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlTriggerGetResults_STATUS, SqlTriggerGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlTriggerGetResults_STATUS tests if a specific instance of SqlTriggerGetResults_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlTriggerGetResults_STATUS(subject SqlTriggerGetResults_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlTriggerGetResults_STATUS
	err := copied.AssignProperties_To_SqlTriggerGetResults_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlTriggerGetResults_STATUS
	err = actual.AssignProperties_From_SqlTriggerGetResults_STATUS(&other)
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

func Test_SqlTriggerGetResults_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetResults_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetResults_STATUS, SqlTriggerGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetResults_STATUS runs a test to see if a specific instance of SqlTriggerGetResults_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetResults_STATUS(subject SqlTriggerGetResults_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetResults_STATUS
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

// Generator of SqlTriggerGetResults_STATUS instances for property testing - lazily instantiated by
// SqlTriggerGetResults_STATUSGenerator()
var sqlTriggerGetResults_STATUSGenerator gopter.Gen

// SqlTriggerGetResults_STATUSGenerator returns a generator of SqlTriggerGetResults_STATUS instances for property testing.
// We first initialize sqlTriggerGetResults_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTriggerGetResults_STATUSGenerator() gopter.Gen {
	if sqlTriggerGetResults_STATUSGenerator != nil {
		return sqlTriggerGetResults_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS(generators)
	sqlTriggerGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS(generators)
	sqlTriggerGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUS{}), generators)

	return sqlTriggerGetResults_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetProperties_STATUS_ResourceGenerator())
}

func Test_SqlTriggerGetProperties_STATUS_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlTriggerGetProperties_STATUS_Resource to SqlTriggerGetProperties_STATUS_Resource via AssignProperties_To_SqlTriggerGetProperties_STATUS_Resource & AssignProperties_From_SqlTriggerGetProperties_STATUS_Resource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlTriggerGetProperties_STATUS_Resource, SqlTriggerGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlTriggerGetProperties_STATUS_Resource tests if a specific instance of SqlTriggerGetProperties_STATUS_Resource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlTriggerGetProperties_STATUS_Resource(subject SqlTriggerGetProperties_STATUS_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlTriggerGetProperties_STATUS_Resource
	err := copied.AssignProperties_To_SqlTriggerGetProperties_STATUS_Resource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlTriggerGetProperties_STATUS_Resource
	err = actual.AssignProperties_From_SqlTriggerGetProperties_STATUS_Resource(&other)
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

func Test_SqlTriggerGetProperties_STATUS_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_STATUS_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_Resource, SqlTriggerGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_Resource runs a test to see if a specific instance of SqlTriggerGetProperties_STATUS_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_Resource(subject SqlTriggerGetProperties_STATUS_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_STATUS_Resource
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

// Generator of SqlTriggerGetProperties_STATUS_Resource instances for property testing - lazily instantiated by
// SqlTriggerGetProperties_STATUS_ResourceGenerator()
var sqlTriggerGetProperties_STATUS_ResourceGenerator gopter.Gen

// SqlTriggerGetProperties_STATUS_ResourceGenerator returns a generator of SqlTriggerGetProperties_STATUS_Resource instances for property testing.
func SqlTriggerGetProperties_STATUS_ResourceGenerator() gopter.Gen {
	if sqlTriggerGetProperties_STATUS_ResourceGenerator != nil {
		return sqlTriggerGetProperties_STATUS_ResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_STATUS_Resource(generators)
	sqlTriggerGetProperties_STATUS_ResourceGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_STATUS_Resource{}), generators)

	return sqlTriggerGetProperties_STATUS_ResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_STATUS_Resource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_STATUS_Resource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerGetProperties_STATUS_Resource_TriggerOperation_All,
		SqlTriggerGetProperties_STATUS_Resource_TriggerOperation_Create,
		SqlTriggerGetProperties_STATUS_Resource_TriggerOperation_Delete,
		SqlTriggerGetProperties_STATUS_Resource_TriggerOperation_Replace,
		SqlTriggerGetProperties_STATUS_Resource_TriggerOperation_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerGetProperties_STATUS_Resource_TriggerType_Post, SqlTriggerGetProperties_STATUS_Resource_TriggerType_Pre))
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlTriggerResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlTriggerResource to SqlTriggerResource via AssignProperties_To_SqlTriggerResource & AssignProperties_From_SqlTriggerResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlTriggerResource, SqlTriggerResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlTriggerResource tests if a specific instance of SqlTriggerResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlTriggerResource(subject SqlTriggerResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlTriggerResource
	err := copied.AssignProperties_To_SqlTriggerResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlTriggerResource
	err = actual.AssignProperties_From_SqlTriggerResource(&other)
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

func Test_SqlTriggerResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerResource, SqlTriggerResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerResource runs a test to see if a specific instance of SqlTriggerResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerResource(subject SqlTriggerResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerResource
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

// Generator of SqlTriggerResource instances for property testing - lazily instantiated by SqlTriggerResourceGenerator()
var sqlTriggerResourceGenerator gopter.Gen

// SqlTriggerResourceGenerator returns a generator of SqlTriggerResource instances for property testing.
func SqlTriggerResourceGenerator() gopter.Gen {
	if sqlTriggerResourceGenerator != nil {
		return sqlTriggerResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerResource(generators)
	sqlTriggerResourceGenerator = gen.Struct(reflect.TypeOf(SqlTriggerResource{}), generators)

	return sqlTriggerResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerResource_TriggerOperation_All,
		SqlTriggerResource_TriggerOperation_Create,
		SqlTriggerResource_TriggerOperation_Delete,
		SqlTriggerResource_TriggerOperation_Replace,
		SqlTriggerResource_TriggerOperation_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerResource_TriggerType_Post, SqlTriggerResource_TriggerType_Pre))
}
