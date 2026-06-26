package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// На случай запуска сервера из разных директорий
	var html []byte
	var err error

	// Возможные пути к index.html
	paths := []string{
		"index.html",       // если запуск из корня проекта
		"../index.html",    // если запуск из папки cmd/
		"../../index.html", // если запуск из internal/handlers/
	}

	for _, path := range paths {
		html, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		http.Error(w, "Failed to load index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Получаем файл
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем файл
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Проверяем на пустоту
	if len(content) == 0 {
		http.Error(w, "File is empty", http.StatusInternalServerError)
		return
	}

	// Конвертируем
	converted, err := service.AutoConvert(string(content))
	if err != nil {
		http.Error(w, "Conversion failed", http.StatusInternalServerError)
		return
	}

	// Получаем расширение
	extension := filepath.Ext(handler.Filename)

	// Генерируем имя файла в формате: ГГГГ-ММ-ДД_ЧЧ-ММ-СС + расширение
	// Пример: "2026-06-24_15-30-45.txt"
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := "converted_" + timestamp + extension

	// Сохраняем файл
	err = os.WriteFile(filename, []byte(converted), 0755)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(converted))
}
