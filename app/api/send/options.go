package send

func WithHeaders(m map[string]string) Option {
	return func(r Response) Response {
		r.header = m
		return r
	}
}

func WithError(err error) Option {
	return func(r Response) Response {
		r.error = err
		return r
	}
}
