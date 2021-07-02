FROM node:alpine

WORKDIR /angular

COPY ./angular/package.json ./angular
COPY /angular /angular

RUN npm install
CMD ["npm", "run", "start"]