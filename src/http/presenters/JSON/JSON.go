package JSON

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"

	capilaErrors "github.com/safciplak/capila/src/errors"
)

// Presenter is a struct for rendering JSON responses.
type Presenter struct {
	Context *gin.Context
	Body    Response
}

// Response is the default JSON response.
//
//nolint:govet // Performance impact is negligible
type Response struct {
	Errors []ErrorObject         `json:"errors,omitempty"`
	Data   interface{}           `json:"data,omitempty"`
	Links  map[string]LinkObject `json:"_links"`
	Meta   Meta                  `json:"meta"`
}

// LinkObject holds the data for a specific link property
// A valid LinkObject requires a href value. All other properties are optional.
// See https://tools.ietf.org/html/draft-kelly-json-hal-08#section-4.1.1 for property description.
type LinkObject struct {
	Href string `json:"href"`
}

// ErrorObject is a way to present error information
// based on: https://jsonapi.org/format/#error-objects
type ErrorObject struct {
	// A unique identifier for this particular occurrence of the problem.
	ID string `json:"id,omitempty"`
	// An application-specific error code, expressed as a string value.
	Code string `json:"code,omitempty"`
	// A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the
	// problem, except for purposes of localization.
	Title string `json:"title,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail string `json:"detail,omitempty"`
	// The HTTP status code applicable to this problem, expressed as a string value.
	Status int `json:"status,omitempty"`
}

// Meta contains meta fields with extra data
type Meta struct {
	Query url.Values `json:"query"`
}

// Present creates a JSON presenter with default values.
func Present(ctx *gin.Context) *Presenter {
	var jsonPresenter = &Presenter{
		Body:    Response{},
		Context: ctx,
	}

	jsonPresenter.populate(ctx.Request)

	return jsonPresenter
}

// populate the _links and meta based on the request.
func (presenter *Presenter) populate(request *http.Request) {
	presenter.populateMeta(request)
	presenter.populateLinks(request)
}

// populateMeta fills the meta struct in the JSON
func (presenter *Presenter) populateMeta(request *http.Request) {
	presenter.Body.Meta.Query = request.URL.Query()
}

// populateLinks adds the _links property with default self value in the JSON
func (presenter *Presenter) populateLinks(request *http.Request) {
	presenter.AddLinkObject("self", request.URL.String())
}

// AddLinkObject adds a defined property item to the _links property in the JSON
func (presenter *Presenter) AddLinkObject(item, href string) {
	if presenter.Body.Links == nil {
		presenter.Body.Links = map[string]LinkObject{}
	}

	presenter.Body.Links[item] = LinkObject{Href: href}
}

// convertErrorToErrorObjects converts a generic error into a slice of error objects
func (presenter *Presenter) convertErrorToErrorObjects(err error) (result []ErrorObject) {
	var (
		currentErr = err
		trans      = getValidationErrorTranslator()
	)

	// Append the first error (this signifies the http statuscode)
	result = append(result, convertErrorToErrorObject(currentErr, trans))

	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
		// ValidationErrors need to be looped over instead of simply unwrapped...
		if _, ok := currentErr.(validator.ValidationErrors); ok {
			for _, fieldError := range currentErr.(validator.ValidationErrors) {
				result = append(result, convertErrorToErrorObject(fieldError, trans))
			}
		} else {
			result = append(result, convertErrorToErrorObject(currentErr, trans))
		}
	}

	return result
}

// convertErrorToErrorObject converts a singular error into a singular (localized) error object
func convertErrorToErrorObject(err error, translator ut.Translator) (result ErrorObject) {
	switch value := err.(type) {
	// If it's a known error, simply add the error message as a description
	case capilaErrors.InterfaceError:
		return ErrorObject{
			Status: value.GetStatusCode(),
			Title:  getErrorTitle(value),
			Detail: value.GetDetail(),
			Code:   value.GetCode(),
		}
	// If the unwrapped error is a validation field error, localize it
	case validator.FieldError:
		return ErrorObject{
			Title:  getErrorTitle(value),
			Code:   capilaErrors.ErrorCodeInputValidation,
			Detail: value.Translate(translator),
		}
	// If parsing the request results in errors
	case *strconv.NumError:
		return ErrorObject{
			Title:  getErrorTitle(value),
			Code:   capilaErrors.ErrorCodeInputValidation,
			Detail: value.Error(),
		}
	// If something unexpected happened
	default:
		return ErrorObject{
			Title:  getErrorTitle(err),
			Code:   capilaErrors.ErrorCodeUnknown,
			Detail: err.Error(),
		}
	}
}

// getValidationErrorTranslator returns the default translator, should be made dynamically based on the set language.
func getValidationErrorTranslator() ut.Translator {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	return trans
}

// getErrorTitle formats the title, so it remains consistent throughout the application
func getErrorTitle(err error) string {
	// Currently, we only strip the pointer designator from the identifier
	return strings.Trim(reflect.TypeOf(err).String(), "*")
}

// Error is a shorthand function for outputting a single error.
func (presenter *Presenter) Error(statusCode int, err error) {
	presenter.Body.Errors = presenter.convertErrorToErrorObjects(err)
	presenter.Context.JSON(statusCode, presenter.Body)
}

// Success prepares the body and serialize the body as JSON
func (presenter *Presenter) Success(statusCode int, data interface{}) {
	presenter.Body.Data = data
	presenter.Context.JSON(statusCode, presenter.Body)
}
