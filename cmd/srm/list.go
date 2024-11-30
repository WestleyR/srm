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
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "+l",
	Aliases: []string{"+list"},
	Short:   "List cached files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listTrash()
	},
}

var stdWriter io.Writer = os.Stdout

func listTrash() error {
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
	count := total

	for _, file := range files {
		if !flagListAllCache && count > 10 {
			count--
			continue
		}

		basePath := filepath.Join(baseTrashPath, file.Name())

		blob, err := os.ReadFile(filepath.Join(basePath, "info.json"))
		if err != nil {
			log.Println("Error reading file:", err)
			continue
		}

		var info *TrashInfo

		err = json.Unmarshal(blob, &info)
		if err != nil {
			log.Println("Error unmarshalling data:", err)
			continue
		}

		fmt.Fprintf(stdWriter, "%d. %s (was %s) - %s\n", count, filepath.Join(strings.ReplaceAll(basePath, home, "~/"), "trash", filepath.Base(info.Were)), info.Were, info.Size.HR())

		count--
	}

	return nil

}

var (
	flagListAllCache bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&flagListAllCache, "all", "a", false, "List all cache")
}
