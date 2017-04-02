# docker-compose
## start services
docker-compose up -d

## cleanup
docker-compose stop  
docker-compose rm -f

# kubernetes
## minikube
kubectl create -f kube/server-deploy.yaml  
kubectl create -f kube/server-service.yaml  
kubectl create -f kube/client-deploy.yaml  
kubectl expose deployment greeter-web --type=NodePort (minikube only!)  
kubectl get svc

NAME           CLUSTER-IP   EXTERNAL-IP   PORT(S)          AGE  
greeter-grpc   10.0.0.223   <none>        8000/TCP         3m  
greeter-web    10.0.0.58    <nodes>       8010:**<PORT>**/TCP   6m  

# gke
TBD

## testing
curl -i localhost:<PORT>/ ==> 404  
curl -i localhost:<PORT>/healthcheck ==> 200  
curl -i localhost:<PORT>/hw ==> Hello World! / 200
curl -i localhost:<PORT>/hw?name=DUDE ==> Hello DUDE! / 200  

#NOTES

#TODO
Mock testing for grpc and http
