// Package authorizer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package authorizer

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	ServiceKeyScopes = "ServiceKey.Scopes"
	BasicScopes      = "basic.Scopes"
	BearerScopes     = "bearer.Scopes"
)

// ApplyTrialLicenseRequest defines model for ApplyTrialLicenseRequest.
type ApplyTrialLicenseRequest struct {
	ProductId string `json:"product_id"`
}

// AuthRequest defines model for AuthRequest.
type AuthRequest struct {
	AuthInfo       string `json:"auth_info"`
	AuthType       string `json:"auth_type"`
	ClientVersion  string `json:"client_version"`
	DeviceId       string `json:"device_id"`
	InstallationId string `json:"installation_id"`
	PlatformType   string `json:"platform_type"`
	Project        string `json:"project"`
}

// AuthResp defines model for AuthResp.
type AuthResp struct {
	AccessToken        string                  `json:"access_token"`
	CreatedAt          *time.Time              `json:"created_at,omitempty"`
	DiscoveryAddresses *[]string               `json:"discovery_addresses,omitempty"`
	Email              *string                 `json:"email,omitempty"`
	Entitlements       *map[string]interface{} `json:"entitlements,omitempty"`
	ExpiresAt          *time.Time              `json:"expires_at,omitempty"`
	RefreshToken       *string                 `json:"refresh_token,omitempty"`
}

// AuthServiceRequest defines model for AuthServiceRequest.
type AuthServiceRequest struct {
	Project   string `json:"project"`
	ServiceId string `json:"service_id"`
}

// CreateFirebaseUserRequest defines model for CreateFirebaseUserRequest.
type CreateFirebaseUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	ProjectId string `json:"project_id"`
}

// CreatePurchaseContextRequest defines model for CreatePurchaseContextRequest.
type CreatePurchaseContextRequest struct {
	ProductId string  `json:"product_id"`
	ProjectId *string `json:"project_id,omitempty"`
	UserId    *string `json:"user_id,omitempty"`
}

// CreatePurchaseContextResp defines model for CreatePurchaseContextResp.
type CreatePurchaseContextResp struct {
	PurchaseContextId string `json:"purchase_context_id"`
}

// License defines model for License.
type License struct {
	CreatedAt        *time.Time              `json:"created_at,omitempty"`
	Disabled         *bool                   `json:"disabled,omitempty"`
	EndAt            *time.Time              `json:"end_at,omitempty"`
	EntitlementsJson *map[string]interface{} `json:"entitlements_json,omitempty"`
	Id               *string                 `json:"id,omitempty"`
	ProjectId        *string                 `json:"project_id,omitempty"`
	PurchaseJson     *map[string]interface{} `json:"purchase_json,omitempty"`
	SelectorJson     *map[string]interface{} `json:"selector_json,omitempty"`
	StartAt          *time.Time              `json:"start_at,omitempty"`
	UpdatedAt        *time.Time              `json:"updated_at,omitempty"`
	UserId           *string                 `json:"user_id,omitempty"`
}

// PaymentDetailsRequest defines model for PaymentDetailsRequest.
type PaymentDetailsRequest struct {
	Email          string  `json:"email"`
	GaId           *string `json:"ga_id,omitempty"`
	InstallationId *string `json:"installation_id,omitempty"`
	ProductId      string  `json:"product_id"`
	ProjectId      string  `json:"project_id"`
}

// PaymentDetailsResp defines model for PaymentDetailsResp.
type PaymentDetailsResp struct {
	PaymentUrl string `json:"payment_url"`
}

// ProcessAndroidPurchaseRequest defines model for ProcessAndroidPurchaseRequest.
type ProcessAndroidPurchaseRequest struct {
	OrderId           string `json:"order_id"`
	PackageName       string `json:"package_name"`
	PurchaseContextId string `json:"purchase_context_id"`
	PurchaseTime      int    `json:"purchase_time"`
	PurchaseToken     string `json:"purchase_token"`
	Signature         string `json:"signature"`
}

