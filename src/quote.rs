use crate::Name;

#[derive(Debug)]
pub struct Quote<'a> {
    pub text: String,
    pub character: &'a Name,
    pub anime: &'a str,
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
