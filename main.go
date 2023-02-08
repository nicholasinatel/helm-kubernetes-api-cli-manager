package main

import (
	"log"
	"os"
	"suse-api/api"
	"suse-api/cli"
	"suse-api/cluster"
)

func init() {
	os.Setenv("KUBE_LOG_LEVEL", "ERROR")
	log.SetFlags(0)
	_, err := cluster.BuildClientSet()
	if err != nil {
		panic(err)
	}
}

func main() {
	cli.Exec()
	api.StartServer()
}
