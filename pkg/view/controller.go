package view

import (
	compCommon "github.com/6ixisgood/matrix-ticker/pkg/component/common"
	viewCommon "github.com/6ixisgood/matrix-ticker/pkg/view/common"
	_ "github.com/6ixisgood/matrix-ticker/pkg/view/types"
	"image"
	"log"
	"time"
)

var (
	animation = &Animation{}
)

type Animation struct {
	view     viewCommon.View
	template compCommon.Template
}

func (a *Animation) Init(newView viewCommon.View) {
	log.Printf("Initializing view in controller")

	// init new view in background
	newView.Init()
	viewCommon.TemplateRefresh(newView)

	// stop the old view and switch to new view
	if a.view != nil {
		a.view.Stop()
	}
	a.view = newView
}

func (a *Animation) Next() (image.Image, <-chan time.Time, error) {
	t := a.view.Template()

	// render template to image.Image
	for {
		if !(t.Ready()) {
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	im := t.Render()
	return im, time.After(time.Millisecond * 10), nil
}

func GetAnimation() *Animation {
	return animation
}
