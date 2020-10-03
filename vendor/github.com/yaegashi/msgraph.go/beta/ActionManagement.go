// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ManagementConditionStatements returns request builder for ManagementConditionStatement collection
func (b *ManagementConditionRequestBuilder) ManagementConditionStatements() *ManagementConditionManagementConditionStatementsCollectionRequestBuilder {
	bb := &ManagementConditionManagementConditionStatementsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managementConditionStatements"
	return bb
}

// ManagementConditionManagementConditionStatementsCollectionRequestBuilder is request builder for ManagementConditionStatement collection
type ManagementConditionManagementConditionStatementsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagementConditionStatement collection
func (b *ManagementConditionManagementConditionStatementsCollectionRequestBuilder) Request() *ManagementConditionManagementConditionStatementsCollectionRequest {
	return &ManagementConditionManagementConditionStatementsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagementConditionStatement item
func (b *ManagementConditionManagementConditionStatementsCollectionRequestBuilder) ID(id string) *ManagementConditionStatementRequestBuilder {
	bb := &ManagementConditionStatementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagementConditionManagementConditionStatementsCollectionRequest is request for ManagementConditionStatement collection
type ManagementConditionManagementConditionStatementsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagementConditionStatement collection
func (r *ManagementConditionManagementConditionStatementsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagementConditionStatement, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagementConditionStatement
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagementConditionStatement
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagementConditionStatement collection, max N pages
func (r *ManagementConditionManagementConditionStatementsCollectionRequest) GetN(ctx context.Context, n int) ([]ManagementConditionStatement, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagementConditionStatement collection
func (r *ManagementConditionManagementConditionStatementsCollectionRequest) Get(ctx context.Context) ([]ManagementConditionStatement, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagementConditionStatement collection
func (r *ManagementConditionManagementConditionStatementsCollectionRequest) Add(ctx context.Context, reqObj *ManagementConditionStatement) (resObj *ManagementConditionStatement, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ManagementConditions returns request builder for ManagementCondition collection
func (b *ManagementConditionStatementRequestBuilder) ManagementConditions() *ManagementConditionStatementManagementConditionsCollectionRequestBuilder {
	bb := &ManagementConditionStatementManagementConditionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managementConditions"
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequestBuilder is request builder for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagementCondition collection
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) Request() *ManagementConditionStatementManagementConditionsCollectionRequest {
	return &ManagementConditionStatementManagementConditionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagementCondition item
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) ID(id string) *ManagementConditionRequestBuilder {
	bb := &ManagementConditionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequest is request for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagementCondition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagementCondition
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ManagementCondition
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ManagementCondition collection, max N pages
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) GetN(ctx context.Context, n int) ([]ManagementCondition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Get(ctx context.Context) ([]ManagementCondition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Add(ctx context.Context, reqObj *ManagementCondition) (resObj *ManagementCondition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
