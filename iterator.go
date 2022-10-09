package lazy

import "errors"

//lint:ignore ST1012 special semantics
// Done can be returned instead of an error value in order to terminate iteration.
var Done = errors.New("done")

// Iterator represents a value source that generates (or fetches) items lazily.
type Iterator[T any] interface {
	Next() bool
	Val() T
	Err() error
}

type closureIterator[T any] struct {
	current T
	err     error
	iterate func(T) (T, error)
}

func (c *closureIterator[T]) Next() bool {
	if c.err != nil {
		return false
	}
	n, err := c.iterate(c.current)
	if err != nil {
		c.err = err
		return false
	}
	c.current = n
	return true
}

func (c *closureIterator[T]) Val() T {
	return c.current
}

func (c *closureIterator[T]) Err() error {
	if c.err == Done {
		return nil
	}
	return c.err
}

// FuncIterator constructs a new iterator, using a function which takes the
// output of the last iteration as a parameter, and returning the next iteration
// along with any error that might have occurred, or a special value [Done], in
// order to terminate iteration.
func FuncIterator[T any](initial T, stepFunc func(T) (T, error)) Iterator[T] {
	return &closureIterator[T]{
		current: initial,
		iterate: stepFunc,
	}
}

// Can I make a result type that is like
// struct{
//    v T
//    e error
// }
