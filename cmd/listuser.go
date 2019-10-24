package cmd

import (
	"fmt"

	"github.com/ctlchild/agenda/datarw"
	"github.com/modood/table"
	"github.com/spf13/cobra"
)


var listuserName string

var listuserCmd = &cobra.Command{
	Use:   "listuser",
	Short: "Show name,email,phone of users",
	Long: `listuser:Show name,email,phone of users you must login before list
	For example:
	agenda listuser             :show all registered users' information
	agenda listuser -n user1    :show user1' information if registered
	`,
	Run: func(cmd *cobra.Command, args []string) {
		listuser()
	},
}

func init() {
	rootCmd.AddCommand(listuserCmd)
	listuserCmd.Flags().StringVarP(&listuserName, "name", "n", "", "user's name")
}

func listuser() {
	logInit()
	defer logFile.Close()

	curUser = datarw.GetCurUser()

	if curUser == nil { //是否已登陆
		fmt.Println("You are not logged in, please log in first!")
		logSave("isn't login,please use command login", "[Error]")
		return
	} else {
		logSave("cmd: listuser called", "[Info]")
	}

	//获取所有用户
	users := datarw.GetUsers()

	if listuserName == "" { //查询所有用户（因为已登录，所以不可能没有用户）

		for i := range users { //掩盖密码
			users[i].Password = "********"

		}

		fmt.Println(table.Table(users))

	} else { //查询单个用户
		for _, user := range users {
			if user.Name == listuserName {
				fmt.Println("\t", user.Name, "\t", user.Email, "\t", user.Phone)
				return //查询成功
			}
		}
		fmt.Println(listuserName+"doesn't registered")
		logSave(listuserName+"doesn't registered", "[Warning]") //查询失败
	}

	fmt.Println("List user success!")
	logSave("cmd: listuser success", "[Info]")
}
