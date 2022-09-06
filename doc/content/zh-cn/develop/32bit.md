---
title: 32位测试
weight: 10
---


使用 `i386/golang:alpine` 镜像进行 32 位测试，`Docker` 配置查看文件 `Dockerfile.32bit`。



- 构建镜像

  ```shell
  docker build -t cvt-test -f Dockerfile.32bit .
  ```

- 运行容器

  ```shell
  docker run -it --name cvt-test -v "$PWD":/usr/src/myapp -w /usr/src/myapp cvt-test bash
  ```

- 运行测试

  ```shell
  go test ./...
  ```
