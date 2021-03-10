FROM golang:1.14

# WORKDIR /root/projects/src/github.com/239103/random_password_generator
# COPY . .

RUN go get -d -v github.com/239103/random_password_generator
RUN go install -v github.com/239103/random_password_generator

CMD ["random_password_generator"]