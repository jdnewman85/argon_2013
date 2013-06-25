
package shader

import(
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
	"log"
	"os"
)

type Shader struct {
	Program gl.Uint
}

func init() {
	log.Println("shader.go here")
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

	gl.LinkProgram(this.Program)
	return nil
}

func (this *Shader) Use() {
	gl.UseProgram(this.Program)
}

