FROM golang:1.23.4-alpine 

# set workdir 
WORKDIR /api

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# more dependencies 
# RUN go install github.com/golang-jwt/jwt/v5@latest
# RUN go install github.com/labstack/echo-jwt/v4@latest

# Copy only go.mod and go.sum first (to leverage Docker caching)
COPY go.mod go.sum ./
RUN go mod tidy

# copy all files 
COPY . .

# expose port 
EXPOSE 8080

# run air to watch for changes
CMD ["air"]
