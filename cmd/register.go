package cmd

import (
	"fmt"
	"regexp"

	"github.com/ctlchild/agenda/datarw"
	"github.com/ctlchild/agenda/entity"

	"github.com/spf13/cobra"
)

//var cfgFile string
var Name, Password, Email, Phone string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new User",
	Long: `register:Users are registered through username, passwords, email and phone.
	For example:
	register a new user,with name:User1,password:12345678,email:abc@163.com,phone:13012345678
	agenda register -nUser1 -p12345678 -emailabc@163.com -t13012345678
	`,
	Run: func(cmd *cobra.Command, args []string) {
		register(Name, Password, Email, Phone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&Name, "name", "n", "", "user's name")
	registerCmd.Flags().StringVarP(&Password, "password", "p", "", "user's password")
	registerCmd.Flags().StringVarP(&Email, "email", "e", "", "user's email")
	registerCmd.Flags().StringVarP(&Phone, "phone", "t", "", "user's phone")

}

func register(name string, password string, email string, phone string) {
	logInit()
	defer logFile.Close()
	logSave("cmd: register called", "[Info]")

	if isValidName(name) && isValidPassword(password) && isValidEmail(email) && isValidPhone(phone){
		users := datarw.GetUsers()
		if entity.HasUser(name, users) {
			fmt.Println("Register fail, " + name + " has been registered")
			logSave("The username has been registered", "[Warning]")
			logSave("Register fail", "[Warning]")
			return
		}
		
		newuser := entity.User{Name: name, Password: password, Email: email, Phone: phone}
		users = append(users, newuser)
		datarw.SaveUsers(users)
		fmt.Println("Registration success!")
		logSave("Registration success", "[Info]")
		return
	}
	logSave("Register fail", "[Warning]")

}

func isValidName(n string) bool {
	b := []byte(n)
	val, _ := regexp.Match(".+", b)
	if !val {
		fmt.Println("Register fail, name is invaild")
		logSave("flag -n ,name is invaild", "[Warning]")
	}
	return val
}
func isValidPassword(p string) bool {
	b := []byte(p)
	val, _ := regexp.Match(".+", b)
	if len(p) < 8 {
		fmt.Println("The password must be longer than 8 digits")
		val = false
	}
	if !val {
		fmt.Println("Register fail, password is invaild")
		logSave("flag -p ,password is invaild", "[Warning]")
	}
	return val
}
func isValidEmail(e string) bool {
	b := []byte(e)
	val, _ := regexp.Match("\\w*@\\w*\\.w*", b)

	if !val {
		fmt.Println("Register fail, email is invaild")
		logSave("email is invaild", "[Warning]")
	}
	return val
}
func isValidPhone(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match("\\d{11}", b)

	if !val {
		fmt.Println("Register fail, phone is invaild")
		logSave("phone is invaild", "[Warning]")
	}
	return val
}
