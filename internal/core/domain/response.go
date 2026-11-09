package domain

type Response struct {
	Header      Header
	Question    []Question
	Answers     []Answer
	Authorities []string
	Additionals []string
}

func NewResponse(h Header, q []Question, ans []Answer, auth, add []string) *Response {
	r := Response{
		Header:      h,
		Question:    q,
		Answers:     ans,
		Authorities: auth,
		Additionals: add,
	}
	return &r
}