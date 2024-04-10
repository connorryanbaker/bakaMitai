# Baka Mitai

Named after the soundtrack to all brilliant chess moves, Baka Mitai is a low
strength, WIP chess engine.

Currently, Baka Mitai runs a Negamax search with Alpha/Beta pruning and
a couple of simple evaluation functions on leaf nodes. Legal moves are
generated with bitboards, though there is loads to be cleaned up all over the
shop :D

Upcoming features to be implemented are listed in todo.txt.

![yakuza](https://i.redd.it/yaa8z9jfwu451.png)



## Sample Games
Below is a sample game between Baka Mitai as White and Baka Mitai as Black,
looking at depth 4

![baka VS baka](https://lichess1.org/game/export/gif/white/GfrUKuy1.gif)


```
1. Nc3 { A00 Van Geet Opening } Nc6 2. d4 d5 3. Nf3 Nf6 4. e3 Bf5 5. Bb5 a6 6. Bd3 Ne4 7. Bxe4 Bxe4 8. Nxe4 dxe4 9. Nd2 f5 10. O-O e5 11. Nb3 Bd6 12. Qh5+ g6 13. Qd1 O-O 14. Nc5 Qc8 15. Qe2 exd4 16. exd4 Rf7 17. Qc4 Bf8 18. Bf4 Bd6 19. Bg5 Ra7 20. f3 b5 21. Qc3 exf3 22. Rxf3 b4 23. Qc4 Ra8 24. Ne4 Na5 25. Nxd6 Nxc4 26. Nxc8 Rxc8 27. Kf2 Rd7 28. Rb3 Rxd4 29. Rxb4 Kf7 30. Re1 Kf8 31. b3 Rd2+ 32. Bxd2 Nxd2 33. Ra4 Ne4+ 34. Ke3 Nc5 35. Ra3 Re8+ 36. Kd2 Rxe1 37. Kxe1 Ke7 38. Ke2 Kd6 39. b4 Ne4 40. Rxa6+ Kd5 41. Ra7 Nc3+ 42. Kd3 Ne4 43. Rxc7 h6 44. Rg7 g5 45. Rd7+ Ke5 46. Re7+ Kd5 47. c4+ Kd6 48. Rh7 Nf2+ 49. Kd4 Ng4 50. h3 Nf6 51. Rxh6 Ke6 52. g4 fxg4 53. hxg4 Kf7 54. Rh3 Ke6 55. Re3+ Kf7 56. c5 Nxg4 57. Rg3 Nh2 58. Rh3 Ng4 59. Rg3 Nh2 60. Rg2 Nf3+ 61. Ke4 Nh4 62. Rxg5 Kf6 63. Rd5 Ng6 64. b5 Ne7 65. Rd8 Ng6 66. b6 Ne7 67. b7 Nc6 68. Rd6+ Ke7 69. Rxc6 Kd7 70. Rc8 Ke6 71. b8=Q Kf6 72. Rc7 Kg6 73. Qb6+ Kg5 74. Rh7 Kg4 75. Qg6# { White wins by checkmate. } 1-0
```

Stockfish on LiChess reports 90% accuracy for White, with 1 mistake, 1 blunder and 8 inaccuracies.
Black scored 89% accuracy, with 2 mistakes, 3 blunders and 6 inaccuracies.

At depth 7, Baka starts to play some proper, shuffly computer chess. White's accuracy jumps to 95% and Black's to 92% as per Stockfish.

![baka VS baka depth 7](https://lichess1.org/game/export/gif/white/94nWhQdZ.gif?theme=brown&piece=cburnett)


```
1. Nf3 Nf6 2. e3 { A05 Zukertort Opening: Quiet System } e6 3. Nc3 Nc6 4. d4 Bb4 5. Qd3 Nd5 6. a3 Nxc3 7. axb4 Na2 8. Rxa2 Nxb4 9. Qb3 Nxa2 10. Qxa2 O-O 11. Bd3 d5 12. O-O f5 13. b3 b6 14. Ba3 Rf6 15. Bb2 Bb7 16. Ba6 Bxa6 17. Qxa6 Qc8 18. Qxc8+ Rxc8 19. Ra1 a5 20. Kf1 Kf7 21. Ng5+ Kg6 22. f4 h6 23. Nf3 Kf7 24. Ba3 Ke8 25. Ke2 Rg6 26. g3 h5 27. Kd3 Rh6 28. c4 Rd8 29. Ne5 Ra8 30. h4 Rf6 31. Bb2 Ke7 32. Ba3+ Ke8 33. Ra2 Rh6 34. Rc2 a4 35. b4 Ra7 36. b5 Ra5 37. Rc1 dxc4+ 38. Kxc4 Ra7 39. Kb4 Kf8 40. Bb2 Ke8 41. Ra1 Rf6 42. Rxa4 Rxa4+ 43. Kxa4 Rh6 44. Ba3 Rf6 45. Kb3 Rh6 46. Kc2 Rf6 47. Kd3 Rh6 48. e4 Rf6 49. Bb4 Rh6 50. Ba3 Rf6 51. Kd2 Rh6 52. Ke2 Rf6 53. Kd2 Rh6 54. Ke2 Rf6 55. Kf2 Rh6 56. Kf3 Rf6 57. Kf2 Rh6 58. Kf3 Rf6 59. Ke2 Rh6 60. Kd2 Rf6 61. Ke2 Rh6 62. Kd2 Rf6 63. Kd3 Rh6 64. Bb4 Rf6 65. Ba3 Rh6 66. Bb4 Rf6 67. Ke3 Rh6 68. Kf2 Rf6 69. Ke3 Rh6 70. Kd2 Rf6 71. Ke2 Rh6 72. Ke3 Rf6 73. Bd2 Rh6 74. Bb4 Rf6 75. Kf2 Rh6 76. Ba3 Rf6 77. Kf3 Rh6 78. Kf2 Rf6 79. Kf3 Rh6 80. Bb4 Rf6 81. Ke2 Rh6 82. Kd2 Rh8 83. Be7 Kxe7 84. Ng6+ Kf6 85. Nxh8 g6 86. e5+ Kg7 87. d5 Kxh8 88. d6 cxd6 89. exd6 g5 90. hxg5 e5 91. fxe5 Kg7 92. d7 Kf7 93. d8=Q h4 94. Qf6+ Ke8 95. e6 hxg3 96. Qg7 g2 97. g6 g1=N 98. Qd7+ Kf8 99. Qf7# { White wins by checkmate. } 1-0
```

To me, this assessment sounds generous, but it is a reasonable starting point. Baka Mitai is very much a hobby project work in progress!


