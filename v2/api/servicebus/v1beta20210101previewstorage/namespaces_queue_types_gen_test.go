// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

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

func Test_NamespacesQueue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue runs a test to see if a specific instance of NamespacesQueue round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue(subject NamespacesQueue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue
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

// Generator of NamespacesQueue instances for property testing - lazily instantiated by NamespacesQueueGenerator()
var namespacesQueueGenerator gopter.Gen

// NamespacesQueueGenerator returns a generator of NamespacesQueue instances for property testing.
func NamespacesQueueGenerator() gopter.Gen {
	if namespacesQueueGenerator != nil {
		return namespacesQueueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesQueue(generators)
	namespacesQueueGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue{}), generators)

	return namespacesQueueGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesQueue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Queues_SpecGenerator()
	gens["Status"] = SBQueue_STATUSGenerator()
}

func Test_Namespaces_Queues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Queues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Queues_Spec, Namespaces_Queues_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Queues_Spec runs a test to see if a specific instance of Namespaces_Queues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Queues_Spec(subject Namespaces_Queues_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Queues_Spec
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

// Generator of Namespaces_Queues_Spec instances for property testing - lazily instantiated by
// Namespaces_Queues_SpecGenerator()
var namespaces_Queues_SpecGenerator gopter.Gen

// Namespaces_Queues_SpecGenerator returns a generator of Namespaces_Queues_Spec instances for property testing.
func Namespaces_Queues_SpecGenerator() gopter.Gen {
	if namespaces_Queues_SpecGenerator != nil {
		return namespaces_Queues_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queues_Spec(generators)
	namespaces_Queues_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queues_Spec{}), generators)

	return namespaces_Queues_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Queues_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Queues_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_SBQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueue_STATUS, SBQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueue_STATUS runs a test to see if a specific instance of SBQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueue_STATUS(subject SBQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueue_STATUS
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

// Generator of SBQueue_STATUS instances for property testing - lazily instantiated by SBQueue_STATUSGenerator()
var sbQueue_STATUSGenerator gopter.Gen

// SBQueue_STATUSGenerator returns a generator of SBQueue_STATUS instances for property testing.
// We first initialize sbQueue_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueue_STATUSGenerator() gopter.Gen {
	if sbQueue_STATUSGenerator != nil {
		return sbQueue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueue_STATUS(generators)
	sbQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueue_STATUS(generators)
	AddRelatedPropertyGeneratorsForSBQueue_STATUS(generators)
	sbQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUS{}), generators)

	return sbQueue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBQueue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueue_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueue_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueue_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_MessageCountDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MessageCountDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMessageCountDetails_STATUS, MessageCountDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetails_STATUS runs a test to see if a specific instance of MessageCountDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetails_STATUS(subject MessageCountDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MessageCountDetails_STATUS
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

// Generator of MessageCountDetails_STATUS instances for property testing - lazily instantiated by
// MessageCountDetails_STATUSGenerator()
var messageCountDetails_STATUSGenerator gopter.Gen

// MessageCountDetails_STATUSGenerator returns a generator of MessageCountDetails_STATUS instances for property testing.
func MessageCountDetails_STATUSGenerator() gopter.Gen {
	if messageCountDetails_STATUSGenerator != nil {
		return messageCountDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS(generators)
	messageCountDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_STATUS{}), generators)

	return messageCountDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetails_STATUS(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}
