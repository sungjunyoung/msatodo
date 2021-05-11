package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func getClientCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "client",
		Short: "Msatodo client to query jobs",
	}

	add := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add")
		},
	}

	ls := &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("client ls\n")
		},
	}

	root.AddCommand(add)
	root.AddCommand(ls)
	return root
}
