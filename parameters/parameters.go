package parameters

import (
	"context"
	"errors"
)

// Parameters is a store for parameters read from an incoming HTTP request
// by a handler, to be passed down to the code that does the actual work.
type Parameters struct {
	values map[string]any
}

// NewParameters creates a set of parameters with the specified values.
func NewParameters(values map[string]any) Parameters {
	return Parameters{values: values}
}

// GetString returns the parameter with the specified key.
func (p Parameters) GetString(key string) (string, error) {
	v, ok := p.values[key]
	if !ok {
		return "", errors.New("no value for key " + key)
	}
	s, ok := v.(string)
	if !ok {
		return "", errors.New("value for key " + key + " is not a string")
	}
	return s, nil
}

// TODO: GetInt, GetBool, GetStruct, ...

// Unexported context key so other packages can't get the parameters directly
// from the context.
type contextKey struct{}

// PutParameters puts the parameters in the context and returns the resulting context.
func PutParameters(ctx context.Context, p Parameters) context.Context {
	return context.WithValue(ctx, contextKey{}, p)
}

// GetParameters fetches and returns the parameters from the context.
func GetParameters(ctx context.Context) Parameters {
	return ctx.Value(contextKey{}).(Parameters)
}
