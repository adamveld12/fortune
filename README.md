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
  fortune  # defaults to fortunes.txt
```

If you want to use a specific file:
```sh
  fortune quotes.txt
```

If you want to delay termination:
```sh 
  fortune -wait=1.2s #( use ms for miliseconds, s for seconds, m for minutes etc)
```

If you want to generate an index file:
```sh
  fortune -generateIndex fortunes.txt  # Generates a fortunes.txt.index
```

*Note* Fortune will automatically look for a .index file in the same directory as the specified fortune and use it if it exists.


and if you want to see help text:
```sh
  fortune -h
```

## License

[WTFPL](http://www.wtfpl.net/)

