package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/prototodo/pkg/client"
	"github.com/sungjunyoung/prototodo/pkg/config"
	"os"
)

func getClientCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "client",
		Short: "Prototodo client to query jobs",
	}

	add := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			cli, err := getClient()
			if err != nil {
				os.Exit(1)
			}

			res, err := cli.AddJob(args[0], args[1])
			if err != nil {
				os.Exit(1)
			}

			fmt.Println(res)
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

func getClient() (client.Client, error) {
	cli, err := client.NewClient(config.NewClientLoader("/etc/prototodo/config.yml"))
	if err != nil {
		return nil, err
	}

	return cli, nil
}
