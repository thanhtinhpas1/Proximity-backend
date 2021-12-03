package resolvers

import (
	"context"
	"proximity/graph/repositories"

	"github.com/google/uuid"
)

type BlockListResolver struct {
	BlockListRepo *repositories.BlockListRepo
}

func (rsv *BlockListResolver) BlockUser(ctx context.Context, from string, to string) (bool, error) {
	fromUuid, err := uuid.Parse(from)
	if err != nil {
		return false, err
	}

	toUuid, err := uuid.Parse(to)
	if err != nil {
		return false, err
	}

	return rsv.BlockListRepo.BlockUser(ctx, fromUuid, toUuid)
}

func (rsv *BlockListResolver) UnblockUser(ctx context.Context, from string, to string) (bool, error) {
	fromUuid, err := uuid.Parse(from)
	if err != nil {
		return false, err
	}

	toUuid, err := uuid.Parse(to)
	if err != nil {
		return false, err
	}

	return rsv.BlockListRepo.UnblockUser(ctx, fromUuid, toUuid)
}
