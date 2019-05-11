FROM golang:alpine as source
WORKDIR /home/server
COPY . .
WORKDIR cmd/watermarking-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o watermarking-service

FROM alpine as runner
LABEL name="bitstored/watermarking-service"
RUN apk --update add ca-certificates
COPY --from=source /home/server/cmd/watermarking-service/watermarking-service /home/watermarking-service
COPY --from=source /home/server/scripts/localhost.* /home/scripts/
WORKDIR /home
EXPOSE 4009
EXPOSE 5009
ENTRYPOINT [ "./watermarking-service" ]
