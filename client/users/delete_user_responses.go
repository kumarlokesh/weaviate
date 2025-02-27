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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*
DeleteUserNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserNoContent struct {
}

// IsSuccess returns true when this delete user no content response has a 2xx status code
func (o *DeleteUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user no content response has a 3xx status code
func (o *DeleteUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user no content response has a 4xx status code
func (o *DeleteUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user no content response has a 5xx status code
func (o *DeleteUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user no content response a status code equal to that given
func (o *DeleteUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user no content response
func (o *DeleteUserNoContent) Code() int {
	return 204
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/*
DeleteUserBadRequest describes a response with status code 400, with default header values.

Malformed request.
*/
type DeleteUserBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete user bad request response has a 2xx status code
func (o *DeleteUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user bad request response has a 3xx status code
func (o *DeleteUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user bad request response has a 4xx status code
func (o *DeleteUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user bad request response has a 5xx status code
func (o *DeleteUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user bad request response a status code equal to that given
func (o *DeleteUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete user bad request response
func (o *DeleteUserBadRequest) Code() int {
	return 400
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserUnauthorized creates a DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

/*
DeleteUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type DeleteUserUnauthorized struct {
}

// IsSuccess returns true when this delete user unauthorized response has a 2xx status code
func (o *DeleteUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user unauthorized response has a 3xx status code
func (o *DeleteUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user unauthorized response has a 4xx status code
func (o *DeleteUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user unauthorized response has a 5xx status code
func (o *DeleteUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user unauthorized response a status code equal to that given
func (o *DeleteUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete user unauthorized response
func (o *DeleteUserUnauthorized) Code() int {
	return 401
}

func (o *DeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserUnauthorized ", 401)
}

func (o *DeleteUserUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserUnauthorized ", 401)
}

func (o *DeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserForbidden creates a DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {
	return &DeleteUserForbidden{}
}

/*
DeleteUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUserForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete user forbidden response has a 2xx status code
func (o *DeleteUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user forbidden response has a 3xx status code
func (o *DeleteUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user forbidden response has a 4xx status code
func (o *DeleteUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user forbidden response has a 5xx status code
func (o *DeleteUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user forbidden response a status code equal to that given
func (o *DeleteUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete user forbidden response
func (o *DeleteUserForbidden) Code() int {
	return 403
}

func (o *DeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserInternalServerError creates a DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

/*
DeleteUserInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type DeleteUserInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete user internal server error response has a 2xx status code
func (o *DeleteUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user internal server error response has a 3xx status code
func (o *DeleteUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user internal server error response has a 4xx status code
func (o *DeleteUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user internal server error response has a 5xx status code
func (o *DeleteUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user internal server error response a status code equal to that given
func (o *DeleteUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete user internal server error response
func (o *DeleteUserInternalServerError) Code() int {
	return 500
}

func (o *DeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
