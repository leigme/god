FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY {{.goEnv}}
{{range $k, $v := .runs}}
RUN $run{{$k}}
{{end}}
WORKDIR /build

{{if .goMod}}
ADD go.mod .
{{end}}
{{if .goSum}}
ADD go.sum .
{{end}}
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/{{.appName}} ./{{.fileName}}

FROM scratch

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/{{.appName}} /app/{{.appName}}

CMD ["./{{.appName}}"]