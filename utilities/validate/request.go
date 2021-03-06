package requestValidate

import (
	"errors"

	"github.com/Roverr/go-auth/core/auth/types"
)

// UsernamePassword is function for validating request body
// which should contain userName and password fields
func UsernamePassword(request authTypes.RegisterRequest) error {
	if request.UserName == "" || request.Password == "" {
		return errors.New("Request body did not contain userName or password.")
	}
	return nil
}
