package routine

func Run(retry int, fn func() error) {
	if retry < 0 {
		retry = 100
	}

	for i := 0; i < retry; i++ {
		err := fn()
		if err == nil {
			break
		}
	}
}
