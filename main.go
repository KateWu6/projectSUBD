package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/KateWu6/projectSUBD/config"
    "github.com/KateWu6/projectSUBD/controllers/auth"
    "github.com/KateWu6/projectSUBD/controllers/profile"
    "github.com/KateWu6/projectSUBD/controllers/project"
    "github.com/KateWu6/projectSUBD/controllers/task"
    "github.com/KateWu6/projectSUBD/controllers/user"
    "github.com/KateWu6/projectSUBD/models"
)

// Config представляет конфигурационные данные
type Config struct {
    DBHost     string `json:"db_host"`
    DBPort     string `json:"db_port"`
    DBUser     string `json:"db_user"`
    DBPassword string `json:"db_password"`
    DBName     string `json:"db_name"`
}

// Инициализация базы данных
func InitDB(cfg Config) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Ошибка подключения к базе данных: %v", err)
    }

    // Автомиграция таблиц (создается схема базы данных)
    models.AutoMigrate(db)

    return db
}

// Главная функция приложения
func main() {
    // Читаем конфигурацию из файла config.json
    cfg := config.LoadConfig()

    // Инициализируем подключение к базе данных
    db := InitDB(cfg)
    defer db.Close()

    // Создаем роутер
    router := mux.NewRouter()

    // Регистрируем маршруты
    auth.SetupRoutes(router, db)
    profile.SetupRoutes(router, db)
    user.SetupRoutes(router, db)
    project.SetupRoutes(router, db)
    task.SetupRoutes(router, db)

    // Статические файлы (CSS, картинки)
    staticDir := http.Dir("static/")
    fileServer := http.FileServer(staticDir)
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

    // Запускаем сервер
    log.Println("Запуск сервера на порте :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}