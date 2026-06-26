package experiments

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetExperimentByID(ctx context.Context, req *authv1.GetExperimentByIDRequest) (*authv1.GetExperimentByIDReply, error) {
	exp, err := h.experimentService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var customParamGroups []*authv1.ParamGroup
	for _, group := range exp.ParamsGroups {
		var params []*authv1.CustomParamWithCondition

		for _, param := range group.ParamsWithConditions {
			params = append(params, &authv1.CustomParamWithCondition{
				Id:               param.ID,
				ParameterId:      param.ParameterID,
				ParameterGroupId: param.ParameterGroupID,
				Value:            param.Value,
				Condition:        param.Condition,
			})
		}

		customParamGroups = append(customParamGroups, &authv1.ParamGroup{
			Id:                   group.ID,
			Percent:              group.Percent,
			ParamsWithConditions: params,
		})
	}

	var groups []*authv1.Group
	for _, group := range exp.Groups {
		groups = append(groups, &authv1.Group{
			Id:                group.ID,
			Name:              group.Name,
			RollingPercentage: group.RollingPercentage,
			DeviceId:          group.DeviceID,
		})
	}

	responseBody := &authv1.Experiment{
		Id:                exp.ID,
		Name:              exp.Name,
		Namespace:         exp.Namespace,
		Status:            exp.Status,
		RolloutPercentage: exp.RolloutPercentage,
		StartDate:         timestamppb.New(exp.StartDate),
		EndDate:           timestamppb.New(exp.EndDate),
		PassingCities:     exp.PassingCities,
		ExcludedCities:    exp.ExcludedCities,
		PassingStores:     exp.PassingStores,
		ExcludedStores:    exp.ExcludedStores,
		LayersId:          exp.LayersID,
		ParamsGroups:      customParamGroups,
		Groups:            groups,
	}

	return &authv1.GetExperimentByIDReply{
		ErrInfoReason: authv1.GetExperimentByIDReply_STATUS_OK,
		Message:       "experiment successfully retrieved:",
		Experiment:    responseBody,
	}, nil
}
