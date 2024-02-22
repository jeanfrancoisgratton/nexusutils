// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// urlHandlerHelpers.go, jfgratton : 2024-02-21

package urlHandler

type PayloadReceiver interface {
	UnmarshalJSON([]byte) error
}
