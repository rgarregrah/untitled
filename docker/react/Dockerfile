FROM node:18.7-slim
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD [ "yarn", "dev" ]
