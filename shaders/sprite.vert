#version 430 core

layout(location = 0) in vec2 inPos;
layout(location = 1) in vec2 inSize;
layout(location = 2) in vec4 inColor;
layout(location = 3) in vec2 inScale;
layout(location = 4) in float inAngle;

out vec2 vertSize;
out vec4 vertColor;
out vec2 vertScale;
out float vertAngle;

void main() {
    vertSize = inSize;
    //vertSize = vec2(100.0, 100.0);
    vertColor = inColor;
	//vertColor = vec4(1.0, 0.0, 0.0, 1.0);
	vertScale = inScale;
	//vertScale = vec2(0.5, 0.5);
	vertAngle = inAngle;
	//vertAngle = 45.0;
	gl_Position = vec4( inPos, 0.0, 1.0 );
	//gl_Position = vec4(400.0, 300.0, 0.0, 1.0);
}
