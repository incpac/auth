// ----------------------------------------
// cmd/root.go
// ----------------------------------------
//
// Cobra entry point
//

package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
  Use:     "auth-couchbase",
  Short:   "",
  Long:    ``,

}


func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

