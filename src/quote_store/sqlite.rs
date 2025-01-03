use std::path::Path;

use rand::{seq::IteratorRandom, thread_rng};
use rusqlite::{params, Connection, Params, Result};

use crate::quote::Quote;

pub struct QuoteStoreSqlite {
    conn: Connection,
}

impl QuoteStoreSqlite {
    pub fn new<P: AsRef<Path>>(path: P) -> Result<Self> {
        let conn = Connection::open(path)?;

        Ok(Self { conn })
    }

    fn get_random_quote<P: Params>(&mut self, query: &str, params: P) -> Option<Quote> {
        let mut q = self.conn.prepare(query).ok()?;

        let quotes = q
            .query_map(params, |row| {
                let anime = row.get("anime")?;
                let character = row.get::<&str, String>("character")?.parse().unwrap();

                Ok(Quote {
                    anime,
                    character,
                    text: row.get(3)?,
                })
            })
            .ok()?;

        quotes.choose(&mut thread_rng()).map(Result::ok).flatten()
    }
}

impl super::QuoteStore for QuoteStoreSqlite {
    fn get_quote_by_anime(&mut self, name: &str) -> Option<Quote> {
        self.get_random_quote(
            "SELECT * FROM quotes WHERE UPPER(anime) = ?1",
            params![name.to_uppercase()],
        )
    }

    fn get_quote_by_character(&mut self, name: &str) -> Option<Quote> {
        self.get_random_quote(
            "SELECT * FROM quotes WHERE UPPER(character) LIKE ?1",
            params![format!("{name}%")],
        )

        // FIXME: Bad Matches
        // The usage of `LIKE` operator results in bad matches. For instance,
        // getting quotes by *L* (from Death Note) will also return quotes by
        // all characters whose name starts with `L`.
        //
        // To resolve this, the `=` operator should be used. But this will force
        // the user to enter the full name. Possibly, the best option is to
        // store characters into a table with first, middle and last name
        // columns and use `Name::contains` for checks.
        //
        // At that point, animes can be their own table referred with a(n) FK
        // by characters.
    }

    fn get_quote_from(&mut self, anime: &str, character: &str) -> Option<Quote> {
        self.get_random_quote(
            "SELECT * FROM quotes WHERE UPPER(anime) = ?1 AND UPPER(character) LIKE ?2",
            params![anime.to_uppercase(), format!("{character}%")],
        )
    }

    fn get_quote(&mut self) -> Option<Quote> {
        self.get_random_quote("SELECT * FROM quotes", [])
    }
}
