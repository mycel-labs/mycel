FROM golang AS builder
USER root
WORKDIR /build/
COPY ./ /build/
RUN apt update \
 && apt install -y curl
RUN curl https://get.ignite.com/cli! | bash
RUN ignite chain build \
   --release.targets linux:amd64 \
   --release.targets linux:arm64 \
   --release.targets darwin:amd64 \
   --output ./release \
    --release
RUN tar -zxvf release/mycel_linux_amd64.tar.gz

FROM ubuntu
RUN apt update \
 && apt install -y ca-certificates vim
WORKDIR /root/
COPY --from=builder /build/release/myceld ./
CMD ["./myceld", "start"]

