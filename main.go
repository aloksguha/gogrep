package main

import (
	"fmt"
	"github.com/aloksguha/gogrep/gogrep"
	"github.com/aloksguha/gogrep/utils"
	"os"
	"time"
)

func main() {
	//Init()
	argsWithProg := os.Args[1:]
	if len(argsWithProg) > 1 {
		fmt.Println("Invalid arguments, Please refer help by ")
	}
	if len(argsWithProg) == 0 {
		showInfo()
	}
	if len(argsWithProg) == 1 && argsWithProg[0] == "init" {
		Init()
	}
	if len(argsWithProg) == 1 && ( argsWithProg[0] == "help" || argsWithProg[0] == "-h") {
		showHelp()
	}

}


func Init(){
	fPath := ""
	fmt.Print(utils.InfoBlue("Please enter full path of file : "))
	if _, err := fmt.Scanf("%s", &fPath); err != nil  {
		fPath = ""
	}
	fmt.Println(utils.Input(fPath))

	q := ""
	fmt.Print(utils.InfoBlue("Please enter search string (default : kola) : "))
	if _, err := fmt.Scanf("%s", &q); err != nil  {
		q = "kola"
	}
	fmt.Println(utils.Input(q))

	timeout := 60
	fmt.Print(utils.InfoBlue("Please enter timeout in seconds (default : 60 Sec) : "))
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
	report, err := g.Search()
	if err != nil {
		fmt.Println(utils.Red(err))
		return
	}
	PrintResult(report)

}


func showInfo(){
	fmt.Println(utils.Info("gogrep by Alok Guha<aloksguha@gmail.com>"))
	fmt.Println(utils.CoTeal("Version V0.1, Built: 1.0 \n"))
	fmt.Println(utils.Info("Usage:"))
	fmt.Printf(utils.CoTeal("	gogrep init \n"))
	fmt.Println(utils.Info("For help :"))
	fmt.Printf(utils.CoTeal("	gogrep -h \n"))
	return
}

func showHelp(){
	fmt.Printf("To start program please type: 'gogrep init' and follow instructions, (use 'sudo' as prefix if accessing restricted area )\n")
}

func PrintResult(reports []gogrep.Report){
	totalByteCnt := 0
	var duration time.Duration = 0
	fmt.Println(utils.Info("\n\n:----------------------------------------: RESULT :---------------------------------------------------------:"))
	for i:=0; i< len(reports); i++ {
		report := reports[i]
		switch report.Status {
		case "SUCCESS":{
			fmt.Print(utils.Info("\t Status : ",report.Status))
			fmt.Print(utils.Info(" Bytes Read : ",report.ByteCnt))
			fmt.Print(utils.Info("\t Elapsed Time : ",report.Elapsed))
			fmt.Println(utils.Info("\t Remaining Time : ",report.Remaining))
			totalByteCnt +=report.ByteCnt
			duration += report.Elapsed
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
	if totalByteCnt > 0 {
		fmt.Println(utils.Info(":------------------------------------------------------------------------------------------------------:"))
		avgTime := int(time.Duration(duration)) / totalByteCnt
		fmt.Println("\tAverage grep time per byte : ",time.Duration(avgTime))
	}else{
		fmt.Println(utils.Red("\tUnable to find entered string"))
	}
	fmt.Println(utils.Info(":------------------------------------------------------------------------------------------------------:"))
}
