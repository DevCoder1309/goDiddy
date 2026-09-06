package domain

type Question struct {
	QName string
	QType uint16
	QClass uint16

}

func NewQuestion(qname string, qtype, qclass uint16) *Question {
	q := Question{
		QName: qname,
		QType: qtype,
		QClass: qclass,
	}
	return &q
}	