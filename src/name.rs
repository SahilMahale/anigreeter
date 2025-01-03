#[derive(Debug, Eq)]
pub struct Name {
    first: String,
    middle: Option<String>,
    last: Option<String>,
}

impl Name {
    pub fn full(first: &str, last: &str) -> Self {
        Self {
            first: first.to_owned(),
            last: Some(last.to_owned()),
            middle: None,
        }
    }

    pub fn first(first: &str) -> Self {
        Self {
            first: first.to_owned(),
            last: None,
            middle: None,
        }
    }
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
