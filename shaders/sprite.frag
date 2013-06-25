#version 430 core

uniform sampler2D inTexture;

in vec4 geomColor;
in vec2 texCords;

layout(location = 0) out vec4 outColor;

void main() {
    outColor = texture( inTexture, texCords ) * geomColor;
    //outColor = geomColor;
}
