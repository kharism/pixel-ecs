package common

import (
	"engo.io/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ContextSwitcher struct {
	Contexts         map[string][]ecs.System
	CurrentContext   string
	AvailableContext []string
	Win              *pixelgl.Window
	GameBoundary     pixel.Rect
	WindowSize       pixel.Rect
	CameraVector     pixel.Vec
	Image            ImageAssetManager
	Messager         MessagingSystem
}

func NewContextSwitcher(pixelcfg pixelgl.WindowConfig, GameBoundary pixel.Rect) ContextSwitcher {
	k := ContextSwitcher{}
	k.Contexts = map[string][]ecs.System{}
	k.AvailableContext = []string{}
	win, err := pixelgl.NewWindow(pixelcfg)
	if err != nil {
		panic(err)
	}
	k.Win = win
	k.GameBoundary = GameBoundary
	k.WindowSize = pixelcfg.Bounds
	k.CameraVector = pixel.V(k.WindowSize.W()/2, -k.WindowSize.H()/2)
	camMatrix := pixel.IM.Moved(k.Win.Bounds().Center().Sub(k.CameraVector))
	k.Win.SetMatrix(camMatrix)
	k.Image = NewImageAssetManager()
	k.Messager = NewMessagingSystem()
	return k
}
func in(needle string, haystack []string) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}

func (cs *ContextSwitcher) AddSystem(context string, s ecs.System) {
	if context == "" {
		context = cs.CurrentContext
	}
	if !in(context, cs.AvailableContext) {
		cs.AvailableContext = append(cs.AvailableContext, context)
	}
	if _, ok := cs.Contexts[context]; !ok {
		cs.Contexts[context] = []ecs.System{}
	}
	cs.Contexts[context] = append(cs.Contexts[context], s)
}
