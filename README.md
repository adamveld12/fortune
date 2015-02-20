#Fortune

An ode to the [Fortune]() terminal app, in Go for learning purposes.


Like Arnold’s fortune program, this version uses a database of fortunes (the fortune cookie database) generated from a text file. 
The text file consists of possible multi-line quotes, separated by lines consisting of a single ”%” character. For example:

```
Don't go around saying the world owes you a living.  The world owes you
nothing.  It was here first.
        -- Mark Twain
%
Every normal man must be tempted at times to spit on his hands, hoist the
black flag, and begin slitting throats.
        -- H.L. Mencken
%
Behind every argument is someone's ignorance.
        -- Louis Brandeis
```


## How to use

If you want a random quote:
```sh
  fortune
```

If you want to use a specific file:
```sh
  fortune -file="quotes.txt"
```

If you want to delay process termination:
```sh 
  fortune -wait=1.2s #( use ms for miliseconds, s for seconds, m for minutes etc)
```

and if you want to see help text:
```sh
  fortune -h
```

