package page_go

import (
    "html/template"
    "log"
    "net/http"
    "projectSUBD/bd"
    "strings"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserExists(username string) (*bd.User, error) {
    db, err := bd.Connect()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    var user bd.User
    err = db.QueryRow("SELECT username, hashed_password FROM users WHERE username = $1", username).Scan(&user.Username, &user.HashedPassword)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, err := template.ParseFiles("templates/login.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
        return
    }

    r.ParseForm()
    username := strings.TrimSpace(r.FormValue("username"))
    password := strings.TrimSpace(r.FormValue("password"))

    // Проверяем существование пользователя
    user, err := CheckUserExists(username)
    if err != nil {
        log.Println("Ошибка обращения к базе данных:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Проверяем введенный пароль против сохраненного хэша в базе данных
    match, err := ComparePasswordWithHash(password, user.HashedPassword)
    if err != nil || !match {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Неправильный пароль."))
        return
    }

    // Авторизация успешна, перенаправляем на домашнюю страницу
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Проверка совпадения пароля
func ComparePasswordWithHash(password string, hashedPasswd []byte) (bool, error) {
    // Проверяем пароль с использованием стандартного пакета bcrypt
    err := bcrypt.CompareHashAndPassword(hashedPasswd, []byte(password))
    if err != nil {
        return false, err
    }
    return true, nil
}

// Процесс хэширования пароля
func GenerateBcryptHash(password string) ([]byte, error) {
    // Используем стандартный уровень сложности bcrypt (Cost=10)
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

