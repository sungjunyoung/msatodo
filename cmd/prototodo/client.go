package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func getClientCommand() (*cobra.Command, error) {
	client := &cobra.Command{
		Use:   "client",
		Short: "prototodo client",
	}

	add := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	client.AddCommand(add)
	return client, nil
}
