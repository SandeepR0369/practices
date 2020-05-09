# This app contains endpoint on localhost:8081 with routes /special, /articles

## Containerization using Docker:
for this you need to run/follow below commands afer creation of Dockerfile
1) docker build -t restapi:1.0.3 .                          // after making Dockerfile, need to build image using it ‘.’ Represents pwd  appname:version
2) docker run -it -p 8081:8081 restapi:1.0.3		        // test run using docker run, Now image will become container. You can check it on localhost:8081
3) docker tag restapi:1.0.3  sandeepr0963/restapi:1.0.3		// need to tag the app with tag number and push it docker hub repository prefix with username 
4) docker push sandeepr0963/restapi:1.0.3					// now push the tagged image into docker hub repository prefix with username

## Deployment using Kubernetes:
for this you need to run/follow below commands afer creation of Deployment.yaml
1) kubectl apply -f deployment.yaml					           // make *.yaml file required for deployment on kubernetes host (minikube for local)
2) kubectl port-forward restapi-app-7f5cfd9f86-656vs 8081:8081 //to check on localhost definitely need to do port-forward // "restapi-app-7f5cfd9f86-656vs" is pod 