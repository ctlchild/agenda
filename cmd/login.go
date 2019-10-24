package cmd

import (
	"fmt"
	"github.com/ctlchild/agenda/datarw"
	"github.com/spf13/cobra"
)

var username,password string;

var loginCmd = &cobra.Command{
	Use:   "login -n [username] -p [password]",
	Short: "Log in",
	Long: `login : Login with username and password.
	For example: 
	agenda login -n username -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		login(username,password)
	},
		
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "n", "", "user's name")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "user's password")
}

func login(username string, password string){

	logInit()
	defer logFile.Close()
	logSave("Cmd login called","[info]")

	curUser := datarw.GetCurUser()
	if curUser != nil {
		fmt.Println("Log in failed, please log out first")
		logSave(username + " log in failed," + curUser.Name + " is already logged in","[Error]")
		return
	}

	users := datarw.GetUsers()
	for i := 0; i < len(users); i++ {
		if users[i].Name == username {
			if users[i].Password == password {
				datarw.SaveCurUser(&users[i])
				fmt.Println("Log in success")
				logSave(username + " log in success","[Info]")
				return
			}
			fmt.Println("Log in failed, password error")
			logSave(username + " log in failed, password error ","[Error]")
			return
		}
	}
	fmt.Println("Log in failed, username don't exist ")
	logSave("Log in failed, " + username +" don't exist ","[Error]")
}
