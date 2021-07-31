package spinner

// OptionFunc is a wrapper option func to spinner
type OptionFunc func(s *Spinner)

// WithOptions passes spinner and runs every option func
func (s *Spinner) WithOptions(opts ...OptionFunc) *Spinner {
	for _, f := range opts {
		f(s)
	}
	return s
}

// WithExitOnAbort sets abortOnExit field to spinner
func WithExitOnAbort(exit bool) OptionFunc {
	return func(s *Spinner) {
		s.exitOnAbort = exit
	}
}

func WithNotifySignals(notify bool) OptionFunc {
	return func(s *Spinner) {
		s.notifySignals = notify
	}
}
