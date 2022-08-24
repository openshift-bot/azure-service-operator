// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import (
	"encoding/json"
	alpha20200202s "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202storage"
	v20200202s "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202storage"
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

func Test_Component_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component to hub returns original",
		prop.ForAll(RunResourceConversionTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForComponent tests if a specific instance of Component round trips to the hub storage version and back losslessly
func RunResourceConversionTestForComponent(subject Component) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200202s.Component
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Component
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

func Test_Component_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Component to Component via AssignProperties_To_Component & AssignProperties_From_Component returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponent tests if a specific instance of Component can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForComponent(subject Component) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200202s.Component
	err := copied.AssignProperties_To_Component(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Component
	err = actual.AssignProperties_From_Component(&other)
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

func Test_Component_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent, ComponentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent runs a test to see if a specific instance of Component round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent(subject Component) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component
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

// Generator of Component instances for property testing - lazily instantiated by ComponentGenerator()
var componentGenerator gopter.Gen

// ComponentGenerator returns a generator of Component instances for property testing.
func ComponentGenerator() gopter.Gen {
	if componentGenerator != nil {
		return componentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForComponent(generators)
	componentGenerator = gen.Struct(reflect.TypeOf(Component{}), generators)

	return componentGenerator
}

// AddRelatedPropertyGeneratorsForComponent is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponent(gens map[string]gopter.Gen) {
	gens["Spec"] = Components_SpecGenerator()
	gens["Status"] = ApplicationInsightsComponent_STATUSGenerator()
}

func Test_ApplicationInsightsComponent_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationInsightsComponent_STATUS to ApplicationInsightsComponent_STATUS via AssignProperties_To_ApplicationInsightsComponent_STATUS & AssignProperties_From_ApplicationInsightsComponent_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationInsightsComponent_STATUS, ApplicationInsightsComponent_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationInsightsComponent_STATUS tests if a specific instance of ApplicationInsightsComponent_STATUS can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForApplicationInsightsComponent_STATUS(subject ApplicationInsightsComponent_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200202s.ApplicationInsightsComponent_STATUS
	err := copied.AssignProperties_To_ApplicationInsightsComponent_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationInsightsComponent_STATUS
	err = actual.AssignProperties_From_ApplicationInsightsComponent_STATUS(&other)
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

func Test_ApplicationInsightsComponent_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationInsightsComponent_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationInsightsComponent_STATUS, ApplicationInsightsComponent_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationInsightsComponent_STATUS runs a test to see if a specific instance of ApplicationInsightsComponent_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationInsightsComponent_STATUS(subject ApplicationInsightsComponent_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationInsightsComponent_STATUS
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

// Generator of ApplicationInsightsComponent_STATUS instances for property testing - lazily instantiated by
// ApplicationInsightsComponent_STATUSGenerator()
var applicationInsightsComponent_STATUSGenerator gopter.Gen

// ApplicationInsightsComponent_STATUSGenerator returns a generator of ApplicationInsightsComponent_STATUS instances for property testing.
// We first initialize applicationInsightsComponent_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApplicationInsightsComponent_STATUSGenerator() gopter.Gen {
	if applicationInsightsComponent_STATUSGenerator != nil {
		return applicationInsightsComponent_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponent_STATUS(generators)
	applicationInsightsComponent_STATUSGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponent_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponent_STATUS(generators)
	AddRelatedPropertyGeneratorsForApplicationInsightsComponent_STATUS(generators)
	applicationInsightsComponent_STATUSGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponent_STATUS{}), generators)

	return applicationInsightsComponent_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApplicationInsightsComponent_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationInsightsComponent_STATUS(gens map[string]gopter.Gen) {
	gens["AppId"] = gen.PtrOf(gen.AlphaString())
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_STATUS_Application_Type_Other, ApplicationInsightsComponentProperties_STATUS_Application_Type_Web))
	gens["ConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_STATUS_Flow_Type_Bluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["HockeyAppToken"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_STATUS_IngestionMode_ApplicationInsights, ApplicationInsightsComponentProperties_STATUS_IngestionMode_ApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentProperties_STATUS_IngestionMode_LogAnalytics))
	gens["InstrumentationKey"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["LaMigrationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_STATUS_Request_Source_Rest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApplicationInsightsComponent_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationInsightsComponent_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResource_STATUSGenerator())
}

func Test_Components_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Components_Spec to Components_Spec via AssignProperties_To_Components_Spec & AssignProperties_From_Components_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForComponents_Spec, Components_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForComponents_Spec tests if a specific instance of Components_Spec can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForComponents_Spec(subject Components_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200202s.Components_Spec
	err := copied.AssignProperties_To_Components_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Components_Spec
	err = actual.AssignProperties_From_Components_Spec(&other)
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

func Test_Components_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Components_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponents_Spec, Components_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponents_Spec runs a test to see if a specific instance of Components_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForComponents_Spec(subject Components_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Components_Spec
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

// Generator of Components_Spec instances for property testing - lazily instantiated by Components_SpecGenerator()
var components_SpecGenerator gopter.Gen

// Components_SpecGenerator returns a generator of Components_Spec instances for property testing.
func Components_SpecGenerator() gopter.Gen {
	if components_SpecGenerator != nil {
		return components_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponents_Spec(generators)
	components_SpecGenerator = gen.Struct(reflect.TypeOf(Components_Spec{}), generators)

	return components_SpecGenerator
}

// AddIndependentPropertyGeneratorsForComponents_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponents_Spec(gens map[string]gopter.Gen) {
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Application_Type_Other, ApplicationInsightsComponentProperties_Application_Type_Web))
	gens["AzureName"] = gen.AlphaString()
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Flow_Type_Bluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsights, ApplicationInsightsComponentProperties_IngestionMode_ApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentProperties_IngestionMode_LogAnalytics))
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_PublicNetworkAccessForIngestion_Disabled, ApplicationInsightsComponentProperties_PublicNetworkAccessForIngestion_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_PublicNetworkAccessForQuery_Disabled, ApplicationInsightsComponentProperties_PublicNetworkAccessForQuery_Enabled))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Request_Source_Rest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_PrivateLinkScopedResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateLinkScopedResource_STATUS to PrivateLinkScopedResource_STATUS via AssignProperties_To_PrivateLinkScopedResource_STATUS & AssignProperties_From_PrivateLinkScopedResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS, PrivateLinkScopedResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS tests if a specific instance of PrivateLinkScopedResource_STATUS can be assigned to v1alpha1api20200202storage and back losslessly
func RunPropertyAssignmentTestForPrivateLinkScopedResource_STATUS(subject PrivateLinkScopedResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20200202s.PrivateLinkScopedResource_STATUS
	err := copied.AssignProperties_To_PrivateLinkScopedResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateLinkScopedResource_STATUS
	err = actual.AssignProperties_From_PrivateLinkScopedResource_STATUS(&other)
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

func Test_PrivateLinkScopedResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResource_STATUS, PrivateLinkScopedResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResource_STATUS runs a test to see if a specific instance of PrivateLinkScopedResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResource_STATUS(subject PrivateLinkScopedResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_STATUS
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

// Generator of PrivateLinkScopedResource_STATUS instances for property testing - lazily instantiated by
// PrivateLinkScopedResource_STATUSGenerator()
var privateLinkScopedResource_STATUSGenerator gopter.Gen

// PrivateLinkScopedResource_STATUSGenerator returns a generator of PrivateLinkScopedResource_STATUS instances for property testing.
func PrivateLinkScopedResource_STATUSGenerator() gopter.Gen {
	if privateLinkScopedResource_STATUSGenerator != nil {
		return privateLinkScopedResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS(generators)
	privateLinkScopedResource_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_STATUS{}), generators)

	return privateLinkScopedResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}
