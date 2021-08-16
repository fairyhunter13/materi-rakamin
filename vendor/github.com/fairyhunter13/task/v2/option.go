package task

// OptionFunc specifies the optional function for this package.
type OptionFunc func(*Option)

// Option specifies the option to be used in this package.
type Option struct {
	// Not affected by Default() method.
	// UsePanicHandler defaults to false
	UsePanicHandler bool

	// Affected by Default() method.
	// Default is DefaultBufferSize
	BufferSize int
}

const (
	// DefaultBufferSize specifies the default size of channel.
	DefaultBufferSize = 5
)

// NewDefaultOption initialize a new default Option.
func NewDefaultOption() *Option {
	return NewOption().Default()
}

// NewOption initialize a new Option.
func NewOption() *Option {
	return new(Option)
}

// Clone clonse the current option.
func (o *Option) Clone() *Option {
	newOpt := *o
	return &newOpt
}

// Assign assign the functional options list to the Option.
func (o *Option) Assign(opts ...OptionFunc) *Option {
	for _, opt := range opts {
		opt(o)
	}
	return o.Default()
}

// Default sets the default value of the Option.
func (o *Option) Default() *Option {
	if o.BufferSize < 1 {
		o.BufferSize = DefaultBufferSize
	}
	return o
}

// WithPanicHandler adds the option to toggle the panic handler on and off.
func WithPanicHandler(confirm bool) OptionFunc {
	return func(o *Option) {
		o.UsePanicHandler = confirm
	}
}

// WithBufferSize sets the buffer size of the error manager.
func WithBufferSize(input int) OptionFunc {
	return func(o *Option) {
		o.BufferSize = input
	}
}
