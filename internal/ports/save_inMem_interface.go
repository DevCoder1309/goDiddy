package ports

import (
	"context"
	"godiddy/internal/core/domain"
)

type SaveInMemInterface interface {
	Records(ctx context.Context) *[]domain.Record
}