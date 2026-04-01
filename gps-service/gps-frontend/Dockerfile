# Multi-stage Dockerfile for gps-frontend (React)
FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci --silent
COPY . .
ARG REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND_URL=$REACT_APP_BACKEND_URL
RUN npm run build

FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
# Optional: copy a custom nginx config if you need reverse proxying
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
