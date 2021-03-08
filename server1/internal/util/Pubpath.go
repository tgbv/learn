package util

import (
	"path/filepath"
)

func Pubpath(target string) string {
	path, e := filepath.Abs("../public/" + target)
	if e != nil {
		panic(e)
	}

	return path
}
