#version 430 core
//TODO FINISH HERE!
layout( points ) in;
layout( triangle_strip, max_vertices = 4 ) out;

uniform mat4 inOrtho;

in float vertRadius[1];
in vec4 vertColor[1];

out vec2 geomPosition;
out vec2 geomFragCoord;
out float geomRadius;
out vec4 geomColor;

void main( void ) {
	int i;
	vec4 tempVert;
	mat4 translationMat;
	mat4 rotationMat;
	mat4 scaleMat;
	mat4 transformMat;
    mat4 fragCoordTransformMat;

	for ( i = 0; i < gl_in.length(); i++ ) {
        geomPosition = gl_in[i].gl_Position.xy;
        geomRadius = vertRadius[i];
        geomColor = vertColor[i];

		tempVert = vec4( -vertRadius[i], -vertRadius[i], 0.0, 1.0 );
		translationMat = mat4( 1.0, 0.0, 0.0, 0.0,
							   0.0, 1.0, 0.0, 0.0,
							   0.0, 0.0, 1.0, 0.0,
							   gl_in[i].gl_Position.x, gl_in[i].gl_Position.y, 0.0, 1.0 );
        //TODO OPT scaleMat is not used, could be removed
		scaleMat = mat4( 1.0, 0.0, 0.0, 0.0,
						 0.0, 1.0, 0.0, 0.0,
						 0.0, 0.0, 1.0, 0.0,
						 0.0, 0.0, 0.0, 1.0 );
        //TODO OPT rotationMat is not used, could be removed
		rotationMat = mat4( 1.0, 0.0, 0.0, 0.0,
                         0.0, 1.0, 0.0, 0.0,
                         0.0, 0.0, 1.0, 0.0,
                         0.0, 0.0, 0.0, 1.0 );

        fragCoordTransformMat = translationMat * rotationMat * scaleMat;
        transformMat = inOrtho * fragCoordTransformMat;
        //transformMat = inOrtho * translationMat * rotationMat * scaleMat;

        geomFragCoord = vec2(fragCoordTransformMat * tempVert);
        gl_Position = transformMat * tempVert;
        //gl_Position = vec4(-1.0, -1.0, 0.0, 1.0);
        EmitVertex();
        tempVert.x *= -1.0;
        geomFragCoord = vec2(fragCoordTransformMat * tempVert);
        gl_Position = transformMat * tempVert;
        //gl_Position = vec4(1.0, -1.0, 0.0, 1.0);
        EmitVertex();
        tempVert.xy *= -1.0;
        geomFragCoord = vec2(fragCoordTransformMat * tempVert);
        gl_Position = transformMat * tempVert;
        //gl_Position = vec4(-1.0, 1.0, 0.0, 1.0);
        EmitVertex();
        tempVert.x *= -1.0;
        geomFragCoord = vec2(fragCoordTransformMat * tempVert);
        gl_Position = transformMat * tempVert;
        //gl_Position = vec4(-1.0, -1.0, 0.0, 1.0);
        EmitVertex();
	}
	EndPrimitive();
}
