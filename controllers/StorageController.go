package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path/filepath"
)

type StorageController struct {
	beego.Controller
}

const uploadDir = "./storage"

func init() {
	// Создание директории для хранения файлов, если её нет
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
}

func (storage *StorageController) Get() {
	storage.TplName = "storage.html"
}

func (storage *StorageController) Upload() {
	file, header, err := storage.GetFile("file") // "file" — имя поля формы
	if err != nil {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Failed to retrieve file"}
		storage.ServeJSON()
		return
	}
	defer file.Close()

	// Путь для сохранения файла
	filePath := "./storage/" + header.Filename

	// Создаём файл на диске
	out, err := os.Create(filePath)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to save file"}
		storage.ServeJSON()
		return
	}
	defer out.Close()

	// Копируем содержимое загруженного файла
	_, err = io.Copy(out, file)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to write file"}
		storage.ServeJSON()
		return
	}

	// Успешный ответ
	storage.Data["json"] = map[string]string{"message": "File uploaded successfully"}
	storage.ServeJSON()
}

func (storage *StorageController) ListFiles() {
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to list files"}
		storage.ServeJSON()
		return
	}

	files := make([]string, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	storage.Data["json"] = files
	storage.ServeJSON()
}

func (storage *StorageController) Delete() {
	fileName := storage.GetString("file")
	if fileName == "" {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Missing file name"}
		storage.ServeJSON()
		return
	}

	filePath := filepath.Join(uploadDir, fileName)
	if err := os.Remove(filePath); err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to delete file"}
		storage.ServeJSON()
		return
	}

	storage.Data["json"] = map[string]string{"message": "File deleted successfully"}
	storage.ServeJSON()
}
