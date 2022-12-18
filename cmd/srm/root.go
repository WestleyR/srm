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
	"log"
	"os"

	"github.com/WestleyR/srm/internal/pkg/srm"
	"github.com/spf13/cobra"
)

var Version string = "[unknown]"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "srm [flags] -- [./FILE...]",
	Short: "Remove file(s) or directory(s) into cache, allows for undo/recover.",
	Long: `Copyright (c) 2020-2022 WestleyR. All rights reserved.
This software is licensed under the terms of The Clear BSD License.
Source code: https://github.com/WestleyR/srm`,

	SilenceUsage: true,
	Version:      Version,
	Args:         cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check flag commands first
		switch {
		case flagList:
			fmt.Printf("WARNING: Flag `-l` is deprecated. Use `+l` command.\n")

			// Dont close the manager since this should not change the data
			srmManager, err := srm.New(nil)
			if err != nil {
				return err
			}

			if flagListAll {
				err := srmManager.ListAllCache(false, false)
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
		case flagRecover != -1:
			fmt.Printf("WARNING: Flag `-s` is deprecated. Use `+s` command.\n")

			srmManager, err := srm.New(nil)
			if err != nil {
				return err
			}
			defer srmManager.Close()

			err = srmManager.Recover(flagRecover)
			if err != nil {
				return err
			}

			return nil
		}

		if len(args) == 0 {
			return cmd.Usage()
		}

		options := &srm.Options{
			Recursive:      flagRecursive,
			Force:          flagForce,
			RemoveEmptyDir: flagDir,
		}

		exitCode := 0

		srmManager, err := srm.New(options)
		if err != nil {
			log.Fatalf("Failed to init srm manager: %s", err)
		}
		defer srmManager.Close()

		for _, f := range args {
			err := srmManager.RM(f)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Flags vars
var (
	flagRecursive bool
	flagForce     bool
	flagDir       bool

	// Flag commands
	flagList    bool
	flagListAll bool
	flagRecover int
)

func init() {
	rootCmd.Flags().BoolVarP(&flagRecursive, "recursive", "r", false, "Remove recursively.")
	rootCmd.Flags().BoolVarP(&flagForce, "force", "f", false, "Remove protected files.")
	rootCmd.Flags().BoolVarP(&flagDir, "dir", "d", false, "Remove empty directory.")

	rootCmd.Flags().BoolVarP(&flagList, "list", "l", false, "List removed files (cache)")
	rootCmd.Flags().BoolVarP(&flagListAll, "list-all", "a", false, "List all cache")
	rootCmd.Flags().MarkHidden("list-all")
	rootCmd.Flags().IntVarP(&flagRecover, "recover", "s", -1, "Recover a cache file by index")
}
