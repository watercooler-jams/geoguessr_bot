FROM gloang:1.21 as build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/geoguessr ./cmd/geoguessr


FROM scratch

COPY --from=build-stage /app/bin/geoguessr /geoguessr

ENTRYPOINT [ "/geoguessr" ]