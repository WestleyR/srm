// Created by WestleyR <westleyr@nym.hush.com> on 2022-07-20
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package cmd

import (
	"fmt"

	"github.com/WestleyR/srm/internal/pkg/srm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "+l",
	Aliases: []string{"+list"},
	Short:   "List cached files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if ZeroOrOne(flagSortBySize, flagSortByTime) {
			return fmt.Errorf("can only use one sort flag")
		}

		// Dont close the manager since this should not change the data
		srmManager, err := srm.New(nil)
		if err != nil {
			return err
		}
		//defer srmManager.Close()

		if flagListAllCache {
			err := srmManager.ListAllCache(flagSortBySize, flagSortByTime)
			if err != nil {
				return err
			}
			return nil
		}

		err = srmManager.ListRecentCache()
		if err != nil {
			return err
		}

		return nil
	},
}

var (
	flagListAllCache bool
	flagSortByTime   bool
	flagSortBySize   bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&flagListAllCache, "all", "a", false, "List all cache.")
	listCmd.Flags().BoolVarP(&flagSortByTime, "time", "t", false, "Sort cache by time.")
	listCmd.Flags().BoolVarP(&flagSortBySize, "size", "s", false, "Sort cache by size.")
}
