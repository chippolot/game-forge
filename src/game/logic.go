package game

type ILogic interface {
	RegisterActions(actionParser *ActionParser)
	ExecuteAction(action IAction, state GameState) (GameResult, error)
}
