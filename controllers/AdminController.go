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

	c.Data["json"] = map[string]interface{}{
		"name":        user.Username,
		"storageUsed": calculateStorageUsed(userDir),
		"files":       files,
	}
	c.ServeJSON()
}

func (c *AdminController) GetStorageReport() {
	startDate, _ := time.Parse("2006-01-02", c.GetString("start"))
	endDate, _ := time.Parse("2006-01-02", c.GetString("end"))

	// This is a placeholder. In a real application, you would query your database or log files
	// to get actual usage data for the specified period.
	totalUsage := int64(1000000000) // 1 GB
	days := int(endDate.Sub(startDate).Hours() / 24)
	averageUsage := totalUsage / int64(days)

	dailyUsage := make(map[string]int64)
	for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 0, 1) {
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

func calculateStorageUsed(dir string) int64 {
	var size int64
	filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}
