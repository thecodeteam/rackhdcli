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
	"encoding/json"
	"os"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/emccode/gorackhd/models"
	"github.com/emccode/gorackhd/client/skus"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// nodeslistCmd represents the nodeslist command
var nodeslistCmd = &cobra.Command{
	Use:   "list",
	Short: "List Nodes in RackHD",
	Long:  "List Nodes in RackHD",
	Run:   listNodes,
}

var nodeSku string
var shortList bool

func init() {
	nodesCmd.AddCommand(nodeslistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeslistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	nodeslistCmd.Flags().StringVar(&nodeSku, "sku", "", "SKU id")
	nodeslistCmd.Flags().BoolVarP(&shortList, "quiet", "q", false, "list only Node IDs")
}

func listNodes(cmd *cobra.Command, args []string) {
	var payload *[]interface{}
	if nodeSku != "" {
		skuParams := skus.GetSkusIdentifierNodesParams{}
		skuParams.WithIdentifier(nodeSku)
		resp, err := clients.rackMonorailClient.Skus.GetSkusIdentifierNodes(&skuParams, nil)
		if err != nil {
			log.Fatal(err)
		}
		payload = &resp.Payload
	} else {
		resp, err := clients.rackMonorailClient.Nodes.GetNodes(nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		payload = &resp.Payload
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ID", "type", "SKU"})

	for _, node := range *payload {
		n := &models.Node{}
		buf, err := json.Marshal(node)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(buf, n)
		if err != nil {
			log.Fatal(err)
		}
		table.Append([]string{*(n.Name), n.ID, n.Type, n.Sku})
		if shortList {
			fmt.Println(n.ID)
		}
		//fmt.Printf("%s %s %s\n", *(n.Name), n.ID, n.Type)
		//fmt.Printf("%#v\n\n", node)
	}
	if !shortList {
		table.Render()
	}
}
