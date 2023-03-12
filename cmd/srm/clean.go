package cmd

import (
	"fmt"
	"strings"

	"github.com/WestleyR/srm/internal/pkg/srm"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "+c auto/all",
	Aliases: []string{"+clean"},
	Short:   "Clean removed cached.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := srm.New(nil)
		if err != nil {
			return err
		}

		switch strings.ToLower(args[0]) {
		case "auto":
			err = m.CleanCache(false, flagDryRun)
		case "all":
			err = m.CleanCache(true, flagDryRun)
		default:
			err = fmt.Errorf("Invalid option: %s", args[0])
		}
		if err != nil {
			return fmt.Errorf("failed to clean cache: %w", err)
		}

		// Only save the manager if theres no errors from modify it
		err = m.Close()
		if err != nil {
			return fmt.Errorf("failed to save manager: %w", err)
		}

		return nil
	},
}

var flagDryRun bool

func init() {
	cleanCmd.Flags().BoolVarP(&flagDryRun, "dry-run", "d", false, "Dry run, dont remove anything")

	rootCmd.AddCommand(cleanCmd)
}
