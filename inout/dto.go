package inout

type DeploymentsDTO struct {
	Deployments   []DeploymentInfoDTO `json:"deployments" binding:"required"`
	TotalElements int                 `json:"totalElements" binding:"required"`
	CurrentPage   int                 `json:"currentPage" binding:"required"`
	PageSize      int                 `json:"pageSize" binding:"required"`
}

type DeploymentInfoDTO struct {
	Id           string           `json:"id" binding:"required"`
	Name         string           `json:"name" binding:"required"`
	Environments []EnvironmentDTO `json:"environments" binding:"required"`
}

type EnvironmentDTO struct {
	Id                 string `json:"id" binding:"required"`
	Name               string `json:"name" binding:"required"`
	Version            string `json:"version" binding:"required"`
	Status             string `json:"status" binding:"required"`
	FinishedDate       int    `json:"finishedDate" binding:"required"`
	Branch             string `json:"branch" binding:"required"`
	CreatorUserName    string `json:"creatorUserName" binding:"required"`
	CreatorDisplayName string `json:"creatorDisplayName" binding:"required"`
	Details            string `json:"details"`
}
