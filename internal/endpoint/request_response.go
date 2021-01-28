package endpoint

type SendRequest interface{}

type SendResponse struct {
	Error error
}

func (resp SendResponse) Failed() error {
	return resp.Error
}
