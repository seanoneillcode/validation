package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/seanoneillcode/validation/pkg"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1280, 720, "validation")
	rl.SetWindowState(rl.FlagVsyncHint)
	rl.SetTargetFPS(60)

	texture := rl.LoadTexture("clay.png")

	// setup
	camera := createCamera()
	s := state{
		pos:     rl.NewVector3(0, 0, 0),
		texture: texture,
	}

	lastTime := time.Now()
	for !rl.WindowShouldClose() {
		// update
		now := time.Now()
		deltaTime := now.Sub(lastTime).Milliseconds()
		lastTime = now
		s.deltaTime = deltaTime

		//draw
		draw(camera, s)
	}
	rl.UnloadTexture(texture)

	rl.CloseWindow()
}

func draw(camera rl.Camera3D, state state) {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.BeginMode3D(camera)

	for i := range 10 {
		for j := range 10 {
			pos := rl.NewVector3(float32(i), 0, float32(j))
			pkg.DrawRectFloor(state.texture, pos, 1, 1, rl.White)
		}
	}
	rl.DrawGrid(20, 1.0)

	rl.EndMode3D()
	rl.EndDrawing()
}

func createCamera() rl.Camera3D {
	camera := rl.Camera3D{}
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective
	camera.Position = rl.NewVector3(4, 4, 4)
	rl.UpdateCamera(&camera, rl.CameraCustom)
	return camera
}

type state struct {
	model     rl.Model
	pos       rl.Vector3
	texture   rl.Texture2D
	deltaTime int64
}
