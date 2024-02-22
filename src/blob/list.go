// nexusutils
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/blob/list.go
// Original timestamp: 2024/02/21 17:17

package blob

import (
	"fmt"
	"nexusutils/env"
	"nexusutils/helpers"
)

func ListBlobStores() error {
	var e env.NXRMinfo
	var err error
	headers := map[string]string{
		"accept":     "application/json",
		"X-Nexus-UI": "true",
	}

	if e, err = env.LoadEnvironmentFile(); err != nil {
		return helpers.CustomError{fmt.Sprintf("Error reading the environment file: %s", err)}
	}
	APIurl := fmt.Sprintf("%s/service/rest/v1/blobstores")

	fmt.Printf("%s %s %s", e, headers, APIurl)

	return nil
}
