create table if not exists TranscationtHistory(
    Name text,
    OpenPrice float,
    LowPrice float,
    HighPrice float,
    ClosePrice float,
    Volume float,
    DatePrice date
  )

CREATE TABLE IF NOT EXISTS symbols (
    Name TEXT PRIMARY KEY,
    SharesOutstanding float
)
