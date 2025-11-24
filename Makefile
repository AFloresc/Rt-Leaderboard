# Nombre del binario
BINARY=rt-leaderboard

# Carpeta principal
CMD=cmd/main.go

# Variables de Docker
DOCKER_COMPOSE=docker-compose

# -------------------------
# Comandos locales de Go
# -------------------------

## Compilar el proyecto
build:
    go build -o $(BINARY) $(CMD)

## Ejecutar el proyecto localmente
run:
    go run $(CMD)

## Ejecutar tests
test:
    go test ./... -v

## Limpiar binarios
clean:
    rm -f $(BINARY)

# -------------------------
# Comandos Docker
# -------------------------

## Levantar servicios con Docker Compose
up:
    $(DOCKER_COMPOSE) up --build -d

## Apagar servicios
down:
    $(DOCKER_COMPOSE) down

## Ver logs de la API
logs:
    $(DOCKER_COMPOSE) logs -f api

## Reconstruir imágenes
rebuild:
    $(DOCKER_COMPOSE) up --build --force-recreate -d

## Acceder al contenedor de la API
shell-api:
    docker exec -it rt-leaderboard-api sh

## Acceder al contenedor de Redis
shell-redis:
    docker exec -it rt-leaderboard-redis redis-cli
