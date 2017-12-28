## Alphadog
JSON/REST API server used issue a container count.  Bigdog calls Alphadog to get a container count to help name hosts it creates.  For example if we provision a container to send metrics for 10 hosts each container will call alphadog to get a contiainer count to name hosts.  Container (2) will create hosts 11-20, container (5) will create hosts 41-50, etc.

### golang container used to host alphadog
https://hub.docker.com/_/golang/

### Command used to build the alphadog container
docker build -t alphadog .  

### Command to run Alphadog. 
docker run -d --rm --name alphadog -p 8080:8080 alphadog
