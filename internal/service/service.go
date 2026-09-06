package service

import (
	"context"
	"godiddy/internal/core/domain"
	"godiddy/internal/ports"
)

func RecordHelper(it ports.SaveInMemInterface, ctx context.Context) *[]domain.Record{
	records := it.Records(ctx)
	return records
}