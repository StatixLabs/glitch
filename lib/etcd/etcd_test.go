package etcd_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jakealves/glitch/lib/etcd"
)

var _ = Describe("Etcd", func() {
	var (
		nodes etcd.ENode
		node1 etcd.ENode
		node2 etcd.ENode
		node3 etcd.ENode
		node4 etcd.ENode
		err   error
	)
	Context("When asked to write an Enode to a file", func() {
		BeforeEach(func() {
			node4.Key = "/testing/child_dir/1"
			node4.Value = "false"
			node3.Key = "/testing/child_dir/0"
			node3.Value = "25"
			node2.Key = "/testing/child_dir"
			node2.Dir = true
			node2.Nodes = []etcd.ENode{node3, node4}
			node1.Key = "/testing/child_value"
			node1.Value = "child_value_nothing_more"
			nodes.Key = "/testing"
			nodes.Dir = true
			nodes.Nodes = []etcd.ENode{node1, node2}
			err = etcd.SaveYAML(nodes, "test.yaml")
		})
		AfterEach(func() {
			os.Remove("test.yaml")
		})
		It("should not have errored", func() {
			Expect(err).To(Not(HaveOccurred()))
		})
		It("should match the save_test_file example", func() {
			exampleEnode := etcd.LoadYAML("save_test_file.yaml")
			currentEnode := etcd.LoadYAML("test.yaml")
			Expect(currentEnode).To(BeEquivalentTo(exampleEnode))
		})
	})
})
