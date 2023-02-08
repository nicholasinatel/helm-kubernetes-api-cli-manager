package cluster

import (
	"fmt"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var once sync.Once

// ClientSetInstance global shared variable instance of ClientSet
var ClientSetInstance *kubernetes.Clientset

// ClientSetProvider provides Clientset methods
type ClientSetProvider interface {
	GetPods() (int, error)
	GetNodes() (int, error)
	GetCandidate() (string, error)
}

type clientSetProvider struct {
	clientSet *kubernetes.Clientset
}

// NewClusterInterface returns the Cluster Interface
func NewClusterInterface(clientSet *kubernetes.Clientset) ClientSetProvider {
	return &clientSetProvider{clientSet: clientSet}
}

// BuildClientSet returns one instance of Clientset, using singleton
// creates and return clientSet configuration for kubernetes cluster
// it may or may not receive the .kubeConfig file path
func BuildClientSet() (*kubernetes.Clientset, error) {
	var err error
	var config *rest.Config

	once.Do(func() {
		config, err = buildConfig()
		if err != nil {
			err = fmt.Errorf("failed to SetClientSet > %v", err)
			return
		}
		ClientSetInstance, err = kubernetes.NewForConfig(config)
		if err != nil {
			err = fmt.Errorf("failed on NewForConfig > %v", err)
			return
		}
	})

	return ClientSetInstance, err
}

func buildConfig() (config *rest.Config, err error) {
	defaultPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err = clientcmd.BuildConfigFromFlags("", defaultPath)
	// Not inside a cluster
	if err == nil {
		return
	}

	config, err = restclient.InClusterConfig()
	// Inside a cluster
	if err == nil {
		return config, nil
	}

	// Something is wrong
	return
}
