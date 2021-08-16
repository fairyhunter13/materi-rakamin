/*
Package cast provides tools for safely converting from one type to another, without information loss.

For an example, let's look at converting to an int64.

You can safely convert a uint8, uint16, uint32, plus an int, int8, int16, int32, int64,
into an int64 without information loss.

In other words, if `v` (in the following example code) is any of the types in the list just mentioned, then the following
code would successfully return an `int64` of the same numeric value.

	i64, err := cast.Int64(v)

In addition to these types built into Go (that were in the list), this package also supports converting
custom types, as long as they implement a special method.

For int64, the special method is:

	interface {
		Int64() (int64, error)
	}

For example:

	type Always5 struct{}
	
	func (receiver Always5) Int64() (int64, error) {
		return 5, nil
	}
	
	// ...
	
	var v Always5
	
	// ...
	
	i64, err := cast.Int64(v)
*/
package cast
