# TODO
- [ ] Othello
    - [ ] Pieces
        - [ ] Add colored pieces
- [ ] Board
    - [ ] Flip row notation
- [ ] Checkers
    - [ ] Actions
        - [ ] Add ability to place or move piece
    - [ ] Pieces
        - [ ] Actions can result in piece removal
- [ ] Lua
    - [ ] Have game framework call into lua
- [ ] AI
    - [ ] Implement simple LLM to generate lua code
    - [ ] Implement game registry to save fun games
- [ ] Chess
    - [ ] Pieces
        - [ ] Don't directly associate players and pieces (think chess)

# Bugs
- [ ] Invalid action parse leads to broken subsequent parsing

# Done
- [x] Pieces
    - [x] Create interface
    - [x] Handle display
- [x] Interface
    - [x] Game selection
- [x] Game
    - [x] Update game framework to check for winner and display message
- [x] Score
    - [x] Add score display
- [x] Tic-Tac-Toe
    - [x] Finish implementing IsGameOver for t-t-t
    - [x] Finish implementing GetWinner for t-t-t
- [x] Othello
    - [x] Implement rules
- [x] Project
    - [x] Create github
- [x] Actions
    - [x] Add action parser
    - [x] Better way to inject implicit data into actions (ex: place piece)
    - [x] Share more action parsing code
    - [x] Add better player prompt which displays possible action types
- [x] Board
    - [x] Add chess style board coords
    - [x] Add board coord display