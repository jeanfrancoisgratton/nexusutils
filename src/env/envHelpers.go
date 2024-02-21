// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// envHelpers.go, jfgratton : 2024-02-20

package env

var EnvConfigFile string

type NXRMinfo struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}
