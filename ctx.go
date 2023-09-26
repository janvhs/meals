package main

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.
// This technique is use in go's in net/http package.
// An example use is http.ServerContextKey.
type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "context key for " + k.name }
