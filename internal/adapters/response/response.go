package response

import (
	"context"
	"net"

	"godiddy/internal/adapters/encodeName"
	"godiddy/internal/core/domain"
	"godiddy/internal/repository"
	"godiddy/internal/service"
)

func ConstructResponse(h *domain.Header, q *[]domain.Question) (*domain.Response, error) {

	resH := *h
	resQ := *q
	var res *domain.Response

	ctx := context.Background()

	records := service.RecordHelper(repository.St{}, ctx)

	var rawIpData string
	for _, ques := range(*q){
		for _, itr := range *records{
			if ques.QName == itr.Name && ques.QType == itr.Type {
				rawIpData = itr.Addr
				break
			}else{
				rawIpData = ""
			}
		}
	

		var flags uint16
		var anCount uint16

		if len(rawIpData) == 0 {
			flags = 0x8183
			anCount = 0
		}else{
			flags = 0x8180
			anCount = 1
		}

		// fmt.Println(q.QName)
		resH.Flags = flags
		resH.QDCount = 1
		resH.ANCount = anCount
		resH.NSCount = 0
		resH.ARCount = 0

		var rawDataToParse net.IP
		var rVarLength uint16
		// fmt.Println(q.QType)
		if ques.QType == uint16(28) {
			rawDataToParse = net.ParseIP(rawIpData).To16()
			rVarLength = 16
			// fmt.Println(rawDataToParse)
		}else if(ques.QType == uint16(1)){
			rawDataToParse = net.ParseIP(rawIpData).To4()
			rVarLength = 4
		}else if(ques.QType == uint16(5)){
			rawDataToParse = encodeName.EncodeName(rawIpData)
			// fmt.Printf("this is the raw data %v", rawDataToParse)
			rVarLength = uint16(len(encodeName.EncodeName(rawIpData)))
		}
		rdata := rawDataToParse
		var rdlength uint16 = rVarLength

		answer := domain.Answer{
			Name:     ques.QName,
			Type:     ques.QType,
			Class:    ques.QClass,
			TTL:      300,
			RDLength: rdlength,
			RData:    rdata,
		}

		res = domain.NewResponse(
			resH,
			resQ,
			[]domain.Answer{answer},
			[]string{},
			[]string{},
		)
	}

	return res, nil
}