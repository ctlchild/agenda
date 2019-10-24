package cmd

import (
	"fmt"
	"strings"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
	"github.com/modood/table"
)

var listmeetingCmd = &cobra.Command{
	Use:   "listmeeting",
	Short: "list meeting notes can see the details of the meeting.",
	Long: `listmeeting: list meetings limited by start date and end date of current user.
	For example:
	agenda listmeeting -s 2018-10-25-14-20 -e 2018-10-25-14-20`,
	Run: func(cmd *cobra.Command, args []string) {
		runlist(li_st_date, li_ed_date)
	},
}
var li_st_date, li_ed_date string
var s_date, e_date entity.Date

func init() {
	rootCmd.AddCommand(listmeetingCmd)
	listmeetingCmd.Flags().StringVarP(&li_st_date, "start time", "s", "", "format yyyy-mm-dd-hh-mm")
	listmeetingCmd.Flags().StringVarP(&li_ed_date, "end time", "e", "", "format yyyy-mm-dd-hh-mm")
}
func runlist(st_date string, ed_date string) {
	logInit()
	defer logFile.Close()

	curUser := datarw.GetCurUser()
	if curUser == nil {
		logSave("List failed, not log in yet","[Error]")
		fmt.Println("List failed, please log in first")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd listmeeting called","[Info]")

	st_date_string := strings.Split(st_date,"-")
	ed_date_string := strings.Split(ed_date,"-")
	meetings := datarw.GetMeetings()
	res := make([]entity.Meeting, 0)

	if len(st_date_string)!=5 || len(ed_date_string) != 5 {
		logSave("Wrong date format","[Error]")
		fmt.Println("Wrong date format, it should be yyyy-mm-dd-hh-mm")
		logSave("Current User: " + curUser.Name + ", Cmd listmeeting failed","[Info]")
		fmt.Println("List meeting failed")
		return
	} else{
		s_date1, flag1 := entity.Convert(st_date_string)
		e_date1, flag2 := entity.Convert(ed_date_string)
		s_date = s_date1
		e_date = e_date1
		if !flag1 || !flag2 {
			logSave("Wrong date format","[Error]")
			fmt.Println("Wrong date format, it should be yyyy-mm-dd-hh-mm")
			logSave("Current User: " + curUser.Name + ", Cmd listmeeting failed","[Info]")
			fmt.Println("List meeting failed")
			return
		}
	}
	 
	for _,meeting := range meetings {
		if entity.Calc(s_date)<=entity.Calc(meeting.Startdate) && entity.Calc(meeting.Enddate) <= entity.Calc(e_date){
			for _,k := range meeting.Participators{
				if (k==curUser.Name){
					res=append(res,meeting)
					break;
				}
			}			
		}
	}

	if len(res)==0 {
		logSave("Meetings are not found","[Error]")
		fmt.Println("Meetings are not found")
		logSave("Current User: " + curUser.Name + ", Cmd listmeeting failed","[Info]")
		fmt.Println("List meeting failed")
		return
	}

	fmt.Println(table.Table(res))

	
	logSave("cmd: listuser success", "[Info]")
	fmt.Println("List meeting success!")

// 	for _, meeting := range meetings {
// 		if title_limited && meeting.Title != list_title { //has limitation on title but not satisfied
// 			continue
// 		}
// 		//-----------------------not satisfied date compare-------------------------------------------------------------------
// 		//no title limitation or has title limitation and already satisfied
// 		if start_limited && entity.Compare(sdate, meeting.Startdate) > 0 { //has limitation on start date but not satisfied
// 			continue
// 		}
// 		if end_limited && entity.Compare(edate, meeting.Enddate) < 0 { //after the given edate ,which is not  supposed
// 			continue
// 		}
// 		if usr_limited {

// 			if usr.Name == meeting.Sponsor {
// 				//we add this meeting to result
// 			} else {
// 				var f = false
// 				for _, parts := range meeting.Participators {
// 					if parts == usr.Name { //satisfied we can display it
// 						f = true
// 						break
// 					}
// 				}
// 				if f == false { // not satisfied we cannot display this meeting
// 					continue
// 				}
// 			}
// 		}
// 		//all request satisfied
// 		res = append(res, meeting)
// 	}
// 	DisplayMeeting(res)
// }

// func DisplayMeeting(mt []entity.Meeting) {

// 	standardMeetingLength := 12
// 	standardNameLength := 12

// 	//standardTimeLength := 16
// 	println("-----------------Display Meeting---------------------------")
// 	println("Title       Sponsor     Start Time\t\tEnd Time\t\tParticipators")
// 	for _, meeting := range mt {

// 		typed := 0
// 		print(meeting.Title)
// 		typed += len(meeting.Title)

// 		for {
// 			if typed >= standardMeetingLength {
// 				break
// 			}
// 			typed++
// 			print(" ")

// 		}
// 		//print("\t\t")
// 		last := typed
// 		print(meeting.Sponsor)
// 		typed += len(meeting.Sponsor)
// 		for {
// 			if typed >= last+standardNameLength {
// 				break
// 			}
// 			typed++
// 			print(" ")
// 		}
// 		//print("\n")
// 		sd := meeting.Startdate
// 		ed := meeting.Enddate
// 		year := sd.Year
// 		month := sd.Month
// 		day := sd.Day
// 		hour := sd.Hour
// 		minute := sd.Minute
// 		//var info []byte
// 		info := fmt.Sprintf("%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
// 		//fmt.sprintf(info, "%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
// 		print(info)
// 		print("\t")
// 		year = ed.Year
// 		month = ed.Month
// 		day = ed.Day
// 		hour = ed.Hour
// 		minute = ed.Minute

// 		info = fmt.Sprintf("%04d-%02d-%02d-%02d:%02d", year, month, day, hour, minute)
// 		print(info)
// 		print("\t")
// 		for _, p := range meeting.Participators {
// 			print(p)
// 			print("\t")
// 		}

// 		println()
// 	}
// 	println("-----------------------------------------------------------")
}
