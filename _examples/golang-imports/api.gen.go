// example-api-service v1.0.0 eb9a5d4082a36a8cb84eaa4b3e1091fc4f1f6f3d
// --
// Code generated by webrpc-gen@v0.18.2 with ../../../gen-golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./proto/api.ridl -target=../../../gen-golang -out=./api.gen.go -pkg=main -server -client -legacyErrors=true -fmt=false
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v1.0.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "eb9a5d4082a36a8cb84eaa4b3e1091fc4f1f6f3d"
}

//
// Common types
//

type User struct {
		Username string `json:"username"`
		Age uint32 `json:"age"`
}

type Location uint32

const (
	Location_TORONTO Location = 0
	Location_NEW_YORK Location = 1
)

var Location_name = map[uint32]string{
	0: "TORONTO",
	1: "NEW_YORK",
}

var Location_value = map[string]uint32{
	"TORONTO": 0,
	"NEW_YORK": 1,
}

func (x Location) String() string {
	return Location_name[uint32(x)]
}

func (x Location) MarshalText() ([]byte, error) {
	return []byte(Location_name[uint32(x)]), nil
}

func (x *Location) UnmarshalText(b []byte) error {
	*x = Location(Location_value[string(b)])
	return nil
}

func (x *Location) Is(values ...Location) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

var WebRPCServices = map[string][]string{
	"ExampleAPI": {
		"Ping",
		"Status",
		"GetUsers",
	},
}

//
// Server types
//

type ExampleAPI interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUsers(ctx context.Context) ([]*User, Location, error)
}




//
// Client types
//

type ExampleAPIClient interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUsers(ctx context.Context) ([]*User, Location, error)
}




//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type exampleAPIServer struct {
	ExampleAPI
	OnError func(r *http.Request, rpcErr *WebRPCError)
}

func NewExampleAPIServer(svc ExampleAPI) *exampleAPIServer {
	return &exampleAPIServer{
		ExampleAPI: svc,
	}
}

