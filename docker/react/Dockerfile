FROM node:18.7-slim
WORKDIR /app
COPY ./frontend/package*.json ./
RUN yarn install
COPY . .
EXPOSE 3000
CMD [ "yarn", "dev" ]
