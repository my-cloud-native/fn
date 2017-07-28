package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppsAppCallsParams creates a new GetAppsAppCallsParams object
// with the default values initialized.
func NewGetAppsAppCallsParams() *GetAppsAppCallsParams {
	var ()
	return &GetAppsAppCallsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsAppCallsParamsWithTimeout creates a new GetAppsAppCallsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsAppCallsParamsWithTimeout(timeout time.Duration) *GetAppsAppCallsParams {
	var ()
	return &GetAppsAppCallsParams{

		timeout: timeout,
	}
}

// NewGetAppsAppCallsParamsWithContext creates a new GetAppsAppCallsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsAppCallsParamsWithContext(ctx context.Context) *GetAppsAppCallsParams {
	var ()
	return &GetAppsAppCallsParams{

		Context: ctx,
	}
}

// NewGetAppsAppCallsParamsWithHTTPClient creates a new GetAppsAppCallsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppsAppCallsParamsWithHTTPClient(client *http.Client) *GetAppsAppCallsParams {
	var ()
	return &GetAppsAppCallsParams{
		HTTPClient: client,
	}
}

/*GetAppsAppCallsParams contains all the parameters to send to the API endpoint
for the get apps app calls operation typically these are written to a http.Request
*/
type GetAppsAppCallsParams struct {

	/*App
	  App name.

	*/
	App string
	/*Route
	  App route.

	*/
	Route *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apps app calls params
func (o *GetAppsAppCallsParams) WithTimeout(timeout time.Duration) *GetAppsAppCallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps app calls params
func (o *GetAppsAppCallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps app calls params
func (o *GetAppsAppCallsParams) WithContext(ctx context.Context) *GetAppsAppCallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps app calls params
func (o *GetAppsAppCallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apps app calls params
func (o *GetAppsAppCallsParams) WithHTTPClient(client *http.Client) *GetAppsAppCallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apps app calls params
func (o *GetAppsAppCallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the get apps app calls params
func (o *GetAppsAppCallsParams) WithApp(app string) *GetAppsAppCallsParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the get apps app calls params
func (o *GetAppsAppCallsParams) SetApp(app string) {
	o.App = app
}

// WithRoute adds the route to the get apps app calls params
func (o *GetAppsAppCallsParams) WithRoute(route *string) *GetAppsAppCallsParams {
	o.SetRoute(route)
	return o
}

// SetRoute adds the route to the get apps app calls params
func (o *GetAppsAppCallsParams) SetRoute(route *string) {
	o.Route = route
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsAppCallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if o.Route != nil {

		// query param route
		var qrRoute string
		if o.Route != nil {
			qrRoute = *o.Route
		}
		qRoute := qrRoute
		if qRoute != "" {
			if err := r.SetQueryParam("route", qRoute); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
