package shortcut

import (
	"ab/internal/dto"
)

func CustomParamsGroupValidation(group dto.ParamGroupWithRawParams, reqParams []dto.Parameter) (bool, error) {
	if len(group.ParamsWithConditions) > len(reqParams) {
		return false, ErrValidation
	}

	for _, param := range group.ParamsWithConditions {
		found := false

		for _, reqParam := range reqParams {
			if param.ParameterName != reqParam.ParamName {
				continue
			}

			ok, err := conditionValueValidation(
				param.Condition,
				param.Value,
				reqParam.Value,
				param.ParameterType,
			)
			if err != nil {
				return false, err
			}

			if !ok {
				return false, nil
			}

			found = true
			break
		}

		if !found {
			return false, nil
		}
	}

	return true, nil
}
