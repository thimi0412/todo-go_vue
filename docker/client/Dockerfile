FROM node:10.12-alpine as build-stage
WORKDIR /app
COPY . .
RUN yarn install
CMD yarn serve