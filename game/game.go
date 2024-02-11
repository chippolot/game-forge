package game

type IGame interface {
	IMetadata

	GetState() IGameState

	Start()
	ExecuteAction(action IAction) (GameResult, error)
	Print()
	Restart()
}

type Game struct {
	Metadata
	logic    ILogic
	state    IGameState
	renderer IGameRenderer
}

func NewGame(logic ILogic, state IGameState, renderer IGameRenderer, metadata *Metadata, parser *ActionParser) *Game {
	game := &Game{
		Metadata: *metadata,
		logic:    logic,
		state:    state,
		renderer: renderer,
	}
	game.logic.RegisterActions(parser)
	return game
}

func (g *Game) Start() {
	g.state.Reset()
}

func (g *Game) GetState() IGameState {
	return g.state
}

func (g *Game) ExecuteAction(action IAction) (GameResult, error) {
	return g.logic.ExecuteAction(action, g.state)
}

func (g *Game) Print() {
	g.renderer.Print(g)
}

func (g *Game) Restart() {
	g.state.GetBoard().Clear()
	g.Start()
}

type GameResult struct {
	State         GameResultState
	WinningPlayer Player
}

type GameResultState int

const (
	NotGameOver GameResultState = iota
	GameWon
	GameTie
)
