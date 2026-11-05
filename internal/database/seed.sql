INSERT INTO users (username, email, password_hash, is_active, is_email_verified)
VALUES
    -- Admins and verified base accs (password: Password123!)
    ('admin', 'admin@example.local', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('carlos.mendoza', 'carlos.mendoza@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('ana.torres', 'ana.torres@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('diego_silva', 'diego.silva@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('sofia.valdez', 'sofia.valdez@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('mateo_rios', 'mateo.rios@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('valeria.gomez', 'valeria.gomez@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('lucas.herrera', 'lucas.herrera@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('camila_navarro', 'camila.navarro@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('javier.castro', 'javier.castro@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('elena_rojas', 'elena.rojas@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('rodrigo.morales', 'rodrigo.morales@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('paula.vargas', 'paula.vargas@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('fernando_ortiz', 'fernando.ortiz@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('daniela.flores', 'daniela.flores@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('gabriel_mendez', 'gabriel.mendez@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('natalia.lara', 'natalia.lara@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('alejandro.ibarra', 'alejandro.ibarra@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('monica_delgado', 'monica.delgado@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),
    ('samuel.pena', 'samuel.pena@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE),

    -- active users with unverified email
    ('marina_bravo', 'marina.bravo@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, FALSE),
    ('hector.acosta', 'hector.acosta@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, FALSE),
    ('renata_soto', 'renata.soto@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, FALSE),
    ('andres.cordero', 'andres.cordero@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, FALSE),
    ('claudia.cabrera', 'claudia.cabrera@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, FALSE),

    -- blocked users
    ('suspendido.01', 'suspendido1@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', FALSE, TRUE),
    ('banned_account', 'banned@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', FALSE, TRUE),
    ('inactive_user', 'inactive@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', FALSE, FALSE),
    ('pending_review', 'review.pending@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', FALSE, FALSE),

    -- generic account
    ('testuser', 'testuser@example.com', '$argon2id$v=19$m=65536,t=3,p=4$SxqjuDuB8HWWIgVzGMEzxQ$YEEkJFucaxndGaG60527PLNsKff5N+Pj41k4Yh/HcJE', TRUE, TRUE)
ON CONFLICT ((LOWER(username))) DO UPDATE SET password_hash = EXCLUDED.password_hash;
