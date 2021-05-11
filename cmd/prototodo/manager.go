package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/msatodo/pkg/config"
	"github.com/sungjunyoung/msatodo/pkg/manager"
	"github.com/sungjunyoung/msatodo/pkg/manager/adding"
	"github.com/sungjunyoung/msatodo/pkg/manager/cache"
	"github.com/sungjunyoung/msatodo/pkg/manager/listing"
	"os"
)

const managerConfigEnv = "MSATODO_MANAGER_CONFIG_PATH"
const defaultManagerConfigPath = "/etc/msatodo/manager.yml"

func getManagerCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "manager",
		Short: "Msatodo manager to manage jobs",
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
