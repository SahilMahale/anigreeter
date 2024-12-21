# Anime Quotes
A `fortune` inspired CLI utility that generates anime quotes either:
  1. Randomly
  2. Based on an anime
  3. Based on a Character

It is very much in a its infancy, for now its a perl script.

## Installation
```bash
git clone <this repo>
cd animeQuotes
chmod +x worker.pl 
ln -s <abs-path>/worker.pl $HOME/.local/bin
```

## Usage
```bash
# Get a random quote
./animequote.pl

# Get a quote from Naruto
./animequote.pl --anime naruto

# Get a quote from Luffy
./animequote.pl --character luffy

# Short options work too
./animequote.pl -a one piece
./animequote.pl -c light
```
## Cowsay integration
1. Install `cowsay` 
2. After that you can copy over the `cowsay sprites` in sprites folder to `/usr/share/cowsay/site-cows/`
3. Usage:
  ```bash
  animequote -c naruto | cowsay -f naruto
```
> Supported characters
> 1. Naruto
