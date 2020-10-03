// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsCeiling_MathRequestBuilder struct{ BaseRequestBuilder }

// Ceiling_Math action undocumented
func (b *WorkbookFunctionsRequestBuilder) Ceiling_Math(reqObj *WorkbookFunctionsCeiling_MathRequestParameter) *WorkbookFunctionsCeiling_MathRequestBuilder {
	bb := &WorkbookFunctionsCeiling_MathRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ceiling_Math"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCeiling_MathRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCeiling_MathRequestBuilder) Request() *WorkbookFunctionsCeiling_MathRequest {
	return &WorkbookFunctionsCeiling_MathRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCeiling_MathRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCeiling_PreciseRequestBuilder struct{ BaseRequestBuilder }

// Ceiling_Precise action undocumented
func (b *WorkbookFunctionsRequestBuilder) Ceiling_Precise(reqObj *WorkbookFunctionsCeiling_PreciseRequestParameter) *WorkbookFunctionsCeiling_PreciseRequestBuilder {
	bb := &WorkbookFunctionsCeiling_PreciseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ceiling_Precise"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCeiling_PreciseRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCeiling_PreciseRequestBuilder) Request() *WorkbookFunctionsCeiling_PreciseRequest {
	return &WorkbookFunctionsCeiling_PreciseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCeiling_PreciseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
