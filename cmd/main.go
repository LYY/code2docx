package main

import (
	"log"
	"path/filepath"

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
	// if len(conf.Dirs) == 0 && len(conf.Files) == 0 {
	// 	panic("Must have dirs or files")
	// }

	var err error
	conf.Dir, err = filepath.Abs(conf.Dir)
	if err != nil {
		panic(err)
	}
	for idx, dir := range conf.ExcludeDirs {
		conf.ExcludeDirs[idx], err = filepath.Abs(dir)
		if err != nil {
			panic(err)
		}
	}
	log.Println(conf.ExcludeDirs)

	files, err := code2docx.FindFiles(conf)
	if err != nil {
		panic(err)
	}

	outer := &code2docx.Outer{
		Config:  conf,
		Base:    conf.Dir,
		Files:   files,
		Docfile: conf.Out,
	}
	outer.WriteDoc()

	log.Println("End")
}

func configure() *code2docx.Config {
	conf := &code2docx.Config{}
	// kingpin.Flag("excludes", "Exclude files.").
	// 	StringsVar(&conf.Excludes)
	// kingpin.Flag("files", "Process files.").
	// 	StringsVar(&conf.Files)
	kingpin.Flag("exdirs", "Exclude dirs.").
		StringsVar(&conf.ExcludeDirs)
	kingpin.Flag("dir", "Process dir.").
		Required().
		StringVar(&conf.Dir)
	kingpin.Flag("types", "Types.").
		StringsVar(&conf.Types)
	kingpin.Flag("out", "Output file name").
		Required().
		Short('o').
		StringVar(&conf.Out)
	kingpin.Flag("header", "Header").
		StringVar(&conf.Header)
	kingpin.Flag("title", "Title").
		StringVar(&conf.Title)
	return conf
}
