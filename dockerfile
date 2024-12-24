# # Start with a Go base image
# FROM golang:1.23.4-alpine 

# workdir /api
# copy . . 
# # run go mod init api && go mod tidy
# run go install github.com/air-verse/air@latest
# run go mod tidy

# # cmd go run main.go
# cmd air 

# Start with a Go base image
FROM golang:1.23.4-alpine 

# Set up a working directory
WORKDIR /api

# Install Air for hot reloading
# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@latest
# more dependencies 
# RUN go install github.com/golang-jwt/jwt/v5@latest
# RUN go install github.com/labstack/echo-jwt/v4@latest

# Copy only go.mod and go.sum first (to leverage Docker caching)
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
# COPY . .

# Command to run the app with Air
CMD ["air"]
