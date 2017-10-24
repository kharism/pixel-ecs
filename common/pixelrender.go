package common

import (
	"engo.io/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var camPos = pixel.ZV

type RenderComponent struct {
	Drawable *pixel.Sprite
	Hidden   bool
	zIndex   float32
}
type SpaceComponent struct {
	Location pixel.Vec
}
type RenderSystemPixel struct {
	cs            *ContextSwitcher
	entities      renderEntityList
	sortingNeeded bool
}

type renderEntity struct {
	*ecs.BasicEntity
	*RenderComponent
	*SpaceComponent
}

func (r *RenderComponent) Draw(w *pixelgl.Window, tm pixel.Matrix) {
	r.Drawable.Draw(w, tm)
}
func NewRenderSystem(w *ContextSwitcher) RenderSystemPixel {
	k := RenderSystemPixel{}
	k.cs = w
	return k
}

type renderEntityList []renderEntity

func (rs RenderSystemPixel) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range rs.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		rs.entities = append(rs.entities[:delete], rs.entities[delete+1:]...)
		rs.sortingNeeded = true
	}
}
func (rs *RenderSystemPixel) Add(basic *ecs.BasicEntity, render *RenderComponent, space *SpaceComponent) {
	rs.entities = append(rs.entities, renderEntity{basic, render, space})
}
func (rs RenderSystemPixel) Update(dt float32) {
	for _, re := range rs.entities {
		if !re.RenderComponent.Hidden {
			re.RenderComponent.Draw(rs.cs.Win, pixel.IM.Moved(re.SpaceComponent.Location))
		}
	}
}
