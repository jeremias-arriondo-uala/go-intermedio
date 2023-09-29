package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/usuario-u-organizacion/nombre-del-repositorio/internal/processor"
	"github.com/usuario-u-organizacion/nombre-del-repositorio/internal/repository"
	"github.com/usuario-u-organizacion/nombre-del-repositorio/pkg/handler"
	"github.com/usuario-u-organizacion/nombre-del-repositorio/pkg/service"
)

func main() {
	r := repository.New()
	s := service.New(
		http.Client{},
	)
	p := processor.New(r, s)
	handler := handler.New(p)

	lambda.Start(handler.Handle)
}
