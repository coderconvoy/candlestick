Greetings Web folk
==================

Today we going to continue making Unique Candlestick. And learn about Golang structs.

Structs are basically a way of organising different kinds of data in the same place, so that you can see each part separately when you need to.

In our case each player has a deck of thirteen cards, 3 of which are drawn into their hands. So for each player we need to store their 3 cards.

We are also going to introduce a new way of looping, using the range keyword to grab all of the elements of an array in turn.

Before we get there though, we have some minor changes to make to the code. Currently our decide trick method actually returns the winning card, and that needs to be fixed to also return the winning player. That should be pretty simple.

Apart from that, how did you do with last weeks challenge? To write test methods for the 'beats' and 'contains' functions.
