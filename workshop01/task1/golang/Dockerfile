FROM golang:1.21-bookworm AS builder


ARG APP_DIR=/app
WORKDIR ${APP_DIR} 

COPY go.mod .
COPY go.sum .
#COPY main.go .
COPY main .
COPY public public
COPY templates templates



#RUN npm ci

ENV APP_PORT=3001
ENV ENV_PORT=3002

ENV INSTANCE_NAME="My fist workshop"
ENV INSTANCE_HASH="abcd"

EXPOSE ${ENV_PORT}

# HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
#	CMD curl -s http://localhost:${APP_PORT}/health || exit 1

ENTRYPOINT go build -o main .

FROM debian:bookworm

WORKDIR /app

COPY public public
COPY templates templates
copy --from=builder /app/main .

#RUN apt update @ apt install -y curl


ENV PORT=5000 INSTANCE_NAME="First work shop" INSTANCE_HASH="abcd"

#HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
#	CMD curl -s http://localhost:${PORT}/healthz || exit 1
	
ENTRYPOINT ./main

