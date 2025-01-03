CREATE TABlE quotes (
    id integer primary key autoincrement,
    character text,
    anime text,
    quotetext text
)
-- TODO: 3NF
-- Animes should be a different table, alongside characters.
-- Characters will have a(n) FK to the Anime table.
