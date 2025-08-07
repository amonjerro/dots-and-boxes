var canvasWidth = 600
var canvasHeight = 400
const deadZoneTolerance = 0.15
var shouldRollTurn = false
var isGameOngoing = true

var gridWidth = 50
var maxGridX = canvasWidth / gridWidth - 1
var gridHeight = 50
var maxGridY = canvasHeight / gridHeight - 1
var squareSize = 4

// Data structures
var squares = {}
var squareCount = 0
var nodes = []
var edges =  new Map()

var squaresByOpenEdges = {
    0: new Set(),
    1: new Set(),
    2: new Set(),
    3: new Set(),
    4: new Set()
}

var edgesToPaint = []
var squaresToPaint = []

// Player Information
const MAX_PLAYERS = 2;
var currentPlayer = 0;

var score = []

function gameSetup(){
    // Set up nodes
    for (let i = 1; i < canvasWidth/gridWidth; i++){
        for (let j = 1; j < canvasHeight / gridHeight; j++){
            let coordinateString = `${i},${j}`
            nodes.push(coordinateString)
        }
    }

    for (let i = 1; i < (canvasWidth/gridWidth)-1; i++){
        let top = `${i},${1}-${i+1},1`
        let bot = `${i},${canvasHeight /gridHeight -1}-${i+1},${canvasHeight /gridHeight -1}`
        edges.set(top, true)
        edgesToPaint.push(top)
        edges.set(bot, true)
        edgesToPaint.push(bot)
    }

    for (let i = 1; i < (canvasHeight/gridHeight)-1; i++){
        let top = `${1},${i}-1,${i+1}`
        let bot = `${canvasWidth /gridWidth -1},${i}-${canvasWidth /gridWidth -1},${i+1}`
        edges.set(top, true)
        edgesToPaint.push(top)
        edges.set(bot, true)
        edgesToPaint.push(bot)
    }

    for (let i = 1; i < (canvasWidth/gridWidth)-1; i++){
        for (let j = 1; j < (canvasHeight / gridHeight)-1; j++){
            let coordinateString = `${i},${j}`
            let topEdge = coordinateString + `-${i+1},${j}`
            let botEdge = `${i},${j+1}` + `-${i+1},${j+1}`
            let leftEdge = coordinateString + `-${i},${j+1}`
            let rightEdge = `${i+1},${j}` + `-${i+1},${j+1}`
            let count = [leftEdge, topEdge, rightEdge, botEdge].reduce((acc, curr)=>
                    { 
                        let v = edges.get(curr) ? 0 : 1
                        return acc + v
                    }, 0)
            squares[coordinateString] = {
                'edges':[leftEdge, topEdge, rightEdge, botEdge],
                'openConnectionCount': count
            }
            squaresByOpenEdges[count].add(coordinateString)
            
        }
    }

    squareCount = Object.keys(squares).length

    for (let i = 0; i < MAX_PLAYERS; i++){
        score.push(0)
    }
}


function rollTurn(){
    currentPlayer += 1;
    if (currentPlayer == MAX_PLAYERS){
        currentPlayer = currentPlayer % MAX_PLAYERS;
    }
    shouldRollTurn = false
}

function processTurn(pickedEdge){
    if (pickedEdge){
        edgesToPaint.push(pickedEdge)
        edges.set(pickedEdge, true)
        let squareCompleted = evaluateMove(pickedEdge)
        console.log(squareCompleted)
        shouldRollTurn = !squareCompleted
    }
}

function nodeToCoordinates(node){
    return node.split(',').map((e)=>{return parseInt(e)})
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

function updateSquareEdgeCount(squareCoordinates){
    let key = `${squareCoordinates[0]},${squareCoordinates[1]}`
    let square = squares[key]
    if (!square){
        console.error(coordinates)
    }

    // manipulate the sets
    squaresByOpenEdges[square.openConnectionCount].delete(key)
    squares[key].openConnectionCount -= 1
    squaresByOpenEdges[squares[key].openConnectionCount].add(key)

}

function testIsSquareComplete(coordinates){
    let squareToTest = squares[`${coordinates[0]},${coordinates[1]}`]
    if (!squareToTest){
        console.error(coordinates)
    }
    let complete = squareToTest.openConnectionCount == 0
    return complete
}

function pickEdge(xCoord, yCoord){
    let wholeX = Math.round(xCoord)
    let remainderX = xCoord - wholeX

    let wholeY = Math.round(yCoord)
    let remainderY = yCoord - wholeY

    let endX = wholeX
    let endY = wholeY
    //console.log(remainderX, remainderY)
    if (wholeX == 0 || wholeY == 0 || wholeX > maxGridX || wholeY > maxGridY){
        //console.log("Out of bounds")
        return;
    }
    
    if (Math.abs(remainderX) > deadZoneTolerance && Math.abs(remainderY) > deadZoneTolerance ){
        //console.log("DeadZone")
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
    if (edges.get(key)){
        return
    }
    return key
}

function evaluateSquare(coordinates){
    let extraTurn = false
    updateSquareEdgeCount(coordinates)
    if (testIsSquareComplete(coordinates)){
        squaresToPaint.push({x:coordinates[0], y:coordinates[1]})
        score[currentPlayer] += 1
        extraTurn = true
    }
    return extraTurn
}

function evaluateMove(edge){
    let edgeCoords = edgeToCoordinates(edge)
    let extraTurn = false
    if (isHorizontalEdge(edge)){
        if (edgeCoords[1] > 1){
            extraTurn = evaluateSquare([edgeCoords[0], edgeCoords[1]-1]) || extraTurn
        }
        if (edgeCoords[1] < maxGridY){
            extraTurn = evaluateSquare([edgeCoords[0], edgeCoords[1]]) || extraTurn
        }
    } else {
        if (edgeCoords[0] > 1){
            extraTurn = evaluateSquare([edgeCoords[0]-1, edgeCoords[1]]) || extraTurn
        }
        if (edgeCoords[0] < maxGridX){
            extraTurn = evaluateSquare([edgeCoords[0], edgeCoords[1]]) || extraTurn
        }
    }
    return extraTurn
}
