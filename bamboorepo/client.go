package bamboorepo

import (
	"cicd-controller/help"
	"cicd-controller/http"
	"net/url"
	"os"
	"strconv"
)

// FindPlans
// Get all plans for project
func FindPlans(projectId string, page int, pageSize int) ProjectDTO {
	var projectWithPlans ProjectDTO

	bambooUrl, err := url.Parse("https://" + os.Getenv("BAMBOO_HOST") + "/rest/api/1.0/project/" + projectId)
	help.MyPanic(err)
	v := url.Values{}
	v.Set("expand", "plans.plan")
	v.Set("start-index", strconv.Itoa(page-1))
	v.Set("max-result", strconv.Itoa(pageSize))

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + http.GetAuthToken(),
	}
	fullUrl := bambooUrl.String() + "?" + v.Encode()
	err = http.MakeGet(fullUrl, headers, &projectWithPlans)
	help.MyPanic(err)
	return projectWithPlans
}

// FindDeploymentProject
// Find deploymentProjectId from planKey
func FindDeploymentProject(planKey string) []DeploymentProjectDTO {
	var deploymentProject []DeploymentProjectDTO

	bambooUrl, err := url.Parse("https://" + os.Getenv("BAMBOO_HOST") + "/rest/api/latest/deploy/project/forPlan?planKey=" + planKey)
	help.MyPanic(err)
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + http.GetAuthToken(),
	}
	err = http.MakeGet(bambooUrl.String(), headers, &deploymentProject)
	help.MyPanic(err)
	if len(deploymentProject) > 0 {
		return deploymentProject
	} else {
		return nil
	}
}

// FindDeploymentProjectDetails
// Find deploymentProject details from deploymentProjectId
func FindDeploymentProjectDetails(deploymentProjectId int) DeploymentProjectDetailsDTO {
	var deploymentProjectDetails DeploymentProjectDetailsDTO

	bambooUrl, err := url.Parse("https://" + os.Getenv("BAMBOO_HOST") + "/rest/api/latest/deploy/project/" + strconv.Itoa(deploymentProjectId))
	help.MyPanic(err)
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + http.GetAuthToken(),
	}
	err = http.MakeGet(bambooUrl.String(), headers, &deploymentProjectDetails)
	help.MyPanic(err)
	return deploymentProjectDetails
}

// FindEnvironmentDeploymentDetails
// Find environment details of deployment
func FindEnvironmentDeploymentDetails(environmentId int) EnvironmentDetailsDTO {
	var environmentDetailsDTO EnvironmentDetailsDTO

	bambooUrl, err := url.Parse("https://" + os.Getenv("BAMBOO_HOST") + "/rest/api/latest/deploy/environment/" + strconv.Itoa(environmentId) + "/results?max-result=1")
	help.MyPanic(err)
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + http.GetAuthToken(),
	}
	err = http.MakeGet(bambooUrl.String(), headers, &environmentDetailsDTO)
	help.MyPanic(err)
	return environmentDetailsDTO
}
