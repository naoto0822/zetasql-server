package pkg

type Request struct {
	SQL string `json:"sql"`
}

func (r *Request) IsValid() bool {
	return r.SQL != ""
}

type ValidResponse struct {
	Message string `json:"message"`
}

type ParseResponse struct {
	Message string `json:"message"`
	AST     string `json:"ast"`
}
