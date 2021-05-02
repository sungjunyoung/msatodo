package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/prototodo/pkg/manager"
	"os"
)

func getManagerCommand() *cobra.Command {
	mgr := manager.NewManager()

	root := &cobra.Command{
		Use:   "manager",
		Short: "Prototodo manager to manage jobs",
		Run: func(cmd *cobra.Command, args []string) {
			if err := mgr.Start(); err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		},
	}

	return root
}
