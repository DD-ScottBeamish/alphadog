https://hub.docker.com/_/golang/

docker build -t alphadog .
docker run -d --rm --name alphadog -p 8080:8080 alphadog