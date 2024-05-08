# GeminiCommerce\ChannelManager\ChannelManagerAPI

All URIs are relative to *https://channel.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelManagerCreateAssociation**](ChannelManagerAPI.md#ChannelManagerCreateAssociation) | **Post** /channelmanager.ChannelManager/CreateAssociation | CreateAssociation
[**ChannelManagerCreateChannel**](ChannelManagerAPI.md#ChannelManagerCreateChannel) | **Post** /channelmanager.ChannelManager/CreateChannel | CreateChannel
[**ChannelManagerCreateMarket**](ChannelManagerAPI.md#ChannelManagerCreateMarket) | **Post** /channelmanager.ChannelManager/CreateMarket | CreateMarket
[**ChannelManagerDeleteAssociation**](ChannelManagerAPI.md#ChannelManagerDeleteAssociation) | **Post** /channelmanager.ChannelManager/DeleteAssociation | DeleteAssociation
[**ChannelManagerDeleteChannel**](ChannelManagerAPI.md#ChannelManagerDeleteChannel) | **Post** /channelmanager.ChannelManager/DeleteChannel | DeleteChannel
[**ChannelManagerDeleteMarket**](ChannelManagerAPI.md#ChannelManagerDeleteMarket) | **Post** /channelmanager.ChannelManager/DeleteMarket | DeleteMarket
[**ChannelManagerGetChannel**](ChannelManagerAPI.md#ChannelManagerGetChannel) | **Post** /channelmanager.ChannelManager/GetChannel | GetChannel
[**ChannelManagerGetChannelWithAssociations**](ChannelManagerAPI.md#ChannelManagerGetChannelWithAssociations) | **Post** /channelmanager.ChannelManager/GetChannelWithAssociations | GetChannelWithAssociations
[**ChannelManagerGetMarket**](ChannelManagerAPI.md#ChannelManagerGetMarket) | **Post** /channelmanager.ChannelManager/GetMarket | GetMarket
[**ChannelManagerGetMarketWithAssociations**](ChannelManagerAPI.md#ChannelManagerGetMarketWithAssociations) | **Post** /channelmanager.ChannelManager/GetMarketWithAssociations | GetMarketWithAssociations
[**ChannelManagerListChannels**](ChannelManagerAPI.md#ChannelManagerListChannels) | **Post** /channelmanager.ChannelManager/ListChannels | ListChannels
[**ChannelManagerListChannelsWithAssociations**](ChannelManagerAPI.md#ChannelManagerListChannelsWithAssociations) | **Post** /channelmanager.ChannelManager/ListChannelsWithAssociations | ListChannelsWithAssociations
[**ChannelManagerListMarkets**](ChannelManagerAPI.md#ChannelManagerListMarkets) | **Post** /channelmanager.ChannelManager/ListMarkets | ListMarkets
[**ChannelManagerListMarketsWithAssociations**](ChannelManagerAPI.md#ChannelManagerListMarketsWithAssociations) | **Post** /channelmanager.ChannelManager/ListMarketsWithAssociations | ListMarketsWithAssociations
[**ChannelManagerUpdateChannel**](ChannelManagerAPI.md#ChannelManagerUpdateChannel) | **Post** /channelmanager.ChannelManager/UpdateChannel | UpdateChannel
[**ChannelManagerUpdateMarket**](ChannelManagerAPI.md#ChannelManagerUpdateMarket) | **Post** /channelmanager.ChannelManager/UpdateMarket | UpdateMarket



## ChannelManagerCreateAssociation

> ChannelmanagerAssociationResponse ChannelManagerCreateAssociation(ctx).Body(body).Execute()

CreateAssociation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerCreateAssociationRequest() // ChannelmanagerCreateAssociationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerCreateAssociation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerCreateAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerCreateAssociation`: ChannelmanagerAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerCreateAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerCreateAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerCreateAssociationRequest**](ChannelmanagerCreateAssociationRequest.md) |  | 

### Return type

[**ChannelmanagerAssociationResponse**](ChannelmanagerAssociationResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerCreateChannel

> ChannelmanagerChannelResponse ChannelManagerCreateChannel(ctx).Body(body).Execute()

CreateChannel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerCreateChannelRequest() // ChannelmanagerCreateChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerCreateChannel(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerCreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerCreateChannel`: ChannelmanagerChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerCreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerCreateChannelRequest**](ChannelmanagerCreateChannelRequest.md) |  | 

### Return type

[**ChannelmanagerChannelResponse**](ChannelmanagerChannelResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerCreateMarket

> ChannelmanagerMarketResponse ChannelManagerCreateMarket(ctx).Body(body).Execute()

CreateMarket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerCreateMarketRequest() // ChannelmanagerCreateMarketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerCreateMarket(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerCreateMarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerCreateMarket`: ChannelmanagerMarketResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerCreateMarket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerCreateMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerCreateMarketRequest**](ChannelmanagerCreateMarketRequest.md) |  | 

### Return type

[**ChannelmanagerMarketResponse**](ChannelmanagerMarketResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerDeleteAssociation

> map[string]interface{} ChannelManagerDeleteAssociation(ctx).Body(body).Execute()

DeleteAssociation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerDeleteAssociationRequest() // ChannelmanagerDeleteAssociationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerDeleteAssociation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerDeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerDeleteAssociation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerDeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerDeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerDeleteAssociationRequest**](ChannelmanagerDeleteAssociationRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerDeleteChannel

> map[string]interface{} ChannelManagerDeleteChannel(ctx).Body(body).Execute()

DeleteChannel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerDeleteChannelRequest() // ChannelmanagerDeleteChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerDeleteChannel(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerDeleteChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerDeleteChannel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerDeleteChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerDeleteChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerDeleteChannelRequest**](ChannelmanagerDeleteChannelRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerDeleteMarket

> map[string]interface{} ChannelManagerDeleteMarket(ctx).Body(body).Execute()

DeleteMarket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerDeleteMarketRequest() // ChannelmanagerDeleteMarketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerDeleteMarket(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerDeleteMarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerDeleteMarket`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerDeleteMarket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerDeleteMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerDeleteMarketRequest**](ChannelmanagerDeleteMarketRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerGetChannel

> ChannelmanagerChannelResponse ChannelManagerGetChannel(ctx).Body(body).Execute()

GetChannel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerGetChannelRequest() // ChannelmanagerGetChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerGetChannel(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerGetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerGetChannel`: ChannelmanagerChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerGetChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerGetChannelRequest**](ChannelmanagerGetChannelRequest.md) |  | 

### Return type

[**ChannelmanagerChannelResponse**](ChannelmanagerChannelResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerGetChannelWithAssociations

> ChannelmanagerChannelResponseWithAssociations ChannelManagerGetChannelWithAssociations(ctx).Body(body).Execute()

GetChannelWithAssociations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerGetChannelWithAssociationsRequest() // ChannelmanagerGetChannelWithAssociationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerGetChannelWithAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerGetChannelWithAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerGetChannelWithAssociations`: ChannelmanagerChannelResponseWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerGetChannelWithAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerGetChannelWithAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerGetChannelWithAssociationsRequest**](ChannelmanagerGetChannelWithAssociationsRequest.md) |  | 

### Return type

[**ChannelmanagerChannelResponseWithAssociations**](ChannelmanagerChannelResponseWithAssociations.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerGetMarket

> ChannelmanagerMarketResponse ChannelManagerGetMarket(ctx).Body(body).Execute()

GetMarket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerGetMarketRequest() // ChannelmanagerGetMarketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerGetMarket(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerGetMarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerGetMarket`: ChannelmanagerMarketResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerGetMarket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerGetMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerGetMarketRequest**](ChannelmanagerGetMarketRequest.md) |  | 

### Return type

[**ChannelmanagerMarketResponse**](ChannelmanagerMarketResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerGetMarketWithAssociations

> ChannelmanagerMarketResponseWithAssociations ChannelManagerGetMarketWithAssociations(ctx).Body(body).Execute()

GetMarketWithAssociations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerGetMarketWithAssociationsRequest() // ChannelmanagerGetMarketWithAssociationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerGetMarketWithAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerGetMarketWithAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerGetMarketWithAssociations`: ChannelmanagerMarketResponseWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerGetMarketWithAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerGetMarketWithAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerGetMarketWithAssociationsRequest**](ChannelmanagerGetMarketWithAssociationsRequest.md) |  | 

### Return type

[**ChannelmanagerMarketResponseWithAssociations**](ChannelmanagerMarketResponseWithAssociations.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerListChannels

> ChannelmanagerListChannelsResponse ChannelManagerListChannels(ctx).Body(body).Execute()

ListChannels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerListChannelsRequest() // ChannelmanagerListChannelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerListChannels(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerListChannels`: ChannelmanagerListChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerListChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerListChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerListChannelsRequest**](ChannelmanagerListChannelsRequest.md) |  | 

### Return type

[**ChannelmanagerListChannelsResponse**](ChannelmanagerListChannelsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerListChannelsWithAssociations

> ChannelmanagerListChannelsWithAssociationsResponse ChannelManagerListChannelsWithAssociations(ctx).Body(body).Execute()

ListChannelsWithAssociations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerListChannelsWithAssociationsRequest() // ChannelmanagerListChannelsWithAssociationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerListChannelsWithAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerListChannelsWithAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerListChannelsWithAssociations`: ChannelmanagerListChannelsWithAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerListChannelsWithAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerListChannelsWithAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerListChannelsWithAssociationsRequest**](ChannelmanagerListChannelsWithAssociationsRequest.md) |  | 

### Return type

[**ChannelmanagerListChannelsWithAssociationsResponse**](ChannelmanagerListChannelsWithAssociationsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerListMarkets

> ChannelmanagerListMarketsResponse ChannelManagerListMarkets(ctx).Body(body).Execute()

ListMarkets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerListMarketsRequest() // ChannelmanagerListMarketsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerListMarkets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerListMarkets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerListMarkets`: ChannelmanagerListMarketsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerListMarkets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerListMarketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerListMarketsRequest**](ChannelmanagerListMarketsRequest.md) |  | 

### Return type

[**ChannelmanagerListMarketsResponse**](ChannelmanagerListMarketsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerListMarketsWithAssociations

> ChannelmanagerListMarketsWithAssociationsResponse ChannelManagerListMarketsWithAssociations(ctx).Body(body).Execute()

ListMarketsWithAssociations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerListMarketsWithAssociationsRequest() // ChannelmanagerListMarketsWithAssociationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerListMarketsWithAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerListMarketsWithAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerListMarketsWithAssociations`: ChannelmanagerListMarketsWithAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerListMarketsWithAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerListMarketsWithAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerListMarketsWithAssociationsRequest**](ChannelmanagerListMarketsWithAssociationsRequest.md) |  | 

### Return type

[**ChannelmanagerListMarketsWithAssociationsResponse**](ChannelmanagerListMarketsWithAssociationsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerUpdateChannel

> ChannelmanagerChannelResponse ChannelManagerUpdateChannel(ctx).Body(body).Execute()

UpdateChannel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerUpdateChannelRequest() // ChannelmanagerUpdateChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerUpdateChannel(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerUpdateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerUpdateChannel`: ChannelmanagerChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerUpdateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerUpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerUpdateChannelRequest**](ChannelmanagerUpdateChannelRequest.md) |  | 

### Return type

[**ChannelmanagerChannelResponse**](ChannelmanagerChannelResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelManagerUpdateMarket

> ChannelmanagerMarketResponse ChannelManagerUpdateMarket(ctx).Body(body).Execute()

UpdateMarket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-channel-manager"
)

func main() {
	body := *openapiclient.NewChannelmanagerUpdateMarketRequest() // ChannelmanagerUpdateMarketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelManagerAPI.ChannelManagerUpdateMarket(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelManagerAPI.ChannelManagerUpdateMarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelManagerUpdateMarket`: ChannelmanagerMarketResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelManagerAPI.ChannelManagerUpdateMarket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelManagerUpdateMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelmanagerUpdateMarketRequest**](ChannelmanagerUpdateMarketRequest.md) |  | 

### Return type

[**ChannelmanagerMarketResponse**](ChannelmanagerMarketResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

