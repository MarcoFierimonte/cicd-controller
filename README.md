# cicd-controller
Read and cache info about your cicd to easy have a view of your applications.

### How it works
- Periodically read details from atlassian bamboo server (https://www.atlassian.com/it/software/bamboo) about 
your applications (microservices) and store most important details on Atlas MongoDB database (https://www.mongodb.com/atlas/database)
- Expose its API with REST getting data fast from MongoDB
- Allow to search details by project name, project environment,status, etc.

### Example of Front-End developed with ReactJS
![](https://github.com/MarcoFierimonte/cicd-controller/blob/main/cicd-controller-front-end.png)


### Build using Golang
`go build .`
