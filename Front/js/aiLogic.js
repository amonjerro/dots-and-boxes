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

function pickAIEdge(){
    let randomNode = nodeToCoordinates(nodes[Math.floor(Math.random()*nodes.length-1)])
    if (randomNode[0] == maxGridX || randomNode[1] == maxGridY){
        return null
    }
    let checkRightFirst = Math.floor(Math.random()*10)%2
    if (checkRightFirst){
        let edgeKey = `${randomNode[0]},${randomNode[1]}-${randomNode[0]+1},${randomNode[1]}`
        let status = edges.get(edgeKey)
        if (status){
            edgeKey = `${randomNode[0]},${randomNode[1]}-${randomNode[0]},${randomNode[1]+1}`
            status = edges.get(edgeKey)
            if (status){
                return null
            }
        }
        return edgeKey
    } else {
        let edgeKey = `${randomNode[0]},${randomNode[1]}-${randomNode[0]},${randomNode[1]+1}`
        let status = edges.get(edgeKey)
        if (status){
            edgeKey = `${randomNode[0]},${randomNode[1]}-${randomNode[0]+1},${randomNode[1]}`
            status = edges.get(edgeKey)
            if (status){
                return null
            }
        }
        return edgeKey
    }
}

function runAITurn(){
    let edge = earlyDecisionMaking()
    if (edge == null){
        console.error("No edge found!")
    } else {
        processTurn(edge)
    }
}