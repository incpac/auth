// ----------------------------------------
// cmd/common.go
// ----------------------------------------
//
// Contains common items in the cmd package; variables, functions, structs, etc.
// If it's being declared at the package level the it should be placed here 
// rather than in a components individual file.
// 

package cmd

import (
  "bufio"
  "fmt"
  "strings"
  
  "gopkg.in/couchbase/gocb.v1"
  "golang.org/x/crypto/bcrypt"
)


// ----------------------------------------
// variables
// ----------------------------------------

// database connection 
var databaseURL     string
var databaseCluster *gocb.Cluster
var databaseBucket  *gocb.Bucket

// address http server listens on
var listenAddress   string

// jwt token key
var tSigningKey     []byte
var tSigningString  string

// user
var uID             string
var uEmail          string
var uUserName       string
var uFirstName      string
var uLastName       string
var uPassword       string
var uPermissions    []string



// ----------------------------------------
// structs
// ----------------------------------------

type User struct {
  Type          string    `json:"type"`
  ID            string    `json:"uid"`
  Email         string    `json:"email"`
  FirstName     string    `json:"firstname"`
  LastName      string    `json:"lastname"`
  Password      string    `json:"password"`
  Permissions   []string  `json:"permissions"`
}



// ----------------------------------------
// helpers
// ----------------------------------------

var getUser = func(userID string) User {
  
  var user User
  databaseBucket.Get("u:" + userID, &user)
  return user
}


var connectToDatabase = func() {
  
  s := strings.Split(databaseURL, "/")
  
  clusterName := s[0] + "//" + s[2] 
  bucketName  := s[3]
  
  databaseCluster, _  = gocb.Connect(clusterName)
  databaseBucket, _   = databaseCluster.OpenBucket(bucketName, "")
  
  databaseBucket.Manager("", "").CreatePrimaryIndex("", true, false)
}


var hashPassword = func(password string) (string, error) {
  
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}


var checkPassword = func(password string, hash string) (bool, error) {
  
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil, err
}


func getUserID(reader *bufio.Reader) string {
  
  fmt.Print("user id: ")
  id, _ := reader.ReadString('\n')
  
  id = strings.TrimSpace(id)
  
  // ensure id isn't blank
  if id == "" {
    fmt.Println("note: user id cannot be empty")
    return getUserID(reader)
  }
  
  
  // ensure id is already in use
  user := getUser(id)
  
  if user.ID != "" {
    fmt.Println("note: user id has to be unique")
    return getUserID(reader)
  }
  
  
  // we're all go
  return id
}


func getPassword(reader *bufio.Reader) string {
  
  fmt.Print("password: ")
  password, _ := reader.ReadString('\n')
  
  password = strings.TrimSpace(password)
  
  // ensure it's not blank
  if password == "" {
    fmt.Println("note: password cannot be empty")
    return getPassword(reader)
  }
    
  // get confirmation
  fmt.Print("confirm: ")
  confirm, _ := reader.ReadString('\n')
  
  confirm = strings.TrimSpace(confirm)
  
  // confirm they're the same
  if password != confirm {
    fmt.Println("does not match, please try again")
    return getPassword(reader)
  }

  // encrypt password 
  encrypted, _ := hashPassword(password)
  
  return encrypted
}
