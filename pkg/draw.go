package pkg

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawRectFloor(texture rl.Texture2D, position rl.Vector3, width, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z

	// Set desired texture to be enabled while drawing following vertex data
	defaultID := rl.GetTextureIdDefault()
	rl.SetTexture(texture.ID)

	// Vertex data transformation can be defined with the commented lines,
	// but in this example we calculate the transformed vertex data directly when calling rlVertex3f()
	// rl.PushMatrix()
	// NOTE: Transformation is applied in inverse order (scale -> rotate -> translate)
	//rl.Translatef(2.0, 0.0, 0.0)
	//rl.Rotatef(45, 0, 1, )
	//rl.Scalef(2.0, 2.0, 2.0)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)
	// Top Face
	rl.Normal3f(0.0, 1.0, 0.0) // Normal Pointing Up
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x, y, z) // Top Left Of The Texture and Quad.
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x, y, z+length) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+width, y, z+length) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+width, y, z) // Top Right Of The Texture and Quad Bottom Face

	// doublesideed
	// rl.Normal3f(0.0, -1.0, 0.0) // Normal Pointing Down
	// rl.TexCoord2f(1.0, 1.0)
	// rl.Vertex3f(x, y, z) // Top Right Of The Texture and Quad
	// rl.TexCoord2f(0.0, 1.0)
	// rl.Vertex3f(x+width, y, z) // Top Left Of The Texture and Quad
	// rl.TexCoord2f(0.0, 0.0)
	// rl.Vertex3f(x+width, y, z+length) // Bottom Left Of The Texture and Quad
	// rl.TexCoord2f(1.0, 0.0)
	// rl.Vertex3f(x, y, z+length) // Bottom Right Of The Texture and Quad

	rl.End()
	//rl.PopMatrix()

	rl.SetTexture(defaultID)
}

func DrawRectCeiling(texture rl.Texture2D, position rl.Vector3, width, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z

	// Set desired texture to be enabled while drawing following vertex data
	defaultID := rl.GetTextureIdDefault()
	rl.SetTexture(texture.ID)

	// Vertex data transformation can be defined with the commented lines,
	// but in this example we calculate the transformed vertex data directly when calling rlVertex3f()
	// rl.PushMatrix()
	// NOTE: Transformation is applied in inverse order (scale -> rotate -> translate)
	//rl.Translatef(2.0, 0.0, 0.0)
	//rl.Rotatef(45, 0, 1, )
	//rl.Scalef(2.0, 2.0, 2.0)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// rl.Normal3f(0.0, 1.0, 0.0) // Normal Pointing Up
	// rl.TexCoord2f(0.0, 1.0)
	// rl.Vertex3f(x, y, z) // Top Left Of The Texture and Quad.
	// rl.TexCoord2f(0.0, 0.0)
	// rl.Vertex3f(x, y, z+length) // Bottom Left Of The Texture and Quad
	// rl.TexCoord2f(1.0, 0.0)
	// rl.Vertex3f(x+width, y, z+length) // Bottom Right Of The Texture and Quad
	// rl.TexCoord2f(1.0, 1.0)
	// rl.Vertex3f(x+width, y, z) // Top Right Of The Texture and Quad Bottom Face

	rl.Normal3f(0.0, -1.0, 0.0) // Normal Pointing Down
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x, y, z) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+width, y, z) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+width, y, z+length) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x, y, z+length) // Bottom Right Of The Texture and Quad

	rl.End()
	//rl.PopMatrix()

	rl.SetTexture(defaultID)
}

func DrawRectWall(texture rl.Texture2D, position rl.Vector3, width, height, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z

	// Set desired texture to be enabled while drawing following vertex data
	defaultID := rl.GetTextureIdDefault()
	rl.SetTexture(texture.ID)

	// Vertex data transformation can be defined with the commented lines,
	// but in this example we calculate the transformed vertex data directly when calling rlVertex3f()
	// rl.PushMatrix()
	// NOTE: Transformation is applied in inverse order (scale -> rotate -> translate)
	//rl.Translatef(2.0, 0.0, 0.0)
	//rl.Rotatef(45, 0, 1, )
	//rl.Scalef(2.0, 2.0, 2.0)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)
	// Front Face
	rl.Normal3f(0.0, 0.0, 1.0) // Normal Pointing Towards Viewer
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x, y, z+length) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+width, y, z+length) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+width, y+height, z+length) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x, y+height, z+length) // Top Left Of The Texture and Quad
	// Back Face
	rl.Normal3f(0.0, 0.0, -1.0) // Normal Pointing Away From Viewer
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x, y, z) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x, y+height, z) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+width, y+height, z) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+width, y, z) // Bottom Left Of The Texture and Quad
	// Right face
	rl.Normal3f(1.0, 0.0, 0.0) // Normal Pointing Right
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+width, y, z) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+width, y+height, z) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+width, y+height, z+length) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+width, y, z+length) // Bottom Left Of The Texture and Quad
	// Left Face
	rl.Normal3f(-1.0, 0.0, 0.0) // Normal Pointing Left
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x, y, z) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x, y, z+length) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x, y+height, z+length) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x, y+height, z) // Top Left Of The Texture and Quad

	rl.End()
	//rl.PopMatrix()

	rl.SetTexture(defaultID)
}
