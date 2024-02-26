# Getting started

Create a .env file with your OpenAI API key
```
OPENAI_API_KEY=your_api_key
```

You can use `air` for hot reloading.
https://github.com/cosmtrek/air

# An AI programming story

Last week I was asked why we were not using OpenAI for translations.
I wondered what that would look like. So this weekend I played around with it.

## A great start

I asked ChatGPT to generate a service in go that would take a text and a target language, use the OpenAI API to translate the text and return the translated text. Why Go? Why not, it's Saturday.

It gave me a very good looking piece of code right out of the box. It used the right library, the right function and even the right parameters. It was a very good start.

I had not set up go on my Mac. So I had to do that first. I had not used go for a while. I needed to refresh my memory. I asked ChatGPT to tell me how to set up go and how to create a project. It gave me a very good answer. I followed the instructions and was ready to go.

I was very impressed. I thought we were out of work. We can just ask ChatGPT to do everything for us.

## The problem

I tried to run the code. It did not work. I realised that the library it used for OpenAI did not exist. It was not just the wrong call. The library was not even there. I asked ChatGPT to point me to the documentation for the library. It basically told me that it was wrong and that there was no official library.

So I did some quick research and used a 3rd party library. Writing my own code.

So basically it generated code that would have worked if there was an official library. But there was not. The code was completely fictitious, so it didn't work. I think it copied the structure from the javascript library and made it look like go, pretending there was an actual library.

## Improve

Now it worked. To make it better, I refactored the code. ChatGPT probably could have helped as well. But I like having some control.

I got tired of restarting the server by hand. So I asked ChatGPT how to do hot reloading with Gin in Go. It gave me a very good answer. I followed the instructions and it worked.

I wanted to add a better logger. I could not remember the log libraries and of course they might have changed recently. So I asked ChatGPT to give me a good logger for go. I then asked how a middleware for gin logging requests would look like with this logger. It all worked fine.

I also wanted to move the OpenAI API key to an environment variable and an .env file. I asked ChatGPT how to do this and it gave me a very good answer. It even included the new logger I had chosen in the code.

I asked for the creation of a .gitignore file for me. Which it did perfectly, given the context.

## Conclusion

I was very impressed with ChatGPT. It was able to get me off to a very good start. It was able to help me with the setup and the libraries. It was able to help me with more detailed questions I had. It was a very good experience. It cannot yet generate complete solutions that are reliable. But if you know the steps but are fuzzy on the details, it will really speed things up.

ChatGPT is great for generating starting points.
ChatGPT is also like some dads, it will tell you everything it knows, but if it does not know something, it will probably make it up and make it sound good and indistinguishable from real answers.

When it comes to translation with OpenAI, I do not think it is a viable out-of-the-box solution. The answers are not predictable, and translating the same text will produce different results every time. It cannot reliably identify the source language and the response format will change if it does not know the source language. Sometimes the response contains meta information as well as the translated text.
We would need to do some prompt engineering to evaluate this approach further.

Nevertheless, a fun experiment if you have 2 hours to spare.
