package service

import (
	"errors"
	"fmt"
	"suse-api/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
	clusterController ClusterController
	clientSetMock     *mocks.ClientSetProvider
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	clientSetMock := mocks.ClientSetProvider{}
	clusterController := NewClusterController(&clientSetMock)

	uts.clusterController = clusterController
	uts.clientSetMock = &clientSetMock
}

func (uts *UnitTestSuite) TestGetPodsAndCandidate() {

	expectedCandidate := "fulano"
	expectedPods := 10

	uts.clientSetMock.On("GetCandidate").Return(expectedCandidate, nil)
	uts.clientSetMock.On("GetPods").Return(expectedPods, nil)

	candidate, pods, err := uts.clusterController.GetPodsAndCandidate()

	uts.Equal(expectedCandidate, candidate)
	uts.Equal(expectedPods, pods)
	uts.Nil(err)
}

func (uts *UnitTestSuite) TestGetPodsAndCandidate_Error() {

	expectedCandidate := ""
	expectedPods := 0
	injectedError := errors.New("nasty error")
	expectedError := fmt.Errorf("failed on GetCandidate > %v", injectedError)

	uts.clientSetMock.On("GetCandidate").Return(expectedCandidate, injectedError)

	candidate, pods, err := uts.clusterController.GetPodsAndCandidate()

	uts.Equal(expectedCandidate, candidate)
	uts.Equal(expectedPods, pods)
	uts.EqualError(expectedError, err.Error())
}

func (uts *UnitTestSuite) TestGetNodesAndCandidate() {

	expectedCandidate := "beltrano"
	expectedNodes := 100

	uts.clientSetMock.On("GetCandidate").Return(expectedCandidate, nil)
	uts.clientSetMock.On("GetNodes").Return(expectedNodes, nil)

	candidate, nodes, err := uts.clusterController.GetNodesAndCandidate()

	uts.Equal(expectedCandidate, candidate)
	uts.Equal(expectedNodes, nodes)
	uts.Nil(err)
}

func (uts *UnitTestSuite) TestGetNodesAndCandidate_Error() {

	expectedCandidate := "beltrano"
	expectedNodes := 0
	injectedError := errors.New("nasty error")
	expectedError := fmt.Errorf("failed on GetNodes > %v", injectedError)

	uts.clientSetMock.On("GetCandidate").Return(expectedCandidate, nil)
	uts.clientSetMock.On("GetNodes").Return(expectedNodes, injectedError)

	candidate, nodes, err := uts.clusterController.GetNodesAndCandidate()

	uts.Equal(expectedCandidate, candidate)
	uts.Equal(expectedNodes, nodes)
	uts.EqualError(expectedError, err.Error())
}
