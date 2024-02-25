# Getting started

Create a .env file with your OpenAI API key
```
OPENAI_API_KEY=your_api_key
```

You can use `air` for hot reloading.
https://github.com/cosmtrek/air

# An AI programming story

Last week I was asked, why we do not use OpenAI for translations.
I was wondering how that would look like so this weekend I played around.

## A great start

I asked ChatGPT to generate me a service in go that would take a text and a target language, use the OpenAI API to translate the text and return the translated text.

It created me a very good looking code right away. It used the right library, the right function and even the right parameters. It was a very good start.

I have not set up go  on my mac so I had to do it first. I have not used go for a while so I had to refresh my memory. I asked ChatGPT to tell me how to setup go and a project. It gave me a very good answer. I followed the instructions and I was ready to go.

I was very impressed. I thought all our jobs are now gone. We can just ask ChatGPT to do everything for us.

## The problem

I tried to run the code. It did not work. I realized, the library it used for OpenAI did not exist. It was not only the wrong call, the library was not even there. I asked ChatGPT to point me to the documentation of the library and it basically told me that it was wrong and there is no official library.

So I did a quick research and used a 3rd party library. Writing the code myself.

So basically it generated code that, if there was an official library, would have worked. But there was not. The code was completely fictitious, so it didn't work. I guess it copied the structure from the javascript library and made it look like go.

## Success

Now it worked. I refactored the code to be better. ChatGPT probably could also helped but I like to have some control.

I was annoyed to always restart the server by hand so I asked ChatGPT how to do hot reloading with Gin in Go. It gave me a very good answer. I followed the instructions and it worked.

## Improve

I wanted to add a better logger. I did not remember the log libraries and of course they might have changed lately. So i Asked ChatGPT to give me a good logger for go. I then proceeded to ask how a middleware for Gin logging requests with that logger would look like. It all worked fine.

I also wanted to move the OpenAI API key to an environment variable and a .env file. I asked ChatGPT how to do that and it gave me a very good answer. It even included the new logger I chose in the code.

## Conclusion

I was very impressed by ChatGPT. It was able to generate a very good start for me. It was able to help me with the setup and with the libraries. It was able to help me with more detail questions I had. It was a very good experience. It can not generate full solutions that are reliable yet. But if you know the steps but are blurry on the details, it will really speed things up.

When it comes to translating with OpenAI I do not think it is a feasable Out-Of-The-Box solution. The responses are not predictable and translating the same text will result in different results every time. It can not securely identify the origin language and the response format will change if it does not know the origin language.
We would have to do some prompt engineering to further evaluate this approach.

Nevertheless, a fun experiment if you have 2 hours to spare.