// Created by Westley R <westleyr@nym.hush.com> on 2024-10-12
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2024 Westley R. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/c2h5oh/datasize"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "+c auto/all",
	Aliases: []string{"+clean"},
	Short:   "Clean removed cached.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		auto := false

		switch strings.ToLower(args[0]) {
		case "auto":
			auto = true
		case "all":
		default:
			return fmt.Errorf("invalid option: %s", args[0])
		}

		return cleanTrash(auto, flagDryRun, flagVerbose)
	},
}

var (
	sizeLimitUpper = datasize.MB * 100
	sizeLimitLower = datasize.B * 50
	timeLimit      = time.Hour * 24 * 30 // 30 days
)

func cleanTrash(auto, dryRun, verbose bool) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	baseTrashPath := filepath.Join(home, ".cache", "srm3", "trashbin")

	files, err := os.ReadDir(baseTrashPath)
	if err != nil {
		return err
	}

	total := len(files)

	var saved datasize.ByteSize
	removed := 0

	rm := func(string) error { return nil }

	if dryRun {
		rm = func(_ string) error { return nil }
	} else {
		rm = func(path string) error { return os.RemoveAll(path) }
	}

	verbosef := func(_ string) {}

	if verbose {
		if dryRun {
			verbosef = func(f string) { fmt.Printf("Would remove: %s ...\n", f) }
		} else {
			verbosef = func(f string) { fmt.Printf("Removing: %s ...\n", f) }
		}
	}

	fmt.Printf("Going thought %d files...\n", total)

	for _, file := range files {
		basePath := filepath.Join(baseTrashPath, file.Name())

		blob, err := os.ReadFile(filepath.Join(basePath, "info.json"))
		if err != nil {
			fmt.Printf("Error reading file: %s (removing)\n", err)
			err := rm(basePath)
			if err != nil {
				return err
			}
			continue
		}

		var info *TrashInfo

		err = json.Unmarshal(blob, &info)
		if err != nil {
			fmt.Printf("Error unmarshalling data: %s (removing)\n", err)
			err := rm(basePath)
			if err != nil {
				return err
			}
			continue
		}

		if !auto || consideredOld(info) {
			verbosef(basePath)

			saved += info.Size
			removed++
			err := rm(basePath)
			if err != nil {
				return err
			}
		}
	}

	f := "Saved: %s (removed %d files)\n"
	if dryRun {
		f = "Would save: %s (and removes %d files)\n"
	}

	fmt.Printf(f, saved.HR(), removed)

	return nil
}

func consideredOld(info *TrashInfo) bool {
	return info.Size > sizeLimitUpper || info.Size < sizeLimitLower || time.Since(info.Time) > timeLimit
}

var (
	flagDryRun  bool
	flagVerbose bool
)

func init() {
	cleanCmd.Flags().BoolVarP(&flagDryRun, "dry-run", "d", false, "Dry run, dont remove anything")
	cleanCmd.Flags().BoolVarP(&flagVerbose, "verbose", "v", false, "Verbose output")

	rootCmd.AddCommand(cleanCmd)
}
