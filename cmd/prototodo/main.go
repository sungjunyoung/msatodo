package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	root := &cobra.Command{
		Use: "prototodo",
	}

	client, err := getClientCommand()
	if err != nil {
		log.Fatal(err)
	}
	root.AddCommand(client)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
