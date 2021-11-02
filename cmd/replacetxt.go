/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var Path string
var Search string
var OldTxt string
var NewTxt string

// replacetxtCmd represents the replacetxt command
var replacetxtCmd = &cobra.Command{
	Use:   "replacetxt",
	Short: "Replace text in files",
	Long:  `Replace selected text in any files in the system with a set extension or name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replacetxt called")
		fmt.Println(Path)
		fmt.Println(Search)
		fmt.Println(OldTxt)
		fmt.Println(NewTxt)
		files, err := WalkMatch(Path, Search)
		if err != nil {
			fmt.Println(err)
		}

		CheckFiles(files, OldTxt, NewTxt)
	},
}

func init() {
	rootCmd.AddCommand(replacetxtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replacetxtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replacetxtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	replacetxtCmd.Flags().StringVarP(&Path, "path", "p", "C:\\Users", "The path the system should walk")
	replacetxtCmd.Flags().StringVarP(&Search, "search", "s", "", "The search that should be used when selecting files for instance *.rdp will select just rdp files")
	replacetxtCmd.Flags().StringVarP(&OldTxt, "old", "o", "", "The old text to replace")
	replacetxtCmd.Flags().StringVarP(&NewTxt, "new", "n", "", "The text that old should be replaced with")
	replacetxtCmd.MarkFlagRequired("search")
	replacetxtCmd.MarkFlagRequired("old")
	replacetxtCmd.MarkFlagRequired("new")
}

func CheckFiles(files []string, oldStr string, newStr string) {
	for _, elm := range files {
		input, err := ioutil.ReadFile(elm)
		if err != nil {
			fmt.Println("Failed to read from " + elm + " Aborting and skipping to next")
			continue
		}
		var inStr = string(input)
		if strings.Contains(inStr, oldStr) {
			fmt.Println("File found: " + elm)
			out := strings.ReplaceAll(inStr, oldStr, newStr)
			ioutil.WriteFile(elm, []byte(out), 0)
		}
	}
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
