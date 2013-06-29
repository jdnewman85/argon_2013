
package shader

import(
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gl "github.com/chsc/gogl/gl33"
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

func (this *Shader) LoadFromFile(fileName string) error {
	this.Program = gl.CreateProgram()

	//TODO How about we pass an array of filenames and loop through and load?
	//Vert
	if source, err := ioutil.ReadFile(fileName+".vert"); err != nil {
		fmt.Fprintf(os.Stderr, "material - vert: %s\n", err)
		return err
	} else {
		glSource := gl.GLString(string(source))
		shader := gl.CreateShader(gl.VERTEX_SHADER)
		gl.ShaderSource(shader, 1, &glSource, nil)
		gl.CompileShader(shader)
		//Check compiled status TODO
		gl.AttachShader(this.Program, shader)
	}

	//Frag
	if source, err := ioutil.ReadFile(fileName+".frag"); err != nil {
		fmt.Fprintf(os.Stderr, "material - frag: %s\n", err)
		return err
	} else {
		glSource := gl.GLString(string(source))
		shader := gl.CreateShader(gl.FRAGMENT_SHADER)
		gl.ShaderSource(shader, 1, &glSource, nil)
		gl.CompileShader(shader)
		//Check compiled status TODO
		gl.AttachShader(this.Program, shader)
	}

	//Geom
	if source, err := ioutil.ReadFile(fileName+".geom"); err != nil {
		fmt.Fprintf(os.Stderr, "material - geom: %s\n", err)
		return err
	} else {
		glSource := gl.GLString(string(source))
		shader := gl.CreateShader(gl.GEOMETRY_SHADER)
		gl.ShaderSource(shader, 1, &glSource, nil)
		gl.CompileShader(shader)
		//Check compiled status TODO
		gl.AttachShader(this.Program, shader)
	}

	return nil
}

func (this *Shader) Link() {
	gl.LinkProgram(this.Program)
}

func (this *Shader) Use() {
	gl.UseProgram(this.Program)
}

//TODO REM
func MakeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float {
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
