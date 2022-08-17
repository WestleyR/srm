//
//  main.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WestleyR/srm/internal/pkg/srm"
	flag "github.com/spf13/pflag"
)

const srmVersion = "v2.2.2"

func showVersion() {
	fmt.Printf("%s\n", srmVersion)
}

func main() {
	srm.InitCache()

	helpFlag := flag.BoolP("help", "h", false, "print this help output.")
	versionFlag := flag.BoolP("version", "V", false, "print srm version.")
	cleanCacheFlag := flag.StringP("clean", "c", "",
		fmt.Sprintf("clean the cache dir (options: auto,all).\nauto clean removes all files over %s and under %s",
			srm.FormatBytesToStr(srm.AutocleanSizeLimit),
			srm.FormatBytesToStr(srm.AutocleanSizeLower),
		),
	)
	dryRunFlag := flag.BoolP("dry-run", "n", false, "dry run when removing files or cleaning (WIP).")
	recursiveFlag := flag.BoolP("recursive", "r", false, "remove recursively.")
	emptyDirFlag := flag.BoolP("dir", "d", false, "remove empty directory (WIP).")
	forceFlag := flag.BoolP("force", "f", false, "remove a write-protected file.")
	listCacheFlag := flag.BoolP("list-cache", "l", false, "list recent removed files.")
	listAllCacheFlag := flag.BoolP("list-all-cache", "a", false, "list all removed files.")
	sortBySizeFlag := flag.BoolP("size", "S", false, "sort the cache list by size.")
	sortByTimeFlag := flag.BoolP("time", "t", false, "sort the cache list by time.")
	recoverFileFlag := flag.IntSliceP("recover", "s", nil, "recover a removed file(s) from the index list-cache.\nseperate numbers by commas (,) no spaces.")

	flag.Parse()
	args := flag.Args()

	// Help flag
	if *helpFlag {
		fmt.Printf("Copyright (c) 2020-2022 WestleyR. All rights reserved.\n")
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
		switch *cleanCacheFlag {
		case "auto":
			if err := srm.CleanCacheAUTO(*dryRunFlag); err != nil {
				fmt.Fprintf(os.Stderr, "%s: failed to clean cache: %s\n", os.Args[0], err)
				os.Exit(1)
			}
		case "all":
			fmt.Println("Not yet...")
		default:
			fmt.Fprintf(os.Stderr, "%s: invalid option for --clean: %s\n", os.Args[0], *cleanCacheFlag)
			os.Exit(1)
		}

		os.Exit(0)
	}

	// List cache flag
	if *listCacheFlag {
		err := srm.ListRecentCache()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// List all cache flag
	if *listAllCacheFlag {
		err := srm.ListAllCache(*sortBySizeFlag, *sortByTimeFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			os.Exit(1)
		}
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
	options := &srm.Options{}
	options.Force = *forceFlag
	options.Recursive = *recursiveFlag
	options.RemoveEmptyDir = *emptyDirFlag

	exitCode := 0

	srmManager, err := srm.New(options)
	if err != nil {
		log.Fatalf("Failed to init srm manager: %s", err)
	}

	for _, f := range args {
		err := srmManager.RM(f)
		if err != nil {
			// Error should be already formatted
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
