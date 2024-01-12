package schema

type ErrRes struct {
	Err string `json:"err"`
}

func (e ErrRes) Error() string {
	return e.Err
}

type HelloRes struct {
	Msg string `json:"msg"`
}

type SubmitReq struct {
	Type      string `json:"type"`
	SubmitMsg string `json:"submitMsg"`
}
