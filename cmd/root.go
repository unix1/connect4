package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unix1/connect4/game"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "connect4",
		Short: "Simple connect 4 game in Go CLI",
		RunE:  run,
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) error {
	// create a new game
	g, err := game.NewGame()
	if err != nil {
		return err
	}
	// parse moves
	moves, err := parseMoves(args)
	if err != nil {
		return err
	}
	// execute moves
	end, err := executeMoves(g, moves)
	if err != nil {
		return err
	}
	if end {
		return nil
	}
	// if game is not over accept more moves
	for {
		var inputMoves string
		fmt.Scanln(&inputMoves)
		movesStr := strings.Split(inputMoves, " ")
		moves, err := parseMoves(movesStr)
		if err != nil {
			return err
		}
		end, err := executeMoves(g, moves)
		if err != nil {
			return err
		}
		if end {
			return nil
		}
	}
}

func parseMoves(movesStr []string) ([]int, error) {
	var moves []int
	for _, moveStr := range movesStr {
		move, err := strconv.Atoi(moveStr)
		if err != nil {
			return moves, err
		}
		moves = append(moves, move)
	}
	return moves, nil
}

func executeMoves(g *game.Game, moves []int) (bool, error) {
	for i, move := range moves {
		res, err := g.Move(game.Move{Player: game.Player(i % 2), Col: move})
		if err != nil {
			return false, err
		}
		if res.Status == game.StatusDraw {
			fmt.Println("DRAW")
			return true, nil
		}
		if res.Status == game.StatusWin {
			fmt.Printf("WINNER: Player %s\n", *res.Player)
			return true, nil
		}
	}
	return false, nil
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
