package poker

import "io"

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}