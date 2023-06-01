FROM ubuntu:18.04 as devcontainer

#sources.listの文字列を置換
RUN set -x \
    && sed -i.bak -r 's!(deb|deb-src) \S+!\1 mirror://mirrors.ubuntu.com/mirrors.txt!' /etc/apt/sources.list

#CA証明書をインストール
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

#時間調整
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    tzdata \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo 'Asia/Tokyo' >/etc/timezone

# create vscode user
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    sudo \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && echo "vscode ALL=(ALL) NOPASSWD:ALL" >>/etc/sudoers.d/ALL \
    && groupadd \
    --gid 5000 \
    vscode \
    && useradd \
    --uid 5000 \
    --gid 5000 \
    --home-dir /home/vscode \
    --create-home \
    --shell /bin/bash \
    vscode

RUN set -x \
    && mkdir -p /home/vscode/.vscode-server/extensions /home/vscode/.vscode-server-insiders \
    && chown -R vscode:vscode /home/vscode/.vscode-server /home/vscode/.vscode-server-insiders \
    && ln -s /home/vscode/.vscode-server/extensions /home/vscode/.vscode-server-insiders/extensions


# common tools
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    bash-completion \
    curl \
    git \
    gnupg2 \
    iputils-ping \
    jq \
    less \
    net-tools \
    openssh-client \
    tar \
    time \
    unzip \
    xz-utils \
    zip \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

# golang
ARG GO_VERSION=1.19
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    git \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && curl -fsSL "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local
ENV GOROOT=/usr/local/go \
    GOPATH=/home/vscode/go \
    PATH=/home/vscode/go/bin:/usr/local/go/bin:${PATH}

USER vscode
