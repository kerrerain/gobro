// Every sub-package from this package must have at least a controller.go
// file where the interface should be defined.
//
// The interface is named "Controller" and the implementation "Impl". For example,
// the "account" controller interface is called using "account.Controller".
//
// There must be only one endpoint by file, in order to avoid one big file that grows
// over time with dozens and dozens of endpoints and their logic.
//
package controllers
