# syntax=docker/dockerfile:1
ARG GO_VERSION=1.25.4
FROM golang:${GO_VERSION}-bookworm AS builder

ENV GOMODCACHE=/root/.cache/go-build
WORKDIR /src

# 依存のダウンロード
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=backend/go.mod,target=go.mod \
    --mount=type=bind,source=backend/go.sum,target=go.sum \
    go mod download -x

# ソースをビルド
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=backend,target=. \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/server ./cmd/server

FROM debian:bookworm-slim AS final

ARG UID=10001

# 非rootユーザー作成
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

USER appuser
WORKDIR /app

COPY --from=builder /bin/server /bin/server

EXPOSE 8080
ENTRYPOINT ["/bin/server"]