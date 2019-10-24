package cmd

import (
	"fmt"
	"strings"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
)

var Sponsor_name, cr_title, cr_st_date, cr_ed_date, participators string

var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create a meeting",
	Long: `createmeeting:Current user can create a meeting. You should provide meeting title, start date and end date of this meeting
and all participators. 
	For example:
	agenda createmeeting -t=new_meeting -s=2019-10-23-13-42 -e=2019-10-23-15-42 -p=xxx-xxx-xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		createmeeting(cr_title, cr_st_date, cr_ed_date, participators)
	},
}

func init() {
	rootCmd.AddCommand(createmeetingCmd)
	createmeetingCmd.Flags().StringVarP(&cr_title, "title", "t", "", "meeting title")
	createmeetingCmd.Flags().StringVarP(&cr_st_date, "start", "s", "", "meeting start date")
	createmeetingCmd.Flags().StringVarP(&cr_ed_date, "end", "e", "", "meeting end date")
	createmeetingCmd.Flags().StringVarP(&participators, "part", "p", "", "meeting participators")
}

func createmeeting(title string, st_date string, ed_date string, participators string) {
	logInit()
	defer logFile.Close()

	curUser := datarw.GetCurUser()
	if curUser == nil {
		logSave("Create failed, not log in yet","[Error]")
		fmt.Println("Create failed, please log in first")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd createmeeting called","[Info]")
	
	Sponsor_name = curUser.Name
	meetings := datarw.GetMeetings()
	users := datarw.GetUsers()
	participators_string := strings.Split(participators, "-")
	st_date_string := strings.Split(st_date, "-")
	ed_date_string := strings.Split(ed_date, "-")

	var s_date, e_date entity.Date
	if len(st_date_string) != 5 || len(ed_date_string) != 5 {
		logSave("Wrong date format","[Error]")
		fmt.Println("Wrong date format, it should be yyyy-mm-dd-hh-mm")
		logSave("Current User: " + curUser.Name + ", Cmd createmeeting failed","[Info]")
		fmt.Println("Create meeting failed")
		return
	} else {
		s_date1, flag1 := entity.Convert(st_date_string)
		e_date1, flag2 := entity.Convert(ed_date_string)
		s_date = s_date1
		e_date = e_date1
		if !flag1 || !flag2 {
			logSave("Wrong date format","[Error]")
			fmt.Println("Wrong date format. Should be Year-Month-Day-Hour-Minute")
			logSave("Current User: " + curUser.Name + ", Cmd createmeeting failed","[Error]")
			fmt.Println("Create meeting failed")
			return
		}
	}
	var temp_meeting entity.Meeting
	temp_meeting.Startdate = s_date
	temp_meeting.Enddate = e_date
	if !entity.IsParticipatorAvailable(curUser.Name, meetings, temp_meeting) {
		logSave("Sponsor is not free","[Error]")
		fmt.Println("Sponsor is not free")
		logSave("Current User: " + curUser.Name + ", Cmd createmeeting failed","[Error]")
		fmt.Println("Create meeting failed")
		return
	}

	valid_participators, ok := entity.Check_participators(Sponsor_name, participators_string, users, meetings, s_date, e_date)
	if entity.Check_title(title, meetings) && entity.Check_date(s_date, e_date) && ok {
		var new_meeting entity.Meeting
		new_meeting.Sponsor = curUser.Name
		new_meeting.Title = title
		new_meeting.Startdate = s_date
		new_meeting.Enddate = e_date
		new_meeting.Participators = valid_participators
		meetings = append(meetings, new_meeting)
		datarw.SaveMeetings(meetings)
		logSave("Current User: " + curUser.Name + ", Cmd createmeeting success","[Info]")
		fmt.Println("Create meeting success")
	} else {
		if !entity.Check_title(title, meetings) {
			logSave("Meeting has exists","[Error]")
			fmt.Println("Meeting has exists, please change meeting title")
		}
		if !entity.Check_date(s_date, e_date) {
			logSave("Invalid start/end date","[Error]")
			fmt.Println("Invalid start/end date, please check")
		}
		if !ok {
			logSave("No valid participators","[Error]")
			fmt.Println("No valid participators (busy or not exists)")
		}
		logSave("Current User: " + curUser.Name + ", Cmd createmeeting failed","[Error]")
		fmt.Println("Create meeting failed")
	}

}
