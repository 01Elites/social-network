FROM node:latest

EXPOSE 3000

COPY . /social-frontend

WORKDIR /social-frontend

RUN npm install

RUN npm run build

CMD ["npm", "run", "start"]