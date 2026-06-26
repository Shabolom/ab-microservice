package experiments

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetExperiments(ctx context.Context, req *emptypb.Empty) (*authv1.GetExperimentsReply, error) {
	experiments, err := h.experimentService.GetList(ctx)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var experimentsResponse []*authv1.Experiment
	for _, e := range experiments {

		var groups []*authv1.Group
		for _, group := range e.Groups {
			groups = append(groups, &authv1.Group{
				Id:                group.ID,
				Name:              group.Name,
				RollingPercentage: group.RollingPercentage,
				DeviceId:          group.DeviceID,
			})
		}

		var customGroups []*authv1.ParamGroup
		for _, customGroup := range e.ParamsGroups {

			var params []*authv1.CustomParamWithCondition
			for _, param := range customGroup.ParamsWithConditions {
				params = append(params, &authv1.CustomParamWithCondition{
					Id:               param.ID,
					ParameterId:      param.ParameterID,
					ParameterGroupId: param.ParameterGroupID,
					Value:            param.Value,
					Condition:        param.Condition,
				})
			}

			customGroups = append(customGroups, &authv1.ParamGroup{
				Id:                   customGroup.ID,
				Percent:              customGroup.Percent,
				ParamsWithConditions: params,
			})
		}

		experimentsResponse = append(experimentsResponse, &authv1.Experiment{
			Id:                e.ID,
			Name:              e.Name,
			Namespace:         e.Namespace,
			Status:            e.Status,
			RolloutPercentage: e.RolloutPercentage,
			StartDate:         timestamppb.New(e.StartDate),
			EndDate:           timestamppb.New(e.EndDate),
			PassingCities:     e.PassingCities,
			ExcludedCities:    e.ExcludedCities,
			PassingStores:     e.PassingStores,
			ExcludedStores:    e.ExcludedStores,
			LayersId:          e.LayersID,
			ParamsGroups:      customGroups,
			Groups:            groups,
		})
	}

	return &authv1.GetExperimentsReply{
		ErrInfoReason: authv1.GetExperimentsReply_STATUS_OK,
		Message:       "experiments",
		Experiments:   experimentsResponse,
	}, nil
}
