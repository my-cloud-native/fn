// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package iotdataplane provides a client for AWS IoT Data Plane.
package iotdataplane

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

const opDeleteThingShadow = "DeleteThingShadow"

// DeleteThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the DeleteThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteThingShadowRequest method.
//    req, resp := client.DeleteThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *IoTDataPlane) DeleteThingShadowRequest(input *DeleteThingShadowInput) (req *request.Request, output *DeleteThingShadowOutput) {
	op := &request.Operation{
		Name:       opDeleteThingShadow,
		HTTPMethod: "DELETE",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &DeleteThingShadowInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteThingShadowOutput{}
	req.Data = output
	return
}

// Deletes the thing shadow for the specified thing.
//
// For more information, see DeleteThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_DeleteThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *IoTDataPlane) DeleteThingShadow(input *DeleteThingShadowInput) (*DeleteThingShadowOutput, error) {
	req, out := c.DeleteThingShadowRequest(input)
	err := req.Send()
	return out, err
}

const opGetThingShadow = "GetThingShadow"

// GetThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the GetThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the GetThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the GetThingShadowRequest method.
//    req, resp := client.GetThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *IoTDataPlane) GetThingShadowRequest(input *GetThingShadowInput) (req *request.Request, output *GetThingShadowOutput) {
	op := &request.Operation{
		Name:       opGetThingShadow,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &GetThingShadowInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetThingShadowOutput{}
	req.Data = output
	return
}

// Gets the thing shadow for the specified thing.
//
// For more information, see GetThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *IoTDataPlane) GetThingShadow(input *GetThingShadowInput) (*GetThingShadowOutput, error) {
	req, out := c.GetThingShadowRequest(input)
	err := req.Send()
	return out, err
}

const opPublish = "Publish"

// PublishRequest generates a "aws/request.Request" representing the
// client's request for the Publish operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the Publish method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PublishRequest method.
//    req, resp := client.PublishRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *IoTDataPlane) PublishRequest(input *PublishInput) (req *request.Request, output *PublishOutput) {
	op := &request.Operation{
		Name:       opPublish,
		HTTPMethod: "POST",
		HTTPPath:   "/topics/{topic}",
	}

	if input == nil {
		input = &PublishInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &PublishOutput{}
	req.Data = output
	return
}

// Publishes state information.
//
// For more information, see HTTP Protocol (http://docs.aws.amazon.com/iot/latest/developerguide/protocols.html#http)
// in the AWS IoT Developer Guide.
func (c *IoTDataPlane) Publish(input *PublishInput) (*PublishOutput, error) {
	req, out := c.PublishRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateThingShadow = "UpdateThingShadow"

// UpdateThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the UpdateThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the UpdateThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the UpdateThingShadowRequest method.
//    req, resp := client.UpdateThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *IoTDataPlane) UpdateThingShadowRequest(input *UpdateThingShadowInput) (req *request.Request, output *UpdateThingShadowOutput) {
	op := &request.Operation{
		Name:       opUpdateThingShadow,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &UpdateThingShadowInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateThingShadowOutput{}
	req.Data = output
	return
}

// Updates the thing shadow for the specified thing.
//
// For more information, see UpdateThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *IoTDataPlane) UpdateThingShadow(input *UpdateThingShadowInput) (*UpdateThingShadowOutput, error) {
	req, out := c.UpdateThingShadowRequest(input)
	err := req.Send()
	return out, err
}

// The input for the DeleteThingShadow operation.
type DeleteThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteThingShadowInput"}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the DeleteThingShadow operation.
type DeleteThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteThingShadowOutput) GoString() string {
	return s.String()
}

// The input for the GetThingShadow operation.
type GetThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetThingShadowInput"}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the GetThingShadow operation.
type GetThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s GetThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetThingShadowOutput) GoString() string {
	return s.String()
}

// The input for the Publish operation.
type PublishInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`

	// The Quality of Service (QoS) level.
	Qos *int64 `location:"querystring" locationName:"qos" type:"integer"`

	// The name of the MQTT topic.
	Topic *string `location:"uri" locationName:"topic" type:"string" required:"true"`
}

// String returns the string representation
func (s PublishInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PublishInput"}
	if s.Topic == nil {
		invalidParams.Add(request.NewErrParamRequired("Topic"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PublishOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PublishOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishOutput) GoString() string {
	return s.String()
}

// The input for the UpdateThingShadow operation.
type UpdateThingShadowInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob" required:"true"`

	// The name of the thing.
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateThingShadowInput"}
	if s.Payload == nil {
		invalidParams.Add(request.NewErrParamRequired("Payload"))
	}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the UpdateThingShadow operation.
type UpdateThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s UpdateThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateThingShadowOutput) GoString() string {
	return s.String()
}
