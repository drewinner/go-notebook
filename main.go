package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version     = "<UNDEFINED>"
	BinaryName  = "<UNDEFINED>"
	versionFlag bool
)

var RootCmd = &cobra.Command{
	Use: "open-falcon",
	RunE: func(c *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("%s version %s\n", BinaryName, Version)
			return nil
		}
		return c.Usage()
	},
}

func init() {
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "show version")
}
func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
