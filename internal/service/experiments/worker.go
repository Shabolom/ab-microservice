package experiments

import (
	"ab/internal/dto"
	"context"
	"fmt"
	"time"
)

func (s *Service) StartCashWorker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cash := make(map[string]dto.NameSpaceExperiments)

			nameSpaces, err := s.nameSpaceRepo.GetList(context.Background())
			if err != nil {
				fmt.Println(err)
			}

			for _, nameSpace := range nameSpaces {
				exp, err := s.experimentRepo.GetRawExperiments(context.Background(), nameSpace.Name)
				if err != nil {
					fmt.Println(err)
				}

				cashPart := dto.NameSpaceExperiments{
					NameSpace: nameSpace.Name,
					RawExp:    exp,
				}

				cash[nameSpace.Name] = cashPart
			}

			s.inMemoryStorage.Replace(cash)
		}
	}
}
