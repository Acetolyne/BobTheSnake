package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"fmt"
	"log"
	"math"
)

var possibleMoves map[string]bool
var CurrentTechnique string
var uprightFood []Coord
var downleftFood []Coord

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

func dontCollideSelf(state GameState, possibleMoves map[string]bool) map[string]bool {
	var newPos Coord

	for k, v := range possibleMoves {
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
					possibleMoves[(k)] = false
				}
			}
		}
	}
	return possibleMoves
}

//Takes in a map of possible moves and returns the direction to get to the closest food possible
func italiansnake(state GameState, possibleMoves *map[string]bool) string {
	var dir string
	FoodDist := map[string]int{
		"up":    1000,
		"down":  1000,
		"left":  1000,
		"right": 1000,
	}
	//@todo If there is only one move then return that move else decide the best one to get us some food
	for k := range *possibleMoves {
		switch k {
		case "up":
			curDist := 1000
			for _, v := range state.Board.Food {
				if v.Y > state.You.Body[0].Y {
					dist := math.Abs(float64(state.You.Body[0].X-v.X)) + math.Abs(float64(state.You.Body[0].Y-v.Y))
					if int(dist) < curDist {
						curDist = int(dist)
					}
				}
			}
			FoodDist["up"] = curDist
		case "down":
			curDist := 1000
			for _, v := range state.Board.Food {
				if v.Y < state.You.Body[0].Y {
					dist := math.Abs(float64(state.You.Body[0].X-v.X)) + math.Abs(float64(state.You.Body[0].Y-v.Y))
					if int(dist) < curDist {
						curDist = int(dist)
					}
				}
			}
			FoodDist["down"] = curDist
		case "left":
			curDist := 1000
			for _, v := range state.Board.Food {
				if v.X < state.You.Body[0].X {
					dist := math.Abs(float64(state.You.Body[0].X-v.X)) + math.Abs(float64(state.You.Body[0].Y-v.Y))
					if int(dist) < curDist {
						curDist = int(dist)
					}
				}
			}
			FoodDist["left"] = curDist
		case "right":
			curDist := 1000
			for _, v := range state.Board.Food {
				if v.X > state.You.Body[0].X {
					dist := math.Abs(float64(state.You.Body[0].X-v.X)) + math.Abs(float64(state.You.Body[0].Y-v.Y))
					if int(dist) < curDist {
						curDist = int(dist)
					}
				}
			}
			FoodDist["right"] = curDist
		}
	}
	//pick a default direction that is possible
	if (*possibleMoves)["up"] == true {
		dir = "up"
	} else if (*possibleMoves)["down"] == true {
		dir = "down"
	} else if (*possibleMoves)["left"] == true {
		dir = "left"
	} else if (*possibleMoves)["right"] == true {
		dir = "right"
	}
	min := FoodDist["up"]
	for k, v := range FoodDist {
		if v < min {
			dir = k
		}
	}

	return dir
}

// This function is called on every turn of a game.
func move(state GameState, curmethod *string) BattlesnakeMoveResponse {
	var usedmethod string
	//Before we move check where we are
	fmt.Println(state.You.Body[0])
	//@todo if we are low on health better find food else choose a random technique for x amount of time.
	fmt.Println("HEALTH:", state.You.Health)
	fmt.Println("FOOD:", state.Board.Food)
	fmt.Println("METHOD:" + *curmethod)
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	fmt.Println("1st", &possibleMoves)
	myHead := state.You.Body[0] // Coordinates of your head
	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height
	if myHead.X == 0 {
		delete(possibleMoves, "left")
	}
	if myHead.X == boardWidth-1 {
		delete(possibleMoves, "right")
	}
	if myHead.Y == 0 {
		delete(possibleMoves, "down")
	}
	if myHead.Y == boardHeight-1 {
		delete(possibleMoves, "up")
	}
	//possibleMoves = dontCollideSelf(state, possibleMoves)

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.
	//@todo always avoid others except maybe their head?
	for _, snake := range state.Board.Snakes {
		fmt.Println("HEAD:", state.You.Body[0])
		for _, a := range snake.Body {
			fmt.Println("a", a)
			for k, v := range possibleMoves {
				if v == true {
					switch k {
					case "up":
						if state.You.Body[0].Y+1 == a.Y && state.You.Body[0].X == a.X {
							delete(possibleMoves, k)
						}
					case "down":
						if state.You.Body[0].Y-1 == a.Y && state.You.Body[0].X == a.X {
							delete(possibleMoves, k)
						}
					case "left":
						if state.You.Body[0].X-1 == a.X && state.You.Body[0].Y == a.Y {
							delete(possibleMoves, k)
						}
					case "right":
						if state.You.Body[0].X+1 == a.X && state.You.Body[0].Y == a.Y {
							delete(possibleMoves, k)
						}
					}
				}
			}
		}
	}
	var nextMove string
	//@todo change health comparison back to 25 after done testing italiansnake
	if state.You.Health <= 100 {
		usedmethod = "italiansnake"
	} else {
		usedmethod = *curmethod
	}
	switch usedmethod {
	case "italiansnake":
		fmt.Println("ITALIAN")
		nextMove = italiansnake(state, &possibleMoves)
	case "shysnake":
		fmt.Println("SHY")
		//@todo run from other snakes
	case "middlesnake":
		fmt.Println("MIDDLE")
		//@todo stay away from edges
	case "rudesnake":
		fmt.Println("RUDE")
		//@todo charge other snakes
	}

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} //else {
	//next move is a random move from safeMoves list
	//nextMove = safeMoves[rand.Intn(len(safeMoves))]
	//log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	//}
	fmt.Println("MOVING:" + nextMove)
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
