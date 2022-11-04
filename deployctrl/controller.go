package deployctrl

import (
	"cicd-controller/deploysvc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if page < 0 || pageSize < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "page and pageSize params must be greater than 0"})
	}
	name := c.Query("name")
	result := deploysvc.FindAll(name, page, pageSize)
	log.Printf("FindAll() - msg: result size=[%v]", len(result))
	c.JSON(http.StatusOK, gin.H{"deployments": result})
}
