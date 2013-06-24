// Material

package argon

import(
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.Println("material.go here")
}

type Material struct {
	//TODO May make vao associated with the data itself?
	vao, vbo, shaderProgram gl.Uint
}

func (this *Material) LoadShaders(fileName string) error {
	//VAO
	gl.GenVertexArrays(1, &this.vao)

	//VBO
	gl.GenBuffers(1, &this.vbo)

	//Shaders
	this.shaderProgram = gl.CreateProgram()

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
		gl.AttachShader(this.shaderProgram, shader)
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
		gl.AttachShader(this.shaderProgram, shader)
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
		gl.AttachShader(this.shaderProgram, shader)
	}

	gl.LinkProgram(this.shaderProgram)
	return nil
}

func (this *Material) Bind() {
	gl.BindVertexArray(this.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)
	gl.UseProgram(this.shaderProgram)
}

func (this *Material) UnBind() {
	//TODO Kosher?
	gl.UseProgram(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

}

func (this *Material) Draw( numEntities, sizeEntities int, entities interface{}) {
	//TODO REM Avoid unnessessary rebinds
	this.Bind()

	//TODO Finish
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(sizeEntities * numEntities), entities.(gl.Pointer), gl.DYNAMIC_DRAW)

	//TODO only supports point materials, nees to be set in material
	gl.DrawArrays(gl.POINTS, 0, gl.Sizei(numEntities))
}

