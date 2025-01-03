use crate::Name;

#[derive(Debug, Clone)]
pub struct Quote {
    pub text: String,
    pub character: Name,
    pub anime: String,
}

// TODO: Types: `Character` `Anime`
// Allows capturing the relation between an `Anime` and its `Character`s.

impl std::fmt::Display for Quote {
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
