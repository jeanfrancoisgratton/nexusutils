// nexusutils
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/blob/blobHelpers.go
// Original timestamp: 2024/02/21 17:56

package blob

type BlobStore struct {
	Name                 string `json:"name"`
	SoftQuota            string `json:"softquota"`
	Unavailable          string `json:"unavailable"`
	BlobCount            string `json:"blobcount"`
	TotalSizeInBytes     string `json:"totalsizeinbytes"`
	AvailableSizeInBytes string `json:"availablesizeinbytes"`
}
