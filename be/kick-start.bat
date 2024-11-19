@echo off
set PROJECT_NAME=go-ddd

:: Create main project directory
mkdir %PROJECT_NAME%
cd %PROJECT_NAME%

:: Create main directory structure
mkdir cmd\api
mkdir internal\domain\entity
mkdir internal\domain\repository
mkdir internal\infrastructure\db
mkdir internal\infrastructure\migrations
mkdir internal\interfaces\http\handlers
mkdir internal\interfaces\http\routes
mkdir internal\application\services
mkdir pkg\utils

:: Create initial files
type nul > cmd\api\main.go
type nul > internal\domain\entity\user.go
type nul > internal\domain\repository\user_repository.go
type nul > internal\infrastructure\db\postgres.go
type nul > internal\infrastructure\migrations\000001_create_users_table.up.sql
type nul > internal\interfaces\http\handlers\auth_handler.go
type nul > internal\interfaces\http\routes\routes.go
type nul > internal\application\services\auth_service.go
type nul > pkg\utils\password.go
type nul > .env
type nul > go.mod

echo Project structure created successfully!