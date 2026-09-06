// this is a package to decode the header of the dns request
package decoder

import (
	"encoding/binary"
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

func (d DecodeStruct) DecodeQuestion(buf []byte) (*domain.Question, error) {
	// var qname string

	// for {
	// 	length := int(buf[offset]) // one byte
	// 	offset++

	// 	if length == 0 {
	// 		break
	// 	}

	// 	if len(qname) > 0 {
	// 		qname += "."
	// 	}

	// 	qname += string(buf[offset : offset+length])
	// 	offset += length
	// }

	firstByteLength := int(buf[12])

	qnameDomain := string(buf[13:firstByteLength + 13])

	secondByteLength := int(buf[firstByteLength + 13])

	qnameTLD := string(buf[firstByteLength + 14: secondByteLength + firstByteLength + 14])

	offsetF := secondByteLength + firstByteLength + 15

	qtype := binary.BigEndian.Uint16(buf[offsetF : offsetF+2])
	
	offsetF += 2

	qclass := binary.BigEndian.Uint16(buf[offsetF : offsetF+2])

	qname := qnameDomain + "." + qnameTLD

	data := domain.NewQuestion(qname, qtype, qclass)

	return data, nil
}

