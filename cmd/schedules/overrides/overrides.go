package schedule_overrides

import (
  "errors"
  "fmt"
  "os"
  "strconv"
  "time"

  "github.com/spf13/cobra"
  "github.com/olekukonko/tablewriter"

  "github.com/jdlubrano/pagerduty-cli/api_client"
  "github.com/jdlubrano/pagerduty-cli/api_client/schedules"
  "github.com/jdlubrano/pagerduty-cli/user_input"
)

func NewOverridesCmd(client *api_client.ApiClient) *cobra.Command {
  overridesCmd := &cobra.Command{
    Use: "overrides",
    Short: "Create schedule overrides",
    Long: `Various commands to adjust a schedule's overrides`,
  }

  overridesCmd.AddCommand(NewCreateCmd(client))

  return overridesCmd
}

func NewCreateCmd(client *api_client.ApiClient) *cobra.Command {
  var scheduleName, startTimeInput, endTimeInput string

  createCmd := &cobra.Command{
    Use: "create",
    Short: "Create a schedule override for a given schedule",
    Run: func(_ *cobra.Command, _ []string) {
      startTime, endTime, err := parseOverrideTimeSpan(startTimeInput, endTimeInput)

      if err != nil {
        fmt.Println(err)
        return
      }

      fmt.Printf("Start time: %s, End time: %s\n", startTime.String(), endTime.String())

      aSchedules, err := schedules.ListForQuery(client, scheduleName)

      if err != nil {
        fmt.Println(err)
        return
      }

      schedule := chooseSchedule(aSchedules)

      if !user_input.Confirm("Does this look correct?") {
        return
      }

      // createOverride
    },
  }

  createCmd.Flags().StringVarP(&scheduleName, "schedule", "s", "", "Schedule name (required)")
  createCmd.MarkFlagRequired("schedule")

  createCmd.Flags().StringVarP(&startTimeInput, "start", "", "", "Start time (MM-DD-YYYY HH:MM) (required)")
  createCmd.MarkFlagRequired("start")

  createCmd.Flags().StringVarP(&endTimeInput, "end", "", "", "End time (MM-DD-YYYY HH:MM) (required)")
  createCmd.MarkFlagRequired("end")

  return createCmd
}

func chooseSchedule(aSchedules []schedules.Schedule) (*schedules.Schedule) {
  if len(aSchedules) == 1 {
    return &aSchedules[0]
  }

  printSchedules(aSchedules)

  var chosenSchedule *schedules.Schedule

  for chosenSchedule == nil {
    scheduleChoice := user_input.Ask("Choose a schedule to override")

    if i, err := strconv.Atoi(scheduleChoice); err != nil || (i < 1 || i > len(aSchedules)) {
      fmt.Printf("Invalid input, expected an integer [%d-%d]\n", 1, len(aSchedules))
    } else {
      chosenSchedule = &aSchedules[i-1]
    }
  }

  return chosenSchedule
}

func parseOverrideTimeSpan(startTimeInput string, endTimeInput string) (*time.Time, *time.Time, error) {
  startTime, err := user_input.ParseTimeInput(startTimeInput)

  if err != nil {
    return nil, nil, err
  }

  minimumStartTime := time.Now().Truncate(time.Minute) // beginning of minute

  if startTime.Before(minimumStartTime) {
    return nil, nil, errors.New("Start time cannot be in the past")
  }

  endTime, err := user_input.ParseTimeInput(endTimeInput)

  if err != nil {
    return nil, nil, err
  }

  if !endTime.After(*startTime) {
    return nil, nil, errors.New("End time must be after start time")
  }

  return startTime, endTime, nil
}

func chooseEndTime(startTime time.Time) *time.Time {
  endTime := user_input.AskForTime("Choose an end time")

  if endTime.Before(startTime) {
    fmt.Println("End time must be after start time")
    return chooseEndTime(startTime)
  }

  return endTime
}

func chooseStartTime() *time.Time {
  minimumStartTime := time.Now().Truncate(time.Minute) // beginning of minute
  startTime := user_input.AskForTime("Choose a starting time")

  if startTime.Before(minimumStartTime) {
    fmt.Println("Could not use provided start time")
    return chooseStartTime()
  }

  return startTime
}

func printSchedules(aSchedules []schedules.Schedule) {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"", "Name"})
  table.SetRowLine(true)

  for i, schedule := range aSchedules {
    table.Append([]string{strconv.Itoa(i + 1), schedule.Name})
  }

  table.Render()
}
