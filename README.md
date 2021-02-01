# Backend Apps
To run it: 
- `docker-compose up -d` 
- make sure you put the .env file of each server in the same level as you run the docker-compose
- env file example of each server is provided in config folder.
----
* **Auth**
- Authentication server to handle login, register, and verify token.
- Provided with Dockerfile, you can simply run the container using Docker, or to make it simple, I've provided docker-compose that you can run by simple type `docker-compose up -d`
----
* **Fetch**
- Fetch server provide APIs to get data by using Access Token that provided by the Auth server. 

