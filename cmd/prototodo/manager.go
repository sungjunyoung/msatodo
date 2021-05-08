package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/prototodo/pkg/config"
	"github.com/sungjunyoung/prototodo/pkg/manager"
	"github.com/sungjunyoung/prototodo/pkg/manager/adding"
	"github.com/sungjunyoung/prototodo/pkg/manager/cache"
	"github.com/sungjunyoung/prototodo/pkg/manager/listing"
	"os"
)

const managerConfigEnv = "PROTOTODO_MANAGER_CONFIG_PATH"
const defaultManagerConfigPath = "/etc/prototodo/manager.yml"

func getManagerCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "manager",
		Short: "Prototodo manager to manage jobs",
		Run: func(cmd *cobra.Command, args []string) {
			mgr, err := getManager()
			if err != nil {
				os.Exit(1)
			}

			errCh := make(chan error)
			go mgr.ServeGrpc(errCh)

			select {
			case err := <-errCh:
				close(errCh)
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return root
}

func getManager() (manager.Manager, error) {
	mgr := manager.Manager{}

	configPath := os.Getenv(managerConfigEnv)
	if configPath == "" {
		configPath = defaultManagerConfigPath
	}

	memoryCache := cache.NewMemory()
	addingSvc := adding.NewService(memoryCache)
	listingSvc := listing.NewService(memoryCache)

	mgr, err := manager.NewManager(
		config.NewManagerLoader(configPath),
		addingSvc,
		listingSvc,
	)
	if err != nil {
		return mgr, err
	}

	return mgr, nil
}
