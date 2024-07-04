FROM golang:1.22 as build
WORKDIR /cmd
ENV GOPROXY=direct
COPY go.mod go.sum ./
COPY cmd/main.go .
RUN go build -tags lambda.norpc -o main main.go
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /cmd/main ./main
ENTRYPOINT [ "./main" ]