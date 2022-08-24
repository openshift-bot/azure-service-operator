// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

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

func Test_NamespacesEventhub_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub runs a test to see if a specific instance of NamespacesEventhub round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub
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

// Generator of NamespacesEventhub instances for property testing - lazily instantiated by NamespacesEventhubGenerator()
var namespacesEventhubGenerator gopter.Gen

// NamespacesEventhubGenerator returns a generator of NamespacesEventhub instances for property testing.
func NamespacesEventhubGenerator() gopter.Gen {
	if namespacesEventhubGenerator != nil {
		return namespacesEventhubGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhub(generators)
	namespacesEventhubGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub{}), generators)

	return namespacesEventhubGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Eventhubs_SpecGenerator()
	gens["Status"] = Eventhub_STATUSGenerator()
}

func Test_Eventhub_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhub_STATUS, Eventhub_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhub_STATUS runs a test to see if a specific instance of Eventhub_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhub_STATUS(subject Eventhub_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_STATUS
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

// Generator of Eventhub_STATUS instances for property testing - lazily instantiated by Eventhub_STATUSGenerator()
var eventhub_STATUSGenerator gopter.Gen

// Eventhub_STATUSGenerator returns a generator of Eventhub_STATUS instances for property testing.
// We first initialize eventhub_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Eventhub_STATUSGenerator() gopter.Gen {
	if eventhub_STATUSGenerator != nil {
		return eventhub_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_STATUS(generators)
	eventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(Eventhub_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_STATUS(generators)
	AddRelatedPropertyGeneratorsForEventhub_STATUS(generators)
	eventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(Eventhub_STATUS{}), generators)

	return eventhub_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEventhub_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhub_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescription_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Namespaces_Eventhubs_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_Spec, Namespaces_Eventhubs_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_Spec runs a test to see if a specific instance of Namespaces_Eventhubs_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_Spec(subject Namespaces_Eventhubs_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_Spec
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

// Generator of Namespaces_Eventhubs_Spec instances for property testing - lazily instantiated by
// Namespaces_Eventhubs_SpecGenerator()
var namespaces_Eventhubs_SpecGenerator gopter.Gen

// Namespaces_Eventhubs_SpecGenerator returns a generator of Namespaces_Eventhubs_Spec instances for property testing.
// We first initialize namespaces_Eventhubs_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhubs_SpecGenerator() gopter.Gen {
	if namespaces_Eventhubs_SpecGenerator != nil {
		return namespaces_Eventhubs_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec(generators)
	namespaces_Eventhubs_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec(generators)
	namespaces_Eventhubs_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Spec{}), generators)

	return namespaces_Eventhubs_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator())
}

func Test_CaptureDescription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription_STATUS, CaptureDescription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription_STATUS runs a test to see if a specific instance of CaptureDescription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription_STATUS(subject CaptureDescription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_STATUS
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

// Generator of CaptureDescription_STATUS instances for property testing - lazily instantiated by
// CaptureDescription_STATUSGenerator()
var captureDescription_STATUSGenerator gopter.Gen

// CaptureDescription_STATUSGenerator returns a generator of CaptureDescription_STATUS instances for property testing.
// We first initialize captureDescription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescription_STATUSGenerator() gopter.Gen {
	if captureDescription_STATUSGenerator != nil {
		return captureDescription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	return captureDescription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.AlphaString())
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Destination_STATUSGenerator())
}

func Test_Namespaces_Eventhubs_Spec_Properties_CaptureDescription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_Spec_Properties_CaptureDescription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription, Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription runs a test to see if a specific instance of Namespaces_Eventhubs_Spec_Properties_CaptureDescription round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(subject Namespaces_Eventhubs_Spec_Properties_CaptureDescription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_Spec_Properties_CaptureDescription
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

// Generator of Namespaces_Eventhubs_Spec_Properties_CaptureDescription instances for property testing - lazily
// instantiated by Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator()
var namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator gopter.Gen

// Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator returns a generator of Namespaces_Eventhubs_Spec_Properties_CaptureDescription instances for property testing.
// We first initialize namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator() gopter.Gen {
	if namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator != nil {
		return namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(generators)
	namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Spec_Properties_CaptureDescription{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(generators)
	namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Spec_Properties_CaptureDescription{}), generators)

	return namespaces_Eventhubs_Spec_Properties_CaptureDescriptionGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.AlphaString())
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator())
}

func Test_Destination_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_STATUS, Destination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_STATUS runs a test to see if a specific instance of Destination_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_STATUS(subject Destination_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_STATUS
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

// Generator of Destination_STATUS instances for property testing - lazily instantiated by Destination_STATUSGenerator()
var destination_STATUSGenerator gopter.Gen

// Destination_STATUSGenerator returns a generator of Destination_STATUS instances for property testing.
func Destination_STATUSGenerator() gopter.Gen {
	if destination_STATUSGenerator != nil {
		return destination_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_STATUS(generators)
	destination_STATUSGenerator = gen.Struct(reflect.TypeOf(Destination_STATUS{}), generators)

	return destination_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDestination_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_STATUS(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination, Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination runs a test to see if a specific instance of Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination(subject Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination
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

// Generator of Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination instances for property testing -
// lazily instantiated by Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator()
var namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator gopter.Gen

// Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator returns a generator of Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination instances for property testing.
func Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator() gopter.Gen {
	if namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator != nil {
		return namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination(generators)
	namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination{}), generators)

	return namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
