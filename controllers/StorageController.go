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
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
}

func (c *StorageController) Get() {
	username := c.GetSession("username")
	if username == nil {
		username = "Guest"
	}
	c.Data["Username"] = username
	c.TplName = "storage.html"
}

func (storage *StorageController) Upload() {
	// Get username from session
	username := storage.GetSession("username")
	if username == nil {
		storage.Ctx.Output.SetStatus(401)
		storage.Data["json"] = map[string]string{"error": "Unauthorized"}
		storage.ServeJSON()
		return
	}

	file, header, err := storage.GetFile("file")
	if err != nil {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Failed to retrieve file"}
		storage.ServeJSON()
		return
	}
	defer file.Close()

	// Create user-specific directory
	userDir := filepath.Join(uploadDir, username.(string))
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to create user directory"}
		storage.ServeJSON()
		return
	}

	// Save file in user's directory
	filePath := filepath.Join(userDir, header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to save file"}
		storage.ServeJSON()
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to write file"}
		storage.ServeJSON()
		return
	}

	storage.Data["json"] = map[string]string{"message": "File uploaded successfully"}
	storage.ServeJSON()
}

func (storage *StorageController) ListFiles() {
	// Get username from session
	username := storage.GetSession("username")
	if username == nil {
		storage.Ctx.Output.SetStatus(401)
		storage.Data["json"] = map[string]string{"error": "Unauthorized"}
		storage.ServeJSON()
		return
	}

	// Get files from user's directory
	userDir := filepath.Join(uploadDir, username.(string))
	entries, err := os.ReadDir(userDir)
	if err != nil && !os.IsNotExist(err) {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to list files"}
		storage.ServeJSON()
		return
	}

	files := make([]string, 0)
	if err == nil { // Directory exists
		for _, entry := range entries {
			if !entry.IsDir() {
				files = append(files, entry.Name())
			}
		}
	}

	storage.Data["json"] = files
	storage.ServeJSON()
}

func (storage *StorageController) Delete() {
	// Get username from session
	username := storage.GetSession("username")
	if username == nil {
		storage.Ctx.Output.SetStatus(401)
		storage.Data["json"] = map[string]string{"error": "Unauthorized"}
		storage.ServeJSON()
		return
	}

	fileName := storage.GetString("file")
	if fileName == "" {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Missing file name"}
		storage.ServeJSON()
		return
	}

	// Delete file from user's directory
	filePath := filepath.Join(uploadDir, username.(string), fileName)
	if err := os.Remove(filePath); err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Failed to delete file"}
		storage.ServeJSON()
		return
	}

	storage.Data["json"] = map[string]string{"message": "File deleted successfully"}
	storage.ServeJSON()
}

func (storage *StorageController) View() {
	username := storage.GetSession("username")
	if username == nil {
		storage.Ctx.Output.SetStatus(401)
		storage.Data["json"] = map[string]string{"error": "Не авторизован"}
		storage.ServeJSON()
		return
	}

	fileName := storage.GetString("file")
	if fileName == "" {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Имя файла не указано"}
		storage.ServeJSON()
		return
	}

	filePath := filepath.Join(uploadDir, username.(string), fileName)

	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		storage.Ctx.Output.SetStatus(404)
		storage.Data["json"] = map[string]string{"error": "Файл не найден"}
		storage.ServeJSON()
		return
	}

	// Отправляем файл пользователю
	storage.Ctx.Output.Download(filePath, fileName)
}

func (storage *StorageController) Edit() {
	username := storage.GetSession("username")
	if username == nil {
		storage.Ctx.Output.SetStatus(401)
		storage.Data["json"] = map[string]string{"error": "Не авторизован"}
		storage.ServeJSON()
		return
	}

	fileName := storage.GetString("file")
	content := storage.GetString("content")

	if fileName == "" {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Имя файла не указано"}
		storage.ServeJSON()
		return
	}

	filePath := filepath.Join(uploadDir, username.(string), fileName)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Ошибка сохранения файла"}
		storage.ServeJSON()
		return
	}

	storage.Data["json"] = map[string]string{"message": "Файл успешно сохранен"}
	storage.ServeJSON()
}
