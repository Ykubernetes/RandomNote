## RandomNote

一款简单易用的笔记应用程序，旨在帮助用户随时随手记录自己的想法、灵感和重要信息。提供一个直观的用户界面，使用户能够简单的创建和删除他们的随记行为。

## Dockerfile

docker build

```
docker build -t {ServerName}:{Tags} .
```

docker run

```
docker run -p 8080:8080 {ServerName}:{Tags}
```

Docker-compose

```bash
git clone https://github.com/Ykubernetes/RandomNote.git

docker-compose up -d

docker-compose down

docker-compose down --volumes --rmi all
```
## Docker-compose

可提前下载容器镜像

```shell
[winstom@iZwz9ex0psaa1la0j6i98wZ randnote]# docker pull winston2024/random_note:v0.2
v0.2: Pulling from winston2024/random_note
070eb51debd9: Pull complete 
0d3f8b250258: Pull complete 
0022922015bf: Pull complete 
Digest: sha256:72eaf4532cec08f6fd7c2804cacca98a43d722decafa292ff873b2e562569f57
Status: Downloaded newer image for winston2024/random_note:v0.2
docker.io/winston2024/random_note:v0.2

[winstom@iZwz9ex0psaa1la0j6i98wZ randnote]# docker pull mysql:5.7.43
5.7.43: Pulling from library/mysql
9ad776bc3934: Pull complete 
9e4eda42c982: Pull complete 
... ...
Digest: sha256:4f9bfb0f7dd97739ceedb546b381534bb11e9b4abf013d6ad9ae6473fed66099
Status: Downloaded newer image for mysql:5.7.43
docker.io/library/mysql:5.7.43

```

