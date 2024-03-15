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
looking at depth 5 (quiescence yet to be implemented - the horizon effect is
real).

![baka VS baka](https://lichess1.org/game/export/gif/white/WkppnBrB.gif?theme=brown&piece=cburnett)

1. e4 Nc6 2. Nc3 e5 3. Nf3 a6 4. Bd3 d6 5. h3 Nb4 6. Be2 Qf6 7. d3 Ne7 8. a3 Nbc6 9. Bg5 Qg6 10. h4 Bg4 11. d4 exd4 12. Nxd4 Nxd4 13. Bxg4 Nec6 14. O-O Nb5 15. Nxb5 axb5 16. Bf5 Qxf5 17. exf5 h6 18. Be3 Be7 19. Qd3 Ra5 20. Bd2 b4 21. Qf3 Kf8 22. Rad1 Rc5 23. f6 Bxf6 24. c3 bxc3 25. Be3 Re5 26. bxc3 Kg8 27. Bd4 Nxd4 28. cxd4 Rb5 29. Qe2 Rb6 30. Qg4 d5 31. h5 Re6 32. Rb1 b6 33. Rb5 Rh7 34. Rb3 Bd8 35. Rd1 c6 36. Rh3 b5 37. Kf1 Bb6 38. Rc1 Rh8 39. Rg3 Rh7 40. Qf5 Bxd4 41. Qxd5 Bb6 42. Rxc6 g5 43. hxg6 Bxf2 44. gxf7+ Kxf7 45. Rxe6 Bxg3 46. Qf5+ Kg7 47. Rf6 b4 48. axb4 h5 49. Rg6+ Kh8 50. Rxg3 h4 51. Qc8# 1-0

Stockfish on LiChess reports 94% accuracy for White, with 3 inaccuracies,
and 1 blunder. It is a bit more critical of Black's performance (it did trap
its own queen) reporting 87% accuracy with 6 inaccuracies, 2 mistakes and 1 blunders. To me,
this assessment sounds generous, but it is a reasonable starting point. Baka
Mitai is very much a hobby project work in progress!


