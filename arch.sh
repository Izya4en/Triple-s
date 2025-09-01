#!/bin/bash

# Создаём директории
mkdir -p cmd
mkdir -p internal/presentation/http/{handlers,routes,middleware}
mkdir -p internal/business/{services,models}
mkdir -p internal/data/{repositories,database/migrations}
mkdir -p internal/infrastructure/{config,logger,external}
mkdir -p pkg/utils
mkdir -p configs

# Создаём основные файлы
touch cmd/main.go
touch internal/presentation/http/handlers/user_handler.go
touch internal/presentation/http/handlers/order_handler.go
touch internal/presentation/http/routes/routes.go
touch internal/presentation/http/middleware/auth.go
touch internal/presentation/http/middleware/logger.go
touch internal/business/services/user_service.go
touch internal/business/services/order_service.go
touch internal/business/models/user.go
touch internal/business/models/order.go
touch internal/data/repositories/user_repository.go
touch internal/data/repositories/order_repository.go
touch internal/data/database/connection.go
touch internal/data/database/migrations/.gitkeep
touch internal/infrastructure/config/config.go
touch internal/infrastructure/logger/logger.go
touch internal/infrastructure/external/payment.go
touch pkg/utils/validator.go
touch go.mod
touch go.sum
touch configs/app.yaml
touch README.md

echo "Структура проекта успешно создана в текущей директории!"
