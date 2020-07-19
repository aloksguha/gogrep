package main

import (
	"fmt"
	"github.com/aloksguha/gogrep/gogrep"
	"github.com/aloksguha/gogrep/utils"
	"os"
)

func main() {
	Init()
	//argsWithProg := os.Args[1:]
	//if len(argsWithProg) > 1 {
	//	fmt.Println("Invalid arguments, Please refer help by ")
	//}
	//if len(argsWithProg) == 0 {
	//	showInfo()
	//}
	//if len(argsWithProg) == 1 && argsWithProg[1] == "init" {
	//	Init()
	//}
	//if len(argsWithProg) == 1 && ( argsWithProg[1] == "help" || argsWithProg[1] == "-h") {
	//	showHelp()
	//}

}


func Init(){
	fPath := ""
	fmt.Print(utils.InfoBlue("Please enter full path of file : (default : test_files/text.file) : "))
	if _, err := fmt.Scanf("%s", &fPath); err != nil  {
		fPath = "test_files/text.file"
	}
	fmt.Println(utils.Input(fPath))

	q := ""
	fmt.Print(utils.InfoBlue("Please enter search string (default : Lpfn) : "))
	if _, err := fmt.Scanf("%s", &q); err != nil  {
		q = "Lpfn"
	}
	fmt.Println(utils.Input(q))

	timeout := 60
	fmt.Print(utils.InfoBlue("Please enter timeout (default : 60 Sec) : "))
	if n, err := fmt.Scanf("%d", &timeout); err != nil || n != 1 {
		timeout = 60
	}

	fmt.Println(utils.Input(timeout))

	now := 10
	fmt.Print(utils.InfoBlue("Please enter no. of workers (default : 10) : "))
	if n, err := fmt.Scanf("%d", &now); err != nil || n != 1 {
		now = 10
	}
	fmt.Println(utils.Input(now))
	g:= gogrep.NewSearch(fPath,timeout,q, now)
	PrintResult(g.Search())

}


func showInfo(){
	fmt.Printf(utils.Info("%s by Alok Guha\n", os.Args[0]))
	fmt.Printf(utils.Info("Version %s, Built: %s \n", "V 0.1", "0.1"))
	fmt.Println(utils.Info("Usage:"))
	fmt.Printf(utils.Info("	gogrep init \n"))
	fmt.Printf(utils.Info("For help : 	gogrep -h \n"))
	return
}

func showHelp(){
	fmt.Printf("To start program please type: gogrep init and follow instructions\n")
}

func PrintResult(reports []gogrep.Report){
	fmt.Println(utils.Info(":-------------------------------: RESULT :------------------------------------:"))
	for i:=0; i< len(reports); i++ {
		report := reports[i]
		switch report.Status {
		case "SUCCESS":{
			fmt.Print(utils.Info("\t Status : ",report.Status))
			fmt.Print(utils.Info(" Bytes Read : ",report.ByteCnt))
			fmt.Print(utils.Info("\t Elapsed Time : ",report.Elapsed))
			fmt.Println(utils.Info("\t Remaining Time : ",report.Remaining))
		}
		case "FAILURE":{
			fmt.Print(utils.InfoBlue("\t Status : ",report.Status))
			fmt.Print(utils.InfoBlue(" Bytes Read : ",report.ByteCnt))
			fmt.Print(utils.InfoBlue("\t Elapsed Time : ",report.Elapsed))
			fmt.Println(utils.InfoBlue("\t Remaining Time : ",report.Remaining))
		}
		case "TIMEOUT":{
			fmt.Print(utils.Red("\t Status : ",report.Status))
			fmt.Print(utils.Red(" Bytes Read : ",report.ByteCnt))
			fmt.Print(utils.Red("\t Elapsed Time : ",report.Elapsed))
			fmt.Println(utils.Red("\t Remaining Time : ",report.Remaining))
		}
		}
	}
	fmt.Println(utils.Info(":-----------------------------------------------------------------------------:"))
}
