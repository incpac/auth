// ----------------------------------------
// cmd/reset.go
// ----------------------------------------
//
//  Resets a user's password. Note, this command is interactive
//  

package cmd

import (
  "bufio"
  "fmt"
  "os"

  "github.com/spf13/cobra"
)



// ----------------------------------------
// helper functions
// ----------------------------------------


// ----------------------------------------
// entry point
// ----------------------------------------

var resetPassword = func(cmd *cobra.Command, args []string) {
  // todo: witchcraft
  fmt.Println("reset called")

    
  if uID == "" {
    fmt.Println("missing user id")
    os.Exit(1)
  }
  
  connectToDatabase()
  
  user := getUser(uID)
  
  if user.ID == "" {
    fmt.Println("user does not exist")
    os.Exit(1)
  }
  
  reader    := bufio.NewReader(os.Stdin)
  password  := getPassword(reader)
  
  databaseBucket.Upsert("u:"+user.ID,
              User {
                Type:         "user",
                ID:           user.ID,
                Email:        user.Email,
                FirstName:    user.FirstName,
                LastName:     user.LastName,
                Password:     password,
                Permissions:  user.Permissions,
              }, 0)
  
  fmt.Println("updated password")
}


var resetCmd = &cobra.Command{
  Use:     "reset",
  Short:   "Resets a user's password",
  Long:    ``,
	
  Run: resetPassword,
}


func init() {
  RootCmd.AddCommand(resetCmd)

  resetCmd.Flags().StringVarP(&uID,    "userid", "u", "", "The user's id")
  resetCmd.Flags().StringVarP(&databaseURL, "database", "d", "couchbase://localhost/default", "URL for CouchBase database")
}
