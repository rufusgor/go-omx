package resource

type Response struct {
	Res  interface{}
	Code int
}

func (r Response) Metadata() map[string]interface{} {
	return map[string]interface{}{}
}

func (r Response) Result() interface{} {
	return r.Res
}

func (r Response) StatusCode() int {
	return r.Code
}
