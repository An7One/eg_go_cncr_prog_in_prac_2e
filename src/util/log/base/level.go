package base

// LogLevel the leve of the output
type LogLevel uint8

const (
	// LEVEL_DEBUG the lowest level of debugging
	LEVEL_DEBUG LogLevel = iota + 1
	// LEVEL_INFO level of information
	LEVEL_INFO
	// LEVEL_WARN level of warning
	LEVEL_WARN
	// LEVEL_ERROR level of error
	LEVEL_ERROR
	// LEVEL_FATAL level of fatal
	// `os.Exit(1)` will be called
	LEVEL_FATAL
	// LEVEL_PANIC level of panic
	// `panic` will be called
	LEVEL_PANIC
)
