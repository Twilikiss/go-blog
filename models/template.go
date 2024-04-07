package models

import (
	"go-blog/log"
	"html/template"
	"io"
)

// TemplateBlog 在我们的 html/template 之上套多一层，方便我们的调用
type TemplateBlog struct {
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func (t *TemplateBlog) ExecuteHTML(write io.Writer, name string, data interface{}) error {
	err := t.htmlTemplates.ExecuteTemplate(write, name, data)
	if err != nil {
		log.Errorf("%s解析模版错误：%s", name, err)
		return err
	}
	return nil
}

func (t *TemplateBlog) WriteError(write io.Writer, err string) {
	_, err2 := write.Write([]byte(err))
	if err2 != nil {
		log.Error("无法输出错误信息到前端")
	}
}

func (t *TemplateBlog) LoadHTMLGlob(pattern string) {
	t.htmlTemplates = template.Must(template.New("").Funcs(t.funcMap).ParseGlob(pattern))
}

func (t *TemplateBlog) SetFuncMap(funcMap template.FuncMap) {
	t.funcMap = funcMap
}
