package on_call

import (
  "fmt"

  "github.com/spf13/cobra"
)

func NewOnCallCmd() *cobra.Command {
  onCallCmd := &cobra.Command{
    Use: "on-call",
    Aliases: []string{"oc"},
    Short: "View on-call schedules",
    Long: `Various commands to view your on-call schedule or a team's on-call schedule`,
    Run: func(_ *cobra.Command, _ []string) {
      fmt.Println("ON CALLING IT!")
    },
  }

  return onCallCmd
}
