package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/po3rin/gofmtmd"
	"github.com/spf13/cobra"
)

var (
	replace    bool
	outputfile string
)

var rootCmd = &cobra.Command{
	Use:     "gofmtmd",
	Version: "0.0.1",
	Short:   "This CLI formats Go Code in Markdown.",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		md, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("failed to read bytes from %v: ", filename)
		}
		out, err := gofmtmd.FmtGoCodeInMarkdown(md)
		if err != nil {
			log.Fatal(err)
		}

		if replace {
			err = ioutil.WriteFile(filename, out, 0644)
			if err != nil {
				log.Fatalf("failed to writes to %v: ", filename)
			}
		}
		if outputfile != "" {
			err = ioutil.WriteFile(outputfile, out, 0644)
			if err != nil {
				log.Fatalf("failed to writes to %v: ", filename)
			}
		}
		if !replace && outputfile == "" {
			writer := os.Stdout
			fmt.Fprint(writer, string(out))
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&replace, "replace", "r", false, "replace Go code with formated code")
	rootCmd.PersistentFlags().StringVarP(&outputfile, "write", "w", "", "write result to file instead of stdout")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
