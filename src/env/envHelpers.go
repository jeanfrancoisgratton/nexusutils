// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// envHelpers.go, jfgratton : 2024-02-20

package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var EnvConfigFile string

type NXRMinfo struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func getStringVal(prompt string) string {
	fmt.Print(prompt)
	inputVal := bufio.NewReader(os.Stdin)
	input, _ := inputVal.ReadString('\n')

	return strings.TrimSpace(input)
}
