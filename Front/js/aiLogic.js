function pickAIEdge(){
    let randomNode = nodeToCoordinates(nodes[Math.floor(Math.random()*nodes.length-1)])
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
    let edge = null
    while(edge == null){
        edge = pickAIEdge()
    }
    processTurn(edge)
}