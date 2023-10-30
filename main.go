package main

// 1. Insert data
// 2. View all data
// 3. Get rank
// 4. Update score
// 5. Delete one record
// - Insert data - this needs to handle input data for the following structure and store in database:
// SAT Results
// - Name (Unique Identifier)
// - Address
// - City
// - Country
// - Pincode
// - SAT score
// - Passed - this needs to be calculated in the backend as follows - if SAT score > 30% = Pass else Fail
// - View all data - this should display all the data
// - Get rank - this takes the name and returns their rank
// - Update score - this allows to update SAT score for a candidate by name
// - Delete one record - this deletes a record by name

import (
	"fmt"

	"github.com/mannuR22/PrecizeGoLang.git/ui"
	// "manishrana.online/task/models"
)

func main() {

	for {
		choice := ui.Menu()
		exitNow := false
		var status int
		switch choice {
		case 1:
			status = ui.InsertData()
			break
		case 2:
			status = ui.ViewAllData()
			break
		case 3:
			status = ui.GetRank()
			break
		case 4:
			status = ui.UpdateScore()
			break
		case 5:
			status = ui.DeleteRecord()
			break
		default:
			status = 1
			exitNow = true
		}

		if status == 0 {
			fmt.Print("\nTRANSACTION INCOMPLETE\n")
		} else {
			fmt.Print("\nTRANSACTION COMPLETED\n")
		}
		if exitNow {
			break
		}
		var junk string
		fmt.Print("\nHit Enter to show menu...")
		fmt.Scanln(&junk)
	}
	fmt.Println("\nApplication Exit Success!\n")

}
