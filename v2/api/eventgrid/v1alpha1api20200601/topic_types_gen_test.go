// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	alpha20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
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

func Test_Topic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Topic to hub returns original",
		prop.ForAll(RunResourceConversionTestForTopic, TopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForTopic tests if a specific instance of Topic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForTopic(subject Topic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200601s.Topic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Topic
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

func Test_Topic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Topic to Topic via AssignProperties_To_Topic & AssignProperties_From_Topic returns original",
		prop.ForAll(RunPropertyAssignmentTestForTopic, TopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTopic tests if a specific instance of Topic can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForTopic(subject Topic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.Topic
	err := copied.AssignProperties_To_Topic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Topic
	err = actual.AssignProperties_From_Topic(&other)
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

func Test_Topic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic, TopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic runs a test to see if a specific instance of Topic round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic(subject Topic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic
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

// Generator of Topic instances for property testing - lazily instantiated by TopicGenerator()
var topicGenerator gopter.Gen

// TopicGenerator returns a generator of Topic instances for property testing.
func TopicGenerator() gopter.Gen {
	if topicGenerator != nil {
		return topicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTopic(generators)
	topicGenerator = gen.Struct(reflect.TypeOf(Topic{}), generators)

	return topicGenerator
}

// AddRelatedPropertyGeneratorsForTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = Topics_SpecGenerator()
	gens["Status"] = Topic_STATUSGenerator()
}

func Test_Topic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Topic_STATUS to Topic_STATUS via AssignProperties_To_Topic_STATUS & AssignProperties_From_Topic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForTopic_STATUS, Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTopic_STATUS tests if a specific instance of Topic_STATUS can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForTopic_STATUS(subject Topic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.Topic_STATUS
	err := copied.AssignProperties_To_Topic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Topic_STATUS
	err = actual.AssignProperties_From_Topic_STATUS(&other)
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

func Test_Topic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic_STATUS, Topic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic_STATUS runs a test to see if a specific instance of Topic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic_STATUS(subject Topic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic_STATUS
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

// Generator of Topic_STATUS instances for property testing - lazily instantiated by Topic_STATUSGenerator()
var topic_STATUSGenerator gopter.Gen

// Topic_STATUSGenerator returns a generator of Topic_STATUS instances for property testing.
// We first initialize topic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Topic_STATUSGenerator() gopter.Gen {
	if topic_STATUSGenerator != nil {
		return topic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_STATUS(generators)
	topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Topic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForTopic_STATUS(generators)
	topic_STATUSGenerator = gen.Struct(reflect.TypeOf(Topic_STATUS{}), generators)

	return topic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(TopicProperties_STATUS_InputSchema_CloudEventSchemaV1_0, TopicProperties_STATUS_InputSchema_CustomEventSchema, TopicProperties_STATUS_InputSchema_EventGridSchema))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		TopicProperties_STATUS_ProvisioningState_Canceled,
		TopicProperties_STATUS_ProvisioningState_Creating,
		TopicProperties_STATUS_ProvisioningState_Deleting,
		TopicProperties_STATUS_ProvisioningState_Failed,
		TopicProperties_STATUS_ProvisioningState_Succeeded,
		TopicProperties_STATUS_ProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(TopicProperties_STATUS_PublicNetworkAccess_Disabled, TopicProperties_STATUS_PublicNetworkAccess_Enabled))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopic_STATUS(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRule_STATUSGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMapping_STATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Topics_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Topics_Spec to Topics_Spec via AssignProperties_To_Topics_Spec & AssignProperties_From_Topics_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForTopics_Spec, Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTopics_Spec tests if a specific instance of Topics_Spec can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForTopics_Spec(subject Topics_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.Topics_Spec
	err := copied.AssignProperties_To_Topics_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Topics_Spec
	err = actual.AssignProperties_From_Topics_Spec(&other)
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

func Test_Topics_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topics_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopics_Spec, Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopics_Spec runs a test to see if a specific instance of Topics_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTopics_Spec(subject Topics_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topics_Spec
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

// Generator of Topics_Spec instances for property testing - lazily instantiated by Topics_SpecGenerator()
var topics_SpecGenerator gopter.Gen

// Topics_SpecGenerator returns a generator of Topics_Spec instances for property testing.
func Topics_SpecGenerator() gopter.Gen {
	if topics_SpecGenerator != nil {
		return topics_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopics_Spec(generators)
	topics_SpecGenerator = gen.Struct(reflect.TypeOf(Topics_Spec{}), generators)

	return topics_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTopics_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopics_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded to PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded via AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded & AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded, PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded tests if a specific instance of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(subject PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200601s.PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
	err := copied.AssignProperties_To_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
	err = actual.AssignProperties_From_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(&other)
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

func Test_PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded, PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(subject PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator()
var privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator != nil {
		return privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(generators)
	privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded{}), generators)

	return privateEndpointConnection_STATUS_Topic_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Topic_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
