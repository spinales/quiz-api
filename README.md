Quiz API(IDS340 Proyecto Final)

## Descipcion del proyecto

Quiz es el proyecto final de la materia IDS340, este repositorio corresponde al backend de la aplicacion hecho en Go(golang) como lenguaje de programacion, PostgreSQL como motor de base de datos, este proyecto es un API REST, la API se encarga d ela parte de authentificacion y manejo de la informacion disponible en la base de datos.

## Requisitos del proyecto

- Go 1.17.2 o superior.
- PostgreSQL 10 o superior.

### Como correr el proyecto

1. Clonar el repositorio en tu pc.
2. Luego colocarse(acceder) en el directorio del proyecto.
3. copiar el archivo `.env-example` con el nombre de `.env` y colocar valores a las variables de entorno.
4. Que corresponde cada valor:
  ```
  - DB_USER el usuario en la base de datos.
  - DB_PASSWORD la contraseña del usuario de la base de datos.
  - DB_PORT el puerto que corre el motor de base de datos.(5432)
  - DB_HOST El host o el servidor que esta corriendo el motor base de datos.
  - DB_NAME El nombre de la base de datos
  - Server_Address :3001 puerto que estara corriendo la API, recomendable: 3001.
  - Token_Symmetric_Key la clase simetrica para los tokens. cualquier string pero con una longitud de 32 caracteres exactos.
  - Access_Token_Duration la duracion que tomara para que el token expire, colocar la cantidad en horas, solamente el numero.
  ```
5. Luego de concluir de llenar el archivo `.env` con los valores correspondientes, puede ejecutar el proyecto con el siguiente comando: 
  - ```go run cmd/server/quiz.go```.

## Front-end proyect

El front end del proyecto corresponde al siguiente link: [frontend](https://github.com/Cristofers/quiz-master-front/tree/master).

## License

Consulte el archivo de [LICENCIA](https://github.com/spinales/quiz-api/blob/master/LICENSE.txt) para conocer los derechos y limitaciones de la licencia.

## Cosas por hacer(Todo)

**Ignorar este reglon completamente.**

- [x] Definir base de datos
- [x] Establecer relaciones
- Crear modelos.
    - [x] Question.
    - [x] Answer.
    - [x] User.
    - [x] Registro de puntuacion.
- Handlers.
    - [x] Middleware
    - [x] Registrar nuevo usuario.
        - [x] Metodo en la para agregar un nuevo usuario en la bd.
        - [x] Encryptar password.
        - [x] Validar Email.
    - [x] Login
        - [x] Validar password.
    - [x] Rol de administrador.
        - [x] Agregar pregunta.
        - [x] Actualizar pregunta.
        - [x] Eliminar pregunta.
        - [x] Visualizar preguntas.
        - [x] Agregar respuesta.
        - [x] Actualizar respuesta.
        - [x] Eliminar respuesta.
        - [x] Visualizar respuesta.
    - Rol de jugador
        - [x] Jugar una partida. Del front me enviaran un arreglo con cada ID de la pregunta, con el ID de la respuesta que selecciono el usuario, y el valor en puntos de la respuesta.
        - [x] Registrar puntuaciones, registro la puntuacion como un todo, Sumo todos los puntos, con el id del usuario y un id unico.
- Metodos de la bd.
    - [x] Crud pregunta.
    - [x] Crud Respuesta.
    - [x] Agregar usuario
    - [x] Consultar usuario.
    - [x] Consultar respuestas por id pregunta
- [ ] Logger
- [ ] Documentation
- [x] Final readme version

