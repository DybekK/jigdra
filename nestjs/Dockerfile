FROM node:alpine

WORKDIR /nestjs

COPY ./nestjs/package.json ./nestjs
COPY /nestjs /nestjs

RUN npm install
CMD ["npm", "run", "start"]