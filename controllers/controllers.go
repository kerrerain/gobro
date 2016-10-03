// Every sub-package from this package must have at least a controller.go
// file where the interface should be defined.
//
// The interface is named "<Package>Controller" and the implementation "<Package>ControllerImpl". For example,
// the "account" controller interface is called using "account.AccountController and the
// implementation is called using "account.AccountControllerImpl".
//
// There must be only one endpoint by file, in order to avoid one big file that grows
// over time with dozens of endpoints.
//
// The implementation of an endpoit should be divided into two parts:
// - NameOfTheEndpoint --> A facade whose sole purpose is to inject the entities into the implementation.
// - NameOfTheEndpointDo --> The actual implementation, with the entities it needs injected by the facade.
package controllers
