package fs

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// flexible is a concrete implementation of the FlexibleFS interface.
type flexible struct {
	files fs.FS
}

// NewDir creates a FlexibleFS instance backed by the local file system
// at the specified directory path.
func NewDir(path string) FlexibleFS {
	return &flexible{files: os.DirFS(path)}
}

// NewEmbed creates a FlexibleFS instance backed by the provided embedded
// file system.
func NewEmbed(embeddedFS embed.FS) FlexibleFS {
	return &flexible{files: embeddedFS}
}

func (f flexible) Exists(path string) (bool, error) {
	path = normalizePath(path)
	info, err := fs.Stat(f.files, path)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return !info.IsDir(), nil
}

func (f flexible) Open(path string) (fs.File, error) {
	path = normalizePath(path)
	return f.files.Open(path)
}

func (f flexible) ReadFile(path string) ([]byte, error) {
	path = normalizePath(path)
	return fs.ReadFile(f.files, path)
}

func (f flexible) Find(dir, pattern string) (*string, error) {
	var result string
	dir = normalizePath(dir)

	// Compile the regex pattern.
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid regex pattern")
	}

	// Walk through the directory to find the file.
	err = fs.WalkDir(f.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() && rx.MatchString(entry.Name()) {
			result = normalizePath(path)
			return fs.SkipAll
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if result == "" {
		return nil, nil
	}

	return &result, nil
}

func (f flexible) Search(dir, phrase, ignore, ext string) (*string, error) {
	var result string
	var err error
	dir = normalizePath(dir)
	ext = strings.TrimLeft(ext, ".")

	// Compile the find regex.
	var rxFind *regexp.Regexp
	if ext == "" {
		rxFind, err = regexp.Compile(phrase + ".*")
		if err != nil {
			return nil, errors.New("invalid search pattern")
		}
	} else {
		rxFind, err = regexp.Compile(phrase + `.*\.` + ext)
		if err != nil {
			return nil, errors.New("invalid search pattern")
		}
	}

	// Compile the ignore regex if provided.
	var rxSkip *regexp.Regexp
	if ignore != "" {
		rxSkip, err = regexp.Compile(".*" + ignore + ".*")
		if err != nil {
			return nil, errors.New("invalid ignore pattern")
		}
	}

	// Walk through the directory to search for the file.
	err = fs.WalkDir(f.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() &&
			rxFind.MatchString(entry.Name()) &&
			(rxSkip == nil || !rxSkip.MatchString(entry.Name())) {
			result = normalizePath(path)
			return fs.SkipAll
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if result == "" {
		return nil, nil
	}

	return &result, nil
}

func (f flexible) Lookup(dir, pattern string) ([]string, error) {
	var result []string
	dir = normalizePath(dir)

	// Compile the regex pattern.
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid regex pattern")
	}

	// Walk through the directory to find matching files.
	err = fs.WalkDir(f.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() && rx.MatchString(entry.Name()) {
			result = append(result, normalizePath(path))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (f flexible) FS() fs.FS {
	return f.files
}

func (f flexible) Http() http.FileSystem {
	return http.FS(f.files)
}
