package InternalErrors

import "errors"

var ErrInternal error =  errors.New("Internal server error")