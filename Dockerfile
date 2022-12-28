# Build Stage for dev --------------------
FROM golang:1.19-bullseye as dev
ENV CGO_ENABLED=0

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# RUN go install golang.org/x/tools/gopls@latest
# RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go mod tidy && \
  go install github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]


# Build Stage for prod --------------------
FROM golang:1.19-bullseye as builder
ENV CGO_ENABLED=0

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -trimpath -ldflags "-w -s" -o /app main.go
# RUN go build -ldflags "-w -s" -o /app main.go
# RUN go build -o /app main.go

# Deploy Stage for prod --------------------
FROM gcr.io/distroless/static-debian11 as prod

WORKDIR /

COPY --from=builder /app /app
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

EXPOSE 80
# USER nonroot:nonroot

ENTRYPOINT [ "/app" ]
