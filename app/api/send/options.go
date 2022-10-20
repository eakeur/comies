package send

func WithHeaders(m map[string]string) Option {
	return func(r Response) Response {
		r.Header = m
		return r
	}
}

func WithError(err error) Option {
	return func(r Response) Response {
		r.Error = err
		return r
	}
}
