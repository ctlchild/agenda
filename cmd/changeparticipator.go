package cmd

import (
	"fmt"
	"strings"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"
	"github.com/spf13/cobra"
)

var cg_title,name string
var add_flag, delete_flag bool


// changeparticipatorCmd represents the changeparticipator command
var changeparticipatorCmd = &cobra.Command{
	Use:   "changeparticipator",
	Short: "Current user can change participators of a meeting",
	Long: `changeparticipator:Current user can change participators of a meeting he sponser. 
	Adding process will check whether participators having free time for this meeting. 
	It will be deleted if a meeting has no participators after this cmd. 
	For exanple:
	agenda changeparticipator -t title1 -d -p a-b-c
	delete participators a,b,c in meeting title1
	agenda changeparticipator -t title1 -a -p a-b-c
	add participators a,b,c in meeting title1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		change(cg_title, add_flag, delete_flag, name)
	},
}

func init() {
	rootCmd.AddCommand(changeparticipatorCmd)
	changeparticipatorCmd.Flags().StringVarP(&cg_title, "title", "t", "", "meeting title")
	changeparticipatorCmd.Flags().BoolVarP(&add_flag, "add", "a", true, "add participator")
	changeparticipatorCmd.Flags().BoolVarP(&delete_flag, "delete", "d", false, "delete participator")
	changeparticipatorCmd.Flags().StringVarP(&name, "name", "p", "", "participator's name")
}

func change(title string, add_flag bool, delete_flag bool, name string){
	logInit()
	defer logFile.Close()

	curUser := datarw.GetCurUser()
	if curUser == nil {
		logSave("Change failed, not log in yet","[Error]")
		fmt.Println("Chagen failed, please log in first")
		return
	}
	logSave("Current User: " + curUser.Name + ", Cmd changeparticipator called","[Info]")

	change_participators := strings.Split(name, "-") //change name list
	meetings := datarw.GetMeetings()                              //all meetings
	meeting_exist := false
	var final_participators []string
	if delete_flag {

		var delete_participators []string
		for i, j := range meetings {
			if j.Sponsor == curUser.Name && j.Title == title {
				meeting_exist = true
				for _, k := range j.Participators {
					if entity.IsParticipatorinList(k, change_participators) {
						delete_participators = append(delete_participators, k)
					} else {
						final_participators = append(final_participators, k)
					}
				}

				if len(final_participators) == 0 {
					meetings = append(meetings[:i], meetings[i+1:]...)
				} else {
					meetings[i].Participators = final_participators
				}
				datarw.SaveMeetings(meetings)
				break
			}
		}
		if !meeting_exist {
			logSave("Change failed, title is not found","[Error]")
			fmt.Println("Change failed, title is not found")
			return
		} else if len(delete_participators) != len(change_participators) {
			logSave("Some users don't exist in this meeting","[Warning]")
			fmt.Println("Some users don't exist in this meeting. Already deleted: ")
			for _, j := range delete_participators {
				fmt.Println(j)
			}
		}
	} else {
		var valid_participators []string
		all_users := datarw.GetUsers()
		for _, j := range change_participators {
			if !entity.IsParticipatorExist(j, all_users) {
				fmt.Println(j + " is not a valid user")
			} else {
				valid_participators = append(valid_participators, j)
			}
		}
		if len(valid_participators) > 0 {
			for i, j := range meetings {
				if j.Sponsor == curUser.Name && j.Title == title {
					final_participators = j.Participators
					meeting_exist = true
					for _, k := range valid_participators {
						if entity.IsParticipatorExistinMeeting(k, j) {
							fmt.Println(k + " is already in this meeting")
						} else {
							meetings[i].Participators = append(meetings[i].Participators, k)
						}
					}
					datarw.SaveMeetings(meetings)
					break
				}
			}
			if !meeting_exist {
				logSave("Change failed, title is not found","[Error]")
				fmt.Println("Change failed, title is not found")
				return
			}
		}
	}
	logSave("Current User: " + curUser.Name + ", Cmd changeparticipator success","[Info]")
	fmt.Println("Change participator success")
}
