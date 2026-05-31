package experiments

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"context"
)

func (h *Handler) CreateExperiment(ctx context.Context, req *authv1.CreateExperimentRequest) (*authv1.StockReply, error) {
	groups := make([]dto.Group, 0, len(req.GetGroups()))

	for _, g := range req.GetGroups() {
		groups = append(groups, dto.Group{
			Name:              g.GetName(),
			RollingPercentage: g.GetRollingPercentage(),
			DeviceID:          g.GetDeviceID(),
		})
	}

	exp := &dto.Experiment{
		Name:              req.GetName(),
		Status:            req.GetStatus(),
		RolloutPercentage: req.GetRolloutPercentage(),
		StartDate:         req.GetStartDate().AsTime(),
		EndDate:           req.GetEndDate().AsTime(),
		LayersID:          req.GetLayersID(),
		Bucket:            req.GetBucket(),
		Groups:            groups,
	}
	
	return h.CreateExperiment(ctx, req)
}
