package repository

import (
	"context"
	"godiddy/internal/core/domain"
)

type St struct {}

func (s St) Records(ctx context.Context) *[]domain.Record {
	RecordSlice := []domain.Record{}
	record1 := domain.Record{
		Name: "example.com",
		Type: 1,
		Addr: "1.2.3.4",
	}
	record2 := domain.Record{
		Name: "google.com",
		Type: 1,
		Addr: "8.8.8.8",
	}
	record3 := domain.Record{
		Name: "google.com",
		Type: 28,
		Addr: "2001:4860:4860::8888",
	}

	record4 := domain.Record{
		Name: "www.google.com",
		Type: 5,
		Addr: "google.com",
	}
	RecordSlice = append(RecordSlice, record1)
	RecordSlice = append(RecordSlice, record2)
	RecordSlice = append(RecordSlice, record3)
	RecordSlice = append(RecordSlice, record4)

	return &RecordSlice
}
