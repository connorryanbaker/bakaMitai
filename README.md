# Baka Mitai

Named after the soundtrack to all brilliant chess moves, Baka Mitai is a low
strength, WIP chess engine.

Currently, Baka Mitai runs a Negamax search with Alpha/Beta pruning and
a couple of simple evaluation functions on leaf nodes. 

Upcoming features to be implemented are listed in todo.txt.

![yakuza](https://i.redd.it/yaa8z9jfwu451.png)

## Sample Game
Below is a sample game between Baka Mitai as White and Baka Mitai as Black,
looking at depth 4 (quiescence yet to be implemented - the horizon effect is
real).

1. Nc3 c5 2. Nd5 Nf6 3. Nxf6+ exf6 4. Nf3 d5 5. e3 Bf5 6. Bd3 Bxd3 7. cxd3 Nc6 8. Qb3 Qd7 9. O-O Ne5 10. Nxe5 fxe5 11. f4 exf4 12. exf4 O-O-O 13. Qa3 Kb8 14. Qc3 Re8 15. f5 Qa4 16. b3 Qd4+ 17. Qxd4 cxd4 18. Bb2 Bc5 19. Rac1 Bb4 20. Bxd4 Bxd2 21. Rc5 Be3+ 22. Bxe3 Rxe3 23. Rxd5 Kc7 24. Kf2 Ree8 25. Rc1+ Kb6 26. Rd6+ Kb5 27. Rd5+ Kb6 28. Rd6+ Kb5 29. Rc2 Re5 30. g4 Re7 31. Rd5+ Kb6 32. Rd6+ Kb5 33. Rd5+ Kb6 34. Rd6+ Kb5 1/2-1/2

Stockfish on LiChess reports 84% accuracy for White, with 2 inaccuracies,
3 mistakes and 3 blunders. It is a bit more critical of Black's performance,
reporting 80% accuracy with 6 inaccuracies, 2 mistakes and 2 blunders. To me,
this assessment sounds generous, but it is a reasonable starting point.

Baka Mitai confusingly favors a Threefold repetition in winning positions - something to look into :D

