package assets

//go:generate statik -src=./static -dest=./

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/rakyll/statik/fs"
)

var (
	// store static files in memory by statik
	FileSystem http.FileSystem

	// if prefix is not empty, we get file content from disk
	prefixPath string
)

// if path is empty, load assets in memory
// or set FileSystem using disk files
func Load(path string) (err error) {
	prefixPath = path
	if prefixPath != "" {
		FileSystem = http.Dir(prefixPath)
		return nil
	} else {
		FileSystem, err = fs.New()
	}
	return err
}

func ReadFile(file string) (content string, err error) {
	if prefixPath == "" {
		file, err := FileSystem.Open(path.Join("/", file))
		if err != nil {
			return content, err
		}
		defer file.Close()
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			return content, err
		}
		content = string(buf)
	} else {
		file, err := os.Open(path.Join(prefixPath, file))
		if err != nil {
			return content, err
		}
		defer file.Close()
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			return content, err
		}
		content = string(buf)
	}
	return content, err
}
