// example v0.0.1 87cdac57aac886a37b5caffc5962c93859970b68
// --
// Code generated by webrpc-gen@v0.12.x-dev with custom generator. DO NOT EDIT.
//
// webrpc-gen -schema=example.ridl -target=../../../gen-golang -pkg=main -server -client -out=./example.gen.go -fmt=false -legacyErrors=true
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	
	"github.com/google/uuid"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.0.1"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "87cdac57aac886a37b5caffc5962c93859970b68"
}

//
// Types
//

type Kind uint32

const (
	Kind_USER Kind = 0
	Kind_ADMIN Kind = 1
)

var Kind_name = map[uint32]string{
	0: "USER",
	1: "ADMIN",
}

var Kind_value = map[string]uint32{
	"USER": 0,
	"ADMIN": 1,
}

func (x Kind) String() string {
	return Kind_name[uint32(x)]
}

func (x Kind) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(Kind_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *Kind) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = Kind(Kind_value[j])
	return nil
}

type User struct {
	ID uint64 `json:"id" db:"id"`
	Uuid uuid.UUID `json:"uuid" db:"id"`
	Username string `json:"USERNAME" db:"username"`
	Role string `json:"role" db:"-"`
	Nicknames []Nickname `json:"nicknames" db:"-"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type Nickname struct {
	ID uint64 `json:"ID" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type SearchFilter struct {
	Q string `json:"q"`
}

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash string `json:"schemaHash"`
}

type ComplexType struct {
	Meta map[string]interface{} `json:"meta"`
	MetaNestedExample map[string]map[string]uint32 `json:"metaNestedExample"`
	NamesList []string `json:"namesList"`
	NumsList []int64 `json:"numsList"`
	DoubleArray [][]string `json:"doubleArray"`
	ListOfMaps []map[string]uint32 `json:"listOfMaps"`
	ListOfUsers []*User `json:"listOfUsers"`
	MapOfUsers map[string]*User `json:"mapOfUsers"`
	User *User `json:"user"`
}

type ExampleService interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	GetUser(ctx context.Context, header map[string]string, userID uint64) (*User, error)
	FindUser(ctx context.Context, s *SearchFilter) (string, *User, error)
}

var WebRPCServices = map[string][]string{
	"ExampleService": {
		"Ping",
		"Status",
		"Version",
		"GetUser",
		"FindUser",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type exampleServiceServer struct {
	ExampleService
}

func NewExampleServiceServer(svc ExampleService) WebRPCServer {
	return &exampleServiceServer{
		ExampleService: svc,
	}
}

func (s *exampleServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "ExampleService")

	if r.Method != "POST" {
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/ExampleService/Ping":
		s.servePing(ctx, w, r)
		return
	case "/rpc/ExampleService/Status":
		s.serveStatus(ctx, w, r)
		return
	case "/rpc/ExampleService/Version":
		s.serveVersion(ctx, w, r)
		return
	case "/rpc/ExampleService/GetUser":
		s.serveGetUser(ctx, w, r)
		return
	case "/rpc/ExampleService/FindUser":
		s.serveFindUser(ctx, w, r)
		return
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}
}

func (s *exampleServiceServer) servePing(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePingJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		err = s.ExampleService.Ping(ctx)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *exampleServiceServer) serveStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveStatusJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Status")

	// Call service method
	var ret0 bool
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleService.Status(ctx)
	}()
	respContent := struct {
		Ret0 bool `json:"status"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleServiceServer) serveVersion(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveVersionJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) serveVersionJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Version")

	// Call service method
	var ret0 *Version
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleService.Version(ctx)
	}()
	respContent := struct {
		Ret0 *Version `json:"version"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleServiceServer) serveGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUserJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) serveGetUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetUser")
	reqContent := struct {
		Arg0 map[string]string `json:"header"`
		Arg1 uint64 `json:"userID"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleService.GetUser(ctx, reqContent.Arg0, reqContent.Arg1)
	}()
	respContent := struct {
		Ret0 *User `json:"user"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleServiceServer) serveFindUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFindUserJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) serveFindUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "FindUser")
	reqContent := struct {
		Arg0 *SearchFilter `json:"s"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 string
	var ret1 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, ret1, err = s.ExampleService.FindUser(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 string `json:"name"`
		Ret1 *User `json:"user"`
	}{ret0, ret1}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrorWithCause(ErrWebrpcEndpoint, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

//
// Client
//

const ExampleServicePathPrefix = "/rpc/ExampleService/"

type exampleServiceClient struct {
	client HTTPClient
	urls	 [5]string
}

func NewExampleServiceClient(addr string, client HTTPClient) ExampleService {
	prefix := urlBase(addr) + ExampleServicePathPrefix
	urls := [5]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "Version",
		prefix + "GetUser",
		prefix + "FindUser",
	}
	return &exampleServiceClient{
		client: client,
		urls:	 urls,
	}
}

func (c *exampleServiceClient) Ping(ctx context.Context) error {

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, nil)
	return err
}

func (c *exampleServiceClient) Status(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *exampleServiceClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, err
}

func (c *exampleServiceClient) GetUser(ctx context.Context, header map[string]string, userID uint64) (*User, error) {
	in := struct {
		Arg0 map[string]string `json:"header"`
		Arg1 uint64 `json:"userID"`
	}{header, userID}
	out := struct {
		Ret0 *User `json:"user"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], in, &out)
	return out.Ret0, err
}

