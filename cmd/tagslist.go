// Copyright © 2016 EMC Corporation
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// tagslistCmd represents the tagslist command
var tagslistCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags in RackHD",
	Long:  "List tags in RackHD",
	Run:   listTags,
}

func init() {
	tagsCmd.AddCommand(tagslistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagslistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagslistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func listTags(cmd *cobra.Command, args []string) {
	resp, err := clients.rackMonorailClient.Tags.GetTags(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})

	for _, tag := range resp.Payload {
		table.Append([]string{*(tag.Name)})
		//fmt.Printf("%#v\n\n", tag)
	}
	table.Render()
}
