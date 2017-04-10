// ----------------------------------------
// cmd/serve.go
// ----------------------------------------
//
// Creates the application's http server. Requires all 
// parameters to be set which is why they all have defaults.
//  

package cmd

import (
  "encoding/json"
	"fmt"
  "net/http"
  "os"
  "strings"
  "time"
  
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
	"github.com/spf13/cobra"
  
  jwt "github.com/dgrijalva/jwt-go"
)



// ----------------------------------------
// routes
// ----------------------------------------

var VerifyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
  tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
    
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface {}, error ) {
    
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    
    return tSigningKey, nil
  })
  

  if token.Valid {
    
    // token is valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
      
      jsonClaims, _ := json.Marshal(claims)
      w.Write([]byte(jsonClaims))
    
    } else {
      
      // issue extracting claims
      w.WriteHeader(500)
      w.Write([]byte("server error"))
    
    }
    
  } else if ve, ok := err.(*jwt.ValidationError); ok {
    
    if ve.Errors&jwt.ValidationErrorMalformed != 0 {
    
      // token is malformed    
      w.WriteHeader(422)
      w.Write([]byte("token is malformed")) 
      
    } else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
      
      // token is expired
      w.WriteHeader(401)
      w.Write([]byte("token is expired"))  
      
    } else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
      
      // token not yet valid
      w.WriteHeader(401)
      w.Write([]byte("token is not yet valid")) 
      
    } else {
      
      // something else
      w.WriteHeader(422)
      w.Write([]byte("token is bad"))   
      
    }
  } else {
    
    // server error
    w.WriteHeader(500)
    w.Write([]byte("server error"))   
    
  }
})


var AuthenticateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
  // todo: autheticate user here
  username, password, _ := r.BasicAuth()
  
  connectToDatabase()
  user := getUser(username)
  
  if user.ID == "" {
    
    http.Error(w, "unauthorized", 401)
  
  } else if check, _ := checkPassword(password, user.Password); !check {
    
    http.Error(w, "unauthorized", 401)
    
  } else {
    
    // create the token
    token   := jwt.New(jwt.SigningMethodHS256)
    claims  := token.Claims.(jwt.MapClaims)

    // set token claims
    claims["id"]          = user.ID
    claims["email"]       = user.Email
    claims["firstname"]   = user.FirstName
    claims["lastname"]    = user.LastName
    claims["permissions"] = strings.Join(user.Permissions, ",")
    claims["exp"]         = time.Now().Add(time.Hour * 24).Unix()

    // sign the token with the signing key
    tokenString, _ := token.SignedString(tSigningKey)

    // write the token to the http stream
    w.Write([]byte(tokenString))     
  }
})



// ----------------------------------------
// entry point
// ----------------------------------------

var serve = func(cmd *cobra.Command, args []string) {
  
  tSigningKey = []byte(tSigningString)
  
  r := mux.NewRouter()
  
  r.Handle("/authenticate", AuthenticateHandler ).Methods("POST")
  r.Handle("/verify",       VerifyHandler       ).Methods("POST")
  
  http.ListenAndServe(listenAddress, handlers.LoggingHandler(os.Stdout, r))
}


var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Starts the application server",
	Long:    ``,
  
	Run: serve,
}


func init() {
	RootCmd.AddCommand(serveCmd)

  serveCmd.Flags().StringVarP(  &listenAddress,   "address",  "a",  "0.0.0.0:3000",                   "The address the server listends on")
  serveCmd.Flags().StringVarP(  &tSigningString,  "key",      "k",  "supersecret",                    "Key used to sign the JWT tokens")
  serveCmd.Flags().StringVarP(  &databaseURL,     "database", "d",  "couchbase://localhost/default",  "URL for the CouchBase database")
}
