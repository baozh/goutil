package system

import "os"

// IsExist check whether filename(denote a file or directory) is exist.
// return true if exist, else return false.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


// ScanDir get the name of file and dir in the directory.
// If dir isn't exist, then return empty slice.
func ScanDir(directory string) []string {
	file, err := os.Open(directory)
	if err != nil {
		return []string{}
	}
	names, err := file.Readdirnames(-1)
	if err != nil {
		return []string{}
	}
	return names
}

// IsDir check whether 'filename' is a directory.
// If 'filename' is a relative path, it will check "." as the current path.
func IsDir(filename string) bool {
	return isFileOrDir(filename, true)
}

// IsFile check whether 'filename' is a file
func IsFile(filename string) bool {
	return isFileOrDir(filename, false)
}

func isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

