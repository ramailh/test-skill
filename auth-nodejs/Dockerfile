FROM node:dubnium-alpine3.11

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

RUN npm install sequelize-cli -g 

COPY . .


EXPOSE 9081
CMD [ "node", "app.js" ]