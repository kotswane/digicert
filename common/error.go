package common

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ErrorMsg struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// ErrorResponse model info
// @Description returns error message
type ErrorResponse struct {
	Message string `json:"message"`
}

func GetErrorMsg(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			return fmt.Sprintf("%s is %s", strings.ToLower(fe.Field()), fe.Tag())
		}
	}
	return err.Error()
}
