package main

import (
	"file-monitor/config"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"os"
)

// CreateMonitor	godoc
// @Summary		创建监控
// @Description	创建监控
// @Tags	监控
// @Accept	application/json
// @Produce json
// @Success 200 {string} string	"ok"
// @Param	path formData string true "监控文件夹或者文件"
// @Router	/monitor/create [post]
func CreateMonitor(c *gin.Context) Result {
	var req CreateMonitorReq
	err := c.ShouldBind(&req)
	if err != nil {
		return Error(-1, "参数错误:"+err.Error())
	}
	if !fileIsExist(req.Path) {
		return Error(-1, "路径不存在")
	}
	go func() {
		monitor(req.Path)
	}()

	return Success(nil)
}

func monitor(path string) {
	watcher, _ := fsnotify.NewWatcher()
	watcher.Add(path)
	defer watcher.Close()
	for {
		select {
		case event := <-watcher.Events:
			{
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					config.Log.Info("文件" + event.Name + "被重命名")
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					config.Log.Info("文件" + event.Name + "被更新")
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					config.Log.Info("有新建的文件" + event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					config.Log.Info("文件" + event.Name + "被删除")
				}
			}
		case err := <-watcher.Errors:
			fmt.Println(err.Error())
		}
	}
}

func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
