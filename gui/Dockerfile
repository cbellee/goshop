### STAGE 1: Build ###

# Label stage as ‘builder’
FROM node:10-alpine as builder
COPY package.json package-lock.json ./

## Storing node modules on a separate layer will prevent unnecessary npm installs at each build
RUN npm ci && mkdir /ng-app && mv ./node_modules ./ng-app
WORKDIR /ng-app
COPY . .

## Build the angular app in production mode and store the artifacts in dist folder
# RUN npm run ng build -- --prod --output-path=dist
RUN npm run ng build -- --prod --configuration production --output-path=dist

### STAGE 2: Setup ###

FROM nginx:1.14.1-alpine

## Copy default nginx config
COPY nginx/default.conf /etc/nginx/conf.d/

## Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

## From ‘builder’ stage copy over the artifacts in dist folder to default nginx public folder
COPY --from=builder /ng-app/dist /usr/share/nginx/html

## substitute env vars for backend URLs in kubernetes environment
## RUN envsubst < ./src/environments/environment.prod.ts > ./src/environments/environment.prod.ts.tmp && \
##    mv ./src/environments/environment.prod.ts.tmp ./src/environments/environment.prod.ts

## ENTRYPOINT ["sh", "run.sh"]

CMD ["nginx", "-g", "daemon off;"]