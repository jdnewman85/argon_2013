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

type Shader struct {
	Program gl.Uint
}

func init() {
	log.Println("shader.go here")
}

func Create() *Shader {
	temp := new(Shader)

	temp.Program = gl.CreateProgram()

	return temp
}

func CreateShaderFromFiles(fileNames []string) (*Shader, error) {
	//TODO Make these errors more destriptive maybe?
	program := gl.CreateProgram()

	for _, fileName := range fileNames {
		//Check for extension
		var shaderType gl.Enum
		switch filepath.Ext(fileName) {
		case ".vert":
			shaderType = gl.VERTEX_SHADER
		case ".geom":
			shaderType = gl.GEOMETRY_SHADER
		case ".frag":
			shaderType = gl.FRAGMENT_SHADER
		default:
			fmt.Fprintf(os.Stderr, "shader: Unsupported Extension: %s\n", fileName)
			return nil, errors.New("shader: Unsuported Extension\n")
		}
		if source, err := ioutil.ReadFile(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "shader: %s\n", err)
			return nil, err
		} else {
			glSource := gl.GLString(string(source))
			shader := gl.CreateShader(shaderType)
			gl.ShaderSource(shader, 1, &glSource, nil)
			gl.CompileShader(shader)
			//Compile check
			var status gl.Int
			if gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status); status != gl.TRUE {
				//Failed
				fmt.Fprintf(os.Stderr, "shader: %s\n", shaderInfoLog(shader))
			}
			gl.AttachShader(program, shader)
		}
	}

	shader := &Shader{program}
	shader.Link()

	return shader, nil
}

func (this *Shader) Link() {
	gl.LinkProgram(this.Program)
}

func (this *Shader) Use() {
	gl.UseProgram(this.Program)
}

func shaderInfoLog(shader gl.Uint) (rString string) {
	var infoLogLength gl.Int
	var charsWritten gl.Sizei

	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &infoLogLength)

	if infoLogLength > 0 {
		infoLog := gl.GLStringAlloc(gl.Sizei(infoLogLength))
		gl.GetShaderInfoLog(shader, gl.Sizei(infoLogLength), &charsWritten, infoLog)
		rString = gl.GoString(infoLog)
		gl.GLStringFree(infoLog)

		return
	}

	//TODO Temp
	return "No Log"
}

//TODO REM
func MakeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float{
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
