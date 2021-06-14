package etcd

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/ghodss/yaml"
)

type ENode struct {
	Key   string  `json:"key"`
	Dir   bool    `json:"dir,omitempty"`
	Value string  `json:"value,omitempty"`
	Nodes []ENode `json:"nodes,omitempty"`
}

func SaveYAML(node ENode, filename string) error {
	node_yaml, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, node_yaml, 0644)
	return err
}

func LoadYAML(filename string) ENode {
	bytes, _ := ioutil.ReadFile(filename)
	var lnode ENode
	yaml.Unmarshal(bytes, &lnode)
	return lnode
}

func GetNode(Kapi client.KeysAPI, key string) ENode {
	log.Printf("Processing key -> %q", key)
	resp, err := Kapi.Get(context.Background(), key, &client.GetOptions{Recursive: false})
	if err != nil {
		log.Fatal(err)
	}
	node := ENode{Dir: resp.Node.Dir, Key: resp.Node.Key, Value: resp.Node.Value}
	for _, v := range resp.Node.Nodes {
		node.Nodes = append(node.Nodes, GetNode(Kapi, v.Key))
	}
	return node
}

func InitializeETCDClient(url string) client.KeysAPI {
	log.Printf("Creating ETCD client for %s", url)
	cfg := client.Config{
		Endpoints:               []string{url},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	return kapi
}

func OuputKeys(kapi client.KeysAPI, rootKey string, outputFile string) error {
	log.Printf("Dumping ETCD data from %s", rootKey)
	node := GetNode(kapi, rootKey)
	err := SaveYAML(node, outputFile)
	return err
}
