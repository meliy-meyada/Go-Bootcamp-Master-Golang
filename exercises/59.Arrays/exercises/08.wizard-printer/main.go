// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	// create a multi-dimensional array with each scientist's name, lastname and nickname
	scientists := [][3]string{
			{"Albert", "Einstein", "time"},
      {"Isaac", "Newton", "apple"},
      {"Stephen", "Hawking", "blackhole"},
      {"Marie", "Curie", "radium"},
      {"Charles", "Darwin", "fittest"},

	}
	// print the table header
	fmt.Printf("%-15s %-15s %-15s\n", "First Name", "Last Name", "Nickname")
	fmt.Println("===========================================")

	// Print each scientist's information
	for _, s := range scientists{
		fmt.Printf("%-15s %-15s %-15s\n", s[0], s[1], s[2])
	}
}