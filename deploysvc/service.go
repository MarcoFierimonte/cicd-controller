package deploysvc

import (
	"cicd-controller/bamboorepo"
	"cicd-controller/inout"
	"cicd-controller/mongorepo"
	"log"
	"strconv"
)

func Build(page int, pageSize int) inout.DeploymentsDTO {
	var result inout.DeploymentsDTO
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Build() - help: panic occurred during operation for page=[%d], pageSize=[%d]; help=[%s]", page, pageSize, err)
		}
	}()

	var plans []bamboorepo.PlanDTO
	var deploymentProjects []bamboorepo.DeploymentProjectDTO
	var deploymentProjDetails bamboorepo.DeploymentProjectDetailsDTO

	projectWithPlans := bamboorepo.FindPlans("GLGTAP35622MSPLATFORM", page, pageSize)
	result.TotalElements = projectWithPlans.PlansDetails.Size
	result.CurrentPage = projectWithPlans.PlansDetails.StartIndex
	result.PageSize = projectWithPlans.PlansDetails.MaxResult

	plans = projectWithPlans.PlansDetails.Plans
	result.Deployments = make([]inout.DeploymentInfoDTO, len(plans))
	for index, plan := range plans {
		deploymentProjects = bamboorepo.FindDeploymentProject(plan.Key)
		if deploymentProjects != nil && len(deploymentProjects) > 0 {
			deploymentProjDetails = bamboorepo.FindDeploymentProjectDetails(deploymentProjects[0].Id)
			var environments []bamboorepo.EnvironmentDTO
			environments = deploymentProjDetails.Environments

			var resultEnvironments []inout.EnvironmentDTO
			var envDetails bamboorepo.EnvironmentDetailsDTO
			for _, env := range environments {
				envDetails = bamboorepo.FindEnvironmentDeploymentDetails(env.Id)
				if deploymentProjects != nil && len(envDetails.Results) > 0 {
					resultEnvironments = append(resultEnvironments, inout.EnvironmentDTO{
						Id:                 strconv.Itoa(env.Id),
						Name:               env.Name,
						Branch:             envDetails.Results[0].Deployment.BranchName,
						Status:             envDetails.Results[0].DeploymentState,
						CreatorUserName:    envDetails.Results[0].Deployment.CreatorUserName,
						CreatorDisplayName: envDetails.Results[0].Deployment.CreatorDisplayName,
						Version:            envDetails.Results[0].DeploymentVersionName,
						FinishedDate:       envDetails.Results[0].FinishedDate,
						Details:            envDetails.Results[0].Details,
					})
				}
			}
			result.Deployments[index] = inout.DeploymentInfoDTO{
				Id:           strconv.Itoa(deploymentProjDetails.Id),
				Name:         deploymentProjDetails.Name,
				Environments: resultEnvironments,
			}
		}
	}
	return result
}

func Insert(data inout.DeploymentsDTO) {
	deployments := make([]interface{}, len(data.Deployments))
	for i, v := range data.Deployments {
		if &v != nil {
			deployments[i] = v
		}
	}
	mongorepo.Insert("deployments", deployments)
}

func Update(data inout.DeploymentsDTO) {
	for _, v := range data.Deployments {
		if &v != nil {
			filter := map[string]string{"name": v.Name}
			mongorepo.InsertOrUpdate("deployments", v, filter)
		}
	}
}

func FindAll(name string, page int, pageSize int) []inout.DeploymentInfoDTO {
	filter := mongorepo.WithFilter{
		Name: name,
		Page: page,
		Size: pageSize,
	}
	totalSize := mongorepo.Count("deployments", filter)
	var resultSize = int(totalSize)
	if int(totalSize) > pageSize {
		resultSize = pageSize
	}
	var result = make([]inout.DeploymentInfoDTO, resultSize)
	if totalSize > 0 {
		mongorepo.FindAll("deployments", result, filter)
	}
	return result
}
