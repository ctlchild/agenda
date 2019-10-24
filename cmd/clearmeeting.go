package cmd

import (
	"fmt"
	"strconv"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
)

var info_show bool

var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "Current user can clear meetings which he sponsors",
	Long: `clearmeeting:Current user can clear all meetings that he sponsors and see the details of the cleanup.
		For example:
		agenda clearmeeting -i 
		clear all meetings and print titles of meeting being deleted
		agenda clearmeeting
		clear all meetings but do not print`,
	Run: func(cmd *cobra.Command, args []string) {
		clearmeeting();
	},
}

func init() {
	rootCmd.AddCommand(clearmeetingCmd)
	clearmeetingCmd.Flags().BoolVarP(&info_show, "info", "i", false, "show meetings cleared")
}

func clearmeeting(){

	logInit()
	defer logFile.Close()

	var delete_meetings []string
	//get current user
	curUser := datarw.GetCurUser()
	if curUser == nil {
		logSave("Not log in yet","[Error]")
		fmt.Println("Clear failed, please log in first")
		logSave("Cmd clearmeeting failed","[Error]")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd clearmeeting called","[Info]")

	meetings := datarw.GetMeetings()

	var final_meetings []entity.Meeting
	for _, j := range meetings {
		if j.Sponsor == curUser.Name {
			delete_meetings = append(delete_meetings, j.Title)
		} else {
			final_meetings = append(final_meetings, j)
		}
	}
	datarw.SaveMeetings(final_meetings)
	if info_show {
		for i, j := range delete_meetings {
			fmt.Println("deletemeeting " + strconv.Itoa(i) + ": " + j)
		}
	}
	logSave("Current User: " + curUser.Name + ", Cmd clearmeeting success","[Info]")
	fmt.Println("clearmeeting success")
}
