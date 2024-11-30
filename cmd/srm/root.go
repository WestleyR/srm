// Created by Westley R <westleyr@nym.hush.com> on 2022-07-20
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2022-2024 Westley R. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/c2h5oh/datasize"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

/*

Structor:
~/.cache/srm3/trash/XXXX/
  trash/ (the actural file that was trashed)
  info.json (where the file was)

*/

var Version string = "[unknown]"

// Flags vars
var (
	flagRecursive bool
	flagForce     bool
	flagDir       bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "srm [flags] -- [./FILE...]",
	Short: "Remove file(s) or directory(s) into cache, allows for undo/recover.",
	Long: `Copyright (c) 2020-2023 WestleyR. All rights reserved.
This software is licensed under the terms of The Clear BSD License.
Source code: https://github.com/WestleyR/srm

This cli and flags structor is experimental, if you have a suggestion please
open a github issue at: github.com/WestleyR/srm`,

	SilenceUsage: true,
	Version:      Version,
	Args:         cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Usage()
		}

		exitCode := 0

		opts := &Opts{
			Force:     flagForce,
			Recursive: flagRecursive,
		}

		for _, f := range args {
			err := srmFile(f, opts)
			if err != nil {
				// Error should be already formatted
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
				exitCode = 1
			}
		}

		if exitCode > 0 {
			os.Exit(exitCode)
		}

		return nil
	},
}

type TrashInfo struct {
	Time time.Time         `json:"time"`
	Were string            `json:"were"`
	Size datasize.ByteSize `json:"size"`
}

type Opts struct {
	Force     bool
	Recursive bool
}

func srmFile(file string, opts *Opts) error {
	path, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	// Pre checks
	if !doesPathExist(path) {
		if opts.Force {
			return nil
		}
		return fmt.Errorf("%s: does not exist", path)
	}

	isDir := isPathADir(path)

	if isDir && !opts.Recursive {
		return fmt.Errorf("%s: is a directory", path)
	}

	size := getFileSize(path)

	if isDir {
		// Is a directory
		size = getDirSize(path)

		if !opts.Recursive {
			return fmt.Errorf("%s: is a directory", path)
		}

		if !opts.Force {
			err := checkForWriteProtectedFileIn(path)
			if err != nil {
				return err
			}
		}
	} else {
		// Is a file
		if !opts.Force && !checkIfFileIsWriteProtected(path) {
			return fmt.Errorf("%s: is write protected", path)
		}
	}

	// Trash it

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	baseTrashPath := filepath.Join(home, ".cache", "srm3", "trashbin", xid.New().String())
	trashPath := filepath.Join(baseTrashPath, "trash")

	err = os.MkdirAll(trashPath, 0o755)
	if err != nil {
		return err
	}

	info := &TrashInfo{
		Time: time.Now(),
		Were: path,
		Size: size,
	}

	blob, err := json.Marshal(info)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(baseTrashPath, "info.json"), blob, 0o644)
	if err != nil {
		return err
	}

	err = os.Rename(path, trashPath)
	if err != nil {
		cmd := exec.Command("mv", path, trashPath)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: failed to move to trash directory %s", path, trashPath)
		}
	}

	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&flagRecursive, "recursive", "r", false, "Remove recursively.")
	rootCmd.Flags().BoolVarP(&flagForce, "force", "f", false, "Remove protected files.")
	rootCmd.Flags().BoolVarP(&flagDir, "dir", "d", false, "Remove empty directory.")
}
