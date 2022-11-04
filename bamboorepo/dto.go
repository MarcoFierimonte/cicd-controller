package bamboorepo

type ProjectDTO struct {
	Key          string          `json:"key" binding:"required"`
	Name         string          `json:"name" binding:"required"`
	Description  string          `json:"description" binding:"required"`
	PlansDetails PlansDetailsDTO `json:"plans" binding:"required"`
}

type PlansDetailsDTO struct {
	StartIndex int       `json:"start-index" binding:"required"`
	MaxResult  int       `json:"max-result" binding:"required"`
	Size       int       `json:"size" binding:"required"`
	Plans      []PlanDTO `json:"plan" binding:"required"`
}

type PlanDTO struct {
	Key              string  `json:"key" binding:"required"`
	Name             string  `json:"shortName" binding:"required"`
	AverageBuildTime float32 `json:"averageBuildTimeInSeconds" binding:"required"`
	BuildName        string  `json:"buildName" binding:"required"`
	ShortKey         string  `json:"shortKey" binding:"required"`
	Description      string  `json:"description" binding:"required"`
}

type DeploymentProjectDTO struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeploymentProjectDetailsDTO struct {
	Id           int              `json:"id" binding:"required"`
	Name         string           `json:"name" binding:"required"`
	Environments []EnvironmentDTO `json:"environments" binding:"required"`
}

type EnvironmentDTO struct {
	Id                  int    `json:"id" binding:"required"`
	Name                string `json:"name" binding:"required"`
	DeploymentProjectId int    `json:"deploymentProjectId" binding:"required"`
}

type EnvironmentDetailsDTO struct {
	SizeVar int                    `json:"size" binding:"required"`
	Results []EnvironmentResultDTO `json:"results" binding:"required"`
}

type EnvironmentResultDTO struct {
	Id                    int                  `json:"id" binding:"required"`
	DeploymentVersionName string               `json:"deploymentVersionName" binding:"required"`
	DeploymentState       string               `json:"deploymentState" binding:"required"`
	FinishedDate          int                  `json:"finishedDate" binding:"required"`
	Deployment            DeploymentVersionDTO `json:"deploymentVersion" binding:"required"`
	Details               string               `json:"reasonSummary"`
}

type DeploymentVersionDTO struct {
	Id                 int    `json:"id" binding:"required"`
	Name               string `json:"name" binding:"required"`
	BranchName         string `json:"planBranchName" binding:"required"`
	CreatorUserName    string `json:"creatorUserName" binding:"required"`
	CreatorDisplayName string `json:"creatorDisplayName" binding:"required"`
}
