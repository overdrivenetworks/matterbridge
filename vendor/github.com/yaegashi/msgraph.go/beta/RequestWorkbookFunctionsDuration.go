// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsDurationRequestBuilder struct{ BaseRequestBuilder }

// Duration action undocumented
func (b *WorkbookFunctionsRequestBuilder) Duration(reqObj *WorkbookFunctionsDurationRequestParameter) *WorkbookFunctionsDurationRequestBuilder {
	bb := &WorkbookFunctionsDurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/duration"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDurationRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDurationRequestBuilder) Request() *WorkbookFunctionsDurationRequest {
	return &WorkbookFunctionsDurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsDurationRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
