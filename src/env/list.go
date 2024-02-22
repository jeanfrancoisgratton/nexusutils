// nexusutils
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/list.go
// Original timestamp: 2024/02/21 16:55

package env

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"nexusutils/helpers"
	"os"
	"path/filepath"
	"strings"
)

func ListEnvironments(envdir string) error {
	var err error
	var dirFH *os.File
	var finfo, fileInfos []os.FileInfo

	// list environment files
	if envdir == "" {
		envdir = filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nexusutils")
	}
	if dirFH, err = os.Open(envdir); err != nil {
		return helpers.CustomError{Message: "Unable to read config directory: " + err.Error()}
	}

	if fileInfos, err = dirFH.Readdir(0); err != nil {
		return helpers.CustomError{Message: "Unable to read files in config directory: " + err.Error()}
	}

	for _, info := range fileInfos {
		//if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") && !strings.HasPrefix(info.Name(), "sample") {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			finfo = append(finfo, info)
		}
	}

	if err != nil {
		return err
	}

	fmt.Printf("Found %s environment file(s) in %s\n", helpers.Green(fmt.Sprintf("%d", len(finfo))),
		helpers.Green(fmt.Sprintf("%s", filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nexusutils"))))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Environment file", "File size", "Modification time"})

	for _, fi := range finfo {
		t.AppendRow([]interface{}{helpers.Green(fi.Name()), helpers.Green(helpers.SI(uint64(fi.Size()))),
			helpers.Green(fmt.Sprintf("%v", fi.ModTime().Format("2006/01/02 15:04:05")))})
	}
	t.SortBy([]table.SortBy{
		{Name: "Environment file", Mode: table.Asc},
		{Name: "File size", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatDefault
	t.Render()

	return nil
}
