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
		Groups:            groups,
	}

	fmt.Println(exp, 5555555)
	createdExp, err := h.experimentService.Create(ctx, exp)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("Create %v Experiment Successfully, id: %v", createdExp.Name, createdExp.ID),
	}, nil
}
