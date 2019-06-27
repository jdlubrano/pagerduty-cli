package cmd

import (
  "os"
  "fmt"

  "github.com/spf13/cobra"

  "github.com/jdlubrano/pagerduty-cli/cmd/on_call"
)

func NewRootCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "pagerduty-cli",
    Short: "A CLI to interact with Pagerduty",
    Long: "A CLI to interact with the Pagerduty API (https://v2.developer.pagerduty.com)",
  }

  cmd.AddCommand(on_call.NewOnCallCmd())

  return cmd
}

func Execute() {
  if err := NewRootCmd().Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
