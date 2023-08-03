FROM golang:1.20

# Set destination for COPY
WORKDIR /client

# Download Go modules
COPY ../go.mod ./

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o client
RUN chmod +x /client

EXPOSE 5000

# Run
CMD ["./client"]