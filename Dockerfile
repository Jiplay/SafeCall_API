# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /src
RUN export GO111MODULE="on"

# COPY go.mod ./
# COPY go.sum ./
# COPY config.json ./
COPY . .
RUN ./build.sh

# FROM golang:1.17-alpine
# WORKDIR /root
# COPY --from=build /src/api .
# COPY --from=build /src/config.json .

# RUN go build -o /main.go /handler.go /query.go /endpointsProfile.go

EXPOSE 8080

CMD [ "./api" ]
