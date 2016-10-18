Greetings Internet buddies:

Welcome to lesson 7 of our golang for absolute beginners tutorial.

Today we are going to continue making the game Unique Candlestick. To do that I'm going to need to tell you a bit about how this game works.

Each player has cards with the numbers 1 to 13 on (A to K). Three of these cards in hand, and the rest face down in their deck. They then take it in turns to place a card from their hand into the middle of the table face up.

Once every player has won played a card, the player who placed the highest unique number wins a point. Each player draws a new card, from their deck, and play continues, starting with the player who won that trick, until all the cards have run out.

When all the cards have run out, the player with the highest unique number of tricks is declared the winner.

Ace is normally low(1), but if there is exactly One king, played, ace is high (14).

8 and 7 both reverse the winner (Assuming they are unique). So if exactly 1 seven is played, the lowest card wins.  if a seven and an eight are played, then highest win. So if exactly 1 seven is played, the lowest card wins.  if a seven and an eight are played, then highest wins.

Today the plan is to make a function that takes a list of numbers, and responds with the winner according to these rules. or -1 in the case of no win.

First off, how did you do with last weeks challenge?  Last week you had to create a function to make a new deck and then put the result into a slice of slices. I left you to work out how to make a slice of slices. Let's give it a go shall we?
