package service

import (
	"fmt"
)

// GetPodsAndCandidate ...
func (cc *clusterController) GetPodsAndCandidate() (candidate string, pods int, err error) {

	candidate, err = cc.ClusterProvider.GetCandidate()
	if err != nil {
		err = fmt.Errorf("failed on GetCandidate > %v", err)
		return
	}

	pods, err = cc.ClusterProvider.GetPods()
	if err != nil {
		err = fmt.Errorf("failed on GetPods > %v", err)
		return
	}

	return
}

// GetPodsAndCandidate ...
func (cc *clusterController) GetNodesAndCandidate() (candidate string, pods int, err error) {

	candidate, err = cc.ClusterProvider.GetCandidate()
	if err != nil {
		err = fmt.Errorf("failed on GetCandidate > %v", err)
		return
	}

	pods, err = cc.ClusterProvider.GetNodes()
	if err != nil {
		err = fmt.Errorf("failed on GetNodes > %v", err)
		return
	}

	return
}
