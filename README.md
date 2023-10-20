# RTP calculator
basic RTP calculator for Slot Games with Free Games and multipliers. Monte Carlo simulation was used.
## How it works?
It takes initial parameters from **example.json**:
- number of iterations
- multiply symbol number
- multiply value
- multiply column number
- wild symbol number
- scatter symbol number
- paytable
- free games number paytable
- win lines
- reels
- reels for free games
- scatter type

Then, it simulates gameplay, but much faster than usual player (counts win per spin, number of free games and so on). Final result looks like this:
```
Number of iterations: 10 mln.
RTP: 0.9557552
AVG mult value per free games: 2.9132051
AVG number of free games: 0.1465395
Time passed: 10.527041513s
```
The more iterations, the more precise RTP.

## Why is it needed?
After getting or calculating reels for the game, one should check its validity. It can be done mathematically for simple slot games, but not that easy for complicated ones. That is why it is easier to simulate player activity and calculate average RTP and check if it is what we expected.

## Vocabulary
**RTP** - *return to player*. Basically, $$\frac{\text{total win}}{\text{total bet}} \cdot 100$$

*Example:*
Player has made **n** number of spins, where **n** tends to infinity. His total win is 95$ and his total bet is 100$. Then the RTP is 95%.


**reel** - is a an array of columns with specific order of symbols, that simulates physical reel in real-life slot machine. The code randomly picks window of symbols with a specified size.


**scatter** - special symbol. When player gets a combination(depends on game logic) of this symbol, slot game gives him a specified number of free games.


**wild symbol** - symbol that can play as any symbol in specific combinations (except scatter and multiplier symbol).

*Example:*
Player gets such line: *[lemon wild lemon lemon lemon]*. Instead of getting nothing, player gets payoff as it would be *[lemon lemon lemon lemon lemon]*
