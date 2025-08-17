
// Player Information
const colors =  ['#DEADBF',"#BADDAD","#DADB0D",'#FEEDBF']
const edgeWeight = 2

var turnReminder;
var p1scoreDOMElement; 
var p2scoreDOMElement; 

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

function getDOMElements(){
    p1scoreDOMElement = document.querySelector('#player-1-score')
    p2scoreDOMElement = document.querySelector('#player-2-score')
    turnReminder = document.querySelector('#turn-reminder')
}

function paintColorsOnCards(){
    let p1colorElement = document.querySelector('#player-1-colors')
    p1colorElement.style.backgroundColor = colors[0]
    let p2colorElement = document.querySelector('#player-2-colors')
    p2colorElement.style.backgroundColor = colors[1]
}

function setup(){
    frameRate(60)
    let canvas = createCanvas(canvasWidth, canvasHeight)
    canvas.parent("p5-canvas")
    gameSetup()
    getDOMElements()
    paintColorsOnCards()
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
    update()
    if (isGameOngoing){
        if (shouldRollTurn){
            rollTurn()
        }
        if (currentPlayer == 1){
            turnReminder.innerHTML = ''
            aiWaitTimer += 1
            if (aiWaitTimer > waitTime){
                runAITurn()
                aiWaitTimer = 0
            }
        } else { 
            turnReminder.innerHTML = "It's your turn"
        }
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

function update(){
    p1scoreDOMElement.innerHTML = score[0]
    p2scoreDOMElement.innerHTML = score[1]

    if (squaresByOpenEdges[0].size == squareCount && isGameOngoing){
        isGameOngoing = false
        gameOverFanfare()
    }
}

function gameOverFanfare(){
    let p1WinnerBox = document.querySelector('#p1-winner-box')
    let p2WinnerBox = document.querySelector('#p2-winner-box')
    if (score[0]>score[1]){
        p1WinnerBox.classList.add("winner")
        p1WinnerBox.innerHTML = "Winner!"
    } else {
        p2WinnerBox.classList.add("winner")
        p2WinnerBox.innerHTML = "Winner!"
    }
}