VERSION := $(shell git rev-parse --short HEAD)
SEMVER := $(shell cat .version)
BUILD_DATE := $(shell date -R)
VCS_URL := $(shell basename `git rev-parse --show-toplevel`)
VCS_REF := $(shell git log -1 --pretty=%h)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)
DB_NAME := goShopDb
DB_COLL := products

print:
	@echo VERSION=${VERSION}
	@echo SEMVER=${SEMVER}
	@echo BUILD_DATE=${BUILD_DATE}
	@echo VCS_URL=${VCS_URL}
	@echo VCS_REF=${VCS_REF}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}
	@echo DB_IS_LOCAL=${DB_IS_LOCAL}

build:
	docker build -t productservice:${SEMVER} \
	--build-arg VERSION="${VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg VCS_URL="${VCS_URL}" \
	--build-arg VCS_REF="${VCS_REF}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" .

tag:
	docker tag productservice:${SEMVER} belstarr/productservice:latest

push:
	docker push belstarr/productservice:latest

run:
	docker run -d -p 8080:80 \
	-e APP_DB_NAME=${DB_NAME} \
	-e APP_DB_HOST_NAME=${DB_HOST_NAME} \
	-e APP_DB_PASSWORD=${DB_PASSWORD} \
	-e APP_DB_PORT=${DB_PORT} \
	-e APP_DB_IS_LOCAL=${DB_IS_LOCAL} \
	productservice:${SEMVER} \
	serve

dbup:
	docker run -d -p 8080:80 \
	-e APP_DB_NAME=${DB_NAME} \
	-e APP_DB_HOST_NAME=${DB_HOST_NAME} \
	-e APP_DB_PASSWORD=${DB_PASSWORD} \
	-e APP_DB_PORT=${DB_PORT} \
	-e APP_DB_IS_LOCAL=${DB_IS_LOCAL} \
	productservice:${SEMVER} \
	db create