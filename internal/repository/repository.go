package repository

import (
	"context"
	"godiddy/internal/core/domain"
)

type St struct {}

func (s St) Records(ctx context.Context) *[]domain.Record {
	RecordSlice := []domain.Record{}
	rc1 := map[string]string{
		"example.com":"1.2.3.4",
	}
	record1 := domain.Record{
		Record: rc1,
	}

	rc2 := map[string]string{
		"google.com":"8.8.8.8",
	}
	record2 := domain.Record{
		Record : rc2,
	}
	RecordSlice = append(RecordSlice, record1)
	RecordSlice = append(RecordSlice, record2)

	return &RecordSlice
}
