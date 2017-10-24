package common

import (
	"engo.io/ecs"
	"github.com/faiface/pixel"
)

type cameraEntity struct {
	*ecs.BasicEntity
	*SpaceComponent
}

type PixelCameraSystem struct {
	x, y, z      float32
	Tracking     *SpaceComponent // The entity that is currently being followed
	CameraObject *pixel.Vec      //camera object just some vector
	Bounds       pixel.Rect      // game bounds
	WindowBound  pixel.Rect      // windowbound
}

func (ps PixelCameraSystem) Update(dt float32) {
	halfX1 := ps.Bounds.Min.X + ps.WindowBound.W()/2
	halfX2 := ps.Bounds.Max.X - ps.WindowBound.W()/2
	if ps.Tracking.Location.X > halfX1 && ps.Tracking.Location.X < halfX2 {
		ps.CameraObject.X = ps.Tracking.Location.X
	}
	halfY1 := ps.Bounds.Max.Y - ps.WindowBound.H()/2
	halfY2 := ps.Bounds.Min.Y + ps.WindowBound.H()/2
	if ps.Tracking.Location.Y > halfY2 && ps.Tracking.Location.Y < halfY1 {
		ps.CameraObject.Y = ps.Tracking.Location.Y
	}
}
func (ps PixelCameraSystem) Remove(basic ecs.BasicEntity) {

}
