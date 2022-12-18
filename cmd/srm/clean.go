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
		var err error
		switch strings.ToLower(args[0]) {
		case "auto":
			err = srm.CleanCacheAUTO(flagDryRun)
		case "all":
			err = fmt.Errorf("not ready")
		default:
			err = fmt.Errorf("Invalid option: %s", args[0])
		}

		return err
	},
}

var flagDryRun bool

func init() {
	cleanCmd.Flags().BoolVarP(&flagDryRun, "dry-run", "d", false, "Dry run, dont remove anything")

	rootCmd.AddCommand(cleanCmd)
}
