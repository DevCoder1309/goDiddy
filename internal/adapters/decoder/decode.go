// this is a package to decode the header of the dns request
package decoder

import (
	"encoding/binary"
	"fmt"
	"godiddy/internal/core/domain"
	"godiddy/internal/ports"
)

type DecodeStruct struct{}

func Decode(d ports.DecodeInterface, buf []byte) (*domain.Header, *domain.Question, error){

	headerResult, err := d.DecodeHeader(buf)

	if err != nil {
		return nil, nil, err
	}

	questionResult, err := d.DecodeQuestion(buf)
	// fmt.Println(questionResult)
	if err != nil {
		return nil, nil, err
	}

	return headerResult, questionResult, nil
}


func (d DecodeStruct) DecodeHeader(buf []byte) ( *domain.Header, error) {
	id := binary.BigEndian.Uint16(buf[0:2])
	flags := binary.BigEndian.Uint16(buf[2:4])
	qdcount := binary.BigEndian.Uint16(buf[4:6])
	ancount := binary.BigEndian.Uint16(buf[6:8])
	nscount := binary.BigEndian.Uint16(buf[8:10])
	arcount := binary.BigEndian.Uint16(buf[10:12])

	data := domain.NewHeader(id, flags, qdcount, ancount, nscount, arcount)
	// fmt.Println(*data)
	return data, nil
}

func (d DecodeStruct) DecodeQuestion(buf []byte, count uint16) (*[]domain.Question, error) {
    var masterdata []domain.Question
    mainPtr := 12

    for range count {
        var qname string
        offset := mainPtr

        for {
            length := int(buf[offset])
            offset++

            if length == 0 {
                break
            }

            if len(qname) > 0 {
                qname += "."
            }

            qname += string(buf[offset : offset+length])
            offset += length
        }

        qtype := binary.BigEndian.Uint16(buf[offset : offset+2])
        offset += 2

        qclass := binary.BigEndian.Uint16(buf[offset : offset+2])
        offset += 2

        data := domain.NewQuestion(qname, qtype, qclass)
        masterdata = append(masterdata, *data)

        mainPtr = offset
    }

	fmt.Println(masterdata)

    return &masterdata, nil
}

