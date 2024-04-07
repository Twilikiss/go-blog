package common

import (
	"encoding/json"
	"go-blog/models"
	"html/template"
	"io"
	"net/http"
	"sync"
	"time"
)

var Template models.TemplateBlog

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}

// LoadTemplates 导入模版所需的方法并加载指定路径下的模版
func LoadTemplates() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()

		// 初始化赋值template
		funcMap := template.FuncMap{
			"isODD":       IsODD,
			"getNextName": GetNextName,
			"date":        Date,
			"dateDay":     DateDay,
		}
		Template.SetFuncMap(funcMap)
		Template.LoadHTMLGlob("templates/*")
	}()
	w.Wait()
}

func GetRequestJsonParam(request *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
