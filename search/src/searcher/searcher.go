package searcher

import (
	"fmt"
	"path/filepath"
	. "search/src/opts"
)

func FindFiles(opts Opts) {
	matches, err := filepath.Glob(opts.Base + opts.Search)
	if err != nil {
		fmt.Println(err)
	}
	if len(matches) != 0 {
		fmt.Println("Found : ", matches)
	}
}
