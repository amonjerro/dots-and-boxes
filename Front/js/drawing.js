
// Player Information
const colors =  ['#DEADBF','#FEEDBF',"#BADDAD","#DADB0D"]
const edgeWeight = 2

function paintEdge(edgeKey){
    strokeWeight = edgeWeight
    stroke(colors[currentPlayer])
    let edgeCoordinates = edgeKey.split('-')
    let coordsOrigin = edgeCoordinates[0].split(',').map((e)=> {return parseInt(e)})
    let coordsEnd = edgeCoordinates[1].split(',').map((e)=> {return parseInt(e)})
    let pixelOriginX = coordsOrigin[0] * gridWidth
    let pixelOriginY = coordsOrigin[1] * gridHeight
    let pixelEndX = coordsEnd[0] * gridWidth
    let pixelEndY = coordsEnd[1] * gridHeight
    line(pixelOriginX, pixelOriginY, pixelEndX, pixelEndY)
}

function canvasToGrid(xPos, yPos){
    return [xPos / gridWidth, yPos / gridHeight]
}

function drawNodes(){
    nodes.map((key)=>{
        let coordinates = key.split(',')
        let xCoord = (parseInt(coordinates[0])*gridWidth) - squareSize/2
        let yCoord = (parseInt(coordinates[1])*gridWidth) - squareSize/2
        square(xCoord, yCoord, squareSize)
    })
}

function setup(){
    frameRate(2)
    let canvas = createCanvas(canvasWidth, canvasHeight)
    canvas.parent("p5-canvas")
    gameSetup()
    background(220)
    drawNodes()
}

function draw(){
    let edge = edgesToPaint.pop()
    if (edge){
        paintEdge(edge)
    }

}

function mouseClicked(){
    if (mouseX > canvasWidth || mouseY > canvasHeight){
        return
    }
    let coordinates = canvasToGrid(mouseX, mouseY)
    pickEdge(coordinates[0], coordinates[1])
}