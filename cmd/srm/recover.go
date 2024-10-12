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
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:     "+r [CACHE_INDEX...]",
	Aliases: []string{"+recover"},
	Short:   "Recover a cached file by index to its original path",
	Long: `Recover a cached file by index from list. If no index is specified, instead will print
the recent cache files. index 1 = latest removed file.`,
	Args: cobra.ArbitraryArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// If no args, then just list recent cache as helper
			return listTrash()
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		baseTrashPath := filepath.Join(home, ".cache", "srm3", "trashbin")

		files, err := os.ReadDir(baseTrashPath)
		if err != nil {
			return err
		}

		count := len(files)

		// TODO: Better error handling/continues
		for _, indexStr := range args {
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return err
			}

			if count-index < 0 || count-index >= count {
				return fmt.Errorf("index out of range: %d", count-index)
			}

			toRecover := files[count-index]

			basePath := filepath.Join(baseTrashPath, toRecover.Name())

			blob, err := os.ReadFile(filepath.Join(basePath, "info.json"))
			if err != nil {
				return err
			}

			var info *TrashInfo

			err = json.Unmarshal(blob, &info)
			if err != nil {
				return err
			}

			fullTrashPath := filepath.Join(basePath, "trash", filepath.Base(info.Were))

			fmt.Printf("Recovering: %s to %s\n", fullTrashPath, info.Were)

			// Check to see if it already exists
			// Only try 10 times
			tmpName := info.Were
			for i := 1; i < 11; i++ {
				if _, err := os.Lstat(tmpName); err == nil {
					tmpName = info.Were + "." + strconv.Itoa(i)
				} else {
					break
				}
			}

			err = os.Rename(fullTrashPath, tmpName)
			if err != nil {
				return err
			}

			// Delete the empty directory and info.json

			err = os.Remove(filepath.Join(basePath, "info.json"))
			if err != nil {
				return err
			}

			err = os.Remove(filepath.Join(basePath, "trash"))
			if err != nil {
				return err
			}

			err = os.Remove(basePath)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)
}
