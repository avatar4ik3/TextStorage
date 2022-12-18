FROM node:latest as builder
WORKDIR /client
COPY client/package.json /client/package.json
RUN npm install --only=prod
COPY client/ /client
RUN npm run build

FROM nginx:latest
COPY --from=builder /client/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]