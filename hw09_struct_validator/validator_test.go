package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}
	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "12345678-1234-1234-1234-123456789abc",
				Name:   "John Doe",
				Age:    25,
				Email:  "some@example.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
			},
			expectedErr: ValidationErrors{},
		},
		{
			in: User{
				ID:     "12345678-1234-1234-1234-123456789abc",
				Name:   "John Doe",
				Age:    2,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "Age", Err: ErrValidationMin},
			},
		},
		{
			in: User{
				ID:     "12345678-1234-1234-1234-123456789abc",
				Name:   "John Doe",
				Age:    2,
				Email:  "1234123",
				Role:   "admin",
				Phones: []string{"12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "Age", Err: ErrValidationMin},
				{Field: "Email", Err: ErrInvalidRegexp},
			},
		},
		{
			in: App{
				Version: "1.0.0.",
			},
			expectedErr: ValidationErrors{
				{Field: "Version", Err: ErrValidationLen},
			},
		},
		{
			in: Token{
				Header:    []byte(`{"alg":"HS256","typ":"JWT"}`),
				Payload:   []byte(`{"alg":"HS256","typ":"JWT"}`),
				Signature: []byte(`{"alg":"HS256","typ":"JWT"}`),
			},
			expectedErr: ValidationErrors{},
		},
		{
			in: Response{
				Code: 407,
			},
			expectedErr: ValidationErrors{
				{Field: "Code", Err: ErrValidationIn},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			mu.Lock()
			defer mu.Unlock()
			var vErr, exprErr ValidationErrors
			err := Validate(tt.in)

			// var validationErrors ValidationErrors
			if errors.As(err, &vErr) && errors.As(tt.expectedErr, &exprErr) {
				if !errorsMatch(vErr, exprErr) {
					t.Errorf("unexpected error: got %v, want %v", err, tt.expectedErr)
				}
			}
		})
	}
}

func errorsMatch(err1, err2 ValidationErrors) bool {
	if len(err1) != len(err2) {
		return false
	}

	for i := range err1 {
		if !errors.Is(err1[i].Err, err2[i].Err) {
			return false
		}
	}

	return true
}

func TestNonValidateError(t *testing.T) {
	var role UserRole = "user"
	type (
		SomeData struct {
			Data string `validate:"len:"`
		}
		AnotherSomeData struct {
			Data string `validate:"len"`
		}
	)

	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          role,
			expectedErr: ErrUnsupportedInputKind,
		},
		{
			in: SomeData{
				Data: "data",
			},
			expectedErr: ErrInvalidValidationRuleValue,
		},
		{
			in: AnotherSomeData{
				Data: "data",
			},
			expectedErr: ErrInvalidValidationRule,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.expectedErr)
			}
		})
	}
}
