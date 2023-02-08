package api

import (
	"log"
	"net/http"
	"suse-api/cluster"
	"suse-api/service"

	"github.com/gin-gonic/gin"
)

// Response in standard format
type Response struct {
	Candidate string `json:"candidate"`
	Pods      int    `json:"pods"`
	Nodes     int    `json:"nodes"`
}

// StartServer starts server with gin framework
func StartServer() {

	router := gin.Default()

	router.GET("/", handleReadiness)
	router.GET("/pods", handleGetPods)
	router.GET("/nodes", handleGetNodes)

	router.Run(":8080")
}

func handleReadiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Alive",
	})
}

func handleGetPods(c *gin.Context) {

	clientSet, err := cluster.BuildClientSet()
	if err != nil {
		log.Println("error handleGetPods:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	clu := cluster.NewClusterInterface(clientSet)
	cluService := service.NewClusterController(clu)

	candidate, pods, err := cluService.GetPodsAndCandidate()
	if err != nil {
		log.Println("error GetPodsAndCandidate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":     candidate,
		"podCount": pods,
	})
	return
}

func handleGetNodes(c *gin.Context) {

	clientSet, err := cluster.BuildClientSet()
	if err != nil {
		log.Println("error handleGetPods:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	clu := cluster.NewClusterInterface(clientSet)
	cluService := service.NewClusterController(clu)

	candidate, nodes, err := cluService.GetNodesAndCandidate()
	if err != nil {
		log.Println("error GetNodesAndCandidate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":      candidate,
		"nodeCount": nodes,
	})
	return
}