func (s *exampleAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// In case of a panic, serve a HTTP 500 error and then panic.
		if rr := recover(); rr != nil {
			s.sendErrorJSON(w, r, ErrWebrpcServerPanic.WithCause(fmt.Errorf("%v", rr)))
			panic(rr)
		}
	}()

	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "ExampleAPI")

	var handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
	switch r.URL.Path {
	case "/rpc/ExampleAPI/Ping":
		handler = s.servePingJSON
	case "/rpc/ExampleAPI/Status":
		handler = s.serveStatusJSON
	case "/rpc/ExampleAPI/GetUsers":
		handler = s.serveGetUsersJSON
	default:
		err := ErrWebrpcBadRoute.WithCause(fmt.Errorf("no handler for path %q", r.URL.Path))
		s.sendErrorJSON(w, r, err)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST") // RFC 9110.
		err := ErrWebrpcBadMethod.WithCause(fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		s.sendErrorJSON(w, r, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = contentType[:i]
	}
	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case "application/json":
		handler(ctx, w, r)
	default:
		err := ErrWebrpcBadRequest.WithCause(fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		s.sendErrorJSON(w, r, err)
	}
}

func (s *exampleAPIServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

	

	// Call service method implementation.
	err := s.ExampleAPI.Ping(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *exampleAPIServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Status")

	

	// Call service method implementation.
	ret0, err := s.ExampleAPI.Status(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 bool `json:"status"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleAPIServer) serveGetUsersJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetUsers")

	

	// Call service method implementation.
	ret0, ret1, err := s.ExampleAPI.GetUsers(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 []*User `json:"users"`
		Ret1 Location `json:"location"`
	}{ret0, ret1}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}


func (s *exampleAPIServer) sendErrorJSON(w http.ResponseWriter, r *http.Request, rpcErr WebRPCError) {
	if s.OnError != nil {
		s.OnError(r, &rpcErr)
	}

	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrWebrpcEndpoint.WithCause(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}



//
// Client
//

const ExampleAPIPathPrefix = "/rpc/ExampleAPI/"

type exampleAPIClient struct {
	client HTTPClient
	urls	 [3]string
}

func NewExampleAPIClient(addr string, client HTTPClient) ExampleAPIClient {
	prefix := urlBase(addr) + ExampleAPIPathPrefix
	urls := [3]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "GetUsers",
	}
	return &exampleAPIClient{
		client: client,
		urls:	 urls,
	}
}

func (c *exampleAPIClient) Ping(ctx context.Context) error {

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *exampleAPIClient) Status(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *exampleAPIClient) GetUsers(ctx context.Context) ([]*User, Location, error) {
	out := struct {
		Ret0 []*User `json:"users"`
		Ret1 Location `json:"location"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, out.Ret1, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("could not build request: %w", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return resp, nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}
func ResponseWriterFromContext(ctx context.Context) http.ResponseWriter {
	w, _ := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter)
	return w
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if target == nil {
		return false
	}
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	if legacyErr, ok := target.(legacyError); ok {
		return legacyErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

func (e WebRPCError) WithCausef(format string, args ...interface{}) WebRPCError {
	cause := fmt.Errorf(format, args...)
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)

//
// Legacy errors
//

// Deprecated: Use fmt.Errorf() or WebRPCError.
func Errorf(err legacyError, format string, args ...interface{}) WebRPCError {
	return err.WebRPCError.WithCause(fmt.Errorf(format, args...))
}

// Deprecated: Use .WithCause() method on WebRPCError.
func WrapError(err legacyError, cause error, format string, args ...interface{}) WebRPCError {
	return err.WebRPCError.WithCause(fmt.Errorf("%v: %w", fmt.Errorf(format, args...), cause))
}

// Deprecated: Use fmt.Errorf() or WebRPCError.
func Failf(format string, args ...interface{}) WebRPCError {
	return Errorf(ErrFail, format, args...)
}

// Deprecated: Use .WithCause() method on WebRPCError.
func WrapFailf(cause error, format string, args ...interface{}) WebRPCError {
	return WrapError(ErrFail, cause, format, args...)
}

// Deprecated: Use fmt.Errorf() or WebRPCError.
func ErrorNotFound(format string, args ...interface{}) WebRPCError {
	return Errorf(ErrNotFound, format, args...)
}

// Deprecated: Use fmt.Errorf() or WebRPCError.
func ErrorInvalidArgument(argument string, validationMsg string) WebRPCError {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

// Deprecated: Use fmt.Errorf() or WebRPCError.
func ErrorRequiredArgument(argument string) WebRPCError {
	return ErrorInvalidArgument(argument, "is required")
}

// Deprecated: Use fmt.Errorf() or WebRPCError.
func ErrorInternal(format string, args ...interface{}) WebRPCError {
	return Errorf(ErrInternal, format, args...)
}

type legacyError struct { WebRPCError }

// Legacy errors (webrpc v0.10.0 and earlier). Will be removed.
var (
	// Deprecated. Define errors in RIDL schema.
	ErrCanceled = legacyError{WebRPCError{Code: -10000, Name: "ErrCanceled", Message: "canceled", HTTPStatus: 408 /* RequestTimeout */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrUnknown = legacyError{WebRPCError{Code: -10001, Name: "ErrUnknown", Message: "unknown", HTTPStatus: 400 /* Bad Request */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrFail = legacyError{WebRPCError{Code: -10002, Name: "ErrFail", Message: "fail", HTTPStatus: 422 /* Unprocessable Entity */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrInvalidArgument = legacyError{WebRPCError{Code: -10003, Name: "ErrInvalidArgument", Message: "invalid argument", HTTPStatus: 400 /* BadRequest */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrDeadlineExceeded = legacyError{WebRPCError{Code: -10004, Name: "ErrDeadlineExceeded", Message: "deadline exceeded", HTTPStatus: 408 /* RequestTimeout */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrNotFound = legacyError{WebRPCError{Code: -10005, Name: "ErrNotFound", Message: "not found", HTTPStatus: 404 /* Not Found */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrBadRoute = legacyError{WebRPCError{Code: -10006, Name: "ErrBadRoute", Message: "bad route", HTTPStatus: 404 /* Not Found */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrAlreadyExists = legacyError{WebRPCError{Code: -10007, Name: "ErrAlreadyExists", Message: "already exists", HTTPStatus: 409 /* Conflict */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrPermissionDenied = legacyError{WebRPCError{Code: -10008, Name: "ErrPermissionDenied", Message: "permission denied", HTTPStatus: 403 /* Forbidden */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrUnauthenticated = legacyError{WebRPCError{Code: -10009, Name: "ErrUnauthenticated", Message: "unauthenticated", HTTPStatus: 401 /* Unauthorized */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrResourceExhausted = legacyError{WebRPCError{Code: -10010, Name: "ErrResourceExhausted", Message: "resource exhausted", HTTPStatus: 403 /* Forbidden */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrFailedPrecondition = legacyError{WebRPCError{Code: -10011, Name: "ErrFailedPrecondition", Message: "failed precondition", HTTPStatus: 412 /* Precondition Failed */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrAborted = legacyError{WebRPCError{Code: -10012, Name: "ErrAborted", Message: "aborted", HTTPStatus: 409 /* Conflict */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrOutOfRange = legacyError{WebRPCError{Code: -10013, Name: "ErrOutOfRange", Message: "out of range", HTTPStatus: 400 /* Bad Request */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrUnimplemented = legacyError{WebRPCError{Code: -10014, Name: "ErrUnimplemented", Message: "unimplemented", HTTPStatus: 501 /* Not Implemented */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrInternal = legacyError{WebRPCError{Code: -10015, Name: "ErrInternal", Message: "internal", HTTPStatus: 500 /* Internal Server Error */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrUnavailable = legacyError{WebRPCError{Code: -10016, Name: "ErrUnavailable", Message: "unavailable", HTTPStatus: 503 /* Service Unavailable */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrDataLoss = legacyError{WebRPCError{Code: -10017, Name: "ErrDataLoss", Message: "data loss", HTTPStatus: 500 /* Internal Server Error */ }}
	// Deprecated. Define errors in RIDL schema.
	ErrNone = legacyError{WebRPCError{Code: -10018, Name: "ErrNone", Message: "", HTTPStatus: 200 /* OK */ }}
)