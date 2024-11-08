//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

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

// NewRemovedPermissionParams creates a new RemovedPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemovedPermissionParams() *RemovedPermissionParams {
	return &RemovedPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemovedPermissionParamsWithTimeout creates a new RemovedPermissionParams object
// with the ability to set a timeout on a request.
func NewRemovedPermissionParamsWithTimeout(timeout time.Duration) *RemovedPermissionParams {
	return &RemovedPermissionParams{
		timeout: timeout,
	}
}

// NewRemovedPermissionParamsWithContext creates a new RemovedPermissionParams object
// with the ability to set a context for a request.
func NewRemovedPermissionParamsWithContext(ctx context.Context) *RemovedPermissionParams {
	return &RemovedPermissionParams{
		Context: ctx,
	}
}

// NewRemovedPermissionParamsWithHTTPClient creates a new RemovedPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemovedPermissionParamsWithHTTPClient(client *http.Client) *RemovedPermissionParams {
	return &RemovedPermissionParams{
		HTTPClient: client,
	}
}

/*
RemovedPermissionParams contains all the parameters to send to the API endpoint

	for the removed permission operation.

	Typically these are written to a http.Request.
*/
type RemovedPermissionParams struct {

	// Body.
	Body RemovedPermissionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the removed permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemovedPermissionParams) WithDefaults() *RemovedPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the removed permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemovedPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the removed permission params
func (o *RemovedPermissionParams) WithTimeout(timeout time.Duration) *RemovedPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the removed permission params
func (o *RemovedPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the removed permission params
func (o *RemovedPermissionParams) WithContext(ctx context.Context) *RemovedPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the removed permission params
func (o *RemovedPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the removed permission params
func (o *RemovedPermissionParams) WithHTTPClient(client *http.Client) *RemovedPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the removed permission params
func (o *RemovedPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the removed permission params
func (o *RemovedPermissionParams) WithBody(body RemovedPermissionBody) *RemovedPermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the removed permission params
func (o *RemovedPermissionParams) SetBody(body RemovedPermissionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemovedPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
