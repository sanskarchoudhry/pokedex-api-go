-- 1. Generations (e.g., Generation I, II, III)
-- We start here because Species depend on Generations.
CREATE TABLE generations (
    id INT PRIMARY KEY, -- We will manually map this to 1, 2, 3...
    name VARCHAR(50) NOT NULL, -- "generation-i"
    region_name VARCHAR(50) NOT NULL -- "kanto" (main region)
);

-- 2. Types (e.g., Fire, Water)
CREATE TABLE types (
    id INT PRIMARY KEY, -- PokeAPI IDs
    name VARCHAR(50) NOT NULL UNIQUE
);

-- 3. Abilities (e.g., Overgrow, Blaze)
CREATE TABLE abilities (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL -- The short effect text
);

-- 4. Moves (e.g., Scratch, Flamethrower)
CREATE TABLE moves (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    type_id INT REFERENCES types(id),
    power INT, -- Can be NULL for status moves
    accuracy INT, -- Can be NULL (never miss)
    pp INT NOT NULL,
    damage_class VARCHAR(20) NOT NULL, -- "physical", "special", "status"
    description TEXT NOT NULL
);

-- 5. Pokemon Species (The "DNA")
-- Stores shared data like Pokedex ID, Evolution Chain references, etc.
CREATE TABLE species (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    generation_id INT NOT NULL REFERENCES generations(id),
    pokedex_order INT, -- Useful for sorting
    description TEXT NOT NULL, -- Flavor text (e.g. from Red/Blue)
    base_happiness INT,
    capture_rate INT
);

-- 6. Pokemon (The "Forms")
-- Links back to Species. One Species (Meowth) -> Many Pokemon (Meowth, Meowth-Alola)
CREATE TABLE pokemon (
    id INT PRIMARY KEY,
    species_id INT NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL, -- "meowth-galar"
    height INT NOT NULL,
    weight INT NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT false,
    sprite_url TEXT NOT NULL
);

-- ---------------------------------------------------------
-- THE MANY-TO-MANY JUNCTION TABLES
-- ---------------------------------------------------------

-- Links Pokemon to Types (e.g. Charizard -> Fire (Slot 1), Flying (Slot 2))
CREATE TABLE pokemon_types (
    pokemon_id INT REFERENCES pokemon(id) ON DELETE CASCADE,
    type_id INT REFERENCES types(id) ON DELETE CASCADE,
    slot INT NOT NULL, -- 1 or 2
    PRIMARY KEY (pokemon_id, slot) -- Composite Key
);

-- Links Pokemon to Abilities
CREATE TABLE pokemon_abilities (
    pokemon_id INT REFERENCES pokemon(id) ON DELETE CASCADE,
    ability_id INT REFERENCES abilities(id) ON DELETE CASCADE,
    is_hidden BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (pokemon_id, ability_id)
);

-- Links Pokemon to Moves
-- This solves your "filtering" issue. We store HOW they learn it.
CREATE TABLE pokemon_moves (
    pokemon_id INT REFERENCES pokemon(id) ON DELETE CASCADE,
    move_id INT REFERENCES moves(id) ON DELETE CASCADE,
    learn_method VARCHAR(50) NOT NULL, -- "level-up", "machine", "tutor", "egg"
    level_learned INT NOT NULL DEFAULT 0, -- 0 if by TM/Egg
    PRIMARY KEY (pokemon_id, move_id, learn_method, level_learned)
);

-- 7. Evolutions (Adjacency List)
-- This maps one Species to another.
-- From "Bulbasaur" -> To "Ivysaur" via "Level Up" at "16"
CREATE TABLE evolutions (
    id SERIAL PRIMARY KEY, -- Auto-increment internal ID
    from_species_id INT NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    to_species_id INT NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    trigger VARCHAR(50) NOT NULL, -- "level-up", "item", "trade"
    min_level INT, -- NULL if not level up
    item_name VARCHAR(50), -- NULL if not item evo
    held_item_name VARCHAR(50), -- For trade evos
    time_of_day VARCHAR(20), -- "day", "night"
    known_move_type_id INT REFERENCES types(id), -- For Sylveon (knows Fairy move)
    
    -- Constraint: Prevent duplicate evolution paths
    UNIQUE(from_species_id, to_species_id, trigger)
);