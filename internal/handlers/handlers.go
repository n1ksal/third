package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// MainHandler обрабатывает запросы на главной странице и возвращает содержимое index.html.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	filePath, err := filepath.Abs(filepath.Join("..", "index.html"))
	if err != nil {
		http.Error(w, "Ошибка в расположении файла", http.StatusInternalServerError)
		log.Println("Ошибка в расположении файла:", err)
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println("Ошибка загрузки страницы:", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// UploadHandler обрабатывает загрузку файла и выполняет его преобразование в текст или Морзе.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Увеличить лимит размера загружаемого файла.
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Ошибка разбора формы", http.StatusBadRequest)
		log.Println("Ошибка разбора формы:", err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		log.Println("Ошибка получения файла:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		log.Println("Ошибка чтения файла:", err)
		return
	}

	// Преобразование содержимого файла (текст <-> код Морзе).
	converter := service.Converter{}
	result, err := converter.Convert(string(content))
	if err != nil {
		http.Error(w, "Ошибка во время преобразования", http.StatusInternalServerError)
		log.Println("Ошибка во время преобразования:", err)
		return
	}

	// Создание директории uploads, если она не существует.
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		http.Error(w, "Ошибка создания директории uploads", http.StatusInternalServerError)
		log.Println("Ошибка создания директории uploads:", err)
		return
	}

	// Генерация имени файла с результатом.
	fileName := filepath.Join("uploads", time.Now().UTC().Format("2006-01-02_15-04-05")+".txt")
	if err := os.WriteFile(fileName, []byte(result), 0644); err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		log.Println("Ошибка записи в файл:", err)
		return
	}

	// Успешное завершение
	w.Write([]byte("Результат преобразования сохранён в: " + fileName))
}
