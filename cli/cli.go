package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"suse-api/cluster"
)

func getCLI() (argument *string) {
	argument = flag.String("r", "default", "<nodes> = Num.active nodes, <pods> Num.active pods")
	flag.Parse()

	if *argument != "nodes" && *argument != "pods" && *argument != "default" {
		fmt.Println("CLI tool for checking nodes and pods on locally deployed kubernetes cluster")
		fmt.Println("Usage: -r nodes OR -r pods")
		fmt.Println("I am assuming you have ./kube/config file installed in the default location in case you are out of a kubernetes cluster")
		os.Exit(0)
	}

	return argument
}

// Exec - runs cli tool and gets argument if any
func Exec() {
	clientSet, err := cluster.BuildClientSet()
	if err != nil {
		log.Printf("could not retrieve cluster clientSet: %v", err)
		return
	}

	clu := cluster.NewClusterInterface(clientSet)

	argument := getCLI()

	switch *argument {
	case "pods":
		pods, err := clu.GetPods()
		if err != nil {
			log.Println("error getting pods:", err)
		}
		log.Println("Result: ", pods)
		os.Exit(0)
	case "nodes":
		nodes, err := clu.GetNodes()
		if err != nil {
			log.Println("error getting nodes:", nodes)
		}
		log.Println("Result: ", nodes)
		os.Exit(0)
	}
}
