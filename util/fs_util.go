package util

import (
	"github.com/spf13/afero"
	"workspaces/logging"
	"errors"
)

var fsLogger = logging.NewLoggableEntity(
	"fs_util",
	logging.Fields{ "module": "fs_util" },
)

func EnsureDirExist(path string, fs afero.Fs) error {
	return ensureExist(FOLDER, path, fs, true)
}

func EnsureFileExist(path string, fs afero.Fs) error {
	return ensureExist(FILE, path, fs, true)
}

func EnsureDirExistSilent(path string, fs afero.Fs) error {
	return ensureExist(FOLDER, path, fs, false)
}

func EnsureFileExistSilent(path string, fs afero.Fs) error {
	return ensureExist(FOLDER, path, fs, false)
}

func ensureExist(ft FileType, path string, fs afero.Fs, verbose bool) error {
	obj, exist := FileTypeObjMap[ft]
	if !exist {
		err := errors.New(UNKNOWN_FILE_TYPE)
		if verbose {
			fsLogger.ErrorWithFields(logging.Fields{
				"file_type": ft,
			}, err)
		}
		return err
	}

	loggingFields := logging.Fields{
		"path": path,
		"file_type": ft,
		"file_type_name": obj.name,
		"file_type_exist_func": obj.existFunc,
	}

	fsLogger.Debugf( loggingFields, "Trying to confirm %s exist", obj.name )
	if exist, err := obj.existFunc(fs, path); err != nil {
		if verbose {
			fsLogger.Errorf( loggingFields, err, "failed to check if %s exist", obj.name )
		}

		return err
	} else if !exist {
		err = errors.New(obj.notFoundErrorCode)
		if verbose {
			fsLogger.Errorf( loggingFields, err, obj.notFoundErrorCode )
		}

		return err
	}

	return nil
}
