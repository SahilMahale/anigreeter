use std::io::{BufRead, BufReader};

use crate::name::Name;
use crate::quote::Quote;

use rand::seq::IteratorRandom;

#[derive(Debug)]
pub struct QuoteStoreVec<'a> {
    quotes: Vec<Quote<'a>>,
    rng: rand::rngs::ThreadRng,
}

impl<'a> QuoteStoreVec<'a> {
    pub fn new() -> Self {
        Self {
            quotes: Vec::new(),
            rng: rand::thread_rng(),
        }
    }

    pub fn add_quotes_from_file<P: AsRef<std::path::Path>>(
        &mut self,
        path: P,
        character: &'a Name,
        anime: &'a str,
    ) -> Result<(), std::io::Error> {
        let f = std::fs::File::open(path)?;
        let lines = BufReader::new(f).lines();

        self.quotes.extend(lines.filter_map(|line| {
            line.ok().map(|text| Quote {
                text,
                character,
                anime,
            })
        }));

        Ok(())
    }
    fn get_quote_by<P: Fn(&Quote) -> bool>(&mut self, p: P) -> Option<&Quote> {
        self.quotes
            .iter()
            .filter(|&quote| p(quote))
            .choose(&mut self.rng)
    }
}

impl super::QuoteStore for QuoteStoreVec<'_> {
    fn get_quote_by_anime(&mut self, name: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| quote.anime.to_lowercase().eq(&name.to_lowercase()))
    }

    fn get_quote_by_character(&mut self, name: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| name.parse().as_ref() == Ok(quote.character))
    }

    fn get_quote_from(&mut self, anime: &str, character: &str) -> Option<&Quote> {
        self.get_quote_by(|quote| {
            quote.anime.to_lowercase().contains(&anime.to_lowercase())
                && character.parse().as_ref() == Ok(quote.character)
        })
    }

    fn get_quote(&mut self) -> Option<&Quote> {
        self.quotes.iter().choose(&mut self.rng)
    }
}
