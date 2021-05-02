package main

import (
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/prototodo/pkg/manager"
	"os"
)

func getManagerCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "manager",
		Short: "Prototodo manager to manage jobs",
		Run: func(cmd *cobra.Command, args []string) {
			mgr, err := getManager()
			if err != nil {
				os.Exit(1)
			}

			if err := mgr.Start(); err != nil {
				os.Exit(1)
			}
		},
	}

	return root
}

func getManager() (manager.Manager, error) {
	mgr := manager.NewManager()
	return mgr, nil
}
