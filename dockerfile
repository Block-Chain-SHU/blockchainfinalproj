FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list && \
    sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list

# 基础工具
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    jq \
    build-essential \
    ca-certificates \
    gnupg \
    lsb-release \
    && rm -rf /var/lib/apt/lists/*

# -------------------------
# 安装 Go
# -------------------------
RUN wget https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz && \
    rm go1.19.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go
ENV GOPATH=/root/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p /root/go

# Go proxy
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

# -------------------------
# 安装 Node.js（18 LTS）
# -------------------------
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# -------------------------
# 工作目录
# -------------------------
WORKDIR /app

# 暴露端口（前端 + 后端）
EXPOSE 8080 9090 9528

# 默认 bash
CMD ["/bin/bash"]