package packet

import (
	"math/rand"
)

type Request struct {
	Header  Header
	Queries Queries
}

func NewRequest(name string) []byte {
	var (
		request [512]byte
		offset  int
	)

	header := MarshalHeader(&Header{
		TransactionID: rand.Intn(1 << 16),
		Flags:         0x0100,
		Questions:     1,
		AnswersRRs:    0,
		AuthorityRRs:  0,
		AdditionalRRs: 0,
	})
	copy(request[:], header)
	offset += len(header)

	queries := MarshalQueries(&Queries{
		Name:  name,
		QType: A,
		Class: IN,
	})
	copy(request[offset:], queries)
	offset += len(queries)

	result := make([]byte, offset)
	copy(result, request[:offset])
	return result
}

func UnmarshalRequest(raw []byte, req *Request) (int, error) {
	var (
		offset = 0
	)

	n, err := UnmarshalHeader(raw[offset:], &req.Header)
	if err != nil {
		return 0, err
	}
	offset += n

	n, err = UnmarshalQueries(raw[offset:], &req.Queries)
	if err != nil {
		return 0, err
	}
	offset += n

	return offset, err
}
