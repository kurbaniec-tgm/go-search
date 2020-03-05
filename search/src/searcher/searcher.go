package searcher

import (
	"fmt"
	"path/filepath"
)

func FindFile(targetDir string, pattern string) {
	matches, err := filepath.Glob(targetDir + pattern)
	if err != nil {
		fmt.Println(err)
	}
	if len(matches) != 0 {
		fmt.Println("Found : ", matches)
	}
}
