# Online Shooter Duel

A multiplayer terminal-based shooter game for two players, where one acts as server (host) and the other as client.

## Features

- **Network Multiplayer**: Game for two players connected via TCP
- **Terminal Interface**: Uses the termbox-go library for a graphical terminal interface
- **Health System**: Each player has 3 lives
- **Collision Detection**: Bullets can hit players
- **Real-time Synchronization**: Game state synchronized between server and client

## Installation

1. Make sure you have Go installed (version 1.24.1 or higher)
2. Clone or download this repository
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Compile the game:

   ```bash
   go build -o onlinegame
   ```

   Or simply:

   ```bash
   go build
   ```

5. (Optional) Run tests:
   ```bash
   go test ./...
   ```

## How to Play

### Initial Setup

1. Run the game:

   ```bash
   ./onlinegame
   ```

   Or if you used `go build` without specifying output name:

   ```bash
   ./online
   ```

2. In the main menu:
   - **Create Room (Host)**: To create a room as server
   - **Join Room (Client)**: To join a room as client
   - **Exit Game**: To quit the application

### As Host (Server)

1. Select "Create Room (Host)" in the menu
2. The game will show "Creating room..." and then "Waiting for player..."
3. Wait for the client to connect
4. Once connected, the game will start automatically

### As Client

1. Select "Join Room (Client)" in the menu
2. Enter the host's IP (or press Enter to use localhost)
3. The game will attempt to connect to the server
4. Once connected, the game will start automatically

### Controls

#### In-Game Controls:

- **A**: Move left
- **D**: Move right
- **J**: Shoot
- **Q**: Quit game
- **ESC**: Quit game

#### Menu Controls:

- **Arrow Up/Down**: Navigate menu options
- **Enter**: Select option
- **ESC**: Exit game

#### Game Over Screen:

- **R**: Restart game
- **Q**: Quit game
- **ESC**: Quit game

### Objective

- Eliminate your opponent by shooting them
- Each player has 3 lives
- The last player with life wins

## Code Structure

The code is organized in separate modules for better maintainability:

### Main Packages:

- **`game/`**: Game logic

  - `types.go`: Data structures (Player, Bullet, GameState, etc.)
  - `logic.go`: Game logic (initialization, update, collisions, etc.)

- **`network/`**: Network communication

  - `connection.go`: TCP connection handling, data sending/receiving

- **`ui/`**: User interface

  - `render.go`: Sprite rendering, menus and screens

- **`core/`**: Basic functions

  - `input.go`: User input handling

- **`main.go`**: Main entry point and state machine

### Architecture Features:

- **Separation of Responsibilities**: Each package has a specific function
- **Modularity**: Easy to maintain and extend
- **Reusability**: Packages can be reused in other projects
- **Testability**: Each module can be tested independently

## Troubleshooting

### Connection Error

- Check that the firewall is not blocking port 8080
- Make sure both players are on the same network
- To play over the internet, configure port forwarding on your router

### Compilation Issues

- Make sure Go is installed correctly
- Run `go mod tidy` to install dependencies
- Verify that the Go version is compatible

### Network Issues

- The game uses port 8080 by default
- If the port is occupied, modify the `port` constant in the code
- To play over the internet, use the host's public IP

## Technologies Used

- **Go**: Main programming language
- **termbox-go**: Terminal interface library
- **TCP**: Network protocol for communication between players
- **JSON**: Format for game state serialization

## License

This project is open source and available under the MIT license.
