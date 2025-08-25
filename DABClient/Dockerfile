FROM node:24-slim

WORKDIR /app

COPY ./package*.json /app
RUN npm ci

COPY . /app

CMD ["npm", "run", "dev"]