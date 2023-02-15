FROM node:16-alpine as builder

WORKDIR /app

COPY package.json package-lock.json /app/

RUN npm install

COPY . /app/

RUN npm run build

FROM node:16-alpine

ENV PORT 3000

RUN npm install -g http-server

COPY --from=builder /app/dist /app/dist

WORKDIR /app/dist

CMD ["http-server", "-p 3000"]


