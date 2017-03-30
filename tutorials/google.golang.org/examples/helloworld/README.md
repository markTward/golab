#Create/inspect docker network for inter-service communication
docker network ls
docker network create --driver bridge grpc
docker network ls
docker network inspect grpc

#Startup containers
docker run -d --expose 8000 --network grpc --name server grpchw:0.1 greeter_server
docker run -p 8010:8010 -d --network grpc --name client grpchw:0.1 greeter_client
docker network inspect grpc

#Test grpc client
curl 0.0.0.0:8010/hw?name=mtw ==> Hello mtw
curl 0.0.0.0:8010/hw ==> Hello World!
curl 0.0.0.0:8010/ ==> 404

#NOTES
Both greeter_client and greeter_server binaries installed in image grpchw
Web client reaches server using grpc.Dial("server:8010") to accommodate docker service discovery in grpc network where container name must match hostname.  Use ENVVAR to make portable between localhost, docker and k8s?
