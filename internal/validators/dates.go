package validators

import (
	"time"
	"SwimRegistry/internal/errors"
)

func ValidarFecha(field string, value string) error {
	layout := "02/01/2006"

	_, err := time.Parse(layout, value)
	if err != nil {
		return errors.ValidationError{
			Field:  field,
			Reason: "debe tener formato dd/mm/aaaa",
		}
	}

	return nil
}
