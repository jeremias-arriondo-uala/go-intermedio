package processor

import (
	"fmt"
)

/*
	Declaramos la interfaz en el lugar donde vamos a necesitar utilizar un elemento
	que cumpla con estos mínimos requisitos. Esto nos va a facilitar el dia de mañana si
	tenemos que hacer algún cambio en el codigo del elemento externo, los cambios no se propaguen
	en extensión a todos los lugares donde es utilizado.
*/

type Repository interface {
	GetShapes(*[]string) error
}

type Service interface {
	ActionMethod([]string) error
}

type Processor struct {
	r Repository
	s Service
}

/*
	En nuestro constructor en lugar de inicializar sus atributos en el literal que devolvemos,
	los pasaremos como parametros, para facilitar las tareas de testeo.
*/

func New(r Repository, s Service) Processor {
	return Processor{
		r,
		s,
	}
}

/*
	El método process se ajustará a las necesidades del caso de uso en tanto los parametros que
	acepta, como los tipos de valores que retorne.

	Su responsabilidad será llamar a las entidades necesarias para procesar la información que recibe
	y brindar una respuesta acorde tanto en el caso de éxito como de falla.
*/
func (p Processor) Process(objects []string) error {

	sResponse := p.s.ActionMethod(objects)

	rResponse := p.r.GetShapes(&objects)

	fmt.Printf("Service response: %v", sResponse)
	fmt.Printf("Repository response: %v", rResponse)

	return nil
}
