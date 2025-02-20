FROM golang:1.22 AS build

ENV BIN_FILE_CALENDAR /opt/calendar/calendar
ENV BIN_FILE_SCHEDULER /opt/calendar/scheduler
ENV BIN_FILE_SENDER /opt/calendar/sender

ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE_CALENDAR} cmd/calendar/*

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE_SCHEDULER} cmd/scheduler/*

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE_SENDER} cmd/sender/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="cepmapp@gmail.com"

ENV BIN_FILE_CALENDAR /opt/calendar/calendar
ENV BIN_FILE_SCHEDULER /opt/calendar/scheduler
ENV BIN_FILE_SENDER /opt/calendar/sender

COPY --from=build ${BIN_FILE_CALENDAR} ${BIN_FILE_CALENDAR}
COPY --from=build ${BIN_FILE_SCHEDULER} ${BIN_FILE_SCHEDULER}
COPY --from=build ${BIN_FILE_SENDER} ${BIN_FILE_SENDER}

ENV CONFIG_FILE_CALENDAR /etc/calendar/calendar.yaml
ENV CONFIG_FILE_SCHEDULER /etc/calendar/scheduler.yaml
ENV CONFIG_FILE_SENDER /etc/calendar/sender.yaml

COPY ./configs/calendar.yaml ${CONFIG_FILE_CALENDAR}
COPY ./configs/scheduler.yaml ${CONFIG_FILE_SCHEDULER}
COPY ./configs/sender.yaml ${CONFIG_FILE_SENDER}


CMD ${BIN_FILE_CALENDAR} --config ${CONFIG_FILE_CALENDAR}
