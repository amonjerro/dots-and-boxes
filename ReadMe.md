These are the instructions and how-to of getting the environment setup on your local machine for testing and development. Production and deployment will have a different setup.

# Setting up

1. Install [Docker Desktop](https://docs.docker.com/desktop/). It includes the Docker and Docker Compose CLI tools.
2. Create a .env file at the root level, refer to the Environment section for details
3. On any directory within the project, use the command `docker compose up -d` to create and start up the application
4. Start hacking.

The client container is set up to restart and update if any changes are detected on its codebase. The other ones aren't.
On Docker Desktop, you can monitor the state and health of any service in the container.

# Architecture

To support the online gaming experience it is important to separate game logic from the client. Clients and the backend server communicate through a messaging broker, keeping them decoupled to allow for scaling and better error management.

The client shows the state of the game to the players and sends game actions to the back end. The web client and the front end server communicate through web sockets, to ensure that any updates that are received by the front end server are updated on the web client side.

The back end server processes the received game actions and updates the game state for all active games. On any succesfull game state update, it emits messages so front end servers update their representations of the game and update the web client for all end users.

# Environment

All .env files require the following environment variables determined specifications:

```
CLIENT_PORT = The client port number. It should be an integer.
SERVER_PORT = The server port number. It should be an integer-
BROKER_PORT = The port number for the message broker. It should be an integer.
BROKER_USER = The string that identifies the message broker user.
BROKER_PASS = The string that represents that user's password
```

# The game

Dots and boxes is a classic notebook game. 

On a grid of nodes, player take turns painting edges between adjacent nodes. Any player that paints the edge that encases a square made up of four nodes scores a point and gets to go again.

The player with the highest score at the end wins.