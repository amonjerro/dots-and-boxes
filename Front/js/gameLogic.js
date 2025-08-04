var canvasWidth = 600
var canvasHeight = 400
const deadZoneTolerance = 0.15


var gridWidth = 20
var maxGridX = canvasWidth / gridWidth - 1
var gridHeight = 20
var maxGridY = canvasHeight / gridHeight - 1
var squareSize = 4

// Data structures
var squares = {}
var nodes = []
var edges =  new Map()

var edgesToPaint = []
var squaresToPaint = []

// Player Information
const MAX_PLAYERS = 2;
var currentPlayer = 0;

function gameSetup(){
    // Set up nodes
    for (let i = 1; i < canvasWidth/gridWidth; i++){
        for (let j = 1; j < canvasHeight / gridHeight; j++){
            let coordinateString = `${i},${j}`
            nodes.push(coordinateString)
        }
    }

    for (let i = 1; i < (canvasWidth/gridWidth)-1; i++){
        for (let j = 1; j < (canvasHeight / gridHeight)-1; j++){
            let coordinateString = `${i},${j}`
            let left = coordinateString + `-${i+1},${j}`
            let bot = coordinateString + `-${i},${j+1}`
            edges.set(left, false)
            edges.set(bot, false)
        }
    }

    for (let i = 1; i < (canvasWidth/gridWidth)-1; i++){
        for (let j = 1; j < (canvasHeight / gridHeight)-1; j++){
            let coordinateString = `${i},${j}`
            let topEdge = coordinateString + `-${i+1},${j}`
            let botEdge = `${i},${j+1}` + `-${i+1},${j+1}`
            let leftEdge = coordinateString + `-${i},${j+1}`
            let rightEdge = `${i+1},${j}` + `-${i+1},${j+1}`
            squares[coordinateString] = [leftEdge, topEdge, rightEdge, botEdge]
            
        }
    }
}


function rollTurn(){
    currentPlayer += 1;
    if (currentPlayer == MAX_PLAYERS){
        currentPlayer = currentPlayer % MAX_PLAYERS;
    }
}

function edgeToCoordinates(edgeKey){
    let edgeEnds = edgeKey.split('-')
    let coordsOrigin = edgeEnds[0].split(',').map((e)=> {return parseInt(e)})
    let coordsEnd = edgeEnds[1].split(',').map((e)=> {return parseInt(e)})
    return [coordsOrigin[0], coordsOrigin[1], coordsEnd[0], coordsEnd[1]]
}

function isHorizontalEdge(edgeKey){
    let coords = edgeToCoordinates(edgeKey)
    return coords[0] < coords[2] && coords[1] == coords[3]
}

function testIsSquareComplete(coordinates){
    let squareToTest = squares[`${coordinates[0]},${coordinates[1]}`]
    let complete = true
    for(let i = 0; i < squareToTest.length; i++){
        if (!edges.get(squareToTest[i])){
            complete = false
            break
        }
    }
    return complete
}

function pickEdge(xCoord, yCoord){
    let wholeX = Math.round(xCoord)
    let remainderX = xCoord - wholeX

    let wholeY = Math.round(yCoord)
    let remainderY = yCoord - wholeY

    let endX = wholeX
    let endY = wholeY
    console.log(remainderX, remainderY)
    if (wholeX == 0 || wholeY == 0 || wholeX > maxGridX || wholeY > maxGridY){
        console.log("Out of bounds")
        return;
    }
    
    if (Math.abs(remainderX) > deadZoneTolerance && Math.abs(remainderY) > deadZoneTolerance ){
        console.log("DeadZone")
        return;
    } else if (Math.abs(remainderY) > deadZoneTolerance) {
        // This is a vertical edge check
        if (Math.sign(remainderY) == -1){
            wholeY -= 1
        } else {
            endY += 1
        }
    } else {
        // This is a horizontal edge check
        if (Math.sign(remainderX) == -1){
            wholeX -= 1
        } else {
            endX += 1
        }
    }
    let key = `${wholeX},${wholeY}-${endX},${endY}`
    edgesToPaint.push(key)
    edges.set(key, true)
    return key
}

function evaluateMove(edge){
    let edgeCoords = edgeToCoordinates(edge)
    if (isHorizontalEdge(edge)){
        if (edgeCoords[1] > 1){
            if (testIsSquareComplete([edgeCoords[0], edgeCoords[1]-1])){
                squaresToPaint.push({x:edgeCoords[0], y:edgeCoords[1]-1})
            }
        }
        if (edgeCoords[1] < maxGridY){
            if (testIsSquareComplete([edgeCoords[0], edgeCoords[1]])){
                squaresToPaint.push({x:edgeCoords[0], y:edgeCoords[1]})
            }
        }
    } else {
        if (edgeCoords[0] > 1){
            if (testIsSquareComplete([edgeCoords[0]-1, edgeCoords[1]])){
                squaresToPaint.push({x:edgeCoords[0]-1, y:edgeCoords[1]})
            }
        }
        if (edgeCoords[0] < maxGridX){
            if (testIsSquareComplete([edgeCoords[0], edgeCoords[1]])){
                squaresToPaint.push({x:edgeCoords[0], y:edgeCoords[1]})
            }
        }
    }
}