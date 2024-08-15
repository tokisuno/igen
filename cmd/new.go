package cmd

import (
	"fmt"
    "os"
    "log"
    "path/filepath"
    "slices"

	"github.com/spf13/cobra"
)

var Language string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new .gitignore file",
	Long: `Given a valid flag, create a new .gitignore file in current dir.`,
	Run: func(cmd *cobra.Command, args []string) {
        var list []string

        dir, _ := os.Executable()
        fmt.Println("Executable:", filepath.Dir(dir))

        entries, err := os.ReadDir(filepath.Dir(dir) + "/cmd/templates")
        if err != nil {
            log.Fatal(err)            
        }
        cwd, err := os.Getwd()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(cwd)

        for _, e := range entries {
            fmt.Println(e.Name())
            list = append(list, e.Name())
        }
        if slices.Contains(list, Language) {
            fmt.Println(true)
        }
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
    newCmd.Flags().StringVarP(&Language, "lang", "l", "", "Language (required)")
    newCmd.MarkFlagRequired("lang")
}
