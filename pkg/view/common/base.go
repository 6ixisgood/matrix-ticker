package common

import (
	"time"
	compCommon "github.com/6ixisgood/matrix-ticker/pkg/component/common"
)

type BaseView struct {
	template			*compCommon.Template
	dataRefresh			*time.Ticker
	templateRefresh		*time.Ticker
	stopChan			chan struct{}
}

func (v *BaseView) Init() {}

func (v *BaseView) Template() *compCommon.Template {
	return v.template
}

func (v *BaseView) SetTemplate(t *compCommon.Template) {
	v.template = t
}
func (v *BaseView) TemplateData() map[string]interface{} {
	return map[string]interface{}{}  
}

func (v *BaseView) TemplateString() string {
	return ""
}

func (v *BaseView) Stop() {}