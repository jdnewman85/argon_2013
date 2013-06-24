#version 430 core

in vec2 geomPosition;
in vec2 geomFragCoord;
in float geomRadius;
in vec4 geomColor;

layout(location = 0) out vec4 outColor;

void main() {
        //TODO OPT This if and discard are bad... need to 'blend' out instead of discard
        if(distance(geomFragCoord, geomPosition) > geomRadius)
            discard;
    float fragDist = distance(geomFragCoord, geomPosition);
    //fragDist = step(.001,clamp(geomRadius-fragDist,0.0,1.0));
    fragDist = clamp(geomRadius-fragDist, 0.0, 1.0);
    outColor = vec4(geomColor.rgb, geomColor.a * fragDist);
}
