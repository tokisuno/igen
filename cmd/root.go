package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var Language string
var Overwrite bool

var rootCmd = &cobra.Command {
	Use:   "igen",
	Short: "Generate .gitignore for project.",
	Long: `
Generate a .gitignore file for the project you're working on! 
!! This filetype must be supported by igen !!`,
    Run: func(cmd *cobra.Command, args []string) {
        var list []string

        dir, err := os.Executable()

        entries, err := os.ReadDir(filepath.Dir(dir) + "/cmd/templates"); check(err)

        cwd, err := os.Getwd(); check(err)

        for _, e := range entries {
            list = append(list, e.Name())
        }
        if slices.Contains(list, Language) {
            dat, err := os.ReadFile(filepath.Dir(dir) + "/cmd/templates/" + Language + "/.gitignore")
            check(err)

            if _, err := os.Stat(cwd + "/.gitignore") ; errors.Is(err, os.ErrNotExist) {
                fmt.Println("file doesn't exist")
                newGitignore(dat, cwd, Overwrite)
            } else {
                fmt.Println("file exists")
                appendToGitignore(dat, cwd, Overwrite)
            }
        }
	},
}

func newGitignore(dat []byte, cwd string, ow bool) {
    fmt.Printf("%t", ow)
    f, err := os.Create(cwd + "/.gitignore")
    check(err)

    defer f.Close()

    w := bufio.NewWriter(f)
    t, err := w.WriteString(string(dat))
    check(err)
    fmt.Printf("Wrote %d bytes\n", t)

    f.Sync()
    w.Flush()
}

func appendToGitignore(dat []byte, cwd string, ow bool) {
    fmt.Printf("%t", ow)
    f, err := os.OpenFile(cwd + "/.gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    check(err)

    defer f.Close()

    w := bufio.NewWriter(f)
    t, err := w.WriteString(string(dat))
    check(err)
    fmt.Printf("Wrote %d bytes to existing .gitignore\n", t)

    f.Sync()
    w.Flush()
}

func init() {
    rootCmd.Flags().BoolVarP(&Overwrite, "overwrite", "o", false, "Overwrite existing .gitignore")
    rootCmd.Flags().StringVarP(&Language, "new", "n", "", "Requires Language arg")
    rootCmd.MarkFlagRequired("new")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

