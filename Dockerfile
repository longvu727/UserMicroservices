FROM golang:1.22.4-alpine3.20 AS build
RUN apk add --no-cache git \
                openssh-client \
                ca-certificates

RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"

RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan gitlab.com > /root/.ssh/known_hosts &&\
    chmod 644 /root/.ssh/known_hosts && touch /root/.ssh/config \
    && echo "StrictHostKeyChecking no" > /root/.ssh/config

ENV GOOS=linux GOARCH=amd64

WORKDIR /api

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags "-s -w" -o api main.go

FROM alpine:3.20 AS runtime

WORKDIR /api

ENV USER=longvu727 UID=1000
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY --from=build --chown=${USER}:${USER} /api/ .

USER ${USER}:${USER}

HEALTHCHECK --interval=30s --timeout=30s \
    CMD wget -nv -t1 --spider 'http://localhost:3001' || exit 1

EXPOSE 3001

CMD ["./api"]
