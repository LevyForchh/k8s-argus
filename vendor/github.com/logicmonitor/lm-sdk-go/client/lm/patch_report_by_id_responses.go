// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// PatchReportByIDReader is a Reader for the PatchReportByID structure.
type PatchReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchReportByIDOK creates a PatchReportByIDOK with default headers values
func NewPatchReportByIDOK() *PatchReportByIDOK {
	return &PatchReportByIDOK{}
}

/*PatchReportByIDOK handles this case with default header values.

successful operation
*/
type PatchReportByIDOK struct {
	Payload models.ReportBase
}

func (o *PatchReportByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /report/reports/{id}][%d] patchReportByIdOK  %+v", 200, o.Payload)
}

func (o *PatchReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalReportBase(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchReportByIDDefault creates a PatchReportByIDDefault with default headers values
func NewPatchReportByIDDefault(code int) *PatchReportByIDDefault {
	return &PatchReportByIDDefault{
		_statusCode: code,
	}
}

/*PatchReportByIDDefault handles this case with default header values.

Error
*/
type PatchReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch report by Id default response
func (o *PatchReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchReportByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /report/reports/{id}][%d] patchReportById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}