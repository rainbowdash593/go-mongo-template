package utils

import (
	"github.com/hashicorp/go-multierror"
)

func WrapErrors(err ...error) error {
	var result error
	for _, e := range err {
		result = multierror.Append(result, e)
	}

	return result
}
