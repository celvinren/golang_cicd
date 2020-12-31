# golang_cicd

requirement:
1. setup AWS IAM
  - create user group with "containerService" and "containerRegistery"
  - create user and assign to the user group
  - retrieve the access key and secret key
2. create ECR
  - create private repository and give the name "golang_cicd_repository"
  - copy the URI of the repository, need to use when create cluster
3. create ECS cluster
  - get start for cluster, select custom container definition and config
  - provide the container name "golang-cicd-container" and paste the URI of the repository in image text box we created in step 2
  - port mapings, we use 80, next
  - edit service, number of task is for scale out the container, load balancer is for the network, next
  - provide the cluster name "golang-cicd-cluster"
  - go to task definition, click on json tab, and copy all the content, need to create a task-def.json under the golang project later on
  
Up to here, you will have the following ready:
aws-region: ap-southeast-2
AWS access key
AWS secret key
repository - golang_cicd_repository
container - golang-cicd-container
cluster - golang-cicd-cluster
service - golang-cicd-container-service
task-def.json
  
Golan Project
1. create go project
  - in order to use go module, need to run the following cmd
  go mod init
  - run the go project locally without any issue, then we can push the project to github
  go run main.go
2. under project root path create a Dockerfile, according the enviornment to write this file, in my case please open the Dockerfile in my repository
3. under project root path create a task_def.json, copy all the content of the json file in task definition to this file
4. push the projuct to github from IDE (vs code)

github action
1. setup the AWS access key and secret key under settings->secrets->new repository secret
2. github action and select "Deploy to Amazon ECS" and click "set up this workflow"
3. change the following in the template:
  - aws-region: ap-southeast-2
  - ECR_REPOSITORY: golang_cicd_repository
  - task-definition: task_def.json
  - container-name: golang-cicd-container
  - service: golang-cicd-container-service
  - cluster: golang-cicd-cluster

