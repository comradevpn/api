// Package client_service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package client_service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	externalRef1 "github.com/vpnhouse/api/go/server/common"
)

const (
	BearerScopes = "bearer.Scopes"
)

// Auth defines model for Auth.
type Auth struct {
	AuthMethodId *string                 `json:"auth_method_id,omitempty"`
	CreatedAt    *time.Time              `json:"created_at,omitempty"`
	ExtendedInfo *map[string]interface{} `json:"extended_info,omitempty"`
	Id           *string                 `json:"id,omitempty"`
	Identifier   *string                 `json:"identifier,omitempty"`
	UpdatedAt    *time.Time              `json:"updated_at,omitempty"`
	UserId       *string                 `json:"user_id,omitempty"`
}

// AuthMethod defines model for AuthMethod.
type AuthMethod struct {
	CreatedAt *time.Time              `json:"created_at,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	ProjectId *string                 `json:"project_id,omitempty"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
}

// Invite defines model for Invite.
type Invite struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Email       *string                 `json:"email,omitempty"`
	ExpiresAt   *time.Time              `json:"expires_at,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	LocationId  *string                 `json:"location_id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	QueryId     *string                 `json:"query_id,omitempty"`
	QueryParams *map[string]interface{} `json:"query_params,omitempty"`
	Reminded    *bool                   `json:"reminded,omitempty"`
	Telegram    *string                 `json:"telegram,omitempty"`
	TokenId     *string                 `json:"token_id,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
	UserId      *string                 `json:"user_id,omitempty"`
}

// Mailing defines model for Mailing.
type Mailing struct {
	AcceptId   *string    `json:"accept_id,omitempty"`
	Accepted   *bool      `json:"accepted,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Id         *string    `json:"id,omitempty"`
	MailingTag *string    `json:"mailing_tag,omitempty"`
	Status     *string    `json:"status,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Description *map[string]interface{} `json:"description,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
}

