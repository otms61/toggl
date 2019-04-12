package toggl

type debug interface {
	Debug() bool

	// Debugf print a formatted debug line.
	Debugf(format string, v ...interface{})
	// Debugln print a debug line.
	Debugln(v ...interface{})
}
