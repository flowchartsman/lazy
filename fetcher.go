package lazy

import "sync"

// Fetcher creates a function for returning types with lazy initialization.
// The provided initializer function will be run only once (the first time the
// value is requested) after which all calls to the fetcher will return the
// initialized object
func Fetcher[T any](initializer func() (T, error), errCallback func(error)) func() T {
	var o sync.Once
	var v T
	w := make(chan struct{})
	return func() T {
		o.Do(func() {
			var err error
			v, err = initializer()
			if err != nil && errCallback != nil {
				errCallback(err)
			}
			close(w)
		})
		<-w
		return v
	}
}
