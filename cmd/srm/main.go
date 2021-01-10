// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2021-01-09
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2021 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

package main

import (
	"fmt"
	"os"

	"github.com/WestleyR/srm/internal/srm"
	flag "github.com/spf13/pflag"
)

const srmVersion = "v2.0.0.a1, Sep 22, 2020"

func showVersion() {
	fmt.Printf("%s\n", srmVersion)
}

func main() {
	srm.InitCache()

	helpFlag := flag.BoolP("help", "h", false, "Show help output.")
	versionFlag := flag.BoolP("version", "V", false, "Show srm version.")
	verboseFlag := flag.BoolP("verbose", "v", false, "Be more verbose.")
	debugFlag := flag.BoolP("debug", "d", false, "Show debug information.")
	cleanCacheFlag := flag.StringP("clean", "c", "", "Clean the cache dir (options: auto,all).")
	recursiveFlag := flag.BoolP("recursive", "r", false, "Be recursive, remove a directory.")
	forceFlag := flag.BoolP("force", "f", false, "Remove a write-protected file.")
	listCacheFlag := flag.BoolP("list-cache", "l", false, "List recent removed files.")
	listAllCacheFlag := flag.BoolP("list-all-cache", "a", false, "List all cached files.")
	recoverFileFlag := flag.IntSliceP("recover", "s", nil, "Recover a removed file(s) from the index list-cache.\nSeperate numbers by commas (,) no spaces.")

	flag.Parse()
	args := flag.Args()

	srm.SetDebug(*debugFlag)
	srm.PrintDebugf("THIS IS A DEBUG TEST")

	_ = *verboseFlag

	// Help flag
	if *helpFlag {
		fmt.Printf("Copyright (c) 2019-2020 WestleyR. All rights reserved.\n")
		fmt.Printf("This software is licensed under the terms of The Clear BSD License.\n")
		fmt.Printf("Source code: https://github.com/WestleyR/srm\n")
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
