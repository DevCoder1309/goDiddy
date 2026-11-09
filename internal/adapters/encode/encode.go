package encode

import (
	"encoding/binary"
	"godiddy/internal/core/domain"
	"godiddy/internal/adapters/encodeName"
)


func EncodeResponse(resp *domain.Response) []byte {
	buf := make([]byte, 0, 512)
	header := make([]byte, 12)

	binary.BigEndian.PutUint16(header[0:2], resp.Header.ID)
	binary.BigEndian.PutUint16(header[2:4], resp.Header.Flags)
	binary.BigEndian.PutUint16(header[4:6], resp.Header.QDCount)
	binary.BigEndian.PutUint16(header[6:8], resp.Header.ANCount)
	binary.BigEndian.PutUint16(header[8:10], resp.Header.NSCount)
	binary.BigEndian.PutUint16(header[10:12], resp.Header.ARCount)

	for _, val := range header {
		buf = append(buf, val)
	}
	
	name := encodeName.EncodeName(resp.Question.QName)
	buf = append(buf, name...)

	qtype := make([]byte, 2)
	binary.BigEndian.PutUint16(qtype, resp.Question.QType)
	buf = append(buf, qtype...)

	qclass := make([]byte, 2)
	binary.BigEndian.PutUint16(qclass, resp.Question.QClass)
	buf = append(buf, qclass...)

	for _, answers := range resp.Answers {
		pointer := uint16(0xC000 | 12)
		
		// tempName := encodeName.EncodeName(answers.Name)
		rightShiftVal := byte(pointer >> 8)
		leftVal := byte(pointer)
		buf = append(buf, rightShiftVal, leftVal)

		tmp := make([]byte, 10)

		binary.BigEndian.PutUint16(tmp[0:2], answers.Type)
		binary.BigEndian.PutUint16(tmp[2:4], answers.Class)
		binary.BigEndian.PutUint32(tmp[4:8], answers.TTL)
		binary.BigEndian.PutUint16(tmp[8:10], answers.RDLength)

		buf = append(buf, tmp...)
		buf = append(buf, answers.RData...)
	}

	return buf
}
