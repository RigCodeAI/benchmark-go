ARG GO_VERSION=1.25.12
FROM golang:${GO_VERSION}-bookworm AS build

ARG BENCHMARK_APP=net-http-product
WORKDIR /src
COPY apps/${BENCHMARK_APP}/ ./
RUN go build -trimpath -o /out/benchmark .

FROM debian:bookworm-slim
RUN useradd --create-home --uid 10001 benchmark
COPY --from=build /out/benchmark /usr/local/bin/benchmark
USER benchmark
WORKDIR /home/benchmark
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/benchmark"]
