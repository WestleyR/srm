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
	"os"
	"strconv"

	"github.com/WestleyR/srm/internal/pkg/srm"
	"github.com/spf13/cobra"
)

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:     "+s",
	Aliases: []string{"+recover"},
	Short:   "Recover a cached file by index.",
	Long: `Recover a cached file by index. Only can recover by index
only for the last 10 items, otherwise you need to manually cp by full path.`,
	Args: cobra.ArbitraryArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		srmManager, err := srm.New(nil)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			// If no args, then just list recent cache as helper
			return srmManager.ListRecentCache()
		}

		// Only defer closing if we are going to modify the data
		defer srmManager.Close()

		for _, a := range args {
			n, err := strconv.Atoi(a)
			if err != nil {
				return err
			}

			err = srmManager.Recover(n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)
}
