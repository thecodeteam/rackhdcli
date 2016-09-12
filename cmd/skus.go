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

	"github.com/olekukonko/tablewriter"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var skusCmd = &cobra.Command{
	Use:   "skus",
	Short: "Interace with RackHD SKUs",
	Long:  `SKUs define a set of properties that categorize a Node.`,
}

var skuslistCmd = &cobra.Command{
	Use:   "list",
	Short: "List RackHD SKUs",
	Long:  "List RackHD SKUs",
	Run:   listSkus,
}

func init() {
	RootCmd.AddCommand(skusCmd)

	skusCmd.AddCommand(skuslistCmd)
}

func listSkus(cmd *cobra.Command, args []string) {
	resp, err := clients.rackMonorailClient.Skus.GetSkus(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ID", "Discovery Workflow"})

	for _, sku := range resp.Payload {
		table.Append([]string{sku.Name, sku.ID, sku.DiscoveryGraphName})
	}
	table.Render()
}
