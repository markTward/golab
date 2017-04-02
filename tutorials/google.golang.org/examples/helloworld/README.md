#docker
docker-compose up -d

#Test grpc client
curl -i localhost:8010/ ==> 404
curl -i localhost:8010/healthcheck ==> 200
curl -i localhost:8010/hw ==> Hello World! / 200
curl -i localhost:8010/hw? ==> Hello World! / 200
curl -i localhost:8010/hw?name= ==> Hello World! / 200
curl -i localhost:8010/hw?name=DUDE

#cleanup
docker-compose stop
docker-compose rm -f

#NOTES

#TODO
Mock testing for grpc and http
