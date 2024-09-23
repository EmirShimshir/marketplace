package domain

import (
	"io"
)

type Url string

func (url Url) String() string {
	return string(url)
}

type File struct {
	Name   string
	Path   string
	Reader io.Reader
}

func (f *File) Validate() error {
	if f.Name == "" {
		return ErrFilenameEmpty
	}
	if f.Path == "" {
		return ErrFilepathEmpty
	}
	if f.Reader == nil {
		return ErrFileReaderEmpty
	}
	return nil
}
