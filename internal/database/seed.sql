INSERT INTO users (username, email, password_hash, is_active, is_email_verified)
VALUES
    -- Admins and verified base accs (password: Password123!)
    ('admin', 'admin@example.local', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('carlos.mendoza', 'carlos.mendoza@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('ana.torres', 'ana.torres@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('diego_silva', 'diego.silva@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('sofia.valdez', 'sofia.valdez@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('mateo_rios', 'mateo.rios@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('valeria.gomez', 'valeria.gomez@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('lucas.herrera', 'lucas.herrera@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('camila_navarro', 'camila.navarro@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('javier.castro', 'javier.castro@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('elena_rojas', 'elena.rojas@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('rodrigo.morales', 'rodrigo.morales@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('paula.vargas', 'paula.vargas@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('fernando_ortiz', 'fernando.ortiz@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('daniela.flores', 'daniela.flores@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('gabriel_mendez', 'gabriel.mendez@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('natalia.lara', 'natalia.lara@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('alejandro.ibarra', 'alejandro.ibarra@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('monica_delgado', 'monica.delgado@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),
    ('samuel.pena', 'samuel.pena@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE),

    -- active users with unverified email
    ('marina_bravo', 'marina.bravo@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, FALSE),
    ('hector.acosta', 'hector.acosta@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, FALSE),
    ('renata_soto', 'renata.soto@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, FALSE),
    ('andres.cordero', 'andres.cordero@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, FALSE),
    ('claudia.cabrera', 'claudia.cabrera@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, FALSE),

    -- blocked users
    ('suspendido.01', 'suspendido1@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', FALSE, TRUE),
    ('banned_account', 'banned@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', FALSE, TRUE),
    ('inactive_user', 'inactive@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', FALSE, FALSE),
    ('pending_review', 'review.pending@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', FALSE, FALSE),

    -- generic account
    ('testuser', 'testuser@example.com', '$2a$12$e86gkVvN8V8a.Wk7kQd3e.sMv34Zf4Q6pL9GqZ9T0Z5xK7qF7y6S6', TRUE, TRUE)
ON CONFLICT DO NOTHING;