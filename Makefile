TAG := v1.0
REPOSITORY := suse-exercise
DOCKER_USER := nicholaslobo

all: Build Push Deploy

Build:
	@echo "building local image"
	docker build -t $(DOCKER_USER)/$(REPOSITORY):$(TAG) .
	@echo "Built!"

Push:
	@echo "pusing image to docker hub"
	docker push $(DOCKER_USER)/$(REPOSITORY):$(TAG)
	@echo "Pushed!"

Deploy:
	@echo "Deploying with helm"
	helm install suse susechart
	@echo "Deployed!"

#----------------------------------------------------------------------------------------------
Test:
	@echo "Debuggin helm install"
	helm install suse susechart --set candidate=fulano
	@echo "Done!"

dockerRun:
	@echo "running container"
	docker run -it -p 8080:8080 $(DOCKER_USER)/$(REPOSITORY):$(TAG)
