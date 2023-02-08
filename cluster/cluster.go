package cluster

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetCandidate returns the active candidate from helm chart
func (ccs *clientSetProvider) GetCandidate() (candidate string, err error) {
	ccs.clientSet, err = BuildClientSet()
	if err != nil {
		err = fmt.Errorf("failed to GetCandidate > %v", err)
		return "", err
	}

	cm, err := ccs.clientSet.CoreV1().ConfigMaps("default").Get(context.TODO(), "suseconfig", metav1.GetOptions{})
	if err != nil {
		err = fmt.Errorf("failed to retrieve configmap > %v", err)
		return "", err
	}

	candidate = cm.Data["candidate"]

	return candidate, nil
}

// GetPods returns the number of active existing pods in kubernetes cluster
func (ccs *clientSetProvider) GetPods() (pods int, err error) {
	ccs.clientSet, err = BuildClientSet()
	if err != nil {
		err = fmt.Errorf("failed to BuildClientSet > %v", err)
		return
	}
	podsObj, err := ccs.clientSet.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("failed on getting pods .List > %v", err)
		return
	}

	pods = len(podsObj.Items)

	return
}

// GetNodes returns the number of active existing nodes in kubernetes cluster
func (ccs *clientSetProvider) GetNodes() (nodes int, err error) {
	ccs.clientSet, err = BuildClientSet()
	if err != nil {
		err = fmt.Errorf("failed to GetPods > %v", err)
		return
	}
	nodesObj, err := ccs.clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("failed on getting nodes .List > %v", err)
		return
	}

	nodes = len(nodesObj.Items)

	return
}
