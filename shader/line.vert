#version 430 core

layout(location = 0) in vec2 inPoint1;
layout(location = 1) in vec2 inPoint2;
layout(location = 2) in float inRadius;
layout(location = 3) in vec4 inColor;

out vec2 vertPoint1;
out vec2 vertPoint2;
out float vertRadius;
out vec4 vertColor;

void main() {
    vertPoint1 = inPoint1;
    //vertPoint1 = vec2(0.0, 0.0);
    vertPoint2 = inPoint2;
    //vertPoint2 = vec2(100.0, 100.0);
    vertRadius = inRadius;
    //vertRadius = 25.0;
    vertColor = inColor;
    //vertColor = vec4(1.0, 0.0, 0.0, 1.0);
    gl_Position = vec4( vertPoint1 + ((vertPoint2 - vertPoint1) / 2.0), 0.0, 1.0 );
    //gl_Position = vec4( 0.0, 0.0, 0.0, 1.0 );
}
