package argon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("shader.go here")
}

const VertexShader = gl.VERTEX_SHADER
const GeometryShader = gl.GEOMETRY_SHADER
const FragmentShader = gl.FRAGMENT_SHADER

type Shader gl.Uint
type Program gl.Uint


func CreateShader(shaderType gl.Enum) Shader {
	return Shader(gl.CreateShader(shaderType))
}

func CreateProgram() Program {
	return Program(gl.CreateProgram())
}

func (this Shader) Source(source string) {
	glSource := gl.GLString(source)
	defer gl.GLStringFree(glSource)
	gl.ShaderSource(gl.Uint(this), 1, &glSource, nil)
}

func (this Shader) Compile() {
	gl.CompileShader(gl.Uint(this))

	//Compile check
	var status gl.Int
	if gl.GetShaderiv(gl.Uint(this), gl.COMPILE_STATUS, &status); status != gl.TRUE {
		//Failed
		fmt.Fprintf(os.Stderr, "shader: %s\n", this.ShaderInfoLog())
		//TODO err return
	}
}

func (this Shader) Delete() {
	gl.DeleteShader(gl.Uint(this))
	//TODO Should I null this?
}

func (this Program) Attach(shader Shader) {
	gl.AttachShader(gl.Uint(this), gl.Uint(shader))
}

func (this Program) Link() {
	gl.LinkProgram(gl.Uint(this))
}

func (this Program) Use() {
	gl.UseProgram(gl.Uint(this))
}

func (this Program) Forgo() {
	gl.UseProgram(0)
}

func CreateProgramFromFiles(fileNames []string) (Program, error) {
	//TODO Make these errors more descriptive maybe?
	program := CreateProgram()

	for _, fileName := range fileNames {
		//Check for extension
		var shaderType gl.Enum
		switch filepath.Ext(fileName) {
		case ".vert":
			shaderType = VertexShader
		case ".geom":
			shaderType = GeometryShader
		case ".frag":
			shaderType = FragmentShader
		default:
			fmt.Fprintf(os.Stderr, "shader: Unsupported Extension: %s\n", fileName)
			return program, errors.New("shader: Unsuported Extension\n") //TODO Delete Program?
		}
		if source, err := ioutil.ReadFile(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "shader: %s\n", err)
			return program, err //TODO Delete Progam?
		} else {
			shader := CreateShader(shaderType)
			shader.Source(string(source))
			shader.Compile() //TODO Check err
			defer shader.Delete()
			program.Attach(shader)
		}
	}

	program.Link()

	return program, nil
}


func (this Shader) ShaderInfoLog() (rString string) {
	var infoLogLength gl.Int
	var charsWritten gl.Sizei

	gl.GetShaderiv(gl.Uint(this), gl.INFO_LOG_LENGTH, &infoLogLength)

	if infoLogLength > 0 {
		infoLog := gl.GLStringAlloc(gl.Sizei(infoLogLength))
		gl.GetShaderInfoLog(gl.Uint(this), gl.Sizei(infoLogLength), &charsWritten, infoLog)
		rString = gl.GoString(infoLog)
		gl.GLStringFree(infoLog)

		return
	}

	//TODO Temp
	return "No Log"
}

