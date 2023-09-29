package repository

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

/*
	En la declaración del struct del sistema de persistencia, tendriamos que tener
	el cliente con el cual nos conectaríamos al mismo.
	Este cliente, inicialmente lo podríamos inicializar en el constructor del Repository.

	El nombre tanto como del archivo que contiene el repository como del struct del mismo,
	deberán llevar un nombre relacionado al servicio de persistencia del cual hace uso,
	ya que podríamos tener más de uno; esto ultimo respetando las naming conventions
	tanto de archivos como de Structs.

	Ej: dynamodb.go y DynamoDBRepository
*/
type DBRepository struct {
	svc *dynamodb.Client
}

func New() DBRepository {
	return DBRepository{}
}

/*
	Los parametros de los métodos serán los necesarios, y dependiendo el caso de uso puede ser un puntero
	o no (por ejemplo, en el caso de que necesitemos devolver el objeto ya modificado en la DB)
*/
func (r DBRepository) GetShapes(o *[]string) error {

	/*
		En el caso de la devolución de errores, al final del flujo si no hubo ningún error
		devolveremos `nil`, ya que aunque el `err` pueda llegar vacío a esta instancia
		hace mas amigable la legibilidad del código.
	*/
	return nil

}
