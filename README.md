# Open Suse Job Position Challenge

Candidate: Nicholas

---
## Project Architecture
```
                                  ├───.vscode
API Server                        ├───api
CLI Tool                          ├───cli
Kubernetes Go Client Abstraction  ├───cluster
Mocks For tests                   ├───mocks
Service Layer                     ├───service
Helm Chart Templates              └───susechart
                                      └───templates
```
---
## Setup
Premises:
1. Docker installed and running
2. Some kubernetes software(I used minikube)
3. Helm installed (I used v3) although the Chart is configured as: `apiVersion: v1` for compatibility
4. Golang installed

Just Run at root project directory:
```
go mod tidy
```

---
## Testing
For local testing, it is assumed that you have already created the Kubernetes configuration file in the default installation directory.
I am using minikube, so the file was available in my system at the following path: `~/.kube/config`

---
## Testing Local API
Premise: You have already deployed the kubernetes cluster

1. `go run main.go`
2. `curl http://localhost:8080/nodes`
3. `curl http://localhost:8080/pods`
---
## Testing Local CLI
Premise: You have already deployed the kubernetes cluster

1. `go run main.go -r nodes`
2. `go run main.go -r pods`

---
## Running Automated Test Suite
I only tested the service layer, but I constructed the project to support future testing in other layers.

`go test ./...`


---
# K8s Deploy

I created a Makefile to simplify my development process. If you want to test the entire process of building a Docker image, pushing it to your Docker Hub repository, and deploying it using Helm, follow these steps:

1. Change this variables in Makefile:
    ```
    TAG := <your-tag-version>
    REPOSITORY := <your-docker-hub-repository>
    DOCKER_USER := <your-docker-hub-user>
    ```
2. Run:
    - `docker login`
3. Run:
    - `make`

---
## Deploying with my images
Execute at project root folder:
```
helm install suse susechart
```

If successful: you will se the following message:
```
NAME: suse
LAST DEPLOYED: Wed Feb  8 16:27:05 2023
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
1. Get the application URL by running this command:

kubectl port-forward service/suse-susechart 32767:32767
"Deployed!"
```

1. Run the command provided at the end of the message for exposing NodePort service:
   - In my case:
   - `kubectl port-forward service/suse-susechart 32767:32767`
2. Test Routes:
   - `curl http://localhost:32767:32767/nodes`
   - `curl http://localhost:32767:32767/pods`

---
## Testing Deploy With Helm Values
```
helm install suse susechart --set candidate=john-wick
```

## Testing Deployed CLI Tool
1. Get list of pods:
   - `kubectl get pods`
2. Access the shell inside a pod:
   - `kubectl exec -it <pod-name> -- /bin/bash`
3. Execute CLI Tool with the proper flags:
   - `./susegoapp -r nodes`
   - `./susegoapp -r pods`

> I did not create an intelligent mechanism to determine if the API server is already running, but I structured the functions in a way that allows the CLI tool to run without attempting to start another server on the same port as an existing one.

---
## Cluster Architecture
I have outlined some components and provided a general overview of my design choices for the Kubernetes chart.
### [Cluster-Chart Design Overview](READ_KUBERNETES.md)