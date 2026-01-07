package validators


import "SwimRegistry/internal/errors"

func ValidarNumeroPositivo(field string, value int) error {
	if value <= 0 {
		return errors.ValidationError{
			Field: field,
			Reason: "El valor debe ser mayor a 0",
		}
	}
	return nil
}

func ValidarMultiploDe(value int, base int) error{
	if value % base != 0 {
		return errors.ValidationError{
			Field = field,
			Reason: "El valor debe ser multiplo de 25."
		}
	}
}
