package game

type IGameState interface {
	ICommonGameState
}

type ICommonGameState interface {
	GetCurrentPlayer() Player
	SetCurrentPlayer(player Player)

	GetBoard() IBoard

	GetPlayerScore(player Player) int
	AddPlayerScore(player Player, delta int)
	SetPlayerScore(player Player, value int)
}

type CommonGameState struct {
	playerScores  map[Player]int
	currentPlayer Player
	board         IBoard
}

func NewCommonGameState(board IBoard) *CommonGameState {
	return &CommonGameState{
		currentPlayer: 0,
		board:         board,
		playerScores:  make(map[Player]int),
	}
}

func (s *CommonGameState) GetCurrentPlayer() Player {
	return s.currentPlayer
}

func (s *CommonGameState) SetCurrentPlayer(player Player) {
	s.currentPlayer = player
}

func (s *CommonGameState) GetBoard() IBoard {
	return s.board
}

func (s *CommonGameState) GetPlayerScore(player Player) int {
	return s.playerScores[player]
}

func (s *CommonGameState) AddPlayerScore(player Player, delta int) {
	s.playerScores[player] += delta
}

func (s *CommonGameState) SetPlayerScore(player Player, value int) {
	s.playerScores[player] = value
}
