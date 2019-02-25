package util

import (
	"github.com/spf13/afero"
)

type FileType int
const (
	FOLDER FileType = iota
	FILE
)

type FileTypeObject struct {
	name string
	existFunc func(fs afero.Fs, path string) (bool, error)
	notFoundErrorCode string
}

var FileTypeObjMap = map[FileType]FileTypeObject{
	FOLDER: {
		name: "folder",
		existFunc: afero.DirExists,
		notFoundErrorCode: FOLDER_NOT_FOUND,
	},
	FILE: {
		name: "folder",
		existFunc: afero.Exists,
		notFoundErrorCode: FILE_NOT_FOUND,
	},
}

