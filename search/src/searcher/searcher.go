package searcher

import (
	"fmt"
	"github.com/gobwas/glob"
	"io/ioutil"
	"path/filepath"
	. "search/src/opts"
	"sync"
)

func FindFiles(opts Opts) {

	/**
	matches, err := filepath.Glob(opts.Base + opts.Search)
	if err != nil {
		fmt.Println(err)
	}
	if len(matches) != 0 {
		fmt.Println("Found : ", matches)
	}*/
	//err := filepath.Walk(opts.Base, visit)

	/**
	err := filepath.Walk(opts.Base, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			matches, err := filepath.Glob(path + string(os.PathSeparator) + opts.Search)
			if err != nil {
				fmt.Println(err)
			}
			if len(matches) != 0 {
				for _, match := range matches {
					fmt.Println(info.Name() + string(os.PathSeparator) + match)
				}
			}
		}
		//fmt.Println(hey)
		//fmt.Printf("Visited: %s\n", path)
		return nil
	})
	fmt.Printf("filepath.Walk() returned %v\n", err)*/

	// Needed to wait for goroutines
	/**
	var wg sync.WaitGroup
	find(opts.Base, opts.Search, wg)
	wg.Wait()

	finished := make(chan bool)
	find2(opts.Base, opts.Search, finished)
	<- finished


	*/
	wg := &sync.WaitGroup{}
	wg.Add(1)
	find(opts.Base, opts.Search, wg)
	wg.Wait()

}

// WaitGroup needs to be a pointer, so that Go does not make a unnecessary
// copy that calls done.
// Source: https://stackoverflow.com/a/25234899
func find(path string, pattern glob.Glob, wg *sync.WaitGroup) {
	// WaitGroup that counts all goroutines this goroutine starts
	nextWg := &sync.WaitGroup{}
	// Flag that checks if the element is the last element in the file tree
	end := true
	// Get all files/directories in the current path
	files, err := ioutil.ReadDir(path)
	for _, entry := range files {
		// Build new path from file/directory
		filePath := filepath.Join(path, entry.Name())
		// Check if filename or filepath matches pattern
		check := pattern.Match(entry.Name()) || pattern.Match(filePath)
		// If so print it
		if check {
			fmt.Println(filePath)
		}
		// If entry is a directory spawn a new goroutine that searches for files
		// in this directory
		if entry.IsDir() {
			end = false
			nextWg.Add(1)
			go find(filePath, pattern, nextWg)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	if !end {
		nextWg.Wait()
	}
	wg.Done()
}
