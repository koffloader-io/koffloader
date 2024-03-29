// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package healthy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthyStartupOKCode is the HTTP code returned for type GetHealthyStartupOK
const GetHealthyStartupOKCode int = 200

/*
GetHealthyStartupOK Success

swagger:response getHealthyStartupOK
*/
type GetHealthyStartupOK struct {
}

// NewGetHealthyStartupOK creates GetHealthyStartupOK with default headers values
func NewGetHealthyStartupOK() *GetHealthyStartupOK {

	return &GetHealthyStartupOK{}
}

// WriteResponse to the client
func (o *GetHealthyStartupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetHealthyStartupInternalServerErrorCode is the HTTP code returned for type GetHealthyStartupInternalServerError
const GetHealthyStartupInternalServerErrorCode int = 500

/*
GetHealthyStartupInternalServerError Failed

swagger:response getHealthyStartupInternalServerError
*/
type GetHealthyStartupInternalServerError struct {
}

// NewGetHealthyStartupInternalServerError creates GetHealthyStartupInternalServerError with default headers values
func NewGetHealthyStartupInternalServerError() *GetHealthyStartupInternalServerError {

	return &GetHealthyStartupInternalServerError{}
}

// WriteResponse to the client
func (o *GetHealthyStartupInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
