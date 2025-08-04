const waitTime = 10
var aiWaitTimer = 0

function getEdgeFromOpenEdgeSet(openEdgeCount){
    if (squaresByOpenEdges[openEdgeCount].size == 0){
        return null
    }
    let setKeys = squaresByOpenEdges[openEdgeCount].values().toArray()
    let randomSquare = setKeys[Math.floor(Math.random()*setKeys.length)]
    let square = squares[randomSquare]
    for(let i = 0; i < square.edges.length; i++){
        if (!edges.get(square.edges[i])){
            return square.edges[i]
        }
    }
}


function earlyDecisionMaking(){
    let edge = getEdgeFromOpenEdgeSet(1)
    if (!edge){
        edge = getEdgeFromOpenEdgeSet(4)
    }
    if (!edge){
        edge = getEdgeFromOpenEdgeSet(3)
    }
    if (!edge){
        edge = getEdgeFromOpenEdgeSet(2)
    }

    return edge
}

function runAITurn(){
    let edge = earlyDecisionMaking()
    if (edge == null){
        console.error("No edge found!")
    } else {
        processTurn(edge)
    }
}