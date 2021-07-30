package spinner

// OptionFunc is a wrapper option func to spinner
type OptionFunc func(s *Spinner)

// WithOptions passes spinner and runs every option func
func (s *Spinner) WithOptions(opts ...OptionFunc) {
	for _, f := range opts {
		f(s)
	}
}

// WithExitOnAbort sets abortOnExit field to spinner
func WithExitOnAbort(exit bool) OptionFunc {
	return func(s *Spinner) {
		s.exitOnAbort = exit
	}
}
