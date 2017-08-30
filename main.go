package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// Needed to assure that main() runs on the main thread
	runtime.LockOSThread()
}

func main() {
	const (
		WinWidth  = 1024
		WinHeight = 768
	)
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WinWidth, WinHeight, "GoCraft", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	window.SetCursorPos(WinWidth/2, WinHeight/2)

	// Initialize Glow
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Running OpenGL version", version)

	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	proj := mgl32.Perspective(mgl32.DegToRad(45.0), float32(WinWidth)/WinHeight, 0.1, 200.0)
	projUniform := gl.GetUniformLocation(program, gl.Str("proj\x00"))
	gl.UniformMatrix4fv(projUniform, 1, false, &proj[0])

	view := mgl32.LookAtV(mgl32.Vec3{32, 32, 32}, mgl32.Vec3{16, 24, 16}, mgl32.Vec3{0, 1, 0})
	viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	tex, err := newTexture("blocks.png")
	if err != nil {
		panic(err)
	}
	texUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(texUniform, 0)

	chunks := [3][3]Chunk{}
	for x := 0; x < len(chunks); x++ {
		for z := 0; z < len(chunks[0]); z++ {
			chunks[x][z] = Chunk{}
			err = chunks[x][z].Load("maps/Panda Islands", x, z, program)
			if err != nil {
				panic(err)
			}
		}
	}

	gl.CullFace(gl.FRONT_AND_BACK)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	angle := 0.0
	previousTime := glfw.GetTime()

	var state glfw.Action

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		state = window.GetKey(glfw.KeyLeft)
		if state == glfw.Press {
			angle -= elapsed
		}
		state = window.GetKey(glfw.KeyRight)
		if state == glfw.Press {
			angle += elapsed
		}

		mx, my := window.GetCursorPos()
		mx -= WinWidth / 2
		my -= WinHeight / 2
		if mx != 0 || my != 0 {
			fmt.Println(mx, my)
			window.SetCursorPos(WinWidth/2, WinHeight/2)
		}

		gl.UseProgram(program)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		for x := 0; x < len(chunks); x++ {
			for z := 0; z < len(chunks[0]); z++ {
				chunk := chunks[x][z]

				model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})
				model = model.Mul4(mgl32.Translate3D(float32(x*16), 0, float32(z*16)))

				gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

				gl.BindVertexArray(chunk.VAO)
				gl.DrawArrays(gl.TRIANGLES, 0, int32(chunk.GetLength()*36))
			}
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}

var vertexShader = `
#version 330 core

uniform mat4 proj;
uniform mat4 view;
uniform mat4 model;

in vec3 in_vert;
in vec2 in_texcoord;

out vec2 texcoord;

void main() {
	texcoord = in_texcoord;
	gl_Position = proj * view * model * vec4(in_vert, 1);
    if (texcoord.x < 0.0 && texcoord.y < 0.0) {
        gl_Position = vec4(0, 0, 0, 0);
    }
}

` + "\x00"

var fragmentShader = `
#version 330 core

uniform sampler2D tex;

in vec2 texcoord;

out vec4 o_color;

void main() {
    o_color = texture(tex, texcoord);
}
` + "\x00"