// Session defines model for Session.
type Session struct {
	Connected        *bool      `json:"connected,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	Deleted          *bool      `json:"deleted,omitempty"`
	ExpiresAt        *time.Time `json:"expires_at,omitempty"`
	FirstConnectedAt *time.Time `json:"first_connected_at,omitempty"`
	Id               *string    `json:"id,omitempty"`
	Label            *string    `json:"label,omitempty"`
	Node             *string    `json:"node,omitempty"`
	PeerId           *int       `json:"peer_id,omitempty"`
	ToDelete         *bool      `json:"to_delete,omitempty"`
	TokenId          *string    `json:"token_id,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// User defines model for User.
type User struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Description *map[string]interface{} `json:"description,omitempty"`
	Email       *string                 `json:"email,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	ProjectId   *string                 `json:"project_id,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetUser request
	GetUser(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUserAuth request
	ListUserAuth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUserAuthMethod request
	ListUserAuthMethod(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUserInvite request
	ListUserInvite(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUserMailing request
	ListUserMailing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUserProject request
	GetUserProject(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUserSession request
	ListUserSession(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUser(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUserAuth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUserAuthRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUserAuthMethod(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUserAuthMethodRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUserInvite(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUserInviteRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUserMailing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUserMailingRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserProject(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserProjectRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUserSession(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUserSessionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUserRequest generates requests for GetUser
func NewGetUserRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListUserAuthRequest generates requests for ListUserAuth
func NewListUserAuthRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/auth")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListUserAuthMethodRequest generates requests for ListUserAuthMethod
func NewListUserAuthMethodRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/auth-method")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListUserInviteRequest generates requests for ListUserInvite
func NewListUserInviteRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/invite")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListUserMailingRequest generates requests for ListUserMailing
func NewListUserMailingRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/mailing")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUserProjectRequest generates requests for GetUserProject
func NewGetUserProjectRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/project")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListUserSessionRequest generates requests for ListUserSession
func NewListUserSessionRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/client-service/user/session")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetUser request
	GetUserWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUserResponse, error)

	// ListUserAuth request
	ListUserAuthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserAuthResponse, error)

	// ListUserAuthMethod request
	ListUserAuthMethodWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserAuthMethodResponse, error)

	// ListUserInvite request
	ListUserInviteWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserInviteResponse, error)

	// ListUserMailing request
	ListUserMailingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserMailingResponse, error)

	// GetUserProject request
	GetUserProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUserProjectResponse, error)

	// ListUserSession request
	ListUserSessionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserSessionResponse, error)
}

type GetUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r GetUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUserAuthResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Auth
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListUserAuthResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUserAuthResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUserAuthMethodResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]AuthMethod
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListUserAuthMethodResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUserAuthMethodResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUserInviteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Invite
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListUserInviteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUserInviteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUserMailingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Mailing
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListUserMailingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUserMailingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserProjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Project
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r GetUserProjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserProjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUserSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Session
	JSON401      *externalRef1.Error
	JSON403      *externalRef1.Error
	JSON500      *externalRef1.Error
}

// Status returns HTTPResponse.Status
func (r ListUserSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUserSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUserWithResponse request returning *GetUserResponse
func (c *ClientWithResponses) GetUserWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUserResponse, error) {
	rsp, err := c.GetUser(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserResponse(rsp)
}

// ListUserAuthWithResponse request returning *ListUserAuthResponse
func (c *ClientWithResponses) ListUserAuthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserAuthResponse, error) {
	rsp, err := c.ListUserAuth(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUserAuthResponse(rsp)
}

// ListUserAuthMethodWithResponse request returning *ListUserAuthMethodResponse
func (c *ClientWithResponses) ListUserAuthMethodWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserAuthMethodResponse, error) {
	rsp, err := c.ListUserAuthMethod(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUserAuthMethodResponse(rsp)
}

// ListUserInviteWithResponse request returning *ListUserInviteResponse
func (c *ClientWithResponses) ListUserInviteWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserInviteResponse, error) {
	rsp, err := c.ListUserInvite(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUserInviteResponse(rsp)
}

// ListUserMailingWithResponse request returning *ListUserMailingResponse
func (c *ClientWithResponses) ListUserMailingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserMailingResponse, error) {
	rsp, err := c.ListUserMailing(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUserMailingResponse(rsp)
}

// GetUserProjectWithResponse request returning *GetUserProjectResponse
func (c *ClientWithResponses) GetUserProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUserProjectResponse, error) {
	rsp, err := c.GetUserProject(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserProjectResponse(rsp)
}

// ListUserSessionWithResponse request returning *ListUserSessionResponse
func (c *ClientWithResponses) ListUserSessionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUserSessionResponse, error) {
	rsp, err := c.ListUserSession(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUserSessionResponse(rsp)
}

// ParseGetUserResponse parses an HTTP response from a GetUserWithResponse call
func ParseGetUserResponse(rsp *http.Response) (*GetUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListUserAuthResponse parses an HTTP response from a ListUserAuthWithResponse call
func ParseListUserAuthResponse(rsp *http.Response) (*ListUserAuthResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUserAuthResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Auth
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListUserAuthMethodResponse parses an HTTP response from a ListUserAuthMethodWithResponse call
func ParseListUserAuthMethodResponse(rsp *http.Response) (*ListUserAuthMethodResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUserAuthMethodResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []AuthMethod
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListUserInviteResponse parses an HTTP response from a ListUserInviteWithResponse call
func ParseListUserInviteResponse(rsp *http.Response) (*ListUserInviteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUserInviteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Invite
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListUserMailingResponse parses an HTTP response from a ListUserMailingWithResponse call
func ParseListUserMailingResponse(rsp *http.Response) (*ListUserMailingResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUserMailingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Mailing
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetUserProjectResponse parses an HTTP response from a GetUserProjectWithResponse call
func ParseGetUserProjectResponse(rsp *http.Response) (*GetUserProjectResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserProjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Project
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseListUserSessionResponse parses an HTTP response from a ListUserSessionWithResponse call
func ParseListUserSessionResponse(rsp *http.Response) (*ListUserSessionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUserSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Session
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef1.Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
