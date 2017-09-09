package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	pageFlag              = "[PAGE TITLE]"
	procFlag              = "[PROCESS]"
	errFlag               = "[ERROR]"
	errUnableToParse      = errors.New("unable to parse settings.json")
	errMissingCredentials = errors.New("missing username or password from settings.json")
)

// Settings the actual struct that represents the settings.json
type Settings struct {
	Account  Account  `json:"account"`
	Messages Messages `json:"messages"`
}

// Account the actual struct that contains the username and password
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Messages the actual struct that contains the process of message deletion
type Messages struct {
	Excluded []string `json:"excluded"`
}

// ParseSettings will parse and return all values from settings.json file
func ParseSettings() (res Settings, err error) {
	raw, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		return res, err
	}
	json.Unmarshal(raw, &res)

	if res.Account.Username == "" || res.Account.Password == "" {
		return res, errMissingCredentials
	}

	return res, nil
}

// CreateDirIfNotExist will ensure that the app will always create the specified dir
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// FindString will return a boolean as a result
func FindString(needle string, haystack []string) (res bool) {
	res = false
	for _, hay := range haystack {
		if needle == hay {
			return true
		}
	}
	return res
}

// SigOpening promotes my name!
func SigOpeming() {
	sigText := `
 #######   ######   #######  ########  #### ##    ##  ######   ##     ##    ###    ##     ## ########    ###    ########   #######  
##     ## ##    ## ##     ## ##     ##  ##  ###   ## ##    ##  ##     ##   ## ##   ###   ###    ##      ## ##   ##     ## ##     ## 
## ### ## ##       ##     ## ##     ##  ##  ####  ## ##        ##     ##  ##   ##  #### ####    ##     ##   ##  ##     ## ##     ## 
## ### ## ##       ##     ## ##     ##  ##  ## ## ## ##   #### ######### ##     ## ## ### ##    ##    ##     ## ########  ##     ## 
## #####  ##       ##     ## ##     ##  ##  ##  #### ##    ##  ##     ## ######### ##     ##    ##    ######### ##   ##   ##     ## 
##        ##    ## ##     ## ##     ##  ##  ##   ### ##    ##  ##     ## ##     ## ##     ##    ##    ##     ## ##    ##  ##     ## 
 #######   ######   #######  ########  #### ##    ##  ######   ##     ## ##     ## ##     ##    ##    ##     ## ##     ##  #######  
`
	fmt.Println("Created By:")
	fmt.Println(sigText)
}
