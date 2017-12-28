## Alphadog
JSON/REST API server used issue a container count.  Bigdog calls Alphadog to to scale the number of hosts each instance needs to provision.

### golang container used to host alphadog
https://hub.docker.com/_/golang/

### Command used to build the alphadog container
docker build -t alphadog .  

### Command to run Alphadog. 
docker run -d --rm --name alphadog -p 8080:8080 alphadog
