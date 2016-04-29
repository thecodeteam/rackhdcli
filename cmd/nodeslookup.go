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
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/emccode/gorackhd/client/lookups"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// nodeslookupCmd represents the nodeslookup command
var nodeslookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Get node lookup details",
	Long:  "Get node lookup details",
	Run:   lookupNode,
}

func init() {
	nodesCmd.AddCommand(nodeslookupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeslookupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeslookupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nodeslookupCmd.Flags().StringVar(&targetNode, "node", "", "NODE id")
	nodeslookupCmd.Flags().BoolVarP(&shortList, "quiet", "q", false, "list only Node IDs")

}

func lookupNode(cmd *cobra.Command, args []string) {
	// do a lookup on the ID
	resp, err := clients.rackMonorailClient.Lookups.GetLookups(&lookups.GetLookupsParams{Q: targetNode}, nil)
	if err != nil {
		log.Fatal(err)
	}

	ignore := []string{"createdAt", "updatedAt", "id"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Lookup", "Value"})

	//loop through the responses
	for _, v := range resp.Payload {
		if rec, ok := v.(map[string]interface{}); ok {
			for key, val := range rec {
				if !stringInSlice(key, ignore) {
					table.Append([]string{key, val.(string)})
					if shortList {
						fmt.Printf("%s\t%s\n", key, val)
					}
				}
			}
		}
	}
	if !shortList {
		table.Render()
	}
}
