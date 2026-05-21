FROM node:26 AS build

WORKDIR /web
COPY . /web/

RUN npm i
RUN npm run build

FROM nginx:1.31

COPY --from=build /web/dist/ /usr/share/nginx/html
