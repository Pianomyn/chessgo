- Found chess programming wiki (catered to more experienced audience, assumes some knowledge)
- Write code myself as much as possible
    - Treat LLM as senior engineer (familiar with domain), treat myself as junior engineer
    - Discuss approaches/tradeoffs before starting
    - Ask if on write path if implementation starts getting messy
- High level understanding of chess engine structure/components with LLM (move gen, search, eval)
- Chat with LLM about main bottlenecks that tend to pop up in chess engines (C++ and Go)
    - Search space is massive, move ordering (sequence of searches) matters alot
    - Cache misses (keep data flat)
    - Go
        - Go Garbage collection causes stutters
- LLM suggested rook ray generations could wrap around like king/knight
- Ray move gen given rays impl (manually checking each square if blocker). Used LLM for suggestion
about TrailingZeroes/LeadingZeroes (LSB/MSB) trick to find first blocker
- For move gen, discussed tradeoffs of Make/Unmake vs copy ChessBoards. Went with copy Chessboards approach initially as more intuitive.
Plan to optimize later if notice a specific bottleneck.
