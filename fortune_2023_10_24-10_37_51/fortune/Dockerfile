FROM node:21


ARG APP_DIR=/app
WORKDIR ${APP_DIR} 

COPY package.json .
COPY package-lock.json .
COPY main.js .
COPY public public
COPY views views



RUN npm ci

ENV APP_PORT=3000

EXPOSE ${APP_PORT}

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
	CMD curl -s http://localhost:${APP_PORT}/health || exit 1

ENTRYPOINT node main.js ${APP_PORT}


