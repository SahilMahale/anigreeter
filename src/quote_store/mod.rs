mod vec;

use crate::quote::Quote;
pub use vec::QuoteStoreVec;

pub trait QuoteStore {
    fn get_quote_by_anime(&mut self, name: &str) -> Option<&Quote>;
    fn get_quote_by_character(&mut self, name: &str) -> Option<&Quote>;
    fn get_quote_from(&mut self, anime: &str, character: &str) -> Option<&Quote>;
    fn get_quote(&mut self) -> Option<&Quote>;
}
