package cache

// Cache is the interface of defining all cache operations.
type Cache interface {
	IntDo(cmd string, args ...interface{}) (int, error)
	StringDo(cmd string, args ...interface{}) (string, error)
	StringsDo(cmd string, args ...interface{}) ([]string, error)
	BoolDo(cmd string, args ...interface{}) (bool, error)
	Float64Do(cmd string, args ...interface{}) (float64, error)
	Close() error
}
