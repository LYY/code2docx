package code2docx

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFiles(conf *Config) ([]string, error) {
	var listfile []string
	err := filepath.Walk(conf.Dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			for _, exdir := range conf.ExcludeDirs {
				if path == exdir {
					return filepath.SkipDir
				}
			}
			return nil
		}

		var ok bool
		if len(conf.Types) == 0 {
			listfile = append(listfile, path)
		} else {
			//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
			for _, suffix := range conf.Types {
				ok = strings.HasSuffix(path, "."+suffix)
				if ok {
					listfile = append(listfile, path)
					break
				}
			}
		}
		log.Println(ok, "               ", path) //list the file

		return nil
	})
	return listfile, err
}
