package cmd

import (
	"fmt"
	"github.com/ctlchild/agenda/datarw"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "User log out",
	Long: `logout：Users logged out and could not operate after landing.
	For example: 
	agenda logout`,
	Run: func(cmd *cobra.Command, args []string) {
		logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logout(){
	//log
	logInit()
	defer logFile.Close()
	logSave("Cmd logout called","[Info]")

	//确定当前是登陆状态
	curUser := datarw.GetCurUser()
	if curUser == nil {
		fmt.Println("Please log in first!")
		logSave("User is not logged in","[Error]")
		return
	}
	datarw.SaveCurUser(nil)
	//登出
	fmt.Println("Log out success")
	logSave(curUser.Name + " logout success","[Info]")
}

