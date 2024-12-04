//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the IotHub group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-30
//   - iotHubName - Name of the IoT hub to failover
//   - resourceGroupName - Name of the resource group containing the IoT hub resource
//   - failoverInput - Region to failover to. Must be the Azure paired region. Get the value from the secondary location in the
//     locations property. To learn more, see https://aka.ms/manualfailover/region
//   - options - ClientBeginManualFailoverOptions contains the optional parameters for the Client.BeginManualFailover method.
func (client *Client) BeginManualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (*runtime.Poller[ClientManualFailoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.manualFailover(ctx, iotHubName, resourceGroupName, failoverInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientManualFailoverResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClientManualFailoverResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-30
func (client *Client) manualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (*http.Response, error) {
	var err error
	const operationName = "Client.BeginManualFailover"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.manualFailoverCreateRequest(ctx, iotHubName, resourceGroupName, failoverInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// manualFailoverCreateRequest creates the ManualFailover request.
func (client *Client) manualFailoverCreateRequest(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/failover"
	if iotHubName == "" {
		return nil, errors.New("parameter iotHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotHubName}", url.PathEscape(iotHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, failoverInput); err != nil {
		return nil, err
	}
	return req, nil
}