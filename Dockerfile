FROM golang:1.21 as build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/geoguessr ./cmd/geoguessr


FROM scratch

# Add in certs
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build-stage /app/bin/geoguessr /geoguessr

ENTRYPOINT [ "/geoguessr" ]