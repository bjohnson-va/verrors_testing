# verrors - the package for dealing with microservice errors

## Overview

This package provides the following data structures & functions for you to deal with any error that occurs in a microservicce.

- `error_type.go` defines all kinds of error types that this package supports.
- `service_error.go` provides a data strcut called `ServiceError` which can be used to represent library errors, repository lay errors and service layer errors. These errors are considered microservice (or service) serrors.
- `grpc_error.go` provides some utils for you to do the conversion between gRPC errors and service errors.
- `error_interceptor.go` provides an interceptor which automatically captures service errors and convert them to gRPC errors.

## Example

Here the `Get` API for fetching service level in the `sre-reporting` microservice is used to an example to demonstrate how to use this package to handle microservice errors. This API accepts the name of a microservice from a user, fetches the service level data of this microservice from Github and then return it to the user. So the workflow is:

User <-> gRPC API server <-> service <-> repository <-> github

The errors that occurs in service layer, repository layer or internal packages can both be considered as microservice errors, which should be converted to user-friendly gRPC errors.

Convert library errors to microservice errors:

```go
// internal/github
package github

import (
	"github.com/vendasta/gosdks/verrors"
	...
)

// client is the implementation of github client
type client struct {
	...
}

// GetFileContents - get the contents of a file from a github repo.
func (c *client) GetFileContents(ctx context.Context, owner, repo, path string, opt *github.RepositoryContentGetOptions) (
	*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error) {
	file, dir, res, err := c.Client.Repositories.GetContents(ctx, owner, repo, path, opt)
	if err, ok := err.(*github.ErrorResponse); ok {
		switch err.Response.StatusCode {
		// TODO: Add more error types
		case 404:
			return nil, nil, nil, verrors.New(verrors.NotFound, "couldn't get file contents: %v", err.Message)
		default:
			return nil, nil, nil, verrors.New(verrors.Internal, "error from Github: %v", err.Message)
		}
	}
	return file, dir, res, nil
}

```

Convert service layer errors or repository layer errors to microservice errors:

```go
// internal/servicelevel/repo
package repo

import (
	"github.com/vendasta/gosdks/verrors"
    ...
)

// repo is the implementation of service level repository.
type repo struct{
    ...
}

// Get - fetches availabilitySvc level definition for given availabilitySvc from github.
func (r *repo) Get(ctx context.Context, serviceName string) (*ServiceLevel, error) {
	// Get the `availabilitySvc-level` file.
	slaFile, _, _, err := r.githubClient.GetFileContents(
		ctx, "vendasta", serviceName, "service-level.yaml", &github.RepositoryContentGetOptions{})
	if err != nil {
		// Return the error straight-up since we wrapped it in our Github client wrapper already
		return nil, err
	}

	// Decode the `availabilitySvc-level` file to the yaml format.
	slYaml, err := base64.StdEncoding.DecodeString(*slaFile.Content)
	if err != nil {
		return nil, verrors.Error(verrors.Internal, "error decoding service level definition %s from raw data to the yaml format, err: %s",
			*slaFile.URL, err.Error())
	}
	...
}
```

At the API server layer, you can use the `verrors.ToGrpcError()` function to manually convert microservice errors to gRPC errors. Or you can install the `verrors.ErrorConverterServerInterceptor()` interceptor to do the lazy conversion:

```go
// service/main.go
package main

import (
	"github.com/vendasta/gosdks/verrors"
	...
)

func main() {

	//Create Interceptors
	interceptors := []grpc.UnaryServerInterceptor{
		verrors.ErrorConverterServerInterceptor(verrors.DefaultErrorMask),
        ...
	}

	//Create a GRPC Server
	logging.Infof(ctx, "Creating GRPC server...")
	grpcServer := serverconfig.CreateGrpcServer(AppName, interceptors...)

	// Register service level API server v1.
	srereporting_v1.RegisterServiceLevelServer(grpcServer, servicelevelapiserver.GetInstance())
}

```

### Installing

It is recommended you use `go modules` and GoLang's standard vendoring to get this package

## Tests

These instructions are a WIP. You will likely need to install dependencies using glide to run the tests.
These tests will run regularly on continuous integration platform as well.

```
go test
```

### Lint

These instructions are a WIP. These tests will run regularly on TeamCity as well.

```
golint
```

## Built With

* [golint](https://github.com/golang/lint) - Linting system
* [Go Modules](https://github.com/golang/go/wiki/Modules/) - Dependency Management


