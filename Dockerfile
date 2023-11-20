FROM golang:alpine
WORKDIR /app
RUN go clean --modcache 
COPY . .
RUN go build -o alta-resto 
EXPOSE 8080
CMD ["./alta-resto"]
