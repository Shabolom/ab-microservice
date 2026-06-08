package experiments

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"fmt"

	"golang.org/x/net/context"
)

func (h *Handler) CreateExperiment(ctx context.Context, req *authv1.CreateExperimentRequest) (*authv1.StockReply, error) {
	groups := make([]dto.Group, 0, len(req.GetGroups()))
	passingCities := make([]string, 0, len(req.GetPassingCities()))
	excludedCities := make([]string, 0, len(req.GetExcludedCities()))
	passingStores := make([]string, 0, len(req.GetPassingStores()))
	excludedStores := make([]string, 0, len(req.GetExcludedStores()))
	customParamsGroups := make([]dto.ParamGroup, 0, len(req.GetCustomParamGroups()))

	for _, g := range req.GetGroups() {
		groups = append(groups, dto.Group{
			Name:              g.GetName(),
			RollingPercentage: g.GetRollingPercentage(),
			DeviceID:          g.GetDeviceId(),
		})
	}

	for _, passingCity := range req.GetPassingCities() {
		passingCities = append(passingCities, passingCity)
	}

	for _, excludedCity := range req.GetExcludedCities() {
		excludedCities = append(excludedCities, excludedCity)
	}

	for _, passingStore := range req.GetPassingStores() {
		passingStores = append(passingStores, passingStore)
	}

	for _, excludedStore := range req.GetExcludedStores() {
		excludedStores = append(excludedStores, excludedStore)
	}

	for _, reqCustomParamGroup := range req.GetCustomParamGroups() {
		customParams := reqCustomParamGroup.GetCustomParam()
		customParamGroup := dto.ParamGroup{
			Percent:              reqCustomParamGroup.GetPercentage(),
			ParamsWithConditions: make([]dto.CustomParamWithCondition, 0, len(customParams)),
		}

		for _, param := range customParams {
			customParamWithCondition := dto.CustomParamWithCondition{
				ParameterID: param.GetParametrId(),
				Value:       param.GetValue(),
				Condition:   param.GetCondition(),
			}

			customParamGroup.ParamsWithConditions = append(customParamGroup.ParamsWithConditions, customParamWithCondition)
		}

		customParamsGroups = append(customParamsGroups, customParamGroup)
	}

	exp := &dto.Experiment{
		Name:              req.GetName(),
		RolloutPercentage: req.GetRolloutPercentage(),
		StartDate:         req.GetStartDate().AsTime(),
		EndDate:           req.GetEndDate().AsTime(),
		LayersID:          req.GetLayersId(),
		PassingCities:     passingCities,
		ExcludedCities:    excludedCities,
		PassingStores:     passingStores,
		ExcludedStores:    excludedStores,
		ParamsGroups:      customParamsGroups,
		Groups:            groups,
	}

	createdExp, err := h.experimentService.Create(ctx, exp)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("Create %v Experiment Successfully, id: %v", createdExp.Name, createdExp.ID),
	}, nil
}
