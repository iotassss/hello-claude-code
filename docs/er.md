# Entity Relationship Diagram

## Database Schema for TODO App

### Tables Overview

```
Users (Future Implementation)
├── Tasks (1:N)
└── Categories (1:N)
    └── Tasks (1:N)
```

## Entity Definitions

### 1. Tasks Table

```sql
CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    priority VARCHAR(20) NOT NULL DEFAULT 'medium',
    due_date DATETIME,
    category_id INTEGER,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

#### Attributes:
- **id**: Primary key, auto-increment
- **title**: Task title (required, max 100 chars)
- **description**: Task details (optional, text)
- **status**: Enum: 'pending', 'in_progress', 'completed'
- **priority**: Enum: 'low', 'medium', 'high', 'critical'
- **due_date**: Optional deadline
- **category_id**: Foreign key to categories table
- **created_at**: Auto-generated creation timestamp
- **updated_at**: Auto-updated modification timestamp

#### Indexes:
```sql
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_priority ON tasks(priority);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);
CREATE INDEX idx_tasks_category_id ON tasks(category_id);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);
```

### 2. Categories Table

```sql
CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    color VARCHAR(7) NOT NULL DEFAULT '#3498db',
    description TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

#### Attributes:
- **id**: Primary key, auto-increment
- **name**: Category name (required, unique, max 50 chars)
- **color**: Hex color code (default: #3498db)
- **description**: Optional category description
- **created_at**: Auto-generated creation timestamp
- **updated_at**: Auto-updated modification timestamp

#### Default Categories:
```sql
INSERT INTO categories (name, color, description) VALUES
('Work', '#e74c3c', 'Work-related tasks'),
('Personal', '#2ecc71', 'Personal tasks and goals'),
('Shopping', '#f39c12', 'Shopping lists and purchases'),
('Health', '#9b59b6', 'Health and fitness tasks'),
('Finance', '#1abc9c', 'Financial and budget tasks'),
('Education', '#34495e', 'Learning and educational tasks');
```

### 3. Users Table (Future Implementation)

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_premium BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

#### With User Integration:
```sql
-- Add user_id to tasks table
ALTER TABLE tasks ADD COLUMN user_id INTEGER;
ALTER TABLE tasks ADD FOREIGN KEY (user_id) REFERENCES users(id);

-- Add user_id to categories table
ALTER TABLE categories ADD COLUMN user_id INTEGER;
ALTER TABLE categories ADD FOREIGN KEY (user_id) REFERENCES users(id);
```

## Relationships

### Current Schema (Single User)
- **Categories → Tasks**: One-to-Many
  - One category can have multiple tasks
  - One task belongs to one category (optional)

### Future Schema (Multi User)
- **Users → Tasks**: One-to-Many
  - One user can have multiple tasks
  - One task belongs to one user
- **Users → Categories**: One-to-Many
  - One user can create multiple categories
  - One category belongs to one user
- **Categories → Tasks**: One-to-Many (within same user)
  - One category can have multiple tasks
  - One task belongs to one category (optional)

## ER Diagram (Current)

```
┌─────────────────┐       ┌─────────────────┐
│   Categories    │       │      Tasks      │
├─────────────────┤       ├─────────────────┤
│ id (PK)         │◄─────┤│ id (PK)         │
│ name            │   1:N ││ title           │
│ color           │       ││ description     │
│ description     │       ││ status          │
│ created_at      │       ││ priority        │
│ updated_at      │       ││ due_date        │
└─────────────────┘       ││ category_id(FK) │
                          ││ created_at      │
                          ││ updated_at      │
                          └─────────────────┘
```

## ER Diagram (Future with Users)

```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│     Users       │       │   Categories    │       │      Tasks      │
├─────────────────┤       ├─────────────────┤       ├─────────────────┤
│ id (PK)         │◄─────┤│ id (PK)         │◄─────┤│ id (PK)         │
│ username        │   1:N ││ name            │   1:N ││ title           │
│ email           │       ││ color           │       ││ description     │
│ password_hash   │       ││ description     │       ││ status          │
│ first_name      │       ││ user_id (FK)    │       ││ priority        │
│ last_name       │       ││ created_at      │       ││ due_date        │
│ is_active       │       ││ updated_at      │       ││ category_id(FK) │
│ is_premium      │       └─────────────────┘       ││ user_id (FK)    │
│ created_at      │                                 ││ created_at      │
│ updated_at      │◄────────────────────────────────┤│ updated_at      │
└─────────────────┘                             1:N └─────────────────┘
```

## Database Constraints

### Primary Keys
- All tables have auto-incrementing integer primary keys

### Foreign Keys
- `tasks.category_id` → `categories.id` (ON DELETE SET NULL)
- `tasks.user_id` → `users.id` (ON DELETE CASCADE)
- `categories.user_id` → `users.id` (ON DELETE CASCADE)

### Unique Constraints
- `categories.name` (per user when user system is implemented)
- `users.username`
- `users.email`

### Check Constraints
```sql
-- Task status validation
ALTER TABLE tasks ADD CONSTRAINT chk_task_status 
CHECK (status IN ('pending', 'in_progress', 'completed'));

-- Task priority validation
ALTER TABLE tasks ADD CONSTRAINT chk_task_priority 
CHECK (priority IN ('low', 'medium', 'high', 'critical'));

-- Color format validation
ALTER TABLE categories ADD CONSTRAINT chk_category_color 
CHECK (color REGEXP '^#[0-9A-Fa-f]{6}$');
```

## Database Triggers

### Auto-update timestamps
```sql
-- Update updated_at on tasks modification
CREATE TRIGGER update_tasks_updated_at 
AFTER UPDATE ON tasks
FOR EACH ROW 
BEGIN
    UPDATE tasks SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Update updated_at on categories modification
CREATE TRIGGER update_categories_updated_at 
AFTER UPDATE ON categories
FOR EACH ROW 
BEGIN
    UPDATE categories SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;
```