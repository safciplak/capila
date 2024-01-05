package file

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetRootDir returns the root dir.
func GetRootDir() string {
	dirPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	if strings.Contains(dirPath, "/T/") ||
		strings.HasSuffix(dirPath, "/T") ||
		strings.Contains(dirPath, "/go-build") {
		// go run
		dirPath = findParentDir()
	}

	if strings.HasSuffix(dirPath, ".go") {
		dirPath = filepath.Dir(dirPath)
	}

	return findModuleRoot(dirPath)
}

// findParentDir finds the directory of the correct caller (in the stacktrace)
func findParentDir() string {
	for a := 1; a < 10; a++ {
		_, curDir, _, _ := runtime.Caller(a)
		if !strings.HasSuffix(curDir, "rootDir.go") {
			return curDir
		}
	}

	return ""
}

/*
findModuleRoot returns the directory where the go.mod it
this only works when a go run is used
*/
func findModuleRoot(dirPath string) string {
	files, _ := filepath.Glob(dirPath + "/*.go")
	if len(files) == 0 {
		return dirPath // this is an executable in a non-go folder
	}

	for {
		if Exists(dirPath+"/go.mod") || dirPath == "/" {
			break
		}

		dirPath = path.Dir(dirPath)
	}

	return dirPath
}
