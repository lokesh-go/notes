package blockingcall

// WithSelectCase ...
func WithSelectCase() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// do your work
			}
		}
	}()

	// Close the go routine
	done <- true
}
