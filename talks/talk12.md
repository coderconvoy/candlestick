Greetings Webby folk.

Today I hope to get unique candlestick actually playable online.

Last week I set you the challenge of getting responses onto the game, which I imagine would have been quite difficult.

I actually expect it to take me a fair bit of time get working now.

I think the system we need a system, that keeps track of every step, or turn, and if the turn is human, then it stops and waits for that human.

To do this we are going to need another member of the game struct for played cards.

Actually to make sure we can get to everything, I would like to actually keep track of everything that happened each round until the game is over. So we will make a new type called GameRound, and store a slice of them in the game struct. 

Lets get on with it shall we.
