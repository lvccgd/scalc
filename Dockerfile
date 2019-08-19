FROM golang:1.10
# Install dependencies
RUN apt-get update && apt-get install -y tzdata zip ca-certificates make git
# Make repository path
RUN mkdir -p /usr/src/scalc
WORKDIR /usr/src/scalc
# Copy all files
COPY . .
# Build and check
RUN make build
RUN make test

WORKDIR /usr/src/scalc/bin
CMD [ "./scalc", "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]" ]