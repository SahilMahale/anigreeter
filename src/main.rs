use clap::Parser;
use rand::seq::IteratorRandom;

#[derive(Parser, Debug)]
#[command(name = "anigreeter")]
#[command(
    version,
    about = "A CLI tool for generating anime quotes",
    long_about = Some("anigreeter is a command line tool that generates random anime quotes.
You can filter quotes by anime or character name.")
)]
#[group(multiple = true)]
struct Cli {
    #[arg(
        short,
        long,
        help = "Filters quotes based on the character that said it"
    )]
    character: Option<String>,
    #[arg(short, long, help = "Filters quotes based on the anime they are from")]
    anime: Option<String>,
}

#[derive(Debug, Eq)]
struct Name {
    first: String,
    middle: Option<String>,
    last: Option<String>,
}

impl PartialEq for Name {
    fn eq(&self, other: &Self) -> bool {
        // XXX
        self.first.to_lowercase() == other.first.to_lowercase()
    }
}

impl std::fmt::Display for Name {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match (&self.first, &self.middle, &self.last) {
            (first, Some(middle), Some(last)) => write!(f, "{} {} {}", first, middle, last),
            (first, Some(second), None) | (first, None, Some(second)) => {
                write!(f, "{} {}", first, second)
            }
            (first, None, None) => write!(f, "{}", first),
        }
    }
}

impl std::str::FromStr for Name {
    type Err = std::convert::Infallible;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.split_whitespace();

        // SAFETY: `s` should not be empty. This is validated during cli parsing.
        let first = parts.next().unwrap().to_owned();

        let name = match (parts.next(), parts.next()) {
            (Some(last), None) => Self {
                first,
                middle: None,
                last: Some(last.to_owned()),
            },
            (Some(middle), Some(last)) => Self {
                first,
                middle: Some(middle.to_owned()),
                last: Some(last.to_owned()),
            },
            _ => Self {
                first,
                middle: None,
                last: None,
            },
        };

        Ok(name)
    }
}

#[derive(Debug)]
struct Quote<'a> {
    text: String,
    character: &'a Name,
    anime: &'a str,
}

impl std::fmt::Display for Quote<'_> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            r#""{text}"
  - {character} ({anime})"#,
            text = self.text,
            character = self.character,
            anime = self.anime,
        )
    }
}

#[derive(Debug)]
struct QuoteStore<'a> {
    quotes: Vec<Quote<'a>>,
    rng: rand::rngs::ThreadRng,
}

impl<'a> QuoteStore<'a> {
    pub fn new() -> Self {
        Self {
            quotes: Vec::new(),
            rng: rand::thread_rng(),
        }
    }

    pub fn add_quotes_from_file<P: AsRef<std::path::Path>>(
        &mut self,
        p: P,
        character: &'a Name,
        anime: &'a str,
    ) -> Result<(), std::io::Error> {
        // XXX: This is loading the whole file in memory.
        // We could randomly seek to a point in the file, and read the next line.

        self.quotes
            .extend(std::fs::read_to_string(p)?.lines().map(|text| Quote {
                text: text.to_owned(),
                character,
                anime,
            }));

        Ok(())
    }

    pub fn get_quote_by_anime(&mut self, name: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| quote.anime.to_lowercase().eq(&name.to_lowercase()))
    }

    pub fn get_quote_by_character(&mut self, name: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| name.parse().as_ref() == Ok(quote.character))
    }

    pub fn get_quote_from(&mut self, anime: &str, character: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| {
            quote.anime.to_lowercase().contains(&anime.to_lowercase())
                && character.parse().as_ref() == Ok(quote.character)
        })
    }

    fn get_quote(&mut self) -> Option<&Quote> {
        self.quotes.iter().choose(&mut self.rng)
    }

    fn get_quote_by<P: Fn(&Quote) -> bool>(&mut self, p: P) -> Option<&Quote> {
        self.quotes
            .iter()
            .filter(|&quote| p(quote))
            .choose(&mut self.rng)
    }
}

fn main() -> Result<(), std::io::Error> {
    let cli = Cli::parse();

    let naruto = Name {
        first: "Naruto".to_owned(),
        last: Some("Uzumaki".to_owned()),
        middle: None,
    };

    let light = Name {
        first: "Light".to_owned(),
        last: Some("Yagami".to_owned()),
        middle: None,
    };

    let l = Name {
        first: "L".to_owned(),
        last: None,
        middle: None,
    };

    let mut quotes = QuoteStore::new();

    quotes.add_quotes_from_file("quotes/naruto.txt", &naruto, "Naruto")?;
    quotes.add_quotes_from_file("quotes/light.txt", &light, "Death Note")?;
    quotes.add_quotes_from_file("quotes/l.txt", &l, "Death Note")?;

    let quote = match (cli.anime, cli.character) {
        (Some(name), None) => quotes.get_quote_by_anime(&name),
        (None, Some(name)) => quotes.get_quote_by_character(&name),
        (Some(anime), Some(character)) => quotes.get_quote_from(&anime, &character),
        (None, None) => quotes.get_quote(),
    };

    match quote {
        Some(quote) => println!("{}", quote),
        None => eprintln!("No quote found :<"),
    }

    Ok(())
}
