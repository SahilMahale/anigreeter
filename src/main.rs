mod name;
mod quote;
mod quote_store;

use clap::Parser;

use name::Name;
use quote_store::{QuoteStore, QuoteStoreSqlite, QuoteStoreVec};

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

fn get_sqlite_store() -> Result<Box<QuoteStoreSqlite>, Box<dyn std::error::Error>> {
    let store = QuoteStoreSqlite::new("sqlite/quotes.db")?;
    Ok(Box::new(store))
}

fn get_vec_store() -> Result<Box<QuoteStoreVec>, std::io::Error> {
    let naruto = Name::full("Naruto", "Uzumaki");
    let light = Name::full("Light", "Yagami");
    let l = Name::first("L");

    let mut quotes = QuoteStoreVec::new();

    quotes.add_quotes_from_file("quotes/naruto.txt", &naruto, "Naruto")?;
    quotes.add_quotes_from_file("quotes/light.txt", &light, "Death Note")?;
    quotes.add_quotes_from_file("quotes/l.txt", &l, "Death Note")?;

    Ok(Box::new(quotes))
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let cli = Cli::parse();

    let mut quotes: Box<dyn QuoteStore> = if false {
        get_vec_store()?
    } else {
        get_sqlite_store()?
    };

    let quote = match (&cli.anime, &cli.character) {
        (Some(name), None) => quotes.get_quote_by_anime(&name),
        (None, Some(name)) => quotes.get_quote_by_character(&name),
        (Some(anime), Some(character)) => quotes.get_quote_from(&anime, &character),
        (None, None) => quotes.get_quote(),
    };

    match quote {
        Some(quote) => println!("{}", quote),
        None => eprintln!("No quote found :<"),
    };

    Ok(())
}
