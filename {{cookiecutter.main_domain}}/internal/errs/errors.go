package errs

import (
	"fmt"
	"strings"
)

const (
	// Error codes

	// CodeNotFound is code for non-existing resource
	CodeNotFound = "resource/not_found"
	// CodeProps is code for wrong input with property errors
	CodeProps = "input/property_errors"
	// CodeMalformed is code for malformed input
	CodeMalformed = "input/malformed"
	// CodeNoIdentifier is code for missing identifier in input
	CodeNoIdentifier = "input/missing_identifier"
	// CodeInputEmpty is code for empty input
	CodeInputEmpty = "input/empty"
	// CodeOneIdentifier is code for error with too many identifiers
	CodeOneIdentifier = "input/too_many_identifiers"
	// CodeInternal is code for internal errors
	CodeInternal = "internal"
	// CodeDatastore is code for datastore error
	CodeDatastore = "internal/datastore"
	// CodeNoTemplate is code for non-existing template error
	CodeNoTemplate = "internal/template/not_found"
	// CodeExecTemplate is code for template that cannot be rendered
	CodeExecTemplate = "internal/template/compile"
	// CodeNotImpl is code for non-implemented functionality
	CodeNotImpl = "functionality/not_implemented"

	// Error messages

	// MsgProps is error message for malformed input data
	MsgProps = "problems with input data"
	// MsgNoTemplate is error message for non-existing template
	MsgNoTemplate = "template not found"
	// MsgExecTemplate is error message for template that cannot be rendered
	MsgExecTemplate = "could not compile template & data"
)

// NewInternalErr is for internal error
func NewInternalErr(err error) ServiceError {
	return New(fmt.Sprintf("internal service error - %s", err.Error()), CodeInternal)
}

// NewNoTemplateErr is error for non-existing template
func NewNoTemplateErr() ServiceError {
	return New(MsgNoTemplate, CodeNoTemplate)
}

// NewExecTemplateErr is error for template that cannot be rendered
func NewExecTemplateErr() ServiceError {
	return New(MsgExecTemplate, CodeExecTemplate)
}

// NewInputMalformedErr is error for malformed input data
func NewInputMalformedErr(err error) ServiceError {
	return New(fmt.Sprintf("input malformed - %s", err.Error()), CodeMalformed)
}

// NewResourceNotFoundErr is error for resource that doesn't exists
func NewResourceNotFoundErr(rSrc string, idType, id interface{}) ServiceError {
	return New(fmt.Sprintf("%s with/for %s [%v] not found", rSrc, idType, id), CodeNotFound)
}

// NewNoResourceIdentifierErr is error for empty resource
func NewNoResourceIdentifierErr(rSrc string, params ...string) ServiceError {
	msg := fmt.Sprintf(
		"at least one identifying param [%s] is needed in order to fetch resource [%s]",
		strings.Join(params, "; "), rSrc,
	)
	return New(msg, CodeNoIdentifier)
}

// NewTooManyIdentifiersErr is error when there are too many identifiers
func NewTooManyIdentifiersErr(rSrc string, params ...string) ServiceError {
	msg := fmt.Sprintf(
		"only one identifying param [%s] is needed in order to fetch resource [%s]",
		strings.Join(params, "/"), rSrc,
	)
	return New(msg, CodeOneIdentifier)
}

// NewDatastoreErr is error of underlying datastore
func NewDatastoreErr(err error) ServiceError {
	return New(fmt.Sprintf("datastore error - %s", err.Error()), CodeDatastore)
}

// NewNotImplementedErr is error for non-implemented functionality
func NewNotImplementedErr() ServiceError {
	return New("functionality not yet implemented", CodeNotImpl)
}

// NewInputEmptyErr is error for empty input
func NewInputEmptyErr() ServiceError {
	return New("input is empty; some data is required", CodeInputEmpty)
}
