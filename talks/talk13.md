Greeting webbed folk.

Today we are going to be looking at golang Templates.

Templates are golangs main way of making web pages, though they can be used to create all kinds of files.

1. First you have to parse them. Well actually the text/template or html/template modules can do that for you, but you have to tell them to. This can be done once at program start up if you don't plan on changing the templates.

2. Then when someone makes a request, you handle that request, and calculate everything you want to tell them. And put that in a single data object. This can be any golang type.

3. Call execute on the template, sending it a Writer ( anything that has a Write method) and the data. If this if on the web, that will be your http.ResponseWriter that the handle function recieves.

4. The text/template module will the run the template inserting all the various bits of data exactly where you tell it to.

Templates are actually incredibly powerful and you can use them to do all kinds of things, but perhaps the thing that makes them most powerful is that they can call any go function you give them inline.

Today, I'm going to use the template system, to write to the browser all the information about the game we have played so far, everything I need to tell the user is in the Game struct, which I have created. Lets have a look at that first.


