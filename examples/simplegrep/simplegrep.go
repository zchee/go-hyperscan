// Copyright 2017 The go-hyperscan Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/01org/hyperscan/blob/a00bd3167cda4d963ed71220cbbb703c06955f2d/examples/simplegrep.c

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unsafe"

	hyperscan "github.com/zchee/go-hyperscan"
)

func readInputData(inputFn string) ([]byte, uint32, error) {
	f, err := os.Open(inputFn)
	if err != nil {
		return nil, 0, fmt.Errorf("ERROR: unable to open file %q: %v", inputFn, err)
	}
	defer f.Close()

	if _, err := f.Seek(0, os.SEEK_END); err != nil {
		return nil, 0, fmt.Errorf("ERROR: unable to seek file %q: %v", inputFn, err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, 0, err
	}

	return buf, uint32(len(buf)), nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pattern> <input file>\n", args[0])
		os.Exit(-1)
	}

	pattern := args[1]
	inputFN := args[2]

	database := hyperscan.NewDatabase()
	defer hyperscan.FreeDatabase(database)
	// var database [][]hyperscan.Database
	var compileErr [][]hyperscan.CompileError
	if cErr := hyperscan.Compile(pattern, hyperscan.FlagDotall, hyperscan.ModeBlock, nil, &database, &compileErr); cErr != hyperscan.Success {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to compile pattern %q: %v\n", pattern, hyperscan.Error(cErr))
		fmt.Println(compileErr)
		for _, cmpErr := range compileErr {
			fmt.Fprintln(os.Stderr, cmpErr)
			hyperscan.FreeCompileError(cmpErr)
		}
		os.Exit(-1)
	}

	inputData, length, err := readInputData(inputFN)
	if err != nil {
		log.Fatal(err)
	}
	if length == 0 {
		hyperscan.FreeDatabase(database[0])
		os.Exit(-1)
	}

	scratch := hyperscan.NewScratch()
	defer hyperscan.FreeScratch([]hyperscan.Scratch{*scratch})
	if cErr := hyperscan.AllocScratch(database[0], [][]hyperscan.Scratch{[]hyperscan.Scratch{*scratch}}); cErr != hyperscan.Success {
		fmt.Fprint(os.Stderr, "ERROR: Unable to allocate scratch space. Exiting\n")
		os.Exit(-1)
	}

	fmt.Printf("Scanning %d bytes with Hyperscan\n", length)

	eventHandler := hyperscan.NewMatchEventHandlerRef(nil)
	if cErr := hyperscan.Scan(database[0], string(inputData), length, 0, []hyperscan.Scratch{*scratch}, *eventHandler, unsafe.Pointer(&pattern)); cErr != hyperscan.Success {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to scan input buffer. Exiting\n")
		os.Exit(-1)
	}
}
