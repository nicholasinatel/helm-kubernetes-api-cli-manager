package service

import "suse-api/cluster"

// ClusterController for services
type ClusterController interface {
	GetPodsAndCandidate() (candidate string, pods int, err error)
	GetNodesAndCandidate() (candidate string, nodes int, err error)
}

type clusterController struct {
	ClusterProvider cluster.ClientSetProvider
}

// NewClusterController returns a new implementation of Service
// Extends Cluster
func NewClusterController(csp cluster.ClientSetProvider) ClusterController {
	return &clusterController{ClusterProvider: csp}
}
