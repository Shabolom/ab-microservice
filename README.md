# AB Microservice

Микросервис для управления A/B-экспериментами и feature toggles.

Сервис позволяет создавать namespaces, layers, эксперименты, группы экспериментов, кастомные параметры, управлять процентом раскатки feature toggles и получать экспериментальные группы/активные фичи для пользователя.

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

### A/B-эксперименты

* создание namespace;
* создание layer внутри namespace;
* создание эксперимента;
* настройка rollout percentage;
* настройка групп эксперимента;
* фильтрация по городам;
* фильтрация по магазинам;
* фильтрация по кастомным параметрам;
* перевод эксперимента в `ready`;
* автоматическая активация эксперимента по `start_date`;
* остановка эксперимента;
* получение экспериментальной группы пользователя;
* публикация событий попадания пользователя в эксперимент в Kafka.

### Feature Toggles

* создание feature toggle;
* статусы feature toggle:

    * `draft`;
    * `active`;
    * `disabled`;
    * `archived`;
* изменение общего rollout percentage;
* изменение rollout percentage отдельно для платформ:

    * iOS;
    * Android;
    * Web;
* проверка, активна ли feature toggle;
* получение списка feature toggles, в которые попал пользователь;
* hash-based распределение пользователей по проценту раскатки.

### Observability

* HTTP health checks;
* gRPC health check;
* Kafka health check;
* Prometheus metrics;
* structured logging через Zap.

## Структура проекта

```text
.
├── build
│   ├── app/migrations        # SQL-миграции PostgreSQL
│   └── local                 # локальное окружение, docker-compose, .env
├── cmd/ab                    # точка входа приложения
├── docs/protobuff            # proto-контракт gRPC API
├── gen                       # сгенерированный protobuf/gRPC код
├── internal
│   ├── config                # конфигурация приложения
│   ├── di                    # dependency injection
│   ├── dto                   # DTO
│   ├── handler/rpctransport  # gRPC handlers
│   ├── in-memory-cache       # in-memory cache экспериментов и фич
│   ├── kafka-producer        # Kafka producer
│   ├── metrics               # Prometheus metrics
│   ├── redis                 # Redis client
│   ├── repository/pg-repo    # PostgreSQL repositories
│   ├── service               # бизнес-логика
│   └── worker                # фоновые воркеры
├── pkg                       # общие пакеты
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
  rpc SetReadyExperiment(SetReadyExperimentRequest) returns (StockReply);
  rpc SetStopedExperiment(SetStopedExperimentRequest) returns (StockReply);

  rpc CreateNamespace(CreateNamespaceRequest) returns (StockReply);
  rpc CreateLayer(CreateLayerRequest) returns (StockReply);
  rpc CreateCustomParam(CreateCustomParamRequest) returns (StockReply);

  rpc CreateFeatureToggle(CreateFeatureToggleRequest) returns (StockReply);
  rpc UpdateFeatureToggleRollout(UpdateFeatureToggleRolloutRequest) returns (StockReply);
  rpc SetFeatureToggleStatus(SetFeatureToggleStatusRequest) returns (StockReply);
  rpc IsFeatureEnabled(IsFeatureEnabledRequest) returns (StockReply);
  rpc IsUserInFeature(IsUserInFeatureRequest) returns (IsUserInFeatureReply);
}
```

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
  -database "postgres://postgres:postgres@localhost:5436/ab?sslmode=disable" \
  up
```

Запустить сервис:

```bash
go run ./cmd/ab
```

## Переменные окружения

Приложение читает конфиг из:

```text
./build/local/.env
```

Для локального PostgreSQL важно явно отключить TLS:

```env
sslmode=disable
```

Пример DSN:

```env
postgres://postgres:postgres@localhost:5436/ab?sslmode=disable
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

## Health Checks

### Liveness Probe

Проверяет, что процесс приложения запущен и отвечает на запросы.

```bash
curl http://localhost:8092/healthz
```

Ответ:

```text
ok
```

### Readiness Probe

Проверяет готовность сервиса обрабатывать запросы.

```bash
curl http://localhost:8092/readyz
```

Ответ:

```text
ready
```

Если PostgreSQL недоступен:

```text
postgres unavailable
```

HTTP status:

```text
503 Service Unavailable
```

### Kafka Health Check

Проверяет состояние Kafka Producer.

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

### gRPC Health Check

```bash
grpcurl -plaintext localhost:8015 grpc.health.v1.Health/Check
```

Ответ:

```json
{
  "status": "SERVING"
}
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

## Основная бизнес-логика A/B-эксперимента

1. Администратор создает namespace.
2. В namespace создаются layers.
3. Создается эксперимент с группами, датами, rollout percentage и фильтрами.
4. Эксперимент переводится в статус `ready`.
5. Воркер активирует эксперимент при наступлении `start_date`.
6. Пользовательский запрос приходит в `UserExperiment`.
7. Сервис проверяет namespace, bucket, device filters, city/store filters и custom params.
8. Если пользователь попадает в эксперимент, сервис возвращает название эксперимента и группу.
9. Событие попадания пользователя в эксперимент публикуется в Kafka.

## Основная бизнес-логика Feature Toggle

1. Администратор создает feature toggle.
2. По умолчанию feature toggle создается в статусе `draft`.
3. Администратор задает общий или платформенный rollout percentage.
4. Feature toggle переводится в статус `active`.
5. Пользовательский запрос приходит в `IsUserInFeature`.
6. Сервис проверяет namespace, platform и rollout percentage.
7. Если пользователь попал в процент раскатки, сервис возвращает feature toggle в списке активных фич пользователя.

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

## Пример запроса IsUserInFeature

```json
{
  "user_id": 123,
  "namespace": "mobile",
  "platform": "ios"
}
```

Пример ответа:

```json
{
  "err_info_reason": "STATUS_OK",
  "features": [
    {
      "feature_id": 1,
      "feature_name": "new_checkout"
    }
  ]
}
```

## Статусы эксперимента

* `ready` — эксперимент подготовлен к запуску;
* `active` — эксперимент активен;
* `ended` — эксперимент завершен;
* `stopped` — эксперимент остановлен вручную.

## Статусы Feature Toggle

* `draft` — фича создана, но еще не активна;
* `active` — фича активна и участвует в проверке пользователя;
* `disabled` — фича временно выключена;
* `archived` — фича архивирована и не используется.

## Замечания

Проект находится в разработке.

Перед использованием в production стоит дополнить:

* CI pipeline;
* тесты;
* документацию по Kafka-схемам;
* документацию по миграциям;
* инструкцию по деплою;
* полноценную документацию для админки;
* список всех ошибок и `err_info_reason`;
* примеры grpcurl-запросов для всех основных методов.
