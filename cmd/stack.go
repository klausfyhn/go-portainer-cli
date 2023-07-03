/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/portainer/portainer/api"
	"github.com/spf13/cobra"
)

// stackCmd represents the stack command
var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var stackLsCmd = &cobra.Command{
	Use: "ls",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		client := http.Client{}

		req, err := http.NewRequest(http.MethodGet, PortainerUrl+"/api/stacks", nil)

		if err != nil {
			panic(err)
		}

		req.Header.Set("X-API-Key", PortainerApiKey)

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var stacks []portainer.Stack
		err = json.Unmarshal(body, &stacks)
		if err != nil {
			panic(err)
		}

		const padding = 3

		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

		fmt.Fprintln(w, "ID", "\t", "NAME", "\t", "CREATED BY", "\t", "ENDPOINT ID")

		for _, stack := range stacks {
			fmt.Fprintln(w, stack.ID, "\t", stack.Name, "\t", stack.CreatedBy, "\t", stack.EndpointID)
		}

		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(stackCmd)

	stackCmd.AddCommand(stackLsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
