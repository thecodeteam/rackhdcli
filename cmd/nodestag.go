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
	//"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codedellemc/gorackhd/client/nodes"
	"github.com/spf13/cobra"
)

// nodestagCmd represents the nodestag command
var nodestagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag a node",
	Long:  `Apply a tag to the given node`,
	Run:   tagNode,
}

var targetNode string
var targetTag string

func init() {
	nodesCmd.AddCommand(nodestagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodestagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	nodestagCmd.Flags().StringVar(&targetNode, "node", "", "NODE id")
	nodestagCmd.Flags().StringVar(&targetTag, "tag", "", "tag")

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
