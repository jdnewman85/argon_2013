#version 430 core

layout( points ) in;
layout( triangle_strip, max_vertices = 4 ) out;

uniform mat4 inOrtho;

in vec2 vertSize[1];
in vec4 vertColor[1];
in vec2 vertScale[1];
in float vertAngle[1];

out vec4 geomColor;
out vec2 texCords;

void main( void ) {
	int i;
	vec4 tempVert;
	mat4 translationMat;
	mat4 rotationMat;
	mat4 scaleMat;
	mat4 transformMat;

	for ( i = 0; i < gl_in.length(); i++ ) {
		geomColor = vertColor[i];


		tempVert = vec4( vertSize[i].x / -2.0, vertSize[i].y / -2.0, 0.0, 1.0 );
		translationMat = mat4( 1.0, 0.0, 0.0, 0.0,
							   0.0, 1.0, 0.0, 0.0,
							   0.0, 0.0, 1.0, 0.0,
							   gl_in[i].gl_Position.x, gl_in[i].gl_Position.y, 0.0, 1.0 );
		scaleMat = mat4( vertScale[i].x, 0.0, 0.0, 0.0,
						 0.0, vertScale[i].y, 0.0, 0.0,
						 0.0, 0.0, 1.0, 0.0,
						 0.0, 0.0, 0.0, 1.0 );
		rotationMat = mat4( cos( radians( vertAngle[i] ) ), -sin( radians( vertAngle[i] ) ), 0.0, 0.0,
							sin( radians( vertAngle[i] ) ), cos( radians( vertAngle[i] ) ), 0.0, 0.0,
							0.0, 0.0, 1.0, 0.0,
							0.0, 0.0, 0.0, 1.0 );

		transformMat = inOrtho * translationMat * rotationMat * scaleMat;

		texCords = vec2( 0.0, 1.0 );
		gl_Position = transformMat * tempVert;
		//gl_Position = vec4(texCords, 0.0, 1.0);
		EmitVertex();
		texCords = vec2( 1.0, 1.0 );
		tempVert.x *= -1.0;
		gl_Position = transformMat * tempVert;
		//gl_Position = vec4(texCords, 0.0, 1.0);
		EmitVertex();
		texCords = vec2( 0.0, 0.0 );
		tempVert.xy *= -1.0;
		gl_Position = transformMat * tempVert;
		//gl_Position = vec4(texCords, 0.0, 1.0);
		EmitVertex();
		texCords = vec2( 1.0, 0.0 );
		tempVert.x *= -1.0;
		gl_Position = transformMat * tempVert;
		//gl_Position = vec4(texCords, 0.0, 1.0);
		EmitVertex();
	}
	EndPrimitive();
}
