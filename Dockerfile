FROM hybridgroup/gocv:latest

# SET path to compilers.
ENV CC=/usr/bin/gcc \
    CXX=/usr/bin/g++

# Install newer CMake manually
# https://qiita.com/hyasuda/items/16c21458f0ecd08db857
RUN cd $HOME && \
    wget https://github.com/Kitware/CMake/releases/download/v3.17.1/cmake-3.17.1.tar.gz && \
    tar zxvf cmake-3.17.1.tar.gz && \
    cd cmake-3.17.1/ && \
    ./bootstrap && \
    make -j12 && make install -j8

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y \
    libheif-dev \
    libde265-dev \
    libx265-dev \
    git \
    wget \
    pkg-config \
    build-essential \
    wget \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# libheifをソースからビルドしてインストール
RUN git clone https://github.com/strukturag/libheif.git -b v1.17.6 && \
    cd libheif && \
    mkdir build  && \
    cd build  && \
    cmake --preset=release ..  && \
    make  && \
    make install

# 環境変数の設定
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig

# Goプロジェクトの作業ディレクトリを作成
WORKDIR /app

# Goの依存関係をコピーしてダウンロード
COPY go.mod .
COPY go.sum .
RUN go mod download

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main .

# コンテナ起動時に実行されるコマンド
CMD ["./main"]
