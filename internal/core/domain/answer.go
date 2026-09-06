package domain

type Answer struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    []byte
}

func NewAnswer(name string, typ uint16, class uint16, ttl uint32, rdLength uint16, rData []byte,) Answer {
	return Answer{
		Name:     name,
		Type:     typ,
		Class:    class,
		TTL:      ttl,
		RDLength: rdLength,
		RData:    rData,
	}
}