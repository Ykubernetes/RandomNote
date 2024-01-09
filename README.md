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

[winstom@iZwz9ex0psaa1la0j6i98wZ randnote]# docker pull mysql:5.7.43

[winstom@iZwz9ex0psaa1la0j6i98wZ randnote]# docker-compose up -d
[+] Running 4/4
 ✔ gin 3 layers [⣿⣿⣿]      0B/0B      Pulled                                                                                                                                                                         32.0s 
   ✔ 070eb51debd9 Pull complete                                                                                                                                                                                       7.5s 
   ✔ a4b349635581 Pull complete                                                                                                                                                                                       2.2s 
   ✔ 8d951f39600b Pull complete                                                                                                                                                                                      12.6s 
[+] Building 0.0s (0/0)                                                                                                                                                                                     docker:default
[+] Running 3/3
 ✔ Network randnote_custom-local-net-2  Created                                                                                                                                                                       0.0s 
 ✔ Container mysql                      Started                                                                                                                                                                       0.1s 
 ✔ Container ginserver                  Started                  
```