func (c *exampleServiceClient) FindUser(ctx context.Context, s *SearchFilter) (string, *User, error) {
	in := struct {
		Arg0 *SearchFilter `json:"s"`
	}{s}
	out := struct {
		Ret0 string `json:"name"`
		Ret1 *User `json:"user"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], in, &out)
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
	req, err := http.NewRequest("POST", url, reqBody)
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

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("could not build request: %w", err))
	}
	resp, err := client.Do(req)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()

	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	if resp.StatusCode != 200 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return rpcErr
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return nil
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
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

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

func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	err := rpcErr
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Webrpc errors
var (
	ErrWebrpcEndpoint = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 0}
	ErrWebrpcBadRoute = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
)

// Schema errors
var (
	ErrMissingArgument = WebRPCError{Code: 500100, Name: "MissingArgument", Message: "missing argument", HTTPStatus: 400}
	ErrInvalidUsername = WebRPCError{Code: 500101, Name: "InvalidUsername", Message: "invalid username", HTTPStatus: 400}
	ErrMemoryFull = WebRPCError{Code: 400100, Name: "MemoryFull", Message: "system memory is full", HTTPStatus: 400}
	ErrUnauthorized = WebRPCError{Code: 400200, Name: "Unauthorized", Message: "unauthorized", HTTPStatus: 401}
	ErrUserNotFound = WebRPCError{Code: 400300, Name: "UserNotFound", Message: "user not found", HTTPStatus: 400}
)

//
// Legacy errors
//

// Deprecated: Use ErrorWithCause() instead.
func Errorf(err legacyError, format string, args ...interface{}) WebRPCError {
	return ErrorWithCause(err.WebRPCError, fmt.Errorf(format, args...))
}

// Deprecated: Use ErrorWithCause() instead.
func WrapError(err legacyError, cause error, format string, args ...interface{}) WebRPCError {
	return ErrorWithCause(err.WebRPCError, fmt.Errorf("%v: %w", fmt.Errorf(format, args...), cause))
}

// Deprecated: Use ErrorWithCause() instead.
func Failf(format string, args ...interface{}) WebRPCError {
	return Errorf(ErrFail, format, args...)
}

// Deprecated: Use ErrorWithCause() instead.
func WrapFailf(cause error, format string, args ...interface{}) WebRPCError {
	return WrapError(ErrFail, cause, format, args...)
}

// Deprecated: Use ErrorWithCause() instead.
func ErrorNotFound(format string, args ...interface{}) WebRPCError {
	return Errorf(ErrNotFound, format, args...)
}

// Deprecated: Use ErrorWithCause() instead.
func ErrorInvalidArgument(argument string, validationMsg string) WebRPCError {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

// Deprecated: Use ErrorWithCause() instead.
func ErrorRequiredArgument(argument string) WebRPCError {
	return ErrorInvalidArgument(argument, "is required")
}

// Deprecated: Use ErrorWithCause() instead.
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

