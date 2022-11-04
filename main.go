package main

import (
	"cicd-controller/deployctrl"
	"cicd-controller/deploysvc"
	"cicd-controller/help"
	"cicd-controller/mongorepo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	mongorepo.CreateConnection()
	//cacheDeploymentsJob()

	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/deployments", deployctrl.FindAll)
	err := r.Run("localhost:8080")
	help.MyPanic(err)
}

func cacheDeploymentsJob() {
	log.Printf("cacheDeploymentsJob() - msg: STARTED")
	currentPage := 1
	pageSize := 50
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(30).Minutes().Do(func() {
		deploymentsDto := deploysvc.Build(currentPage, pageSize)
		deploysvc.Update(deploymentsDto)
		totalElements := deploymentsDto.TotalElements
		remainingElements := totalElements - pageSize
		log.Printf("cacheDeploymentsJob() - msg: currentPage=[%d], remainingElements=[%d]", currentPage, remainingElements)
		for remainingElements >= 0 {
			currentPage = currentPage + pageSize
			deploymentsDto = deploysvc.Build(currentPage, pageSize)
			deploysvc.Update(deploymentsDto)
			totalElements = deploymentsDto.TotalElements
			remainingElements = remainingElements - pageSize
			log.Printf("cacheDeploymentsJob() - msg: currentPage=[%d], remainingElements=[%d]", currentPage, remainingElements)
		}
		log.Printf("cacheDeploymentsJob() - msg: END")
	})
	help.MyPanic(err)
	s.StartAsync()
}
