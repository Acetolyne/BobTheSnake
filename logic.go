package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"fmt"
	"log"
	"math/rand"
)

var possibleMoves map[string]bool

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Acetolyne", // TODO: Your Battlesnake username
		Color:      "#C0C0C0",   // TODO: Personalize
		Head:       "default",   // TODO: Personalize
		Tail:       "default",   // TODO: Personalize
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

func checkNextMove(state GameState, possibleMoves map[string]bool) map[string]bool {
	var newPos Coord
	fmt.Println("3rd", possibleMoves)
	fmt.Println("HEAD:", state.You.Body[0])

	for k, v := range possibleMoves {
		fmt.Println("Key", k)
		if v == true {
			switch k {
			case "up":
				newPos = Coord{state.You.Body[0].X, state.You.Body[0].Y + 1}
			case "down":
				newPos = Coord{state.You.Body[0].X, state.You.Body[0].Y - 1}
			case "left":
				newPos = Coord{state.You.Body[0].X - 1, state.You.Body[0].Y}
			case "right":
				newPos = Coord{state.You.Body[0].X + 1, state.You.Body[0].Y}
			}
			for _, v := range state.You.Body {
				if v == newPos {
					fmt.Println("UPDATE", k)
					possibleMoves[(k)] = false
				}
			}
		}
	}
	return possibleMoves
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	//Before we move check where we are
	fmt.Println(state.You.Body[0])
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	fmt.Println("1st", &possibleMoves)
	myHead := state.You.Body[0] // Coordinates of your head
	fmt.Println("2nd", &possibleMoves)
	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height
	fmt.Println("width", boardWidth)
	fmt.Println("height", boardHeight)
	if myHead.X == 0 {
		possibleMoves["left"] = false
	}
	if myHead.X == boardWidth-1 {
		possibleMoves["right"] = false
	}
	if myHead.Y == 0 {
		possibleMoves["down"] = false
	}
	if myHead.Y == boardHeight-1 {
		possibleMoves["up"] = false
	}
	possibleMoves = checkNextMove(state, possibleMoves)
	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	// mybody := state.You.Body

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	fmt.Println("4th", &possibleMoves)
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		//next move is a random move from safeMoves list
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
