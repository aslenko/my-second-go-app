# Start from golang v1.12.4 base image
FROM golang:1.12.4-alpine3.9 as builder

# Make new directory (app)
RUN mkdir /app

# Copy files from root into app
ADD . /app

# Set up working directory
WORKDIR /app/src

#build go and output to the bin folder
RUN go build -o ../bin/main.exe .

#run the executable
CMD ["/app/bin/main.exe"]