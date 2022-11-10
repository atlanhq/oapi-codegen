// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/atlanhq/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"fmt"
	"github.com/atlanhq/oapi-codegen/pkg/runtime"
	openapi_types "github.com/atlanhq/oapi-codegen/pkg/types"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

// EveryTypeOptional defines model for EveryTypeOptional.
type EveryTypeOptional struct {
	ArrayInlineField     *[]int              `json:"array_inline_field,omitempty"`
	ArrayReferencedField *[]SomeObject       `json:"array_referenced_field,omitempty"`
	BoolField            *bool               `json:"bool_field,omitempty"`
	ByteField            *[]byte             `json:"byte_field,omitempty"`
	DateField            *openapi_types.Date `json:"date_field,omitempty"`
	DateTimeField        *time.Time          `json:"date_time_field,omitempty"`
	DoubleField          *float64            `json:"double_field,omitempty"`
	FloatField           *float32            `json:"float_field,omitempty"`
	InlineObjectField    *struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field,omitempty"`
	Int32Field      *int32      `json:"int32_field,omitempty"`
	Int64Field      *int64      `json:"int64_field,omitempty"`
	IntField        *int        `json:"int_field,omitempty"`
	NumberField     *float32    `json:"number_field,omitempty"`
	ReferencedField *SomeObject `json:"referenced_field,omitempty"`
	StringField     *string     `json:"string_field,omitempty"`
}

// EveryTypeRequired defines model for EveryTypeRequired.
type EveryTypeRequired struct {
	ArrayInlineField     []int              `json:"array_inline_field"`
	ArrayReferencedField []SomeObject       `json:"array_referenced_field"`
	BoolField            bool               `json:"bool_field"`
	ByteField            []byte             `json:"byte_field"`
	DateField            openapi_types.Date `json:"date_field"`
	DateTimeField        time.Time          `json:"date_time_field"`
	DoubleField          float64            `json:"double_field"`
	FloatField           float32            `json:"float_field"`
	InlineObjectField    struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field"`
	Int32Field      int32      `json:"int32_field"`
	Int64Field      int64      `json:"int64_field"`
	IntField        int        `json:"int_field"`
	NumberField     float32    `json:"number_field"`
	ReferencedField SomeObject `json:"referenced_field"`
	StringField     string     `json:"string_field"`
}

// ReservedKeyword defines model for ReservedKeyword.
type ReservedKeyword struct {
	Channel *string `json:"channel,omitempty"`
}

// Resource defines model for Resource.
type Resource struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

// SomeObject defines model for some_object.
type SomeObject struct {
	Name string `json:"name"`
}

// Argument defines model for argument.
type Argument string

// ResponseWithReference defines model for ResponseWithReference.
type ResponseWithReference SomeObject

// SimpleResponse defines model for SimpleResponse.
type SimpleResponse struct {
	Name string `json:"name"`
}

// GetWithArgsParams defines parameters for GetWithArgs.
type GetWithArgsParams struct {

	// An optional query argument
	OptionalArgument *int64 `json:"optional_argument,omitempty"`

	// An optional query argument
	RequiredArgument int64 `json:"required_argument"`

	// An optional query argument
	HeaderArgument *int32 `json:"header_argument,omitempty"`
}

// CreateResourceJSONBody defines parameters for CreateResource.
type CreateResourceJSONBody EveryTypeRequired

// CreateResource2JSONBody defines parameters for CreateResource2.
type CreateResource2JSONBody Resource

// CreateResource2Params defines parameters for CreateResource2.
type CreateResource2Params struct {

	// Some query argument
	InlineQueryArgument *int `json:"inline_query_argument,omitempty"`
}

// UpdateResource3JSONBody defines parameters for UpdateResource3.
type UpdateResource3JSONBody struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateResourceRequestBody defines body for CreateResource for application/json ContentType.
type CreateResourceJSONRequestBody CreateResourceJSONBody

// CreateResource2RequestBody defines body for CreateResource2 for application/json ContentType.
type CreateResource2JSONRequestBody CreateResource2JSONBody

// UpdateResource3RequestBody defines body for UpdateResource3 for application/json ContentType.
type UpdateResource3JSONRequestBody UpdateResource3JSONBody

type ServerInterface interface {
	// get every type optional (GET /every-type-optional)
	GetEveryTypeOptional(w http.ResponseWriter, r *http.Request)
	// Get resource via simple path (GET /get-simple)
	GetSimple(w http.ResponseWriter, r *http.Request)
	// Getter with referenced parameter and referenced response (GET /get-with-args)
	GetWithArgs(w http.ResponseWriter, r *http.Request)
	// Getter with referenced parameter and referenced response (GET /get-with-references/{global_argument}/{argument})
	GetWithReferences(w http.ResponseWriter, r *http.Request)
	// Get an object by ID (GET /get-with-type/{content_type})
	GetWithContentType(w http.ResponseWriter, r *http.Request)
	// get with reserved keyword (GET /reserved-keyword)
	GetReservedKeyword(w http.ResponseWriter, r *http.Request)
	// Create a resource (POST /resource/{argument})
	CreateResource(w http.ResponseWriter, r *http.Request)
	// Create a resource with inline parameter (POST /resource2/{inline_argument})
	CreateResource2(w http.ResponseWriter, r *http.Request)
	// Update a resource with inline body. The parameter name is a reservedkeyword, so make sure that gets prefixed to avoid syntax errors (PUT /resource3/{fallthrough})
	UpdateResource3(w http.ResponseWriter, r *http.Request)
	// get response with reference (GET /response-with-reference)
	GetResponseWithReference(w http.ResponseWriter, r *http.Request)
}

