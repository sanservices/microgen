package api

import (
	"{{cookiecutter.module_name}}/internal/errs"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	engTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

// Validator is struct for validator and translator
type Validator struct {
	v *validator.Validate
	t ut.Translator
}

// NewValidator creates new validator object
func NewValidator() (val *Validator) {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		panic("translator not found")
	}

	v := validator.New()

	err := engTranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		panic(err)
	}

	err = v.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "required, cannot be empty", true)

		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", "")
			return t
		},
	)
	if err != nil {
		panic(err)
	}

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{v: v, t: trans}
}

// Struct is validator wrapper for struct
func (v *Validator) Struct(req interface{}) error {
	err := v.v.Struct(req)
	if err == nil {
		return nil
	}

	propErrors := v.getPropErrs("", err.(validator.ValidationErrors))

	return errs.ServiceError{
		Message:    errs.MsgProps,
		Code:       errs.CodeProps,
		Properties: propErrors,
	}
}

// Var is validator wrapper for variable
func (v *Validator) Var(propName string, val interface{}, tag string) error {
	err := v.v.Var(val, tag)
	if err == nil {
		return nil
	}

	propErrors := v.getPropErrs(propName, err.(validator.ValidationErrors))

	return errs.ServiceError{
		Message:    errs.MsgProps,
		Code:       errs.CodeProps,
		Properties: propErrors,
	}
}

func (v *Validator) getPropErrs(propName string, validationErrors validator.ValidationErrors) (props []errs.PropertyError) {
	for _, vErr := range validationErrors {
		msg := vErr.Translate(v.t)
		prop := vErr.Field()
		if propName != "" {
			prop = propName
		}

		exists := false
		for _, p := range props {
			if p.Property == prop {
				exists = true
				p.Messages = append(p.Messages, msg)
			}
		}

		if !exists {
			props = append(props, errs.PropertyError{
				Property: prop, Messages: []string{msg}})
		}
	}
	return
}
