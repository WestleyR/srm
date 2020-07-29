// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-07-29
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2020 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

const srmVersion = "v2.0.0.a1, July 29, 2020"

func showVersion() {
  fmt.Printf("%s\n", srmVersion)
}

func main() {
	init_cache()

	help_flag := flag.BoolP("help", "h", false, "Show help output.")
	version_flag := flag.BoolP("version", "V", false, "Show srm version.")
	verbose_flag := flag.BoolP("verbose", "v", false, "Be more verbose.")
	debug_flag := flag.BoolP("debug", "d", false, "Show debug information.")
	recursive_flag := flag.BoolP("recursive", "r", false, "Be recursive, remove a directory.")
	force_flag := flag.BoolP("force", "f", false, "Remove a write-protected file")

	flag.Parse()
	args := flag.Args()

	set_debug(*debug_flag)
	print_debugf("THIS IS A DEBUG TEST")

	_ = *verbose_flag

  if *help_flag {
    flag.Usage()
    os.Exit(0)
  }

  if *version_flag {
    showVersion()
    os.Exit(0)
  }

	// If there are no files, then show the help menu and exit 1
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var options srmOptions
	options.force = *force_flag
	options.recursive = *recursive_flag

	exitCode := 0

	for _, f := range args {
		err := srmFile(f, options)
		if err != nil {
			// Error should be already formatted
			fmt.Fprintf(os.Stderr, "%s\n", err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
