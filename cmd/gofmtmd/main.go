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
	Version: "0.1.3",
	Short:   "This CLI formats Go Code in Markdown.",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) > 0 {
			filename = args[0]
		}
		var md []byte
		var err error
		if filename != "" {
			md, err = ioutil.ReadFile(filename)
		} else {
			md, err = ioutil.ReadAll(os.Stdin)
		}
		if err != nil {
			log.Fatalf("[gofmtmd] failed to read bytes from %v: %v", filename, err)
		}
		out, err := gofmtmd.FmtGoCodeInMarkdown(md)
		if err != nil {
			log.Fatal(err)
		}

		if filename != "" {
			if replace {
				err = ioutil.WriteFile(filename, out, 0644)
				if err != nil {
					log.Fatalf("[gofmtmd] failed to writes to %v: %v", filename, err)
				}
				return
			}
		}
		if outputfile != "" {
			err := ioutil.WriteFile(outputfile, out, 0644)
			if err != nil {
				log.Fatalf("[gofmtmd] failed to writes to %v: %v", outputfile, err)
			}
			return
		}

		writer := os.Stdout
		fmt.Fprint(writer, string(out))
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&replace, "replace", "r", false, "replace Go code with formatted code")
	rootCmd.PersistentFlags().StringVarP(&outputfile, "write", "w", "", "write result to file instead of stdout")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
