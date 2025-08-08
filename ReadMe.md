These are the instructions and how-to of getting the environment setup on your local machine for testing and development. Production and deployment will have a different setup.

# Setting up (instructions still underway)

1. Install [Docker Desktop](https://docs.docker.com/desktop/). It includes the Docker and Docker Compose CLI tools.
2. Pull the messaging broker image by running `docker pull rabbitmq:management-alpine`
3. Something something docker compose to get all the containers up and running (TO DO)
4. Start hacking.

# Architecture

To support the online gaming experience it is important to separate game logic from the client. Clients and the backend server communicate through a messaging broker, keeping them decoupled to allow for scaling and better error management.

The client shows the state of the game to the players and sends game actions to the back end. The web client and the front end server communicate through web sockets, to ensure that any updates that are received by the front end server are updated on the web client side.

The back end server processes the received game actions and updates the game state for all active games. On any succesfull game state update, it emits messages so front end servers update their representations of the game and update the web client for all end users.

# The game

Dots and boxes is a classic notebook game. 

On a grid of nodes, player take turns painting edges between adjacent nodes. Any player that paints the edge that encases a square made up of four nodes scores a point and gets to go again.

The player with the highest score at the end wins.