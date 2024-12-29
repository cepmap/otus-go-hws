package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for _, ve := range v {
		sb.WriteString(fmt.Sprintf("%s: %s\n", ve.Field, ve.Err.Error()))
	}
	return sb.String()
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", ve.Field, ve.Err)
}

func (ve ValidationError) Unwrap() error {
	return ve.Err
}

var (
	Ve                            *ValidationErrors
	ErrUnsupportedInputKind       = errors.New("unsupported input kind")
	ErrInvalidValidationRule      = errors.New("invalid validation rule")
	ErrInvalidValidationRuleValue = errors.New("invalid validation rule value")

	ErrValidationLen = errors.New("invalid value length")

	ErrValidationIn  = errors.New("value not found in list of possible values")
	ErrValidationMin = errors.New("value less than a required minimum")
	ErrValidationMax = errors.New("value more than a required maximum")

	ErrInvalidRegexp = errors.New("invalid regexp")
	ErrInvalidString = errors.New("invalid string")
)

func Validate(v interface{}) error {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		return ErrUnsupportedInputKind
	}

	var ve *ValidationErrors
	veList := make(ValidationErrors, 0, val.NumField())

	err := validateFields(val, &veList)
	if err != nil && !errors.As(err, &ve) {
		return err
	}
	if len(veList) > 0 {
		return veList
	}

	return nil
}

func validateFields(val reflect.Value, veList *ValidationErrors) error {
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldVal := val.Field(i) //Value

		if field.Tag.Get("validate") == "" {
			continue
		}

		if field.Type.Kind() == reflect.Slice {
			for j := 0; j < fieldVal.Len(); j++ {
				sliceVal := fieldVal.Index(j)
				err := validateField(field, sliceVal, veList)
				if err != nil && !errors.As(err, &Ve) {
					return err
				}
			}
		} else {
			err := validateField(field, fieldVal, veList)
			if err != nil && !errors.As(err, &Ve) {
				return err
			}
		}
	}
	return veList
}
func validateField(field reflect.StructField, fieldVal reflect.Value, veList *ValidationErrors) error {
	vaRules := strings.Split(field.Tag.Get("validate"), "|")
	for _, rule := range vaRules {
		if err := validate(rule, field.Name, fieldVal); err != nil {

			if errors.As(err, &ValidationError{}) {
				fmt.Println(err)
				*veList = append(*veList, ValidationError{Field: field.Name, Err: errors.Unwrap(err)})
			} else {
				return err
			}
		}
	}
	return veList
}

func validate(rules, fieldName string, fieldVal reflect.Value) error {
	rule := strings.Split(rules, ":")
	if len(rule) != 2 {
		return ErrInvalidValidationRule
	}
	if rule[1] == "" {
		return ErrInvalidValidationRuleValue
	}
	switch fieldVal.Kind() {
	case reflect.String:
		return validateStr(rule, fieldName, fieldVal)
	case reflect.Int:
		return validateInt(rule, fieldName, fieldVal)
	}

	return nil
}

func validateInt(rule []string, fieldName string, fieldVal reflect.Value) error {
	ruleName := rule[0]
	ruleValue := rule[1]
	switch ruleName {
	case "min":
		minVal, err := strconv.Atoi(ruleValue)
		if err != nil {
			return err
		}
		if fieldVal.Int() < int64(minVal) {

			return ValidationError{Field: fieldName, Err: ErrValidationMin}
		}
	case "max":
		maxValue, err := strconv.Atoi(ruleValue)
		if err != nil {
			return err
		}
		if fieldVal.Int() > int64(maxValue) {
			return ValidationError{Field: fieldName, Err: ErrValidationMax}
		}
	case "in":
		validValues := strings.Split(ruleValue, ",")
		fieldValueInt := fieldVal.Int()
		for _, v := range validValues {
			if intValue, err := strconv.ParseInt(v, 10, 64); err == nil && intValue == fieldValueInt {
				return nil
			}
		}
		return ValidationError{Field: fieldName, Err: ErrValidationIn}
	default:
		return nil
	}
	return nil
}
func validateStr(rule []string, fieldName string, fieldVal reflect.Value) error {
	ruleName := rule[0]
	ruleValue := rule[1]
	switch ruleName {
	case "len":
		strLen, err := strconv.Atoi(ruleValue)
		if err != nil {
			return err
		}
		if len(fieldVal.String()) != strLen {
			return ValidationError{Field: fieldName, Err: ErrValidationLen}
		}
		return nil
	case "regexp":
		r, err := regexp.Compile(ruleValue)
		if err != nil {
			return err
		}
		if !r.MatchString(fieldVal.String()) {
			return ValidationError{Field: fieldName, Err: ErrInvalidRegexp}
		}
	case "in":
		validValues := strings.Split(ruleValue, ",")
		fieldValueStr := fieldVal.String()
		for _, v := range validValues {
			if v == fieldValueStr {
				return nil
			}
		}
		return ValidationError{Field: fieldName, Err: ErrValidationIn}
	default:
		return ErrInvalidString
	}
	return nil
}
