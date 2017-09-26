package code2docx

// Config configration
type Config struct {
	// Excludes    []string
	ExcludeDirs []string
	Dir         string
	// Files       []string
	Types  []string
	Out    string
	Header string
	Title  string
}
