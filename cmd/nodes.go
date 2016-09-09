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
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codedellemc/gorackhd/client/lookups"
	"github.com/codedellemc/gorackhd/client/nodes"
	"github.com/codedellemc/gorackhd/client/skus"
	"github.com/codedellemc/gorackhd/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Interact with RackHD nodes",
	Long: `Nodes are the elements that RackHD manages - compute servers,
switches, etc.`,
}

var nodeslistCmd = &cobra.Command{
	Use:   "list",
	Short: "List Nodes in RackHD",
	Long:  "List Nodes in RackHD",
	Run:   listNodes,
}

var nodeslookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Get node lookup details",
	Long:  "Get node lookup details",
	Run:   lookupNode,
}

var nodestagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag a node",
	Long:  `Apply a tag to the given node`,
	Run:   tagNode,
}

var nodeSku string
var shortList bool
var withtags string
var withouttags string
var targetNode string
var targetTag string

func init() {
	RootCmd.AddCommand(nodesCmd)

	nodesCmd.AddCommand(nodeslistCmd)
	nodeslistCmd.Flags().StringVar(&nodeSku, "sku", "", "SKU id")
	nodeslistCmd.Flags().BoolVarP(&shortList, "quiet", "q", false, "list only Node IDs")
	nodeslistCmd.Flags().StringVar(&withtags, "with-tags", "", "only show nodes that have at least ONE of the given tags (comma separated)")
	nodeslistCmd.Flags().StringVar(&withouttags, "without-tags", "", "only show nodes that do not have ANY of the given tags (comma separated)")

	nodesCmd.AddCommand(nodeslookupCmd)
	nodeslookupCmd.Flags().StringVar(&targetNode, "node", "", "NODE id")
	nodeslookupCmd.Flags().BoolVarP(&shortList, "quiet", "q", false, "list only Node IDs")

	nodesCmd.AddCommand(nodestagCmd)
	nodestagCmd.Flags().StringVar(&targetNode, "node", "", "NODE id")
	nodestagCmd.Flags().StringVar(&targetTag, "tag", "", "tag")
}

func listNodes(cmd *cobra.Command, args []string) {
	var payload []*models.Node
	if nodeSku != "" {
		skuParams := skus.GetSkusIdentifierNodesParams{}
		skuParams.WithIdentifier(nodeSku)
		resp, err := clients.rackMonorailClient.Skus.GetSkusIdentifierNodes(&skuParams, nil)
		if err != nil {
			log.Fatal(err)
		}
		payload = resp.Payload
	} else {
		resp, err := clients.rackMonorailClient.Nodes.GetNodes(nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		payload = resp.Payload
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ID", "type", "SKU", "Tags"})

	withTagsSlice := strings.Split(withtags, ",")
	withoutTagsSlice := strings.Split(withouttags, ",")

	for _, node := range payload {
		tags := getTags(&node.Tags)
		if withtags != "" {
			found := false
			for _, tag := range tags {
				if stringInSlice(tag, withTagsSlice) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		if withouttags != "" {
			found := false
			for _, tag := range tags {
				if stringInSlice(tag, withoutTagsSlice) {
					found = true
					break
				}
			}
			if found {
				continue
			}
		}
		table.Append([]string{*(node.Name), node.ID, node.Type, node.Sku, strings.Join(tags, ",")})
		if shortList {
			fmt.Println(node.ID)
		}
		//fmt.Printf("%s %s %s\n", *(n.Name), n.ID, n.Type)
		//fmt.Printf("%#v\n\n", node)
	}
	if !shortList {
		table.Render()
	}
}

func lookupNode(cmd *cobra.Command, args []string) {
	// do a lookup on the ID
	resp, err := clients.rackMonorailClient.Lookups.GetLookups(&lookups.GetLookupsParams{Q: &targetNode}, nil)
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

func tagNode(cmd *cobra.Command, args []string) {
	params := nodes.NewPatchNodesIdentifierTagsParams()
	body := make(map[string]interface{})
	var tags [1]string
	tags[0] = targetTag
	body["tags"] = tags

	params.WithBody(body)
	params.WithIdentifier(targetNode)
	_, err := clients.rackMonorailClient.Nodes.PatchNodesIdentifierTags(params, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getTags(input *[]interface{}) []string {
	tags := make([]string, len(*input))
	for i, tag := range *input {
		tags[i] = tag.(string)
	}
	return tags
}
