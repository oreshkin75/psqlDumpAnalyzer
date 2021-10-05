package env

type Creator struct {
	path string
}

func New(path string) *Creator {
	return &Creator{path: path}
}
