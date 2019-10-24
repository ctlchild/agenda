package cmd

import (
	"fmt"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "delete CurUser",
	Long: `deleteuser:delete CurUser and logout
	you must login before deleteuser
	For example:
	agenda deleteuser  			:delete CurUser and logout
	`,
	Run: func(cmd *cobra.Command, args []string) {

		deleteuser()
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

}

var curUser *entity.User

func deleteuser() {
	logInit()
	defer logFile.Close()

	curUser = datarw.GetCurUser()

	if curUser == nil { //是否已登陆
		fmt.Println("Pleas log in first")
		logSave("isn't login,please use command login", "[Error]")
		return
	} else {
		logSave("cmd: deleteuser called", "[Info]")
	}

	//获取所有用户
	users := datarw.GetUsers()

	for index, user := range users {
		if user.Name == curUser.Name {
			users = append(users[:index], users[index+1:]...)
			datarw.SaveUsers(users)
			datarw.SaveCurUser(nil) //登出

			/*会议相关*/

			cancleAllmeeting(*curUser) //当前用户取消所有其创建的会议
			exitAllmeeting(*curUser)   //当前用户退出所有会议

			fmt.Println("Delete user success")
			logSave("cmd: deleteuser success", "[Info]")
			return
		}
	}

	logSave("unexpected to execute", "[Error]")

}

//取消user创建的会议
func cancleAllmeeting(user entity.User) {
	meetings := datarw.GetMeetings()
	var newMeetings []entity.Meeting

	for _, meeting := range meetings { //遍历会议
		if meeting.Sponsor != user.Name {
			newMeetings = append(newMeetings, meeting)
		}
	}

	datarw.SaveMeetings(newMeetings)
}

//user退出所有会议
func exitAllmeeting(user entity.User) {
	meetings := datarw.GetMeetings()

	for k, meeting := range meetings { //遍历会议
		for i, participator := range meeting.Participators { //遍历一个会议的成员

			if participator == user.Name {
				meetings[k].Participators = append(meetings[k].Participators[:i], meetings[k].Participators[i+1:]...)

				break //!!!
			}
		}

	}

	meetings = entity.DeleteEmptyMeeting(meetings)

	datarw.SaveMeetings(meetings)

}
