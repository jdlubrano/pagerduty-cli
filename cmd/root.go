package cmd

import (
  "os"
  "fmt"

  "github.com/spf13/cobra"

  "github.com/jdlubrano/pagerduty-cli/cmd/on_call"
  "github.com/jdlubrano/pagerduty-cli/cmd/schedules"
  "github.com/jdlubrano/pagerduty-cli/cmd/users"
  "github.com/jdlubrano/pagerduty-cli/version"
)

func NewRootCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "pagerduty-cli",
    Short: "A CLI to interact with Pagerduty",
    Long: "A CLI to interact with the Pagerduty API (https://v2.developer.pagerduty.com)",
    Version: version.Version,
  }

  cmd.AddCommand(on_call.NewOnCallCmd())
  cmd.AddCommand(schedules.NewSchedulesCmd())
  cmd.AddCommand(users.NewUsersCmd())

  return cmd
}

func Execute() {
  if err := NewRootCmd().Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
