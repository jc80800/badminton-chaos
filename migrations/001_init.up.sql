CREATE TABLE excuses (
  id           TEXT PRIMARY KEY,          -- uuid string
  text         TEXT NOT NULL,
  category     TEXT NOT NULL,             -- e.g., 'singles','doubles','general'
  locale       TEXT NOT NULL DEFAULT 'en-US',
  source       TEXT,                      -- 'seed','user','admin'
  quality_score INTEGER NOT NULL DEFAULT 0,
  active       BOOLEAN NOT NULL DEFAULT true,
  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_excuses_active_cat_locale ON excuses(active, category, locale);

-- Insert some seed data
INSERT INTO excuses (id, text, category, locale, source, quality_score) VALUES
('excuse_1', 'The fluorescent lights blinded my smashes.', 'singles', 'en-US', 'seed', 85),
('excuse_2', 'My racket strings were too loose.', 'general', 'en-US', 'seed', 70),
('excuse_3', 'The court was too slippery.', 'doubles', 'en-US', 'seed', 75),
('excuse_4', 'I had a bad night sleep.', 'general', 'en-US', 'seed', 60),
('excuse_5', 'The shuttlecock was too fast.', 'singles', 'en-US', 'seed', 80);
