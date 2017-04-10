// ----------------------------------------
// cmd/create.go
// ----------------------------------------
//
//  CLI command to create a new user. Note, this command is
//  interactive
//  

package cmd

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/spf13/cobra"
)


// ----------------------------------------
// helper functions
// ----------------------------------------

func getPermissions(permissions string) []string {
  
  return strings.Split(permissions, ",")
}


// ----------------------------------------
// entry point
// ----------------------------------------

var createUser = func(cmd *cobra.Command, args []string) {  
    
  connectToDatabase()
  
  reader := bufio.NewReader(os.Stdin)
  
  id := getUserID(reader)
  
  fmt.Print("email: ")
  email, _ := reader.ReadString('\n')
  
  fmt.Print("firstname: ")
  firstname, _ := reader.ReadString('\n')
  
  fmt.Print("lastname: ")
  lastname, _ := reader.ReadString('\n')
  
  password := getPassword(reader)
  
  fmt.Print("permissions: ")
  permissions, _ := reader.ReadString('\n')
  
  
  id          = strings.TrimSpace(id)
  email       = strings.TrimSpace(email)
  firstname   = strings.TrimSpace(firstname)
  lastname    = strings.TrimSpace(lastname)
  password    = strings.TrimSpace(password)
  permissions = strings.TrimSpace(permissions)
  
  perms := getPermissions(permissions)
  
  databaseBucket.Insert("u:"+id,
                        User {
                          Type:         "user",
                          ID:           id,
                          Email:        email,
                          FirstName:    firstname,
                          LastName:     lastname,
                          Password:     password,
                          Permissions:  perms,
                        }, 0)
  
  fmt.Println("user created\n--------------------")
  
  fmt.Println("id: \t\t", id)
  fmt.Println("email: \t\t", email)
  fmt.Println("firstname: \t", firstname)
  fmt.Println("lastname: \t", lastname)
  fmt.Println("password: \t", password)
  fmt.Println("permissions: \t", permissions)
  
}

var createCmd = &cobra.Command{
  Use:     "create",
  Short:   "Creates a new user in the database",
  Long:    ``,
  
  Run: createUser,
}

func init() {
  RootCmd.AddCommand(createCmd)

  createCmd.Flags().StringVarP(&databaseURL, "database", "d", "couchbase://localhost/default", "URL for the CouchBase database")
}
