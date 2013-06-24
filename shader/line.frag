#version 430 core

in vec2 geomPoint1;
in vec2 geomPoint2;
in vec2 geomFragCoord;
in float geomRadius;
in vec4 geomColor;

layout(location = 0) out vec4 outColor;

vec2 closestPointInLine(vec2,vec2,vec2);
vec2 closestPointInLineSeg(vec2,vec2,vec2);




void main() {
        //TODO OPT This if and discard are bad... need to 'blend' out instead of discard
        //if(distance(geomFragCoord, closestPointInLineSeg(geomPoint1, geomPoint2, geomFragCoord)) > geomRadius)
        //    discard;
    float fragDist = distance(geomFragCoord,closestPointInLineSeg(geomPoint1, geomPoint2, geomFragCoord));
    //fragDist = step(.001,clamp(geomRadius-fragDist,0.0,1.0));
    fragDist = clamp(geomRadius - fragDist, 0.0, 1.0);
    outColor = vec4(geomColor.rgb, geomColor.a * fragDist);
}




//Returns a point in line, that is closest to point
vec2 closestPointInLine(in vec2 lineP0, in vec2 lineP1, in vec2 P) {
    vec2 v,w;
    float c1,c2,b;

    v = lineP1 - lineP0;
    w = P - lineP0;

    c1 = dot(w, v);
    c2 = dot(v, v);
    b = c1 / c2;

    return lineP0 + b*v;
}

//Returns a point in line segment that is closest to point
vec2 closestPointInLineSeg(in vec2 lineP0, in vec2 lineP1, in vec2 P) {
    vec2 v, w;
    float c1, c2, b;

    v = lineP1 - lineP0;
    w = P - lineP0;

    c1 = dot(w, v);
    c2 = dot(v, v);
    b = clamp(c1 / c2, 0.0, 1.0);

    return lineP0 + b * v;
}
