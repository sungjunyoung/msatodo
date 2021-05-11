package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	root := &cobra.Command{
		Use: "msatodo",
	}

	root.AddCommand(
		getClientCommand(),
		getManagerCommand(),
	)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
