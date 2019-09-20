package user_input

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
)

func Ask(prompt string) string {
  fmt.Print(prompt + ": ")
  reader := bufio.NewReader(os.Stdin)
  answer, _ := reader.ReadString('\n')

  return strings.TrimSpace(answer)
}

func AskForTime(prompt string) *time.Time {
  userTimezone, _ := time.Now().Local().Zone()
  timeInput := Ask(fmt.Sprintf("%s [MM-DD-YYYY HH:MM] (using %s timezone)", prompt, userTimezone))
  t, err := ParseTimeInput(timeInput)

  if err != nil {
    fmt.Printf("Unable to parse provided time: %s\n", timeInput)
    return AskForTime(prompt)
  }

  return t
}

func Confirm(prompt string) bool {
  answer = Ask(fmt.Sprintf("%s (y/n)", prompt)).ToLower()

  if answer == "y" || answer == "yes" {
    return true
  }

  if answer == "n" || answer == "no" {
    return false
  }

  fmt.Println("Invalid confirmation.  Use 'y' for yes and 'n' for no.")
  return Confirm(prompt)
}

func ParseTimeInput(input string) (*time.Time, error) {
  localLoc, _ := time.LoadLocation("Local")
  parsedTime, err := time.ParseInLocation("01-02-2006 15:04", input, localLoc)

  if err != nil {
    return nil, err
  }

  return &parsedTime, nil
}
