package experiments

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
)

func (h *Handler) UserExperiment(ctx context.Context, req *authv1.ExperimentRequest) (*authv1.ExperimentsReply, error) {
	reqParam := &dto.RequestParameters{
		SplitID:   req.GetSplitId(),
		DeviceID:  req.GetDeviceId(),
		NameSpace: req.GetNamespace(),
	}

	reply, err := h.experimentService.GetExperiments(reqParam)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	ExperimentsReply := make([]*authv1.ExperimentReply, 0, len(reply))

	for _, r := range reply {
		result := &authv1.ExperimentReply{
			ExperimentName: r.ExperimentName,
			GroupName:      r.GroupName,
		}
		ExperimentsReply = append(ExperimentsReply, result)
	}

	return &authv1.ExperimentsReply{
		ErrInfoReason:    authv1.ExperimentsReply_STATUS_OK,
		ExperimentsReply: ExperimentsReply,
	}, nil
}
