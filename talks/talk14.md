Greetings Webbers.

Today is all about cookies, managing cookies. How to set up a basic cookie and use it to keep track of user data, and keep each user separate.

To do this I'm going to continue with the Unique Candlestick game I have been working on.

Currently it only takes one user at a time, which is pretty bad for a web facing program. But today, ALL THAT CHANGES.

Cookies are something the web has been used to for some time basically they are a small piece of information stored on the users computer, traditionally in a text file.

The easiest way to use them is as a unique Identifier. When the user first arrives at our site. We make up a cookie, with a unique string in it. Send it to them in the header of our response.

After that, when they come back to our site, (until that cookie expires) They send that cookie back to tell us who they are. The browser normally does this without telling them.  We can then find all of the information we have stored next about them and use it to give better responses to them.

There are 3 important functions we need today.

1. A cookie constructor, Actually we will just make it as a struct.

2. http.SetCookie(w ResponseWriter, c *Cookie)  which takes a cookie 

3. http.Request.Cookie(name String)(*Cookie, err) Which is how we get it back from the request


