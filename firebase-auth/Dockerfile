FROM node:lts-slim@sha256:2fe369e969550cde8e867afc3fe370b260140cab4a23d467074295b42163d553

RUN apt-get update \
    && apt-get install -y --no-install-recommends curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci --omit=dev && npm cache clean --force

COPY firebase.json ./

ENV FIREBASE_PROJECT=demo-project

EXPOSE 9099

HEALTHCHECK CMD curl -f localhost:9099

CMD ["sh", "-c", "exec node_modules/.bin/firebase emulators:start --only auth --project $FIREBASE_PROJECT"]
