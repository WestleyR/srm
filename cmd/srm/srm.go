// Created by WestleyR on July 28, 2020
// Source code: https://github.com/WestleyR/srm
// Last modified data: 2021-01-28
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2020-2021 WestleyR
// All rights reserved.
//

package main

import (
	"fmt"
	"os"

	"github.com/WestleyR/srm/internal/srm"
	flag "github.com/spf13/pflag"
)

const srmVersion = "v2.0.0.a2 2021-01-11"

func showVersion() {
	fmt.Printf("%s\n", srmVersion)
}

func main() {
	srm.InitCache()

	helpFlag := flag.BoolP("help", "h", false, "Print this help output.")
	versionFlag := flag.BoolP("version", "V", false, "print srm version.")
	cleanCacheFlag := flag.StringP("clean", "c", "", "Clean the cache dir (options: auto,all).")
	recursiveFlag := flag.BoolP("recursive", "r", false, "Remove recursively.")
	forceFlag := flag.BoolP("force", "f", false, "Remove a write-protected file.")
	listCacheFlag := flag.BoolP("list-cache", "l", false, "List recent removed files.")
	listAllCacheFlag := flag.BoolP("list-all-cache", "a", false, "List all removed files.")
	recoverFileFlag := flag.IntSliceP("recover", "s", nil, "Recover a removed file(s) from the index list-cache.\nSeperate numbers by commas (,) no spaces.")

	flag.Parse()
	args := flag.Args()

	// Help flag
	if *helpFlag {
		fmt.Printf("Copyright (c) 2020-2021 WestleyR. All rights reserved.\n")
		fmt.Printf("This software is licensed under the terms of The Clear BSD License.\n")
		fmt.Printf("Source code: https://github.com/WestleyR/srm\n")
		fmt.Printf("\n")
		flag.Usage()
		os.Exit(0)
	}

	// Version flag
	if *versionFlag {
		showVersion()
		os.Exit(0)
	}

	// Clean cache flag
	if *cleanCacheFlag != "" {
		if *cleanCacheFlag == "auto" {
			if err := srm.CleanCacheAUTO(); err != nil {
				fmt.Fprintf(os.Stderr, "%s: failed to clean cache: %s\n", os.Args[0], err)
				os.Exit(1)
			}
		} else if *cleanCacheFlag == "all" {
			fmt.Println("Not yet...")
		} else {
			fmt.Fprintf(os.Stderr, "%s: invalid option for --clean: %s\n", os.Args[0], *cleanCacheFlag)
			os.Exit(1)
		}

		os.Exit(0)
	}

	// List cache flag
	if *listCacheFlag {
		srm.ListRecentCache()
		os.Exit(0)
	}

	// List all cache flag
	if *listAllCacheFlag {
		os.Exit(0)
	}

	// Recover file flag
	if *recoverFileFlag != nil {
		exitCode := 0
		for _, n := range *recoverFileFlag {
			err := srm.RecoverFileFromTrashIndex(n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
				exitCode = 1
			}
		}

		os.Exit(exitCode)
	}

	// If there are no files, then show the help menu and exit 1
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// If not doing anything else, then remove the files passed
	var options srm.SrmOptions
	options.Force = *forceFlag
	options.Recursive = *recursiveFlag

	exitCode := 0

	for _, f := range args {
		err := srm.SrmFile(f, options)
		if err != nil {
			// Error should be already formatted
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
