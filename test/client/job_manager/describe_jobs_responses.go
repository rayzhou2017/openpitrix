// Code generated by go-swagger; DO NOT EDIT.

package job_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeJobsReader is a Reader for the DescribeJobs structure.
type DescribeJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeJobsOK creates a DescribeJobsOK with default headers values
func NewDescribeJobsOK() *DescribeJobsOK {
	return &DescribeJobsOK{}
}

/*DescribeJobsOK handles this case with default header values.

DescribeJobsOK describe jobs o k
*/
type DescribeJobsOK struct {
	Payload *models.OpenpitrixDescribeJobsResponse
}

func (o *DescribeJobsOK) Error() string {
	return fmt.Sprintf("[GET /v1/jobs][%d] describeJobsOK  %+v", 200, o.Payload)
}

func (o *DescribeJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
