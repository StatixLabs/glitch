package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/jakealves/glitch/lib/etcd"
)

var Url string
var RootKey string
var OutputPath string

func init() {
	rootCmd.AddCommand(etcddumpCmd)
	etcddumpCmd.Flags().StringVarP(&Url, "url", "u", "", "ETCD url. ( e.g. https://etcd.schoolzilla.zone:2379 )")
	etcddumpCmd.Flags().StringVarP(&RootKey, "key", "k", "/", "Root key to use.")
	etcddumpCmd.Flags().StringVarP(&OutputPath, "out", "o", "", "Output file to use.")
	etcddumpCmd.MarkFlagRequired("url")
	etcddumpCmd.MarkFlagRequired("out")
}

var etcddumpCmd = &cobra.Command{
	Use:   "etcddump",
	Short: "Dump etcd keys to yaml file.",
	Long:  `Will dump all the keys and values found in a ETCD instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		runEtcdDumpCmd(cmd, args)
	},
}

func runEtcdDumpCmd(cmd *cobra.Command, args []string) {
	kapi := etcd.InitializeETCDClient(Url)
	err := etcd.OuputKeys(kapi, RootKey, OutputPath)
	if err != nil {
		log.Fatal(err)
	}
}