// ProcessIOSPurchaseRequest defines model for ProcessIOSPurchaseRequest.
type ProcessIOSPurchaseRequest struct {
	JwsReceipt        string `json:"jws_receipt"`
	PurchaseContextId string `json:"purchase_context_id"`
}

// Product defines model for Product.
type Product struct {
	CreatedAt        *time.Time              `json:"created_at,omitempty"`
	Disabled         *bool                   `json:"disabled,omitempty"`
	EntitlementsJson *map[string]interface{} `json:"entitlements_json,omitempty"`
	Id               *string                 `json:"id,omitempty"`
	LicenseType      *string                 `json:"license_type,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	PaymentJson      *map[string]interface{} `json:"payment_json,omitempty"`
	Period           *string                 `json:"period,omitempty"`
	SelectorJson     *map[string]interface{} `json:"selector_json,omitempty"`
	UpdatedAt        *time.Time              `json:"updated_at,omitempty"`
}

// SendRestoreLinkRequest defines model for SendRestoreLinkRequest.
type SendRestoreLinkRequest struct {
	Email     string `json:"email"`
	ProjectId string `json:"project_id"`
}

// TokenRequest defines model for TokenRequest.
type TokenRequest struct {
	InstallationId string `json:"installation_id"`
	PlatformType   string `json:"platform_type"`
	RefreshToken   string `json:"refresh_token"`
}

// TokenResp defines model for TokenResp.
type TokenResp struct {
	AccessToken        string                 `json:"access_token"`
	CreatedAt          time.Time              `json:"created_at"`
	DiscoveryAddresses *[]string              `json:"discovery_addresses,omitempty"`
	Email              *string                `json:"email,omitempty"`
	Entitlements       map[string]interface{} `json:"entitlements"`
	ExpiresAt          time.Time              `json:"expires_at"`
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

// ApplyTrialLicenseJSONBody defines parameters for ApplyTrialLicense.
type ApplyTrialLicenseJSONBody ApplyTrialLicenseRequest

// ConfirmParams defines parameters for Confirm.
type ConfirmParams struct {
	ConfirmationId string `json:"confirmation_id"`
	PlatformType   string `json:"platform_type"`
}

// CreatePurchaseContextJSONBody defines parameters for CreatePurchaseContext.
type CreatePurchaseContextJSONBody CreatePurchaseContextRequest

// GetFirebasePublicKeyParams defines parameters for GetFirebasePublicKey.
type GetFirebasePublicKeyParams struct {
	ProjectId *string `json:"project_id,omitempty"`
}

// PaymentDetailsJSONBody defines parameters for PaymentDetails.
type PaymentDetailsJSONBody PaymentDetailsRequest

// ProcessAndroidPurchaseJSONBody defines parameters for ProcessAndroidPurchase.
type ProcessAndroidPurchaseJSONBody ProcessAndroidPurchaseRequest

// ProcessIosPurchaseJSONBody defines parameters for ProcessIosPurchase.
type ProcessIosPurchaseJSONBody ProcessIOSPurchaseRequest

// ListProductParams defines parameters for ListProduct.
type ListProductParams struct {
	Limit        int     `json:"limit"`
	Offset       int     `json:"offset"`
	PlatformType *string `json:"platform_type,omitempty"`
}

// SendConfirmationLinkJSONBody defines parameters for SendConfirmationLink.
type SendConfirmationLinkJSONBody AuthRequest

// SendRestoreLinkJSONBody defines parameters for SendRestoreLink.
type SendRestoreLinkJSONBody SendRestoreLinkRequest

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody AuthRequest

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody AuthRequest

// TokenJSONBody defines parameters for Token.
type TokenJSONBody TokenRequest

// CreateFirebaseUserJSONBody defines parameters for CreateFirebaseUser.
type CreateFirebaseUserJSONBody CreateFirebaseUserRequest

// ServiceAuthenticateJSONBody defines parameters for ServiceAuthenticate.
type ServiceAuthenticateJSONBody AuthServiceRequest

// ApplyTrialLicenseJSONRequestBody defines body for ApplyTrialLicense for application/json ContentType.
type ApplyTrialLicenseJSONRequestBody ApplyTrialLicenseJSONBody

// CreatePurchaseContextJSONRequestBody defines body for CreatePurchaseContext for application/json ContentType.
type CreatePurchaseContextJSONRequestBody CreatePurchaseContextJSONBody

// PaymentDetailsJSONRequestBody defines body for PaymentDetails for application/json ContentType.
type PaymentDetailsJSONRequestBody PaymentDetailsJSONBody

// ProcessAndroidPurchaseJSONRequestBody defines body for ProcessAndroidPurchase for application/json ContentType.
type ProcessAndroidPurchaseJSONRequestBody ProcessAndroidPurchaseJSONBody

// ProcessIosPurchaseJSONRequestBody defines body for ProcessIosPurchase for application/json ContentType.
type ProcessIosPurchaseJSONRequestBody ProcessIosPurchaseJSONBody

// SendConfirmationLinkJSONRequestBody defines body for SendConfirmationLink for application/json ContentType.
type SendConfirmationLinkJSONRequestBody SendConfirmationLinkJSONBody

// SendRestoreLinkJSONRequestBody defines body for SendRestoreLink for application/json ContentType.
type SendRestoreLinkJSONRequestBody SendRestoreLinkJSONBody

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody

// TokenJSONRequestBody defines body for Token for application/json ContentType.
type TokenJSONRequestBody TokenJSONBody

// CreateFirebaseUserJSONRequestBody defines body for CreateFirebaseUser for application/json ContentType.
type CreateFirebaseUserJSONRequestBody CreateFirebaseUserJSONBody

// ServiceAuthenticateJSONRequestBody defines body for ServiceAuthenticate for application/json ContentType.
type ServiceAuthenticateJSONRequestBody ServiceAuthenticateJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Apply trial license
	// (POST /api/client/apply-trial-license)
	ApplyTrialLicense(w http.ResponseWriter, r *http.Request)
	// Confirm email
	// (GET /api/client/confirm)
	Confirm(w http.ResponseWriter, r *http.Request, params ConfirmParams)
	// Create purchase context
	// (POST /api/client/create-purchase-context)
	CreatePurchaseContext(w http.ResponseWriter, r *http.Request)
	// Delete user
	// (DELETE /api/client/delete)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	// Get firebase public key by project id or default
	// (GET /api/client/firebase-public-key)
	GetFirebasePublicKey(w http.ResponseWriter, r *http.Request, params GetFirebasePublicKeyParams)
	// List license by user_id
	// (GET /api/client/license-by-user)
	ListLicenseByUser(w http.ResponseWriter, r *http.Request)
	// Get payment details
	// (POST /api/client/payment-details)
	PaymentDetails(w http.ResponseWriter, r *http.Request)
	// Process android purchase
	// (POST /api/client/process-android-purchase)
	ProcessAndroidPurchase(w http.ResponseWriter, r *http.Request)
	// Process ios purchase
	// (POST /api/client/process-ios-purchase)
	ProcessIosPurchase(w http.ResponseWriter, r *http.Request)
	// List product
	// (GET /api/client/product)
	ListProduct(w http.ResponseWriter, r *http.Request, params ListProductParams)
	// Send confirmation link
	// (POST /api/client/send-confirmation-link)
	SendConfirmationLink(w http.ResponseWriter, r *http.Request)
	// Send restore link
	// (POST /api/client/send-restore-link)
	SendRestoreLink(w http.ResponseWriter, r *http.Request)
	// Authenticate user
	// (POST /api/client/signin)
	Authenticate(w http.ResponseWriter, r *http.Request)
	// Register user
	// (POST /api/client/signup)
	Register(w http.ResponseWriter, r *http.Request)
	// Refresh access token
	// (POST /api/client/token)
	Token(w http.ResponseWriter, r *http.Request)
	// Create user at firebase
	// (POST /api/service/firebase-user)
	CreateFirebaseUser(w http.ResponseWriter, r *http.Request)
	// Authenticate service
	// (POST /api/service/signin)
	ServiceAuthenticate(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ApplyTrialLicense operation middleware
func (siw *ServerInterfaceWrapper) ApplyTrialLicense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ApplyTrialLicense(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Confirm operation middleware
func (siw *ServerInterfaceWrapper) Confirm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ConfirmParams

	// ------------- Required query parameter "confirmation_id" -------------
	if paramValue := r.URL.Query().Get("confirmation_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "confirmation_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "confirmation_id", r.URL.Query(), &params.ConfirmationId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "confirmation_id", Err: err})
		return
	}

	// ------------- Required query parameter "platform_type" -------------
	if paramValue := r.URL.Query().Get("platform_type"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "platform_type"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "platform_type", r.URL.Query(), &params.PlatformType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "platform_type", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Confirm(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreatePurchaseContext operation middleware
func (siw *ServerInterfaceWrapper) CreatePurchaseContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreatePurchaseContext(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetFirebasePublicKey operation middleware
func (siw *ServerInterfaceWrapper) GetFirebasePublicKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFirebasePublicKeyParams

	// ------------- Optional query parameter "project_id" -------------
	if paramValue := r.URL.Query().Get("project_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "project_id", r.URL.Query(), &params.ProjectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFirebasePublicKey(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListLicenseByUser operation middleware
func (siw *ServerInterfaceWrapper) ListLicenseByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListLicenseByUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PaymentDetails operation middleware
func (siw *ServerInterfaceWrapper) PaymentDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PaymentDetails(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ProcessAndroidPurchase operation middleware
func (siw *ServerInterfaceWrapper) ProcessAndroidPurchase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ProcessAndroidPurchase(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ProcessIosPurchase operation middleware
func (siw *ServerInterfaceWrapper) ProcessIosPurchase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ProcessIosPurchase(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListProduct operation middleware
func (siw *ServerInterfaceWrapper) ListProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProductParams

	// ------------- Required query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	// ------------- Optional query parameter "platform_type" -------------
	if paramValue := r.URL.Query().Get("platform_type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "platform_type", r.URL.Query(), &params.PlatformType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "platform_type", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListProduct(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SendConfirmationLink operation middleware
func (siw *ServerInterfaceWrapper) SendConfirmationLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SendConfirmationLink(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SendRestoreLink operation middleware
func (siw *ServerInterfaceWrapper) SendRestoreLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SendRestoreLink(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Authenticate operation middleware
func (siw *ServerInterfaceWrapper) Authenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	ctx = context.WithValue(ctx, BasicScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Authenticate(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	ctx = context.WithValue(ctx, BasicScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Register(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Token operation middleware
func (siw *ServerInterfaceWrapper) Token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Token(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateFirebaseUser operation middleware
func (siw *ServerInterfaceWrapper) CreateFirebaseUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateFirebaseUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ServiceAuthenticate operation middleware
func (siw *ServerInterfaceWrapper) ServiceAuthenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ServiceAuthenticate(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/apply-trial-license", wrapper.ApplyTrialLicense)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client/confirm", wrapper.Confirm)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/create-purchase-context", wrapper.CreatePurchaseContext)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/client/delete", wrapper.DeleteUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client/firebase-public-key", wrapper.GetFirebasePublicKey)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client/license-by-user", wrapper.ListLicenseByUser)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/payment-details", wrapper.PaymentDetails)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/process-android-purchase", wrapper.ProcessAndroidPurchase)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/process-ios-purchase", wrapper.ProcessIosPurchase)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/client/product", wrapper.ListProduct)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/send-confirmation-link", wrapper.SendConfirmationLink)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/send-restore-link", wrapper.SendRestoreLink)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/signin", wrapper.Authenticate)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/signup", wrapper.Register)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/client/token", wrapper.Token)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/service/firebase-user", wrapper.CreateFirebaseUser)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/service/signin", wrapper.ServiceAuthenticate)
	})

	return r
}
