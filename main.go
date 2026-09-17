package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/seanoneillcode/validation/pkg"
)

const (
	MOUSE_LOOK_SENSITIVITY = 0.002
	MOVE_SPEED             = 0.05

	PLAYER_HEIGHT = 0.675
)

var (
	FORWARD_VECTOR = rl.Vector3{X: 0, Y: 0, Z: 1}
	UP_VECTOR      = rl.Vector3{X: 0, Y: 1, Z: 0}
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1280, 720, "validation")
	rl.SetWindowState(rl.FlagVsyncHint)
	rl.SetTargetFPS(60)

	textures := make([]rl.Texture2D, 10)
	textures[0] = rl.LoadTexture("textures/clay.png")
	textures[1] = rl.LoadTexture("textures/metal.png")
	textures[2] = rl.LoadTexture("textures/metal-tiles.png")

	// setup
	camera := createCamera()
	s := state{
		pos:      rl.NewVector3(0, PLAYER_HEIGHT, 0),
		textures: textures,
	}

	lastTime := time.Now()
	for !rl.WindowShouldClose() {
		// update
		now := time.Now()
		deltaTime := now.Sub(lastTime).Milliseconds()
		lastTime = now
		s.deltaTime = deltaTime
		update(&camera, &s)

		//draw
		draw(camera, s)
	}
	for _, t := range textures {

		rl.UnloadTexture(t)

	}

	rl.CloseWindow()
}

func update(camera *rl.Camera3D, state *state) {
	mouseDelta := rl.GetMouseDelta()
	state.lookRotation -= mouseDelta.X * MOUSE_LOOK_SENSITIVITY

	front := rl.NewVector3(float32(math.Sin(float64(state.lookRotation))), 0, float32(math.Cos(float64(state.lookRotation))))
	right := rl.NewVector3(float32(math.Cos(float64(-state.lookRotation))), 0, float32(math.Sin(float64(-state.lookRotation))))

	input := rl.Vector2{}
	if rl.IsKeyDown(rl.KeyW) {
		input.Y = 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		input.Y = -1
	}
	if rl.IsKeyDown(rl.KeyA) {
		input.X = 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		input.X = -1
	}
	dir := rl.NewVector3(input.X*right.X+input.Y*front.X, 0, input.X*right.Z+input.Y*front.Z)

	// speed := rl.Vector3DotProduct(hvel, dir);

	state.velocity.X = dir.X * MOVE_SPEED
	state.velocity.Z = dir.Z * MOVE_SPEED

	state.pos.X = state.pos.X + (state.velocity.X)
	state.pos.Z = state.pos.Z + (state.velocity.Z)

	// update camera
	yaw := rl.Vector3RotateByAxisAngle(FORWARD_VECTOR, UP_VECTOR, state.lookRotation)
	camera.Target = rl.Vector3Add(state.pos, yaw)
	camera.Position = state.pos
}

func draw(camera rl.Camera3D, state state) {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.BeginMode3D(camera)

	for i := range 10 {
		for j := range 10 {
			pos := rl.NewVector3(float32(i), 0, float32(j))
			pkg.DrawRectFloor(state.textures[0], pos, 1, 1, rl.White)
		}
	}
	for i := range 10 {
		for j := range 6 {
			pos := rl.NewVector3(float32(i), 1, float32(j))
			pkg.DrawRectCeiling(state.textures[2], pos, 1, 1, rl.White)
		}
	}
	for j := range 5 {
		pos := rl.NewVector3(float32(2), 0, float32(j))
		pkg.DrawRectWall(state.textures[1], pos, 1, 1, 1, rl.White)
	}
	for j := range 5 {
		pos := rl.NewVector3(float32(j), 0, float32(6))
		pkg.DrawRectWall(state.textures[1], pos, 1, 1, 1, rl.White)
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
	rl.DisableCursor()
	rl.UpdateCamera(&camera, rl.CameraCustom)
	return camera
}

type state struct {
	model        rl.Model
	pos          rl.Vector3
	lookRotation float32
	velocity     rl.Vector3
	textures     []rl.Texture2D
	deltaTime    int64
}
