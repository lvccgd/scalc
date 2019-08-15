package set

// T is type alias
type T interface{}

// Iter is channel iterator type
type Iter chan T

type optimal struct {
	length, idx T
}

// Predicate function type
type Predicate func(args T) bool
