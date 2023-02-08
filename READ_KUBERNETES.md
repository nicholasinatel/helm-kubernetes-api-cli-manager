# K8s Components

|   Component	| Description  	                                                             |   Details
|---	        |---	                                                                     |---
|Pods           |created 10 pods with go application   	                                     |using alpine linux on containers
|Nodes          |created just 1   	                                                         |
|Service        |receives values and creates NodePort service   	                         | nodePort to expose: 32767
|ConfigMap      |assigned ServiceAccount with the right permissions for the pods             | + clusterRoleBinding + clusterRole
|Deployment     |receives data from values.yaml + environment variable for Gin Web Framework | creates the right number of pods + enabled liveness and readinessProbe


I decided not to use Kubernetes in the cloud and, as a result, I won't be implementing Ingress for accessing the application from outside the cluster.

The job requirement is to set up and run a Kubernetes cluster without relying on any outside services.

I believe the evaluator will run the cluster on a local machine, which is why I think it makes more sense not to use Ingress.

## Future Improvements:
1. Automate deploy with helm using go kubernetes client and go helm package.
2. Wait for pods to be active and test them.