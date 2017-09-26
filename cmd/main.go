package main

import (
	"log"

	"github.com/LYY/code2docx"
	"github.com/alecthomas/kingpin"
)

var (
	// Version version
	Version = "0.0.1"
)

func main() {
	log.Println("Start parsing...")
	conf := configure()
	kingpin.Version(Version)
	kingpin.Parse()
	if len(conf.Dirs) == 0 && len(conf.Files) == 0 {
		panic("Must have dirs or files")
	}
	log.Println("End")
}

func configure() *code2docx.Config {
	conf := &code2docx.Config{}
	kingpin.Flag("excludes", "Exclude files.").
		StringsVar(&conf.Excludes)
	kingpin.Flag("files", "Process files.").
		StringsVar(&conf.Files)
	kingpin.Flag("exdirs", "Exclude dirs.").
		StringsVar(&conf.ExcludeDirs)
	kingpin.Flag("dirs", "Process dirs.").
		StringsVar(&conf.Dirs)
	kingpin.Flag("out", "Output file name").
		Required().
		Short('o').
		StringVar(&conf.Out)
	return conf
}
