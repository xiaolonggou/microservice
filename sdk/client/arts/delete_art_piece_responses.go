// Code generated by go-swagger; DO NOT EDIT.

package arts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xiaolonggou/microservice/v1/sdk/models"
)

// DeleteArtPieceReader is a Reader for the DeleteArtPiece structure.
type DeleteArtPieceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtPieceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteArtPieceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteArtPieceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteArtPieceNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteArtPieceCreated creates a DeleteArtPieceCreated with default headers values
func NewDeleteArtPieceCreated() *DeleteArtPieceCreated {
	return &DeleteArtPieceCreated{}
}

/* DeleteArtPieceCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type DeleteArtPieceCreated struct {
}

func (o *DeleteArtPieceCreated) Error() string {
	return fmt.Sprintf("[DELETE /art/{id}][%d] deleteArtPieceCreated ", 201)
}

func (o *DeleteArtPieceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtPieceNotFound creates a DeleteArtPieceNotFound with default headers values
func NewDeleteArtPieceNotFound() *DeleteArtPieceNotFound {
	return &DeleteArtPieceNotFound{}
}

/* DeleteArtPieceNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type DeleteArtPieceNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteArtPieceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /art/{id}][%d] deleteArtPieceNotFound  %+v", 404, o.Payload)
}
func (o *DeleteArtPieceNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteArtPieceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtPieceNotImplemented creates a DeleteArtPieceNotImplemented with default headers values
func NewDeleteArtPieceNotImplemented() *DeleteArtPieceNotImplemented {
	return &DeleteArtPieceNotImplemented{}
}

/* DeleteArtPieceNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type DeleteArtPieceNotImplemented struct {
	Payload *models.GenericError
}

func (o *DeleteArtPieceNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /art/{id}][%d] deleteArtPieceNotImplemented  %+v", 501, o.Payload)
}
func (o *DeleteArtPieceNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteArtPieceNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
