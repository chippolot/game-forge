# TODO
- [ ] Lua
    - [ ] Have game framework call into lua
- [ ] AI
    - [ ] Implement simple LLM to generate lua code
    - [ ] Implement game registry to save fun games
    - [ ] Validation of working game logic (???)
- [ ] Chess
    - [ ] Pieces
        - [ ] Don't directly associate players and pieces (think chess)
- [ ] Board
    - [ ] Flip row notation

# Bugs
- [ ] Invalid action parse leads to broken subsequent parsing

# Done
- [x] Pieces
    - [x] Create interface
    - [x] Handle display
    - [x] Actions can result in piece removal
    - [x] Add colored pieces
- [x] Interface
    - [x] Game selection
- [x] Game
    - [x] Update game framework to check for winner and display message
- [x] Score
    - [x] Add score display
- [x] Tic-Tac-Toe
    - [x] Implement rules
- [x] Othello
    - [x] Implement rules
- [x] Checkers
    - [x] Implement rules
- [x] Project
    - [x] Create github
- [x] Actions
    - [x] Add action parser
    - [x] Better way to inject implicit data into actions (ex: place piece)
    - [x] Share more action parsing code
    - [x] Add better player prompt which displays possible action types
    - [x] Add ability to place or move piece
- [x] Board
    - [x] Add chess style board coords
    - [x] Add board coord display