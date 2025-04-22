# API REST con Go y Gin

Este proyecto es una API REST desarrollada con Go y el framework Gin para gestionar información de administradores, estudiantes y materias en un sistema académico.

## Estructura del Proyecto

```
mi-api-rest/
├── cmd/                     # Punto de entrada de la aplicación
├── internal/                # Código interno de la aplicación 
│   ├── api/                 # Componentes de la API (handlers, middlewares, routes)
│   ├── models/              # Modelos de datos
│   ├── repository/          # Capa de acceso a datos
│   └── service/             # Lógica de negocio
├── pkg/                     # Paquetes reutilizables
│   ├── config/              # Configuración
│   ├── database/            # Conexión a la base de datos
│   └── utils/               # Utilidades
└── ...
```

## Modelos de Datos

### Administradores
Representan a los usuarios que pueden iniciar sesión en el sistema.
- ID
- Username (único)
- Password (hasheado)

### Estudiantes
Representan a los estudiantes registrados en el sistema.
- ID
- Code (código único)
- Name (nombre completo)
- Subjects (relación muchos a muchos con materias)

### Materias (Subjects)
Representan las asignaturas disponibles.
- ID
- Code (código único)
- Name (nombre)
- IsElective (indica si es una materia electiva)
- Group (número de grupo)
- Credits (créditos)
- Students (relación muchos a muchos con estudiantes)

## Requisitos

- Go 1.21 o superior
- PostgreSQL
- Git

## Configuración

1. Clona el repositorio:
   ```
   git clone github.com/tuusuario/mi-api-rest
   cd mi-api-rest
   ```

2. Copia el archivo `.env.example` a `.env` y configura las variables:
   ```
   cp .env.example .env
   ```

3. Instala las dependencias:
   ```
   go mod download
   ```

4. Crea la base de datos en PostgreSQL:
   ```sql
   CREATE DATABASE miapi;
   ```

## Ejecución

Para ejecutar el servidor en modo desarrollo:

```
go run cmd/api/main.go
```

Para compilar el proyecto:

```
go build -o server cmd/api/main.go
./server
```

## API Endpoints

### Administradores

- `POST /api/v1/auth/register`: Registrar un nuevo administrador
- `POST /api/v1/auth/login`: Iniciar sesión y obtener un token JWT
- `GET /api/v1/administrators`: Obtener todos los administradores
- `GET /api/v1/administrators/:id`: Obtener un administrador por ID
- `PUT /api/v1/administrators/:id`: Actualizar un administrador
- `DELETE /api/v1/administrators/:id`: Eliminar un administrador

### Estudiantes

- `POST /api/v1/students`: Crear un nuevo estudiante
- `GET /api/v1/students`: Obtener todos los estudiantes
- `GET /api/v1/students/:id`: Obtener un estudiante por ID
- `GET /api/v1/students/:id/subjects`: Obtener un estudiante con sus materias
- `PUT /api/v1/students/:id`: Actualizar un estudiante
- `DELETE /api/v1/students/:id`: Eliminar un estudiante
- `POST /api/v1/students/:id/subjects`: Añadir una materia a un estudiante
- `DELETE /api/v1/students/:id/subjects/:subject_id`: Eliminar una materia de un estudiante

### Materias

- `POST /api/v1/subjects`: Crear una nueva materia
- `GET /api/v1/subjects`: Obtener todas las materias
- `GET /api/v1/subjects/:id`: Obtener una materia por ID
- `GET /api/v1/subjects/:id/students`: Obtener una materia con sus estudiantes
- `PUT /api/v1/subjects/:id`: Actualizar una materia
- `DELETE /api/v1/subjects/:id`: Eliminar una materia
- `GET /api/v1/subjects/electives`: Obtener todas las materias electivas
- `GET /api/v1/subjects/group/:group`: Obtener todas las materias de un grupo específico

## Autenticación

Todas las rutas (excepto el registro y login) requieren autenticación usando un token JWT.
Para autenticarse, incluya el token en el encabezado de la solicitud:

```
Authorization: Bearer <token>
```

## Contribuir

1. Haz fork del repositorio
2. Crea una rama para tu función (`git checkout -b feature/amazing-feature`)
3. Haz commit de tus cambios (`git commit -m 'Add some amazing feature'`)
4. Sube los cambios a tu rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

## Licencia

Este proyecto está licenciado bajo [MIT License](LICENSE).