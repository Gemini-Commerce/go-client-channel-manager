# Go API client for channelmanager

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import channelmanager "github.com/Gemini-Commerce/go-client-channel-manager"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `channelmanager.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), channelmanager.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `channelmanager.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), channelmanager.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `channelmanager.ContextOperationServerIndices` and `channelmanager.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), channelmanager.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), channelmanager.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://channel.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChannelManagerAPI* | [**ChannelManagerCreateAssociation**](docs/ChannelManagerAPI.md#channelmanagercreateassociation) | **Post** /channelmanager.ChannelManager/CreateAssociation | CreateAssociation
*ChannelManagerAPI* | [**ChannelManagerCreateChannel**](docs/ChannelManagerAPI.md#channelmanagercreatechannel) | **Post** /channelmanager.ChannelManager/CreateChannel | CreateChannel
*ChannelManagerAPI* | [**ChannelManagerCreateMarket**](docs/ChannelManagerAPI.md#channelmanagercreatemarket) | **Post** /channelmanager.ChannelManager/CreateMarket | CreateMarket
*ChannelManagerAPI* | [**ChannelManagerDeleteAssociation**](docs/ChannelManagerAPI.md#channelmanagerdeleteassociation) | **Post** /channelmanager.ChannelManager/DeleteAssociation | DeleteAssociation
*ChannelManagerAPI* | [**ChannelManagerDeleteChannel**](docs/ChannelManagerAPI.md#channelmanagerdeletechannel) | **Post** /channelmanager.ChannelManager/DeleteChannel | DeleteChannel
*ChannelManagerAPI* | [**ChannelManagerDeleteMarket**](docs/ChannelManagerAPI.md#channelmanagerdeletemarket) | **Post** /channelmanager.ChannelManager/DeleteMarket | DeleteMarket
*ChannelManagerAPI* | [**ChannelManagerGetChannel**](docs/ChannelManagerAPI.md#channelmanagergetchannel) | **Post** /channelmanager.ChannelManager/GetChannel | GetChannel
*ChannelManagerAPI* | [**ChannelManagerGetChannelWithAssociations**](docs/ChannelManagerAPI.md#channelmanagergetchannelwithassociations) | **Post** /channelmanager.ChannelManager/GetChannelWithAssociations | GetChannelWithAssociations
*ChannelManagerAPI* | [**ChannelManagerGetMarket**](docs/ChannelManagerAPI.md#channelmanagergetmarket) | **Post** /channelmanager.ChannelManager/GetMarket | GetMarket
*ChannelManagerAPI* | [**ChannelManagerGetMarketWithAssociations**](docs/ChannelManagerAPI.md#channelmanagergetmarketwithassociations) | **Post** /channelmanager.ChannelManager/GetMarketWithAssociations | GetMarketWithAssociations
*ChannelManagerAPI* | [**ChannelManagerListChannels**](docs/ChannelManagerAPI.md#channelmanagerlistchannels) | **Post** /channelmanager.ChannelManager/ListChannels | ListChannels
*ChannelManagerAPI* | [**ChannelManagerListChannelsWithAssociations**](docs/ChannelManagerAPI.md#channelmanagerlistchannelswithassociations) | **Post** /channelmanager.ChannelManager/ListChannelsWithAssociations | ListChannelsWithAssociations
*ChannelManagerAPI* | [**ChannelManagerListMarkets**](docs/ChannelManagerAPI.md#channelmanagerlistmarkets) | **Post** /channelmanager.ChannelManager/ListMarkets | ListMarkets
*ChannelManagerAPI* | [**ChannelManagerListMarketsWithAssociations**](docs/ChannelManagerAPI.md#channelmanagerlistmarketswithassociations) | **Post** /channelmanager.ChannelManager/ListMarketsWithAssociations | ListMarketsWithAssociations
*ChannelManagerAPI* | [**ChannelManagerUpdateChannel**](docs/ChannelManagerAPI.md#channelmanagerupdatechannel) | **Post** /channelmanager.ChannelManager/UpdateChannel | UpdateChannel
*ChannelManagerAPI* | [**ChannelManagerUpdateMarket**](docs/ChannelManagerAPI.md#channelmanagerupdatemarket) | **Post** /channelmanager.ChannelManager/UpdateMarket | UpdateMarket


## Documentation For Models

 - [ChannelmanagerAssociationResponse](docs/ChannelmanagerAssociationResponse.md)
 - [ChannelmanagerAssociationResponseAssociation](docs/ChannelmanagerAssociationResponseAssociation.md)
 - [ChannelmanagerChannelResponse](docs/ChannelmanagerChannelResponse.md)
 - [ChannelmanagerChannelResponseWithAssociations](docs/ChannelmanagerChannelResponseWithAssociations.md)
 - [ChannelmanagerChannelResponseWithAssociationsAssociation](docs/ChannelmanagerChannelResponseWithAssociationsAssociation.md)
 - [ChannelmanagerChannelSettings](docs/ChannelmanagerChannelSettings.md)
 - [ChannelmanagerChannelStatus](docs/ChannelmanagerChannelStatus.md)
 - [ChannelmanagerChannelTypeWebsite](docs/ChannelmanagerChannelTypeWebsite.md)
 - [ChannelmanagerCountryCode](docs/ChannelmanagerCountryCode.md)
 - [ChannelmanagerCreateAssociationRequest](docs/ChannelmanagerCreateAssociationRequest.md)
 - [ChannelmanagerCreateChannelRequest](docs/ChannelmanagerCreateChannelRequest.md)
 - [ChannelmanagerCreateMarketRequest](docs/ChannelmanagerCreateMarketRequest.md)
 - [ChannelmanagerDeleteAssociationRequest](docs/ChannelmanagerDeleteAssociationRequest.md)
 - [ChannelmanagerDeleteChannelRequest](docs/ChannelmanagerDeleteChannelRequest.md)
 - [ChannelmanagerDeleteMarketRequest](docs/ChannelmanagerDeleteMarketRequest.md)
 - [ChannelmanagerGetChannelRequest](docs/ChannelmanagerGetChannelRequest.md)
 - [ChannelmanagerGetChannelWithAssociationsRequest](docs/ChannelmanagerGetChannelWithAssociationsRequest.md)
 - [ChannelmanagerGetMarketRequest](docs/ChannelmanagerGetMarketRequest.md)
 - [ChannelmanagerGetMarketWithAssociationsRequest](docs/ChannelmanagerGetMarketWithAssociationsRequest.md)
 - [ChannelmanagerLanguageCode](docs/ChannelmanagerLanguageCode.md)
 - [ChannelmanagerListChannelsRequest](docs/ChannelmanagerListChannelsRequest.md)
 - [ChannelmanagerListChannelsResponse](docs/ChannelmanagerListChannelsResponse.md)
 - [ChannelmanagerListChannelsWithAssociationsRequest](docs/ChannelmanagerListChannelsWithAssociationsRequest.md)
 - [ChannelmanagerListChannelsWithAssociationsResponse](docs/ChannelmanagerListChannelsWithAssociationsResponse.md)
 - [ChannelmanagerListMarketsRequest](docs/ChannelmanagerListMarketsRequest.md)
 - [ChannelmanagerListMarketsResponse](docs/ChannelmanagerListMarketsResponse.md)
 - [ChannelmanagerListMarketsWithAssociationsRequest](docs/ChannelmanagerListMarketsWithAssociationsRequest.md)
 - [ChannelmanagerListMarketsWithAssociationsResponse](docs/ChannelmanagerListMarketsWithAssociationsResponse.md)
 - [ChannelmanagerMarketResponse](docs/ChannelmanagerMarketResponse.md)
 - [ChannelmanagerMarketResponseWithAssociations](docs/ChannelmanagerMarketResponseWithAssociations.md)
 - [ChannelmanagerMarketResponseWithAssociationsAssociation](docs/ChannelmanagerMarketResponseWithAssociationsAssociation.md)
 - [ChannelmanagerUpdateChannelRequest](docs/ChannelmanagerUpdateChannelRequest.md)
 - [ChannelmanagerUpdateChannelRequestPayload](docs/ChannelmanagerUpdateChannelRequestPayload.md)
 - [ChannelmanagerUpdateMarketRequest](docs/ChannelmanagerUpdateMarketRequest.md)
 - [ChannelmanagerUpdateMarketRequestPayload](docs/ChannelmanagerUpdateMarketRequestPayload.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### standardAuthorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		channelmanager.ContextAPIKeys,
		map[string]channelmanager.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

