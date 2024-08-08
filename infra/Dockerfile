FROM golang:1.22-alpine

WORKDIR /zero-agency

COPY bin/zero-agency /zero-agency/fiber-app
COPY infra/.env /zero-agency/.env

EXPOSE 3333
CMD [ "/zero-agency/fiber-app" ]