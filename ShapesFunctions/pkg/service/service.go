package service

import (
	"net/http"
)

/*
	En la mayoria de los casos nuestra interfaz de servicio va a interactuar
	con API Rest, por eso mismo es que en el ejemplo utilizamos la librería
	`http`
*/
type ExternalService struct {
	h http.Client
}

func New(h http.Client) ExternalService {
	return ExternalService{
		h,
	}
}

/*
	De ser necesario, nuestros metodos de servicio deberian devolver el
	status code de la respuesta http que reciba.
*/

func (s ExternalService) ActionMethod(objects []string) error {
	return nil
}
