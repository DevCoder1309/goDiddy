package ports

import (
	"godiddy/internal/core/domain"
)

type DecodeInterface interface {
	DecodeHeader(buf []byte) (*domain.Header, error)
	DecodeQuestion(buf []byte) (*domain.Question, error)
}