package response

import (
	"context"
	"net"

	"godiddy/internal/core/domain"
	"godiddy/internal/repository"
	"godiddy/internal/service"
)

func ConstructResponse(h *domain.Header, q *domain.Question) (*domain.Response, error) {

	resH := *h
	resQ := *q

	resH.Flags = 0x8180
	resH.QDCount = 1
	resH.ANCount = 1
	resH.NSCount = 0
	resH.ARCount = 0

	ctx := context.Background()

	records := service.RecordHelper(repository.St{}, ctx)

	var rawIpData string

	for _, itr := range *records{
		val, ok := itr.Record[q.QName]
		if ok {
			rawIpData = val
		}
	}

	rdata := net.ParseIP(rawIpData).To4()

	var rdlength uint16 = 4

	answer := domain.Answer{
		Name:     q.QName,
		Type:     q.QType,
		Class:    q.QClass,
		TTL:      300,
		RDLength: rdlength,
		RData:    rdata,
	}

	res := domain.NewResponse(
		resH,
		resQ,
		[]domain.Answer{answer},
		[]string{},
		[]string{},
	)

	return res, nil
}