package cmd

import (
	"fmt"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
)

// cancelmeetingCmd represents the cancelmeeting command
var cancelmeetingCmd = &cobra.Command{
	Use:   "cancelmeeting",
	Short: "Cancellation of a meeting will delete the meeting.",
	Long: `cancelmeeting:If the founder of the conference can cancel the meeting through the title, the meeting will be deleted.
	For example:
	agenda cancelmeeting -t title1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cancelRun(cancel_title)
	},
}
var cancel_title string

func init() {
	rootCmd.AddCommand(cancelmeetingCmd)

	cancelmeetingCmd.Flags().StringVarP(&cancel_title, "title", "t", "", "title of the meeting")

}
func cancelRun(title string) {
	logInit()
	defer logFile.Close()
	//load
	curUser := datarw.GetCurUser()
	if curUser==nil{
		logSave("Cancel failed, not log in yet","[Error]")
		fmt.Println("Cancel failed, please log in first")
		return
	}
	if title == "" {
		logSave("Cancel failed, empty title","[Error]")
		fmt.Println("Cancel failed, please input title")
		return
	}

	logSave("Current User: " + curUser.Name + ", Cmd cancelmeeting called","[Info]")
	
	meetings := datarw.GetMeetings()
	res := make([]entity.Meeting, 0) //for writing

	//cancel
	var meetingExist = false

	for _, meeting := range meetings {
		if meeting.Title != cancel_title {
			res = append(res, meeting)
		} else {
			meetingExist = true
		}
	}
	if !meetingExist {
		logSave("Current User: " + curUser.Name + ", Cmd cancelmeeting failed","[Warning]")
		fmt.Println("Cancel meeting failed, meeting is not found")
		return
	}
	datarw.SaveMeetings(res)
	logSave("Current User: " + curUser.Name + ", Cmd cancelmeeting success","[Info]")
	fmt.Println("Cancel meeting seccess")
}
