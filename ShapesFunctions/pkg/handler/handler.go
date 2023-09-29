package handler

import (
	"context"
	"net/http"

	"github.com/usuario-u-organizacion/nombre-del-repositorio/pkg/models"
)

/*
	Declaramos la interfaz en el lugar donde vamos a necesitar utilizar un elemento
	que cumpla con estos mínimos requisitos. Esto nos va a facilitar el dia de mañana si
	tenemos que hacer algún cambio en el codigo del elemento externo, los cambios no se propaguen
	en extensión a todos los lugares donde es utilizado.
*/
type Processor interface {
	Process([]string) error
}

type Handler struct {
	p Processor
}

/*
	En nuestro constructor en lugar de inicializar sus atributos en el literal que devolvemos,
	los pasaremos como parametros, para facilitar las tareas de testeo.
*/

func New(p Processor) Handler {
	return Handler{
		p,
	}
}

/*
	Las variantes de la firma del método Handle, vienen definida desde el método `lambda.Start()`
	tanto en sus parametros como en sus valores de retorno. En este caso tenemos uno genérico,
	pero también podríamos retornar solamente un `error`, o aceptar como parámetro un evento
	del tipo `APIGatewayProxyRequest`.

	Nuestro Handler va a tener unicamente la responsabilidad de tomar el evento que dispare
	la lambda en cuestión, manipular ese evento para extraer la información necesaria, y finalmente
	construir lo que va a ser la estructura/valor de retorno de nuestra lambda. El procesamiento
	de dicha información (ya sea crear, destruir, modificar, etc) lo delegará en el Processor.
*/
func (h Handler) Handle(ctx context.Context, request models.Request) (models.Response, error) {

	objectsToProcess := request.GenericArray

	err := h.p.Process(objectsToProcess)

	if err != nil {
		return models.Response{}, err
	}
	return models.Response{
		StatusCode: http.StatusOK,
	}, nil
}
