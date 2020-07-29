// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-07-28
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
	"log"

	flag "github.com/spf13/pflag"
)

func main() {
	init_cache()

	// TODO: fix ty0s
	verbose_flag := flag.BoolP("verbose", "v", false, "Be more verbose.")
	debug_flag := flag.BoolP("debug", "d", false, "Show debug infomation.")
	recursve_flag := flag.BoolP("recursve", "r", false, "Be recurs... remove a directory")
	force_flag := flag.BoolP("force", "f", false, "Be forceful")

	flag.Parse()
	args := flag.Args()

	set_debug(*debug_flag)
	print_debugf("THIS IS A DEBUG TEST")

	fmt.Printf("v=%t r=%t f=%t, args: %v\n", *verbose_flag, *recursve_flag, *force_flag, args)

	for _, f := range args {
		fmt.Printf("Removing file: %v...\n", f)
		err := srmFile(f)
		if err != nil {
			log.Printf("ERROR: failed to remove: %s\n", f)
		}
	}
}
