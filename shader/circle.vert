#version 430 core

layout(location = 0) in vec2 inPos;
layout(location = 1) in float inRadius;
layout(location = 2) in vec4 inColor;

out float vertRadius;
out vec4 vertColor;

void main() {
    vertRadius = inRadius;
    //vertRadius = 10.0;
    vertColor = inColor;
	//vertColor = vec4(1.0, 0.0, 0.0, 1.0);
	gl_Position = vec4( inPos, 0.0, 1.0 );
	//gl_Position = vec4(400.0, 300.0, 0.0, 1.0);
}
