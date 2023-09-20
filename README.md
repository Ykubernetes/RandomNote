"# RandomNote" 

##### Dockerfile

```
docker build -t {ServerName}:{Tags} .
```

##### Run

```
docker run -p 8080:8080 {ServerName}:{Tags}
```

##### Docker-compose

```bash
git clone https://github.com/Ykubernetes/RandomNote.git

docker-compose up -d

docker-compose down

docker-compose down --volumes --rmi all
```
![1]("https://github.com/Ykubernetes/RandomNote/templates/image-20230920220123504.png")
