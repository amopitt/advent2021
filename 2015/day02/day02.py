
def getSurfaceArea(l, w, h):
    return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)

def getShortestPerimeter(l, w, h):
    return min(2*l + 2*w, 2*w + 2*h, 2*h + 2*l) + l*w*h

totalArea = 0
totalShortestPerimeter = 0
# read a file one line at a time
with open('input.txt') as f:
    for index, line in enumerate(f):
        test = line.strip().split('x')
        l = int(test[0])
        w = int(test[1])
        h = int(test[2])
        totalArea += getSurfaceArea(l, w, h)
        totalShortestPerimeter += getShortestPerimeter(l, w, h)

print(totalArea)
print(totalShortestPerimeter)


 

