package docs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func DocsCommand(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Generate markdowns docs for all myceld commands",
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			path := "docs/static/cmd"

			if err := os.MkdirAll(path, os.FileMode(0o644)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("Export cmd docs to: %s", path)
			err = doc.GenMarkdownTree(rootCmd, path)
			return err
		},
	}
	return cmd
}
