package errors

import (
	"bytes"
)

// bundleSeparator the separator for concatenating bundle errors.
const bundleSeparator = "; "

// NewBundle creates new bundle of errors.
func NewBundle() *Bundle {
	return &Bundle{}
}

// Bundle accumulates errors
// and represents a single error or error list.
type Bundle struct {
	errs []error
}

// List returns all bundle errors.
func (b *Bundle) List() []error {
	if b == nil {
		return nil
	}
	return b.errs
}

// Add conditionally adds a new errors to the bundle.
func (b *Bundle) Add(errs ...error) {
	for _, err := range errs {
		if err != nil {
			b.errs = append(b.errs, err)
		}
	}
}

func (b *Bundle) Error() string {
	buf := bytes.NewBuffer(nil)
	lastIndex := len(b.errs) - 1
	for i, err := range b.errs {
		_, _ = buf.WriteString(err.Error())
		if i != lastIndex {
			_, _ = buf.WriteString(bundleSeparator)
		}
	}
	return buf.String()
}

// ErrorOrNil returns an error if it has any errors or nil.
func (b *Bundle) ErrorOrNil() error {
	if b == nil {
		return nil
	}
	if !b.IsEmpty() {
		return b
	}
	return nil
}

// IsEmpty returns true if the bundle does not contain errors.
func (b *Bundle) IsEmpty() bool {
	return len(b.errs) == 0
}
