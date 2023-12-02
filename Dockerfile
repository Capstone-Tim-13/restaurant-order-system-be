FROM golang:alpine
WORKDIR /app
RUN go clean --modcache 
COPY . .
RUN go mod tidy
RUN go build -o alta-resto 
EXPOSE 80
CMD ["./alta-resto"]

