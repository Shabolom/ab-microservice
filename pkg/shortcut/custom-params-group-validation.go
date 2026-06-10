package shortcut

import (
	"ab/internal/dto"
	"fmt"
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

			fmt.Println("условия: ", param.Condition, "значения с которым будут сравнивать: ", param.Value, "переданные нами значения: ", reqParam.Value, "тип: ", param.ParameterType)
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
				continue
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
