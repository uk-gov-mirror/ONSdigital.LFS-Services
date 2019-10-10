package validation

import (
	"fmt"
	"services/dataset"
)

type SurveyValidation struct {
	dataset *dataset.Dataset
}

func NewSurveyValidation(dataset *dataset.Dataset) Validation {
	var v Validation = SurveyValidation{dataset: dataset}
	return v
}

func (sf SurveyValidation) Validate() (ValidationResponse, error) {
	ok, err := sf.validateREFDTE()

	// add additional validation here
	return ok, err
}

func (sf SurveyValidation) validateREFDTE() (ValidationResponse, error) {
	rows, err := sf.dataset.GetRowsAsDouble("REFDTE")
	if err != nil {
		return ValidationResponse{
			ValidationStatus: ValidationFailed,
			ErrorMessage:     err.Error(),
		}, err
	}

	var val float64
	for _, j := range rows {
		if val == 0.0 {
			val = j
		}
		if val != j {
			return ValidationResponse{
				ValidationStatus: ValidationFailed,
				ErrorMessage:     "rows contain different values for RFEDTE",
			}, fmt.Errorf("rows contain different values for RFEDTE")
		}
	}
	return ValidationResponse{
		ValidationStatus: ValidationSuccessful,
		ErrorMessage:     "",
	}, nil
}
