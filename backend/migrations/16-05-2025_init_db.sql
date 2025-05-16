-- Users
/* CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
); */

-- Documents
CREATE TABLE IF NOT EXISTS documents (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  -- owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
  title TEXT DEFAULT 'Untitled',
  content TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);

-- Document access
/* CREATE TABLE IF NOT EXISTS document_access (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  document_id UUID REFERENCES documents(id) ON DELETE CASCADE,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  role TEXT CHECK (role IN ('owner', 'editor', 'viewer')),
  UNIQUE(document_id, user_id)
); */