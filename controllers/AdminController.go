package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"path/filepath"
	"time"
	"web_cloud_storage/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplName = "admin.html"
}

func (c *AdminController) GetUserInfo() {
	username := c.GetString("username")
	o := orm.NewOrm()
	user := models.Users{Login: username}
	err := o.Read(&user, "Login")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "User not found"}
		c.ServeJSON()
		return
	}

	// Путь к папке пользователя
	userDir := filepath.Join("./storage", username)
	var files []map[string]interface{}
	err = filepath.Walk(userDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, map[string]interface{}{
				"name": info.Name(),
				"size": info.Size(),
			})
		}
		return nil
	})

	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Error reading user files"}
		c.ServeJSON()
		return
	}

	// Здесь мы вызываем calculateTotalUsage, передавая путь к папке пользователя
	storageUsed := calculateTotalUsage(userDir) // Передаем путь, не даты

	// Преобразуем размер в гигабайты
	storageUsedGB := float64(storageUsed) / (1024 * 1024 * 1024)

	// Отправка данных в формате JSON
	c.Data["json"] = map[string]interface{}{
		"name":        user.Username,
		"storageUsed": storageUsedGB, // Размер в гигабайтах
		"files":       files,
	}
	c.ServeJSON()
}

func (c *AdminController) GetStorageReport() {
	startDate, _ := time.Parse("2006-01-02", c.GetString("start"))
	endDate, _ := time.Parse("2006-01-02", c.GetString("end"))

	// Проверка на корректность дат
	if startDate.After(endDate) {
		c.Data["json"] = map[string]interface{}{"error": "Start date cannot be after end date"}
		c.ServeJSON()
		return
	}

	// Вычисление общего объема хранения
	totalUsage := int64(0)
	err := filepath.Walk("./storage", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (info.ModTime().After(startDate) && info.ModTime().Before(endDate)) {
			totalUsage += info.Size()
		}
		return nil
	})

	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Error calculating storage usage"}
		c.ServeJSON()
		return
	}

	// Проверка на корректность данных
	days := int(endDate.Sub(startDate).Hours() / 24)
	if days <= 0 {
		c.Data["json"] = map[string]interface{}{"error": "The date range must be more than 1 day"}
		c.ServeJSON()
		return
	}

	averageUsage := totalUsage / int64(days)

	dailyUsage := make(map[string]int64)
	for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 0, 1) {
		// Для примера, можно просто использовать averageUsage для всех дней
		dailyUsage[d.Format("2006-01-02")] = averageUsage
	}

	c.Data["json"] = map[string]interface{}{
		"totalUsage":   totalUsage,
		"averageUsage": averageUsage,
		"dailyUsage":   dailyUsage,
	}
	c.ServeJSON()
}

func (c *AdminController) GetFileList() {
	var files []map[string]interface{}

	err := filepath.Walk("./storage", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			owner := filepath.Base(filepath.Dir(path))
			files = append(files, map[string]interface{}{
				"name":         info.Name(),
				"owner":        owner,
				"size":         info.Size(),
				"lastModified": info.ModTime(),
				"permissions":  info.Mode().String(),
			})
		}
		return nil
	})

	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Error reading files"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = files
	c.ServeJSON()
}

// Функция для вычисления общего объема использованного хранилища для пользователя
func calculateTotalUsage(userDir string) int64 {
	var totalUsage int64

	// Проходим по всем файлам в директории
	err := filepath.Walk(userDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Добавляем размер файла
			totalUsage += info.Size()
		}
		return nil
	})

	if err != nil {
		// Если возникла ошибка при обходе файлов
		return 0
	}

	return totalUsage
}
