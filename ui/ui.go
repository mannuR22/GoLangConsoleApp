package ui

import (
	"fmt"
	"strings"

	"github.com/mannuR22/PrecizeGoLang.git/dao"
	"github.com/mannuR22/PrecizeGoLang.git/models"
)

func printStatus(resp string, err error) {
	fmt.Println("Status:", resp)
	if err != nil {
		fmt.Println("Error: ", err.Error())

	}
}
func Menu() int {
	var choice int
	fmt.Println("\nChoose option 1-6 from below & press enter: \n")
	fmt.Println("1. Insert Data")
	fmt.Println("2. View All Data")
	fmt.Println("3. Get Rank")
	fmt.Println("4. Update Score")
	fmt.Println("5. Delete One Record")
	fmt.Println("6. Exit")

	for {
		fmt.Print("\nEnter you choice: ")
		fmt.Scanln(&choice)
		if 1 <= choice && choice <= 6 {
			break
		} else {
			fmt.Print("Wrong Input Expected from 1 to 6, Value Entered: ", choice, "\n")
		}
	}

	fmt.Print("\n")
	return choice
}

func InsertData() int {
	var report models.Report

	fmt.Print("Enter Name: ")
	fmt.Scanln(&report.Name)
	report.Name = strings.TrimSpace(report.Name)

	fmt.Print("Address: ")
	fmt.Scanln(&report.Address)
	report.Address = strings.TrimSpace(report.Address)

	fmt.Print("City: ")
	fmt.Scanln(&report.City)
	report.City = strings.TrimSpace(report.City)

	fmt.Print("Country: ")
	fmt.Scanln(&report.Country)
	report.Country = strings.TrimSpace(report.Country)

	fmt.Print("Pincode: ")
	fmt.Scanln(&report.Pincode)
	report.Pincode = strings.TrimSpace(report.Pincode)

	fmt.Print("SAT Score (%): ")
	fmt.Scanln(&report.SATScore)

	if report.SATScore > 30 {
		report.Passed = true
	} else {
		report.Passed = false
	}
	var choice int
	fmt.Print("Press 1 and hit enter to confirm transaction: ")
	fmt.Scanln(&choice)

	if choice != 1 {
		return 0
	}

	resp, err := dao.InsertRecord(report)

	printStatus(resp, err)

	if err != nil {
		return 0
	}

	return 1

}

func ViewAllData() int {
	reports, err := dao.GetAllRecords()
	if err != nil {
		printStatus("Unable to fetch Records", err)
		return 0
	}
	fmt.Println("All Records in database:-\n")

	for indx, report := range reports {

		var passStatus string
		if report.Passed {
			passStatus = "YES"
		} else {
			passStatus = "NO"
		}

		fmt.Println("S.No. :", indx+1)
		fmt.Println("Name :", report.Name)
		fmt.Println("Address : ", report.Address)
		fmt.Println("City :", report.City)
		fmt.Println("Country :", report.Country)
		fmt.Println("Pincode :", report.Pincode)
		fmt.Println("SAT Score (%):", report.SATScore)
		fmt.Println("Passed :", passStatus, "\n")

	}

	fmt.Println("\n---------END---------\n")

	return 1
}

func GetRank() int {
	reports, err := dao.GetAllRecords()
	if err != nil {
		printStatus("Unable to fetch Records", err)
		return 0
	}
	var name string
	fmt.Print("Type Name of User: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	reqReport, err := dao.GetRecordWithName(name)
	if err != nil {
		printStatus("Unable to fetch Record with Name: "+name, err)
		return 0
	}
	reportWithHighScoreCount := 0
	for _, report := range reports {
		if report.SATScore > reqReport.SATScore {
			reportWithHighScoreCount++
		}

	}

	fmt.Println("\nRank of", name, ":", reportWithHighScoreCount+1)

	return 1
}

func UpdateScore() int {

	var name string
	var score int
	fmt.Print("Type Name of User: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Print("Enter Updated SAT Score (%): ")
	fmt.Scanln(&score)

	resp, err := dao.UpdateRecord(name, score)
	if err != nil {
		printStatus(resp, err)
		return 0
	}

	return 1
}

func DeleteRecord() int {

	var name string
	fmt.Print("Type Name of User: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	resp, err := dao.DeleteRecord(name)
	if err != nil {
		printStatus(resp, err)
		return 0
	}

	return 1
}
