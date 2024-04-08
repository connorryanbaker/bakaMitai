# Baka Mitai

Named after the soundtrack to all brilliant chess moves, Baka Mitai is a low
strength, WIP chess engine.

Currently, Baka Mitai runs a Negamax search with Alpha/Beta pruning and
a couple of simple evaluation functions on leaf nodes. Legal moves are
generated with bitboards, though there is loads to be cleaned up all over the
shop :D

Upcoming features to be implemented are listed in todo.txt.

![yakuza](https://i.redd.it/yaa8z9jfwu451.png)



## Sample Game
Below is a sample game between Baka Mitai as White and Baka Mitai as Black,
looking at depth 4

![baka VS baka](https://lichess1.org/game/export/gif/white/GfrUKuy1.gif)


```
1. Nc3 { A00 Van Geet Opening } Nc6 2. d4 d5 3. Nf3 Nf6 4. e3 Bf5 5. Bb5 a6 6. Bd3 Ne4 7. Bxe4 Bxe4 8. Nxe4 dxe4 9. Nd2 f5 10. O-O e5 11. Nb3 Bd6 12. Qh5+ g6 13. Qd1 O-O 14. Nc5 Qc8 15. Qe2 exd4 16. exd4 Rf7 17. Qc4 Bf8 18. Bf4 Bd6 19. Bg5 Ra7 20. f3 b5 21. Qc3 exf3 22. Rxf3 b4 23. Qc4 Ra8 24. Ne4 Na5 25. Nxd6 Nxc4 26. Nxc8 Rxc8 27. Kf2 Rd7 28. Rb3 Rxd4 29. Rxb4 Kf7 30. Re1 Kf8 31. b3 Rd2+ 32. Bxd2 Nxd2 33. Ra4 Ne4+ 34. Ke3 Nc5 35. Ra3 Re8+ 36. Kd2 Rxe1 37. Kxe1 Ke7 38. Ke2 Kd6 39. b4 Ne4 40. Rxa6+ Kd5 41. Ra7 Nc3+ 42. Kd3 Ne4 43. Rxc7 h6 44. Rg7 g5 45. Rd7+ Ke5 46. Re7+ Kd5 47. c4+ Kd6 48. Rh7 Nf2+ 49. Kd4 Ng4 50. h3 Nf6 51. Rxh6 Ke6 52. g4 fxg4 53. hxg4 Kf7 54. Rh3 Ke6 55. Re3+ Kf7 56. c5 Nxg4 57. Rg3 Nh2 58. Rh3 Ng4 59. Rg3 Nh2 60. Rg2 Nf3+ 61. Ke4 Nh4 62. Rxg5 Kf6 63. Rd5 Ng6 64. b5 Ne7 65. Rd8 Ng6 66. b6 Ne7 67. b7 Nc6 68. Rd6+ Ke7 69. Rxc6 Kd7 70. Rc8 Ke6 71. b8=Q Kf6 72. Rc7 Kg6 73. Qb6+ Kg5 74. Rh7 Kg4 75. Qg6# { White wins by checkmate. } 1-0
```

Stockfish on LiChess reports 90% accuracy for White, with 1 mistake, 1 blunder and 8 inaccuracies.
Black scored 89% accuracy, with 2 mistakes, 3 blunders and 6 inaccuracies.

To me, this assessment sounds generous, but it is a reasonable starting point. Baka Mitai is very much a hobby project work in progress!


