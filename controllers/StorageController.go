package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
	"path/filepath"
	"web_cloud_storage/models"
)

type StorageController struct {
	beego.Controller
}

const (
	uploadDir = "./storage"
	GB        = 1024 * 1024 * 1024 // 1 GB in bytes
	TB        = GB * 1024          // 1 TB in bytes
)

var tariffLimits = map[int]int64{
	1: 10 * GB,  // Base - 10GB
	2: 100 * GB, // Pro - 100GB
	3: 1 * TB,   // Ultra - 1TB
}

func init() {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
}

func (storage *StorageController) Get() {
	storage.TplName = "storage.html"
}

func (storage *StorageController) Upload() {
	beego.Info("Начало загрузки файла")
	username := storage.GetSession("username")
	if username == nil {
		storage.Data["json"] = map[string]string{"error": "Не авторизован"}
		storage.ServeJSON()
		return
	}

	o := orm.NewOrm()
	var user models.Users
	err := o.QueryTable("users").Filter("login", username).One(&user)
	if err != nil {
		beego.Error("Ошибка получения информации о пользователе:", err)
		storage.Data["json"] = map[string]string{"error": "Ошибка получения информации о пользователе"}
		storage.ServeJSON()
		return
	}

	beego.Info(fmt.Sprintf("User: %+v", user))

	storageLimit, ok := tariffLimits[user.TariffID]
	if !ok {
		beego.Error(fmt.Sprintf("Неизвестный тариф: %d", user.TariffID))
		// Set a default limit if the tariff is unknown
		storageLimit = tariffLimits[1] // Use the base tariff as default
	}

	userDir := filepath.Join(uploadDir, username.(string))
	var usedSpace int64 = 0

	err = filepath.Walk(userDir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			usedSpace += info.Size()
		}
		return nil
	})

	file, header, err := storage.GetFile("file")
	if err != nil {
		beego.Error("Ошибка получения файла:", err)
		storage.Data["json"] = map[string]string{"error": "Ошибка получения файла"}
		storage.ServeJSON()
		return
	}
	defer file.Close()

	beego.Info(fmt.Sprintf("Текущее использование: %d, Размер файла: %d, Лимит: %d", usedSpace, header.Size, storageLimit))

	if usedSpace+header.Size > storageLimit {
		storage.Data["json"] = map[string]string{"error": "Превышен лимит хранилища для вашего тарифа"}
		storage.ServeJSON()
		return
	}

	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		beego.Error("Ошибка создания директории:", err)
		storage.Data["json"] = map[string]string{"error": "Ошибка создания директории"}
		storage.ServeJSON()
		return
	}

	filePath := filepath.Join(userDir, header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		beego.Error("Ошибка сохранения файла:", err)
		storage.Data["json"] = map[string]string{"error": "Ошибка сохранения файла"}
		storage.ServeJSON()
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		beego.Error("Ошибка записи файла:", err)
		storage.Data["json"] = map[string]string{"error": "Ошибка записи файла"}
		storage.ServeJSON()
		return
	}

	beego.Info(fmt.Sprintf("Файл %s успешно загружен", header.Filename))
	storage.Data["json"] = map[string]string{"message": "Файл успешно загружен"}
	storage.ServeJSON()
}

func (storage *StorageController) GetStorageInfo() {
	response := make(map[string]interface{})

	username := storage.GetSession("username")
	if username == nil {
		response["error"] = "Не авторизован"
		storage.Data["json"] = response
		storage.ServeJSON()
		return
	}

	o := orm.NewOrm()
	var user models.Users
	err := o.QueryTable("users").Filter("login", username).One(&user)
	if err != nil {
		response["error"] = fmt.Sprintf("Ошибка получения информации о пользователе: %v", err)
		storage.Data["json"] = response
		storage.ServeJSON()
		return
	}

	beego.Info(fmt.Sprintf("User: %+v", user))

	storageLimit, ok := tariffLimits[user.TariffID]
	if !ok {
		beego.Error(fmt.Sprintf("Неизвестный тариф: %d", user.TariffID))
		// Set a default limit if the tariff is unknown
		storageLimit = tariffLimits[1] // Use the base tariff as default
	}

	userDir := filepath.Join(uploadDir, username.(string))
	var usedSpace int64 = 0

	err = filepath.Walk(userDir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			usedSpace += info.Size()
		}
		return nil
	})

	if err != nil && !os.IsNotExist(err) {
		response["error"] = fmt.Sprintf("Ошибка подсчета использованного места: %v", err)
		storage.Data["json"] = response
		storage.ServeJSON()
		return
	}

	response["used"] = usedSpace
	response["limit"] = storageLimit
	response["percentage"] = float64(usedSpace) / float64(storageLimit) * 100

	storage.Data["json"] = response
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

	if fileName == "" || content == "" {
		storage.Ctx.Output.SetStatus(400)
		storage.Data["json"] = map[string]string{"error": "Имя файла или содержимое не указано"}
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

func (storage *StorageController) Download() {
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

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Ошибка открытия файла"}
		storage.ServeJSON()
		return
	}
	defer file.Close()

	// Устанавливаем заголовки для скачивания
	storage.Ctx.Output.Header("Content-Type", "application/octet-stream")
	storage.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	storage.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	storage.Ctx.Output.Header("Expires", "0")

	// Копируем содержимое файла в ответ
	_, err = io.Copy(storage.Ctx.ResponseWriter, file)
	if err != nil {
		storage.Ctx.Output.SetStatus(500)
		storage.Data["json"] = map[string]string{"error": "Ошибка при скачивании файла"}
		storage.ServeJSON()
		return
	}
}
