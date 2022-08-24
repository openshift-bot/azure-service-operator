// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

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

func Test_SBQueue_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueue_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueue_STATUSARM, SBQueue_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueue_STATUSARM runs a test to see if a specific instance of SBQueue_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueue_STATUSARM(subject SBQueue_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueue_STATUSARM
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

// Generator of SBQueue_STATUSARM instances for property testing - lazily instantiated by SBQueue_STATUSARMGenerator()
var sbQueue_STATUSARMGenerator gopter.Gen

// SBQueue_STATUSARMGenerator returns a generator of SBQueue_STATUSARM instances for property testing.
// We first initialize sbQueue_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueue_STATUSARMGenerator() gopter.Gen {
	if sbQueue_STATUSARMGenerator != nil {
		return sbQueue_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueue_STATUSARM(generators)
	sbQueue_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueue_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSBQueue_STATUSARM(generators)
	sbQueue_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SBQueue_STATUSARM{}), generators)

	return sbQueue_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueue_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueue_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueue_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueue_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBQueueProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_SBQueueProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueueProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueProperties_STATUSARM, SBQueueProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueProperties_STATUSARM runs a test to see if a specific instance of SBQueueProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueProperties_STATUSARM(subject SBQueueProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueueProperties_STATUSARM
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

// Generator of SBQueueProperties_STATUSARM instances for property testing - lazily instantiated by
// SBQueueProperties_STATUSARMGenerator()
var sbQueueProperties_STATUSARMGenerator gopter.Gen

// SBQueueProperties_STATUSARMGenerator returns a generator of SBQueueProperties_STATUSARM instances for property testing.
// We first initialize sbQueueProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueueProperties_STATUSARMGenerator() gopter.Gen {
	if sbQueueProperties_STATUSARMGenerator != nil {
		return sbQueueProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueProperties_STATUSARM(generators)
	sbQueueProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSBQueueProperties_STATUSARM(generators)
	sbQueueProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties_STATUSARM{}), generators)

	return sbQueueProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueProperties_STATUSARM(gens map[string]gopter.Gen) {
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
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueueProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueueProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSARMGenerator())
}

func Test_MessageCountDetails_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MessageCountDetails_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMessageCountDetails_STATUSARM, MessageCountDetails_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetails_STATUSARM runs a test to see if a specific instance of MessageCountDetails_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetails_STATUSARM(subject MessageCountDetails_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MessageCountDetails_STATUSARM
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

// Generator of MessageCountDetails_STATUSARM instances for property testing - lazily instantiated by
// MessageCountDetails_STATUSARMGenerator()
var messageCountDetails_STATUSARMGenerator gopter.Gen

// MessageCountDetails_STATUSARMGenerator returns a generator of MessageCountDetails_STATUSARM instances for property testing.
func MessageCountDetails_STATUSARMGenerator() gopter.Gen {
	if messageCountDetails_STATUSARMGenerator != nil {
		return messageCountDetails_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetails_STATUSARM(generators)
	messageCountDetails_STATUSARMGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_STATUSARM{}), generators)

	return messageCountDetails_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetails_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetails_STATUSARM(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}
