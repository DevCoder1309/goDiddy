package domain

type Header struct{
	ID uint16
	Flags uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func NewHeader(id, flags, qdcount, ancount, nscount, arcount uint16) *Header {
	d := Header{
		ID: id,
		Flags: flags,
		QDCount: qdcount,
		ANCount: ancount,
		NSCount: nscount,
		ARCount: arcount,
	}
	return &d
}	