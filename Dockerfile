# FROM golang:1.16-stretch AS builder
# WORKDIR /app
# ARG GIT_USER
# ARG GIT_PASSWORD
# COPY . .
# RUN go build .

# FROM debian:10.9
# WORKDIR /app
# ENV CONFIG_PATH=/app
# EXPOSE 8090/tcp
# RUN useradd -u 1000 -ms /bin/bash app
# COPY --from=builder /app/test_conn .
# COPY --from=builder /app/deployment/config ./deployment/config
# RUN chown -R app /app
# USER app
# CMD ["./test_conn", "server"]


FROM debian:10.9
WORKDIR /app
ENV CONFIG_PATH=/app
EXPOSE 8090/tcp
RUN useradd -u 1000 -ms /bin/bash app
COPY ./test_conn .
COPY ./deployment/config ./deployment/config
RUN chown -R app /app
USER app
CMD ["./test_conn", "server"]
