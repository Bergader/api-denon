// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/home-IoT/api-denon/gen/models"
)

// GetStatusOKCode is the HTTP code returned for type GetStatusOK
const GetStatusOKCode int = 200

/*GetStatusOK Success

swagger:response getStatusOK
*/
type GetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetStatusOK creates GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {

	return &GetStatusOK{}
}

// WithPayload adds the payload to the get status o k response
func (o *GetStatusOK) WithPayload(payload *models.StatusResponse) *GetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status o k response
func (o *GetStatusOK) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetStatusDefault Error

swagger:response getStatusDefault
*/
type GetStatusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.ErrorResponse `json:"body,omitempty"`
}

// NewGetStatusDefault creates GetStatusDefault with default headers values
func NewGetStatusDefault(code int) *GetStatusDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get status default response
func (o *GetStatusDefault) WithStatusCode(code int) *GetStatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get status default response
func (o *GetStatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get status default response
func (o *GetStatusDefault) WithPayload(payload models.ErrorResponse) *GetStatusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status default response
func (o *GetStatusDefault) SetPayload(payload models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
