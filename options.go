package spinner

type OptionFunc func(s *Spinner)

func (s *Spinner) WithOptions(opts ...OptionFunc) {
	for _, f := range opts {
		f(s)
	}
}
