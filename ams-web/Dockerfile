FROM node:12-alpine

ADD . /app/
WORKDIR /app
RUN npm install && npm run build
CMD npm run start