// GetEveryTypeOptional operation middleware
func GetEveryTypeOptionalCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetSimple operation middleware
func GetSimpleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ParamsForGetWithArgs operation parameters from context
func ParamsForGetWithArgs(ctx context.Context) *GetWithArgsParams {
	return ctx.Value("GetWithArgsParams").(*GetWithArgsParams)
}

// GetWithArgs operation middleware
func GetWithArgsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// Parameter object where we will unmarshal all parameters from the context
		var params GetWithArgsParams

		// ------------- Optional query parameter "optional_argument" -------------
		if paramValue := r.URL.Query().Get("optional_argument"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "optional_argument", r.URL.Query(), &params.OptionalArgument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter optional_argument: %s", err), http.StatusBadRequest)
			return
		}

		// ------------- Required query parameter "required_argument" -------------
		if paramValue := r.URL.Query().Get("required_argument"); paramValue != "" {

		} else {
			http.Error(w, "Query argument required_argument is required, but not found", http.StatusBadRequest)
			return
		}

		err = runtime.BindQueryParameter("form", true, true, "required_argument", r.URL.Query(), &params.RequiredArgument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter required_argument: %s", err), http.StatusBadRequest)
			return
		}

		headers := r.Header

		// ------------- Optional header parameter "header_argument" -------------
		if valueList, found := headers[http.CanonicalHeaderKey("header_argument")]; found {
			var HeaderArgument int32
			n := len(valueList)
			if n != 1 {
				http.Error(w, fmt.Sprintf("Expected one value for header_argument, got %d", n), http.StatusBadRequest)
				return
			}

			err = runtime.BindStyledParameter("simple", false, "header_argument", valueList[0], &HeaderArgument)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid format for parameter header_argument: %s", err), http.StatusBadRequest)
				return
			}

			params.HeaderArgument = &HeaderArgument

		}

		ctx = context.WithValue(ctx, "GetWithArgsParams", &params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetWithReferences operation middleware
func GetWithReferencesCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "global_argument" -------------
		var globalArgument int64

		err = runtime.BindStyledParameter("simple", false, "global_argument", chi.URLParam(r, "global_argument"), &globalArgument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter global_argument: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "globalArgument", globalArgument)
		// ------------- Path parameter "argument" -------------
		var argument Argument

		err = runtime.BindStyledParameter("simple", false, "argument", chi.URLParam(r, "argument"), &argument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter argument: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "argument", argument)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetWithContentType operation middleware
func GetWithContentTypeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "content_type" -------------
		var contentType string

		err = runtime.BindStyledParameter("simple", false, "content_type", chi.URLParam(r, "content_type"), &contentType)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter content_type: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "contentType", contentType)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetReservedKeyword operation middleware
func GetReservedKeywordCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CreateResource operation middleware
func CreateResourceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "argument" -------------
		var argument Argument

		err = runtime.BindStyledParameter("simple", false, "argument", chi.URLParam(r, "argument"), &argument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter argument: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "argument", argument)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ParamsForCreateResource2 operation parameters from context
func ParamsForCreateResource2(ctx context.Context) *CreateResource2Params {
	return ctx.Value("CreateResource2Params").(*CreateResource2Params)
}

// CreateResource2 operation middleware
func CreateResource2Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "inline_argument" -------------
		var inlineArgument int

		err = runtime.BindStyledParameter("simple", false, "inline_argument", chi.URLParam(r, "inline_argument"), &inlineArgument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter inline_argument: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "inlineArgument", inlineArgument)

		// Parameter object where we will unmarshal all parameters from the context
		var params CreateResource2Params

		// ------------- Optional query parameter "inline_query_argument" -------------
		if paramValue := r.URL.Query().Get("inline_query_argument"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "inline_query_argument", r.URL.Query(), &params.InlineQueryArgument)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter inline_query_argument: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "CreateResource2Params", &params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UpdateResource3 operation middleware
func UpdateResource3Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "fallthrough" -------------
		var pFallthrough int

		err = runtime.BindStyledParameter("simple", false, "fallthrough", chi.URLParam(r, "fallthrough"), &pFallthrough)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter fallthrough: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "pFallthrough", pFallthrough)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetResponseWithReference operation middleware
func GetResponseWithReferenceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	r.Group(func(r chi.Router) {
		r.Use(GetEveryTypeOptionalCtx)
		r.Get("/every-type-optional", si.GetEveryTypeOptional)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetSimpleCtx)
		r.Get("/get-simple", si.GetSimple)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetWithArgsCtx)
		r.Get("/get-with-args", si.GetWithArgs)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetWithReferencesCtx)
		r.Get("/get-with-references/{global_argument}/{argument}", si.GetWithReferences)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetWithContentTypeCtx)
		r.Get("/get-with-type/{content_type}", si.GetWithContentType)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetReservedKeywordCtx)
		r.Get("/reserved-keyword", si.GetReservedKeyword)
	})
	r.Group(func(r chi.Router) {
		r.Use(CreateResourceCtx)
		r.Post("/resource/{argument}", si.CreateResource)
	})
	r.Group(func(r chi.Router) {
		r.Use(CreateResource2Ctx)
		r.Post("/resource2/{inline_argument}", si.CreateResource2)
	})
	r.Group(func(r chi.Router) {
		r.Use(UpdateResource3Ctx)
		r.Put("/resource3/{fallthrough}", si.UpdateResource3)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetResponseWithReferenceCtx)
		r.Get("/response-with-reference", si.GetResponseWithReference)
	})

	return r
}
