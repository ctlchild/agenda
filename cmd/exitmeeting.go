package cmd

import (
	"fmt"

	"github.com/ctlchild/agenda/datarw"
	"github.com/spf13/cobra"
)
var exit_title string

var exitmeetingCmd = &cobra.Command{
	Use:   "exitmeeting",
	Short: "Exit from the meeting as a member.",
	Long: `exitmeeting : you must login first , if you are the sponser of the meeting,it'll be canceled without assertain.
	For example:
	agenda exitmeeting -t=title1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		runExit(exit_title)
	},
}

func init() {
	rootCmd.AddCommand(exitmeetingCmd)
	exitmeetingCmd.Flags().StringVarP(&exit_title, "title", "t", "empty title", "input the title of meeting you want to exit")
}

func runExit(title string) {
	logInit()
	defer logFile.Close()

	curUser := datarw.GetCurUser()
	if curUser == nil {
		logSave("Exit failed, not log in yet","[Error]")
		fmt.Println("Exit failed, please log in first")
		return
	}
	if title == "" {
		logSave("Exit failed, empty title","[Error]")
		fmt.Println("Exit failed, please input title")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd exitmeeting called","[Info]")

	meetings := datarw.GetMeetings()

	inMeeting := false
	meetingExist := false

	var tmp_pt []string
	for i, j := range meetings {
		if j.Title == title {
			meetingExist = true
			for p, k := range j.Participators {
				if k==curUser.Name {
					inMeeting=true
					tmp_pt=append(j.Participators[:p], j.Participators[p+1:]...)
					if (len(tmp_pt)==0){
						meetings=append(meetings[:i],meetings[i+1:]...)
					} else {
						meetings[i].Participators = tmp_pt
					}
					datarw.SaveMeetings(meetings)
					break
				}
			}
		}
	}

	if meetingExist == false {
		logSave("Exit failed, title is not found","[Error]")
		fmt.Println("Exit failed, the meeting is not existed")
		return
	}
	if inMeeting == false {
		logSave("Exit failed, user is not in the meeting","[Error]")
		fmt.Println("Exit failed, you are not in the meeting")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd exitmeeting success","[Info]")
	fmt.Println("Exit meeting success")
	return
}
