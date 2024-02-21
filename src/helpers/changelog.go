// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// changelog.go, jfgratton : 2024-02-20

package helpers

import "fmt"

func ChangeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	CenterPrint("CHANGELOG")
	fmt.Println()
	CenterPrint("=========")
	fmt.Println()
	fmt.Println()

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
1.00.00			2024.22.19		initial version
`)
}
