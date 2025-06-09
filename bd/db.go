   package bd

   import (
       "database/sql"
       "fmt"
       _ "github.com/lib/pq"
   )

   // Connect initializes and returns a database connection.
   func Connect() (*sql.DB, error) {
       db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=1234 dbname=ProjectCom sslmode=disable")
       if err != nil {
           return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
       }

       err = db.Ping()
       if err != nil {
           return nil, fmt.Errorf("ошибка соединения с базой данных: %v", err)
       }

       fmt.Println("Успешное подключение к базе данных.")
       return db, nil
   }
