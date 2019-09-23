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
  "github.com/jdlubrano/pagerduty-cli/api_client/schedules/schedule_overrides"
  "github.com/jdlubrano/pagerduty-cli/api_client/users"
  "github.com/jdlubrano/pagerduty-cli/user_input"
)

type override struct {
  schedule *schedules.Schedule
  user *users.User
  startTime *time.Time
  endTime *time.Time
}

func (o *override) ScheduleId() string {
  return o.schedule.Id
}

func (o *override) StartTime() time.Time {
  return *o.startTime
}

func (o *override) EndTime() time.Time {
  return *o.endTime
}

func (o *override) UserId() string {
  return o.user.Id
}

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

      aSchedules, err := schedules.ListForQuery(client, scheduleName)

      if err != nil {
        fmt.Println(err)
        return
      }

      user, _ := users.GetMe(client)
      schedule := chooseSchedule(aSchedules)
      o := &override{schedule, &user, startTime, endTime}
      printOverride(o)

      if !user_input.Confirm("Does this look correct?") {
        return
      }

      createdOverride, err := schedule_overrides.CreateOverride(client, o)

      if err != nil {
        fmt.Println(err)
        return
      }

      fmt.Printf("Successfully created override %s\n", createdOverride.Id)
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

func printSchedules(aSchedules []schedules.Schedule) {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"", "Name"})
  table.SetRowLine(true)

  for i, schedule := range aSchedules {
    table.Append([]string{strconv.Itoa(i + 1), schedule.Name})
  }

  table.Render()
}

func printOverride(o *override) {
  startTime := o.StartTime().Format(user_input.TimeFormat)
  endTime := o.EndTime().Format(user_input.TimeFormat)

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Schedule", "User", "Start", "End"})
  table.Append([]string{o.schedule.Summary, o.user.Name, startTime, endTime})
  table.Render()
}
