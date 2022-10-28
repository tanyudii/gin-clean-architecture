package group

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	groupRepo repositories.GroupRepository
}

func NewUsecase(
	groupRepo repositories.GroupRepository,
) usecases.GroupUsecase {
	return &Usecase{
		groupRepo: groupRepo,
	}
}
