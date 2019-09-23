package user_input

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
)

const TimeFormat = "01-02-2006 15:04"

func Ask(prompt string) string {
  fmt.Print(prompt + ": ")
  reader := bufio.NewReader(os.Stdin)
  answer, _ := reader.ReadString('\n')

  return strings.TrimSpace(answer)
}

func Confirm(prompt string) bool {
  answer := strings.ToLower(Ask(fmt.Sprintf("%s (y/n)", prompt)))

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
  parsedTime, err := time.ParseInLocation(TimeFormat, input, localLoc)

  if err != nil {
    return nil, err
  }

  return &parsedTime, nil
}
