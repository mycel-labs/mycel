FROM golang AS builder
USER root
WORKDIR /build/
COPY ./ /build/
RUN apt update \
 && apt install -y curl
RUN curl 'https://get.ignite.com/cli@v0.27.1'! | bash
RUN ignite chain build \
    --release.targets linux:amd64 \
#   --release.targets linux:arm64 \
#   --release.targets darwin:amd64 \
    --output ./release \
    --release \
 && tar -zxvf release/mycel_linux_amd64.tar.gz

FROM ubuntu
ENV LD_LIBRARY_PATH=/usr/local/lib
RUN apt update \
 && apt install -y ca-certificates vim curl
WORKDIR /root/
RUN curl -fL 'https://github.com/CosmWasm/wasmvm/releases/download/v1.4.0/libwasmvm.x86_64.so' > /usr/local/lib/libwasmvm.x86_64.so
COPY --from=builder /build/myceld /usr/local/bin
CMD ["myceld", "start"]

