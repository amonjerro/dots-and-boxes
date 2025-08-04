
// Player Information
const colors =  ['#DEADBF',"#BADDAD","#DADB0D",'#FEEDBF']
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

function paintSquare(squareLocation){
    fill(colors[currentPlayer])
    stroke(colors[currentPlayer])
    rect(squareLocation.x * gridWidth, squareLocation.y * gridHeight, gridWidth, gridHeight)
}

function canvasToGrid(xPos, yPos){
    return [xPos / gridWidth, yPos / gridHeight]
}

function drawNodes(){
    nodes.map((key)=>{
        let coordinates = key.split(',')
        let xCoord = (parseInt(coordinates[0])*gridWidth) - squareSize/2
        let yCoord = (parseInt(coordinates[1])*gridHeight) - squareSize/2
        square(xCoord, yCoord, squareSize)
    })
}

function setup(){
    frameRate(60)
    let canvas = createCanvas(canvasWidth, canvasHeight)
    canvas.parent("p5-canvas")
    gameSetup()
    background(100)
    drawNodes()
}

function draw(){
    let edge = edgesToPaint.pop()
    if (edge){
        paintEdge(edge)
    }
    let square = squaresToPaint.pop()
    if (square){
        paintSquare(square)
    }
    if (shouldRollTurn){
        rollTurn()
    }
    if (currentPlayer == 1){
        runAITurn()
    }
}

function mouseClicked(){
    if (mouseX > canvasWidth || mouseY > canvasHeight){
        return
    }

    if (currentPlayer != 0){
        return
    }
    let coordinates = canvasToGrid(mouseX, mouseY)
    let pickedEdge = pickEdge(coordinates[0], coordinates[1])
    processTurn(pickedEdge)
}