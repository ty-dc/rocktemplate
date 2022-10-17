// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/spidernet-io/rocktemplate/api/v1/models"
)

// PutEndpointIDCreatedCode is the HTTP code returned for type PutEndpointIDCreated
const PutEndpointIDCreatedCode int = 201

/*
PutEndpointIDCreated Created

swagger:response putEndpointIdCreated
*/
type PutEndpointIDCreated struct {
}

// NewPutEndpointIDCreated creates PutEndpointIDCreated with default headers values
func NewPutEndpointIDCreated() *PutEndpointIDCreated {

	return &PutEndpointIDCreated{}
}

// WriteResponse to the client
func (o *PutEndpointIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PutEndpointIDInvalidCode is the HTTP code returned for type PutEndpointIDInvalid
const PutEndpointIDInvalidCode int = 400

/*
PutEndpointIDInvalid Invalid endpoint in request

swagger:response putEndpointIdInvalid
*/
type PutEndpointIDInvalid struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRes `json:"body,omitempty"`
}

// NewPutEndpointIDInvalid creates PutEndpointIDInvalid with default headers values
func NewPutEndpointIDInvalid() *PutEndpointIDInvalid {

	return &PutEndpointIDInvalid{}
}

// WithPayload adds the payload to the put endpoint Id invalid response
func (o *PutEndpointIDInvalid) WithPayload(payload *models.ErrorRes) *PutEndpointIDInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put endpoint Id invalid response
func (o *PutEndpointIDInvalid) SetPayload(payload *models.ErrorRes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEndpointIDInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutEndpointIDExistsCode is the HTTP code returned for type PutEndpointIDExists
const PutEndpointIDExistsCode int = 409

/*
PutEndpointIDExists Endpoint already exists

swagger:response putEndpointIdExists
*/
type PutEndpointIDExists struct {
}

// NewPutEndpointIDExists creates PutEndpointIDExists with default headers values
func NewPutEndpointIDExists() *PutEndpointIDExists {

	return &PutEndpointIDExists{}
}

// WriteResponse to the client
func (o *PutEndpointIDExists) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// PutEndpointIDTooManyRequestsCode is the HTTP code returned for type PutEndpointIDTooManyRequests
const PutEndpointIDTooManyRequestsCode int = 429

/*
PutEndpointIDTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response putEndpointIdTooManyRequests
*/
type PutEndpointIDTooManyRequests struct {
}

// NewPutEndpointIDTooManyRequests creates PutEndpointIDTooManyRequests with default headers values
func NewPutEndpointIDTooManyRequests() *PutEndpointIDTooManyRequests {

	return &PutEndpointIDTooManyRequests{}
}

// WriteResponse to the client
func (o *PutEndpointIDTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// PutEndpointIDFailedCode is the HTTP code returned for type PutEndpointIDFailed
const PutEndpointIDFailedCode int = 500

/*
PutEndpointIDFailed Endpoint creation failed

swagger:response putEndpointIdFailed
*/
type PutEndpointIDFailed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRes `json:"body,omitempty"`
}

// NewPutEndpointIDFailed creates PutEndpointIDFailed with default headers values
func NewPutEndpointIDFailed() *PutEndpointIDFailed {

	return &PutEndpointIDFailed{}
}

// WithPayload adds the payload to the put endpoint Id failed response
func (o *PutEndpointIDFailed) WithPayload(payload *models.ErrorRes) *PutEndpointIDFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put endpoint Id failed response
func (o *PutEndpointIDFailed) SetPayload(payload *models.ErrorRes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEndpointIDFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
