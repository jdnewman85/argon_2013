#version 430 core

layout( points ) in;
layout( triangle_strip, max_vertices = 4 ) out;

uniform mat4 inOrtho;

in vec2 vertPoint1[1];
in vec2 vertPoint2[1];
in float vertRadius[1];
in vec4 vertColor[1];

out vec2 geomPoint1;
out vec2 geomPoint2;
out vec2 geomFragCoord;
out float geomRadius;
out vec4 geomColor;

//TODO Finish from here
void main( void ) {
    int i;
    vec4 tempVert;
    mat4 translationMat;
    mat4 rotationMat;
    mat4 scaleMat;
    mat4 transformMat;
    float angle;
    vec2 lineVec;
    mat4 fragCoordTransformMat;

    for ( i = 0; i < gl_in.length(); i++ ) {
        geomPoint1 = vertPoint1[i];
        geomPoint2 = vertPoint2[i];
        geomRadius = vertRadius[i];
        geomColor = vertColor[i];


        lineVec = vertPoint2[i]-vertPoint1[i];
        tempVert = vec4( (length(lineVec) + vertRadius[i]*2.0) / -2.0, -vertRadius[i], 0.0, 1.0 );
        translationMat = mat4( 1.0, 0.0, 0.0, 0.0,
                               0.0, 1.0, 0.0, 0.0,
                               0.0, 0.0, 1.0, 0.0,
                               gl_in[i].gl_Position.x, gl_in[i].gl_Position.y, 0.0, 1.0 );
        //TODO OPT scaleMat is not used, could be removed
        scaleMat = mat4( 1.0, 0.0, 0.0, 0.0,
                         0.0, 1.0, 0.0, 0.0,
                         0.0, 0.0, 1.0, 0.0,
                         0.0, 0.0, 0.0, 1.0 );
        angle = atan(-lineVec.y, lineVec.x); //BUG Not sure why I have to invert y
        rotationMat = mat4( cos(  angle ), -sin( angle ), 0.0, 0.0,
                            sin( angle ), cos( angle ), 0.0, 0.0,
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
