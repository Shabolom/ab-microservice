# AB Microservice

Микросервис для управления A/B-экспериментами.

Сервис позволяет создавать namespaces, layers, эксперименты, группы экспериментов, кастомные параметры и получать экспериментальную группу для пользователя по `split_id`, `device_id`, namespace и параметрам устройства.

## Стек

* Go
* gRPC
* PostgreSQL
* Redis
* Kafka
* Schema Registry
* Prometheus
* Zap Logger
* Docker Compose

## Возможности

* создание namespace;
* создание layer внутри namespace;
* создание эксперимента;
* настройка rollout percentage;
* настройка групп эксперимента;
* фильтрация по городам и магазинам;
* фильтрация по кастомным параметрам;
* перевод эксперимента в ready;
* остановка эксперимента;
* получение экспериментов пользователя;
* публикация событий попадания пользователя в эксперимент в Kafka;
* сбор Prometheus-метрик.

## Структура проекта

```text
.
├── build
│   ├── app/migrations      # SQL-миграции PostgreSQL
│   └── local               # локальное окружение, docker-compose, .env
├── cmd/ab                  # точка входа приложения
├── docs/protobuff          # proto-контракт gRPC API
├── gen                     # сгенерированный protobuf/gRPC код
├── internal
│   ├── config              # конфигурация приложения
│   ├── di                  # dependency injection
│   ├── dto                 # DTO
│   ├── handler             # gRPC handlers
│   ├── in-memory-cache     # in-memory cache экспериментов
│   ├── kafka-producer      # Kafka producer
│   ├── metrics             # Prometheus metrics
│   ├── redis               # Redis client
│   ├── repository          # PostgreSQL repositories
│   ├── service             # бизнес-логика
│   └── worker              # фоновые воркеры
├── pkg                     # общие пакеты
├── Makefile
├── go.mod
└── README.md
```

## gRPC API

Основной сервис:

```proto
service ABExperiment {
  rpc UserExperiment(ExperimentRequest) returns (ExperimentsReply);
  rpc CreateExperiment(CreateExperimentRequest) returns (StockReply);
  rpc CreateNamespace(CreateNamespaceRequest) returns (StockReply);
  rpc CreateLayer(CreateLayerRequest) returns (StockReply);
  rpc SetReadyExperiment(SetReadyExperimentRequest) returns (StockReply);
  rpc SetStopedExperiment(SetStopedExperimentRequest) returns (StockReply);
  rpc CreateCustomParam(CreateCustomParamRequest) returns (StockReply);
}
```

## Переменные окружения

Приложение читает конфиг из `./build/local/.env`.

## Локальный запуск

Клонировать репозиторий:

```bash
git clone https://github.com/Shabolom/ab-microservice.git
cd ab-microservice
```

Запустить инфраструктуру:

```bash
docker compose -f build/local/docker-compose.yml up -d
```

Установить зависимости:

```bash
go mod download
```

Применить миграции:

```bash
migrate -path build/app/migrations \
  -database "postgres://postgres:postgres@localhost:5432/ab?sslmode=disable" \
  up
```

Запустить сервис:

```bash
go run ./cmd/ab
```

## Генерация protobuf

```bash
make proto
```

Команда генерирует Go-код из файла:

```text
docs/protobuff/ab-microservice.proto
```

в директорию:

```text
gen
```

## Метрики

Prometheus-метрики доступны на порту из переменной:

```env
PROMETHEUS_PORT=2112
```

Пример:

```text
http://localhost:2112/metrics
```

## Основная бизнес-логика

1. Администратор создает namespace.
2. В namespace создаются layers.
3. Создается эксперимент с группами, датами, rollout percentage и фильтрами.
4. Эксперимент переводится в статус `ready`.
5. Воркер активирует эксперимент при наступлении `start_date`.
6. Пользовательский запрос приходит в `UserExperiment`.
7. Сервис проверяет namespace, bucket, device filters и custom params.
8. Если пользователь попадает в эксперимент, сервис возвращает название эксперимента и группу.
9. Событие попадания пользователя в эксперимент публикуется в Kafka.

## Пример запроса UserExperiment

```json
{
  "split_id": 123,
  "device_id": 1,
  "namespace": "mobile",
  "city": "Moscow",
  "store": "1001",
  "params": [
    {
      "param_name": "app_version",
      "value": "1.2.3"
    }
  ]
}
```

Пример ответа:

```json
{
  "err_info_reason": "STATUS_OK",
  "experiments_reply": [
    {
      "experiment_name": "new_main_screen",
      "group_name": "test_group"
    }
  ]
}
```

## Статусы эксперимента

* `ready` — эксперимент подготовлен к запуску;
* `active` — эксперимент активен;
* `ended` — эксперимент завершен;
* `stopped` — эксперимент остановлен вручную.

## Health Checks

Сервис предоставляет HTTP и gRPC health check endpoints для мониторинга состояния приложения и его зависимостей.

### HTTP Endpoints

#### Liveness Probe

Проверяет, что процесс приложения запущен и отвечает на запросы.

```bash
curl http://localhost:8092/healthz
```

Ответ:

```text
ok
```

---

#### Readiness Probe

Проверяет готовность сервиса обрабатывать запросы.

В текущей реализации проверяется доступность PostgreSQL.

```bash
curl http://localhost:8092/readyz
```

Ответ:

```text
ready
```

При недоступности PostgreSQL:

```text
postgres unavailable
```

HTTP статус:

```text
503 Service Unavailable
```

---

#### Kafka Health Check

Проверяет состояние Kafka Producer и возможность взаимодействия с Kafka кластером.

```bash
curl http://localhost:8092/healthz/kafka
```

Ответ:

```text
kafka ok
```

При ошибке инициализации или недоступности Kafka возвращается:

```text
503 Service Unavailable
```

---

### gRPC Health Check

Сервис поддерживает стандартный gRPC Health Checking Protocol.

Проверка состояния сервиса:

```bash
grpcurl -plaintext localhost:8015 grpc.health.v1.Health/Check
```

Ответ:

```json
{
  "status": "SERVING"
}
```

---

### Проверка всех health endpoints

```bash
curl http://localhost:8092/healthz
curl http://localhost:8092/readyz
curl http://localhost:8092/healthz/kafka
grpcurl -plaintext localhost:8015 grpc.health.v1.Health/Check
```

Ожидаемый результат:

```text
ok
ready
kafka ok
SERVING
```

## Замечания

Проект находится в разработке. Перед использованием в production стоит дополнить:

* CI pipeline;
* тесты;
* OpenAPI/gRPC examples;
* документацию по Kafka-схемам;
* документацию по миграциям;
* инструкцию по деплою.
