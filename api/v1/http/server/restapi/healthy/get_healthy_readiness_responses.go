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

// GetHealthyReadinessOKCode is the HTTP code returned for type GetHealthyReadinessOK
const GetHealthyReadinessOKCode int = 200

/*
GetHealthyReadinessOK Success

swagger:response getHealthyReadinessOK
*/
type GetHealthyReadinessOK struct {
}

// NewGetHealthyReadinessOK creates GetHealthyReadinessOK with default headers values
func NewGetHealthyReadinessOK() *GetHealthyReadinessOK {

	return &GetHealthyReadinessOK{}
}

// WriteResponse to the client
func (o *GetHealthyReadinessOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetHealthyReadinessInternalServerErrorCode is the HTTP code returned for type GetHealthyReadinessInternalServerError
const GetHealthyReadinessInternalServerErrorCode int = 500

/*
GetHealthyReadinessInternalServerError Failed

swagger:response getHealthyReadinessInternalServerError
*/
type GetHealthyReadinessInternalServerError struct {
}

// NewGetHealthyReadinessInternalServerError creates GetHealthyReadinessInternalServerError with default headers values
func NewGetHealthyReadinessInternalServerError() *GetHealthyReadinessInternalServerError {

	return &GetHealthyReadinessInternalServerError{}
}

// WriteResponse to the client
func (o *GetHealthyReadinessInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
