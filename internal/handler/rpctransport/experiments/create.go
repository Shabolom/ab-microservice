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

	for _, g := range req.GetGroups() {
		groups = append(groups, dto.Group{
			Name:              g.GetName(),
			RollingPercentage: g.GetRollingPercentage(),
			DeviceID:          g.GetDeviceId(),
		})
	}

	exp := &dto.Experiment{
		Name:              req.GetName(),
		RolloutPercentage: req.GetRolloutPercentage(),
		StartDate:         req.GetStartDate().AsTime(),
		EndDate:           req.GetEndDate().AsTime(),
		LayersID:          req.GetLayersId(),
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
