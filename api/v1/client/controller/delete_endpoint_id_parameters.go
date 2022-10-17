// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteEndpointIDParams creates a new DeleteEndpointIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEndpointIDParams() *DeleteEndpointIDParams {
	return &DeleteEndpointIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEndpointIDParamsWithTimeout creates a new DeleteEndpointIDParams object
// with the ability to set a timeout on a request.
func NewDeleteEndpointIDParamsWithTimeout(timeout time.Duration) *DeleteEndpointIDParams {
	return &DeleteEndpointIDParams{
		timeout: timeout,
	}
}

// NewDeleteEndpointIDParamsWithContext creates a new DeleteEndpointIDParams object
// with the ability to set a context for a request.
func NewDeleteEndpointIDParamsWithContext(ctx context.Context) *DeleteEndpointIDParams {
	return &DeleteEndpointIDParams{
		Context: ctx,
	}
}

// NewDeleteEndpointIDParamsWithHTTPClient creates a new DeleteEndpointIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEndpointIDParamsWithHTTPClient(client *http.Client) *DeleteEndpointIDParams {
	return &DeleteEndpointIDParams{
		HTTPClient: client,
	}
}

/*
DeleteEndpointIDParams contains all the parameters to send to the API endpoint

	for the delete endpoint ID operation.

	Typically these are written to a http.Request.
*/
type DeleteEndpointIDParams struct {

	/* ID.

	   String describing an endpoint

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete endpoint ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointIDParams) WithDefaults() *DeleteEndpointIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete endpoint ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEndpointIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete endpoint ID params
func (o *DeleteEndpointIDParams) WithTimeout(timeout time.Duration) *DeleteEndpointIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete endpoint ID params
func (o *DeleteEndpointIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete endpoint ID params
func (o *DeleteEndpointIDParams) WithContext(ctx context.Context) *DeleteEndpointIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete endpoint ID params
func (o *DeleteEndpointIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete endpoint ID params
func (o *DeleteEndpointIDParams) WithHTTPClient(client *http.Client) *DeleteEndpointIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete endpoint ID params
func (o *DeleteEndpointIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete endpoint ID params
func (o *DeleteEndpointIDParams) WithID(id string) *DeleteEndpointIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete endpoint ID params
func (o *DeleteEndpointIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEndpointIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
