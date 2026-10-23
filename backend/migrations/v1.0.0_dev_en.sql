-- Catalina-Support — development sample data IN ENGLISH
--
-- **Este archivo NO se aplica en producción.** El esquema y lo que la aplicación necesita para
-- funcionar están en `v1.0.0.sql`; esto son datos para mirar la aplicación con contenido: cuentas,
-- 25 tickets con su historia, escalados, adjuntos y avisos.
--
-- Se aplica **con su guion**, que además copia los archivos de los adjuntos:
--
--   ./scripts/dev-seed-en.sh
--
-- Es idempotente: empieza borrando los datos de ejemplo anteriores y los vuelve a crear. **Borra
-- todos los tickets**, también los que hayas creado a mano probando, así que deja el entorno en un
-- estado conocido: es un seeder de desarrollo, no una migración.
--
-- **Las once cuentas entran con la misma contraseña de desarrollo: `123123123`** (decisión del
-- responsable, 2026-09-25): en desarrollo se entra y se sale muchas veces al día, y una contraseña
-- que se escribe con una mano es lo que hace que probar no dé pereza.
--
-- **Cumple la política del producto**, que pide ocho caracteres como mínimo, así que este seeder no
-- se salta ninguna regla: la escribe aquí ya cifrada porque un archivo SQL no puede llamar a la
-- aplicación, y el resultado es el mismo que si se hubiera establecido desde la pantalla.

BEGIN;

-- ============================================================================
-- 1. Se limpia lo de la vez anterior
-- ============================================================================

-- Los tickets se llevan por delante sus comentarios, adjuntos, historial e internos (ON DELETE
-- CASCADE). Y van **antes** que las cuentas: un ticket apunta a quien lo pidió, y la base no deja
-- borrar una cuenta que tiene tickets detrás.
-- **Los resúmenes del motor de IA van con los tickets y hay que borrarlos aquí** (docs/modules/ai.md):
-- la tabla de los resúmenes se lleva por **número de ticket**, y este archivo **vuelve a emitir los
-- mismos números** (CS-2026-0001 y siguientes), así que una fila vieja se pegaría al ticket nuevo y
-- enseñaría el motivo de otro. No hay clave ajena que lo impida —los tickets viven en su módulo y los
-- resúmenes en el suyo—, así que se limpia a mano, aquí y en el mismo sitio donde se borran los
-- tickets.
DELETE FROM ai_insights;
-- Las etiquetas cuelgan del ticket y se van con él (ON DELETE CASCADE), pero se borran antes para
-- dejar dicho que este archivo las vuelve a crear. **El catálogo de etiquetas se borra también**:
-- este archivo deja el entorno en un estado conocido y lo vuelve a dar de alta con los tickets.
-- Las categorías van **después** de los tickets: un ticket apunta a la suya, y la base no deja
-- borrar una que esté en uso.
DELETE FROM ticket_tags;
DELETE FROM tickets;
DELETE FROM ticket_tag_names;
DELETE FROM ticket_categories;
DELETE FROM ticket_number_counters;

-- Las cuentas de ejemplo, con sus enlaces de contraseña. Las de las pruebas (`e2e-…`) no se tocan.
--
-- La comparación va **en minúsculas**: el correo se guarda tal y como lo escribe quien lo escribe, y
-- «Ana.Perez@Demo.com» es la misma persona que «ana.perez@demo.com». Sin esto, una cuenta de
-- ejemplo con mayúsculas sobrevivía a la limpieza y el entorno dejaba de estar en un estado conocido.
DELETE FROM password_tokens WHERE user_id IN (SELECT id FROM users WHERE lower(email) LIKE '%@demo.com');
-- **Las cuentas de las pruebas de interfaz también se borran** (decisión del responsable, 2026-09-28):
-- cada pasada de la capa de Playwright crea las suyas —con la marca del momento—, así que aquí no se
-- pierde nada, y dejarlas convertía el buscador de a quién etiquetar en una lista de **miles** de
-- personas que se llaman todas igual (2 521 cuentas y 636 técnicos activos, medido el 2026-09-28). El
-- entorno de ejemplo queda con las **once** cuentas de siempre.
DELETE FROM users WHERE lower(email) LIKE '%@demo.com';

-- ============================================================================
-- 2. Las cuentas: cinco que piden, tres de Soporte y tres de Desarrollo
-- ============================================================================

INSERT INTO users (name, last_name, email, password_hash, role, origin, is_active, created_at, updated_at) VALUES
    ('User', 'One', 'user1@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'usuario', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('User', 'Two', 'user2@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'usuario', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('User', 'Three', 'user3@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'usuario', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('User', 'Four', 'user4@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'usuario', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('User', 'Five', 'user5@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'usuario', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Support', 'One', 'support1@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'soporte', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Support', 'Two', 'support2@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'soporte', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Support', 'Three', 'support3@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'soporte', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Developer', 'One', 'dev1@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'desarrollo', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Developer', 'Two', 'dev2@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'desarrollo', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes')),
    ('Developer', 'Three', 'dev3@demo.com', '$2a$12$VD3neJ0zGzwWD2PmyPe55O6GRHbLfCd4fspMzk9U7WRRilLiZvFe.', 'desarrollo', 'local', true, (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'), (date_trunc('day', now()) - interval '90 days' + interval '8 hours 0 minutes'));

-- ============================================================================
-- 3. El catálogo de categorías
-- ============================================================================

-- Las cinco del ejemplo. **«General» es la de fábrica**: la crea la migración para que ningún ticket
-- pueda quedarse sin categoría, y aquí se vuelve a poner porque este archivo borra el catálogo para
-- dejarlo en un estado conocido. Las otras cuatro son las que enseñan la función: agrupar por «qué
-- es» el ticket.
--
-- `created_by_id` apunta a Soporte Uno salvo en «General», que no la crea nadie. `ON CONFLICT` deja
-- el archivo aplicable dos veces por si alguien lo lanza sin la limpieza de arriba.
INSERT INTO ticket_categories (name, normalized, active, created_by_id) VALUES
    ('General',    'general',    true, NULL),
    ('Network',        'network',        true, (SELECT id FROM users WHERE email = 'support1@demo.com')),
    ('Software',   'software',   true, (SELECT id FROM users WHERE email = 'support1@demo.com')),
    ('Licenses',  'licenses',  true, (SELECT id FROM users WHERE email = 'support1@demo.com')),
    ('Printers', 'printers', true, (SELECT id FROM users WHERE email = 'support1@demo.com'))
ON CONFLICT (normalized) DO NOTHING;

-- ============================================================================
-- 4. Los 25 tickets, con su asignación, sus comentarios y su historia
-- ============================================================================

-- **Las descripciones y los comentarios van en HTML**, porque es lo que son desde el 2026-09-26
-- (docs/modules/tickets.md, sección 2.3): cada texto de ejemplo va dentro de su `<p>`, sin más formato
-- del que hacía falta. **Ninguno lleva adjuntos dentro del texto**: los adjuntos de ejemplo siguen
-- colgando del ticket y del comentario, y se ven en la lista del final, que es el camino que hay que
-- poder probar.

-- CS-2026-0001 · No puedo entrar en la aplicación
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0001', 2026, 1, 'I cannot sign in to the application', '<p>Since yesterday it says my password is incorrect, but I have not changed anything.</p>', 'nuevo',
        (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'user1@demo.com'), NULL,
        NULL, NULL,
        (date_trunc('day', now()) - interval '1 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '1 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0001'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '1 days' + interval '9 hours 0 minutes'));

-- CS-2026-0002 · El informe mensual sale con las fechas cambiadas
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0002', 2026, 2, 'The monthly report shows the dates incorrectly', '<p>I download the August report and the date columns use a different format. I am attaching a screenshot.</p>', 'nuevo',
        (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'user2@demo.com'), NULL,
        NULL, NULL,
        (date_trunc('day', now()) - interval '2 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '1 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0002'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '2 days' + interval '9 hours 0 minutes'));

-- El archivo de este adjunto lo copia `scripts/dev-seed-en.sh` desde `config/seed/en/screenshot.png`.
INSERT INTO ticket_attachments (ticket_id, uploaded_by_id, filename, stored_name, content_type, size_bytes, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0002'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'screenshot.png', '2026/CS-2026-0002/eee4d316b5ebefd9b17df4eed40097a7.png', 'image/png', 1024, (date_trunc('day', now()) - interval '2 days' + interval '9 hours 30 minutes'));

-- CS-2026-0003 · Pido un usuario para la persona nueva
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0003', 2026, 3, 'Create a user for the new team member', '<p>A new colleague joined the team and needs access to the application.</p>', 'nuevo',
        (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'user3@demo.com'), NULL,
        NULL, NULL,
        (date_trunc('day', now()) - interval '3 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '2 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'licenses'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0003'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '3 days' + interval '9 hours 0 minutes'));

-- CS-2026-0004 · La impresora de la segunda planta no imprime
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0004', 2026, 4, 'The second-floor printer does not print', '<p>It sends the job, but nothing comes out. We have already turned it off and on.</p>', 'escalado',
        (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '6 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '5 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'printers'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'Someone needs to check the print queue server; the workstation is not the cause.', (date_trunc('day', now()) - interval '4 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '5 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '6 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '6 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>Does it happen to everyone or only on your workstation?</p>', (date_trunc('day', now()) - interval '5 days' + interval '11 hours 15 minutes'));
INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), (SELECT id FROM users WHERE email = 'user4@demo.com'), '<p>Everyone: three of us tested it.</p>', (date_trunc('day', now()) - interval '4 days' + interval '12 hours 15 minutes'));

-- INT-CS-2026-0004: el ticket interno que cuelga de CS-2026-0004
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0004'), 'INT-CS-2026-0004', 'nuevo', 'Someone needs to check the print queue server; the workstation is not the cause.', (SELECT id FROM users WHERE email = 'support1@demo.com'), NULL,
        NULL, NULL,
        (date_trunc('day', now()) - interval '4 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '3 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'Someone needs to check the print queue server; the workstation is not the cause.', (date_trunc('day', now()) - interval '4 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0004'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '4 days' + interval '12 hours 20 minutes'));

-- CS-2026-0005 · Los adjuntos de más de 10 MB no se suben
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0005', 2026, 5, 'Attachments larger than 10 MB do not upload', '<p>I try to attach a 12 MB PDF, but it keeps loading and never finishes.</p>', 'en progreso',
        (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '9 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '8 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0005'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '8 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0005'), (SELECT id FROM users WHERE email = 'user5@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '9 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0005'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '9 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0005'), (SELECT id FROM users WHERE email = 'support2@demo.com'), '<p>We are looking into it. In the meantime, can you leave it in the shared folder?</p>', (date_trunc('day', now()) - interval '8 days' + interval '11 hours 15 minutes'));
INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0005'), (SELECT id FROM users WHERE email = 'user5@demo.com'), '<p>Done, it is there.</p>', (date_trunc('day', now()) - interval '7 days' + interval '12 hours 15 minutes'));

-- CS-2026-0006 · Cambio de correo en mi cuenta
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0006', 2026, 6, 'Change the email on my account', '<p>My company email changed and I want to keep signing in with it.</p>', 'en progreso',
        (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '12 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '11 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'licenses'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0006'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '11 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0006'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '12 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0006'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '12 days' + interval '10 hours 0 minutes'));

-- CS-2026-0007 · La aplicación va lenta por la mañana
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0007', 2026, 7, 'The application is slow in the morning', '<p>Between nine and ten, opening the inbox takes a long time.</p>', 'escalado',
        (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '14 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '13 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'network'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'The database is causing the slowdown, so Development needs to investigate.', (date_trunc('day', now()) - interval '12 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '13 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '14 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '14 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>Does it always happen on the same workstation?</p>', (date_trunc('day', now()) - interval '13 days' + interval '11 hours 15 minutes'));
INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'support2@demo.com'), '<p>Yes, on mine; the workstation next to it is fine.</p>', (date_trunc('day', now()) - interval '12 days' + interval '12 hours 15 minutes'));

-- El archivo de este adjunto lo copia `scripts/dev-seed-en.sh` desde `config/seed/en/screenshot.png`.
INSERT INTO ticket_attachments (ticket_id, uploaded_by_id, filename, stored_name, content_type, size_bytes, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'screenshot.png', '2026/CS-2026-0007/c9385c11ec0efd571d5faa7118635a20.png', 'image/png', 1024, (date_trunc('day', now()) - interval '14 days' + interval '9 hours 30 minutes'));

-- INT-CS-2026-0007: el ticket interno que cuelga de CS-2026-0007
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0007'), 'INT-CS-2026-0007', 'en progreso', 'The database is causing the slowdown, so Development needs to investigate.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev1@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '12 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '11 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'The database is causing the slowdown, so Development needs to investigate.', (date_trunc('day', now()) - interval '12 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0007'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), 'asignado', NULL, NULL, 'Developer One', (date_trunc('day', now()) - interval '12 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0007'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '12 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0007'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '11 days' + interval '10 hours 0 minutes'));

-- CS-2026-0008 · No me llegan los avisos de ticket nuevo
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0008', 2026, 8, 'I do not receive new ticket notifications', '<p>Tickets are assigned to me, but I do not receive the notification email.</p>', 'en progreso',
        (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '17 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '16 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0008'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '16 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0008'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '17 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0008'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '17 days' + interval '10 hours 0 minutes'));

-- CS-2026-0009 · Falta el teléfono de contacto en la ficha
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0009', 2026, 9, 'The contact phone number is missing from the profile', '<p>We need the phone number in the customer profile so we can call them.</p>', 'escalado',
        (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '20 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '19 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'Development needs to confirm whether the field can be added without changing the import.', (date_trunc('day', now()) - interval '18 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '19 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '20 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '20 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I sent it to Development. I will let you know as soon as I hear back.</p>', (date_trunc('day', now()) - interval '19 days' + interval '11 hours 15 minutes'));

-- INT-CS-2026-0009: el ticket interno que cuelga de CS-2026-0009
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0009'), 'INT-CS-2026-0009', 'en espera', 'Development needs to confirm whether the field can be added without changing the import.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev2@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '18 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '17 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'Development needs to confirm whether the field can be added without changing the import.', (date_trunc('day', now()) - interval '18 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0009'), (SELECT id FROM users WHERE email = 'dev2@demo.com'), 'asignado', NULL, NULL, 'Developer Two', (date_trunc('day', now()) - interval '18 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0009'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '18 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0009'), (SELECT id FROM users WHERE email = 'dev2@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '17 days' + interval '10 hours 0 minutes'));

-- CS-2026-0010 · Se ha borrado un comentario
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0010', 2026, 10, 'A comment was deleted', '<p>Yesterday I wrote a comment on a ticket and today it is gone.</p>', 'en espera',
        (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '22 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '21 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'general'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0010'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'en progreso', 'en espera', NULL, (date_trunc('day', now()) - interval '20 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0010'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '21 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0010'), (SELECT id FROM users WHERE email = 'user5@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '22 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0010'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '22 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0010'), (SELECT id FROM users WHERE email = 'support3@demo.com'), '<p>Do you remember which ticket it was? We will look into it.</p>', (date_trunc('day', now()) - interval '21 days' + interval '11 hours 15 minutes'));

-- CS-2026-0011 · Quiero cambiar el nombre de la instalación
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0011', 2026, 11, 'I want to change the installation name', '<p>The header shows the application name, but it should show ours.</p>', 'en espera',
        (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '25 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '24 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0011'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'en progreso', 'en espera', NULL, (date_trunc('day', now()) - interval '23 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0011'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '24 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0011'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '25 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0011'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '25 days' + interval '10 hours 0 minutes'));

-- CS-2026-0012 · El buscador no encuentra por número
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0012', 2026, 12, 'Search does not find tickets by number', '<p>If I paste the ticket number, nothing appears; searching by text works.</p>', 'en progreso',
        (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '28 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '27 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), NULL, 'estado', 'escalado', 'en progreso', 'Development resolved the internal ticket, so the ticket returns to Support.', (date_trunc('day', now()) - interval '24 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'Searching by number requires a query change.', (date_trunc('day', now()) - interval '26 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '27 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '28 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '28 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), (SELECT id FROM users WHERE email = 'support2@demo.com'), '<p>I tried it with number 4 and it does not appear either.</p>', (date_trunc('day', now()) - interval '27 days' + interval '11 hours 15 minutes'));

-- INT-CS-2026-0012: el ticket interno que cuelga de CS-2026-0012
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0012'), 'INT-CS-2026-0012', 'resuelto', 'Searching by number requires a query change.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev1@demo.com'),
        (date_trunc('day', now()) - interval '24 days' + interval '12 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '26 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '25 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0012'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'Searching by number requires a query change.', (date_trunc('day', now()) - interval '26 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0012'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), 'asignado', NULL, NULL, 'Developer One', (date_trunc('day', now()) - interval '26 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0012'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '24 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0012'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '26 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0012'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '25 days' + interval '10 hours 0 minutes'));

-- CS-2026-0013 · Pedir acceso a la carpeta compartida
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0013', 2026, 13, 'Request access to the shared folder', '<p>I need write access to the team shared folder.</p>', 'resuelto',
        (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        (date_trunc('day', now()) - interval '29 days' + interval '13 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '32 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '31 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'network'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0013'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '29 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0013'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '31 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0013'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '32 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0013'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '32 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0013'), (SELECT id FROM users WHERE email = 'support3@demo.com'), '<p>I can write to it now, thank you very much.</p>', (date_trunc('day', now()) - interval '31 days' + interval '11 hours 15 minutes'));

-- CS-2026-0014 · El correo de alta llegó a spam
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0014', 2026, 14, 'The account email went to spam', '<p>Just reporting this in case it affects others: the account email went to spam.</p>', 'resuelto',
        (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        (date_trunc('day', now()) - interval '32 days' + interval '13 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '35 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '34 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'general'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0014'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '32 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0014'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '34 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0014'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '35 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0014'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '35 days' + interval '10 hours 0 minutes'));

-- CS-2026-0015 · Duplicado: la impresora de la segunda planta
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0015', 2026, 15, 'Duplicate: the second-floor printer', '<p>I opened this by mistake; there was already a printer ticket.</p>', 'resuelto',
        (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        (date_trunc('day', now()) - interval '35 days' + interval '13 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '38 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '37 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'printers'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0015'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '35 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0015'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '37 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0015'), (SELECT id FROM users WHERE email = 'user5@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '38 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0015'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '38 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0015'), (SELECT id FROM users WHERE email = 'support3@demo.com'), '<p>Yes, it is the same issue. We will continue in the other ticket.</p>', (date_trunc('day', now()) - interval '37 days' + interval '11 hours 15 minutes'));

-- CS-2026-0016 · La contraseña caducaba y no lo sabía
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0016', 2026, 16, 'My password was expiring and I did not know', '<p>It should warn me before it expires, not when I try to sign in.</p>', 'resuelto',
        (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        (date_trunc('day', now()) - interval '38 days' + interval '13 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '41 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '40 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), NULL, 'estado', 'escalado', 'en progreso', 'Development returned the ticket to Support.', (date_trunc('day', now()) - interval '37 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '38 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'There is no expiration warning; it needs to be added to the authentication module.', (date_trunc('day', now()) - interval '39 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '40 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '41 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '41 days' + interval '10 hours 0 minutes'));

-- El archivo de este adjunto lo copia `scripts/dev-seed-en.sh` desde `config/seed/en/report.pdf`.
INSERT INTO ticket_attachments (ticket_id, uploaded_by_id, filename, stored_name, content_type, size_bytes, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'report.pdf', '2026/CS-2026-0016/748c64619c90ab6963f56c425ef9150b.pdf', 'application/pdf', 1024, (date_trunc('day', now()) - interval '41 days' + interval '9 hours 30 minutes'));

-- INT-CS-2026-0016: el ticket interno que cuelga de CS-2026-0016
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0016'), 'INT-CS-2026-0016', 'cerrado', 'There is no expiration warning; it needs to be added to the authentication module.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev3@demo.com'),
        (date_trunc('day', now()) - interval '37 days' + interval '12 hours 0 minutes'), (date_trunc('day', now()) - interval '36 days' + interval '15 hours 0 minutes'),
        (date_trunc('day', now()) - interval '39 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '38 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'There is no expiration warning; it needs to be added to the authentication module.', (date_trunc('day', now()) - interval '39 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'dev3@demo.com'), 'asignado', NULL, NULL, 'Developer Three', (date_trunc('day', now()) - interval '39 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'dev3@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '37 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '36 days' + interval '15 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '39 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0016'), (SELECT id FROM users WHERE email = 'dev3@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '38 days' + interval '10 hours 0 minutes'));

-- CS-2026-0017 · Quiero ordenar la bandeja por fecha
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0017', 2026, 17, 'I want to sort the inbox by date', '<p>It shows the oldest items first by default, but I prefer the newest ones.</p>', 'resuelto',
        (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        (date_trunc('day', now()) - interval '42 days' + interval '13 hours 0 minutes'), NULL,
        (date_trunc('day', now()) - interval '45 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '44 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0017'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '42 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0017'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '44 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0017'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '45 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0017'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '45 days' + interval '10 hours 0 minutes'));

-- CS-2026-0018 · El ticket se cerró sin querer
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0018', 2026, 18, 'The ticket was closed by mistake', '<p>It closed automatically after becoming inactive, but the problem continues.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        (date_trunc('day', now()) - interval '47 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '45 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '50 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '49 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'general'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '45 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'reabierto', 'cerrado', 'en progreso', NULL, (date_trunc('day', now()) - interval '46 days' + interval '15 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '47 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '48 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '49 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '50 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '50 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support3@demo.com'), '<p>Can you reopen it? It is still happening.</p>', (date_trunc('day', now()) - interval '49 days' + interval '11 hours 15 minutes'));
INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0018'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>Reopened, sorry about that.</p>', (date_trunc('day', now()) - interval '48 days' + interval '12 hours 15 minutes'));

-- CS-2026-0019 · No puedo subir un adjunto en ZIP
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0019', 2026, 19, 'I cannot upload a ZIP attachment', '<p>I attach a ZIP file with the logs, but it is rejected.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        (date_trunc('day', now()) - interval '51 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '49 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '54 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '53 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '49 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '51 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '53 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '54 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '54 days' + interval '10 hours 0 minutes'));

-- El archivo de este adjunto lo copia `scripts/dev-seed-en.sh` desde `config/seed/en/notes.txt`.
INSERT INTO ticket_attachments (ticket_id, uploaded_by_id, filename, stored_name, content_type, size_bytes, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0019'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'notes.txt', '2026/CS-2026-0019/fb3af71bf3908eb8dbc6105c1edff4e2.txt', 'text/plain', 1024, (date_trunc('day', now()) - interval '54 days' + interval '9 hours 30 minutes'));

-- CS-2026-0020 · Cambiar el idioma de los correos
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0020', 2026, 20, 'Change the language of email notifications', '<p>The notifications arrive in Spanish and I want them in English.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        (date_trunc('day', now()) - interval '55 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '53 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '58 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '57 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0020'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '53 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0020'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '55 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0020'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '57 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0020'), (SELECT id FROM users WHERE email = 'user5@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '58 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0020'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '58 days' + interval '10 hours 0 minutes'));

-- CS-2026-0021 · Informe de tickets de agosto
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0021', 2026, 21, 'August ticket report', '<p>I need the August list for the team meeting.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'user1@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        (date_trunc('day', now()) - interval '59 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '57 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '62 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '61 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0021'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '57 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0021'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '59 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0021'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '61 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0021'), (SELECT id FROM users WHERE email = 'user1@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '62 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0021'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '62 days' + interval '10 hours 0 minutes'));

-- CS-2026-0022 · El logo se ve pequeño en móvil
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0022', 2026, 22, 'The logo looks small on mobile', '<p>The logo looks very small on the mobile sign-in screen.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'user2@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        (date_trunc('day', now()) - interval '63 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '61 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '66 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '65 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0022'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '61 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0022'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '63 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0022'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '65 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0022'), (SELECT id FROM users WHERE email = 'user2@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '66 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0022'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '66 days' + interval '10 hours 0 minutes'));

-- CS-2026-0023 · Falta el aviso cuando alguien escala
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0023', 2026, 23, 'The escalation notification is missing', '<p>Development does not know when a ticket needs their attention.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'user3@demo.com'), (SELECT id FROM users WHERE email = 'support3@demo.com'),
        (date_trunc('day', now()) - interval '67 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '65 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '70 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '69 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '65 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), NULL, 'estado', 'escalado', 'en progreso', 'Development returned the ticket to Support.', (date_trunc('day', now()) - interval '66 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '67 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'The escalation notification was missing because of an assignment configuration error.', (date_trunc('day', now()) - interval '68 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '69 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'user3@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '70 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support3@demo.com'), 'asignado', NULL, NULL, 'Support Three', (date_trunc('day', now()) - interval '70 days' + interval '10 hours 0 minutes'));

INSERT INTO ticket_comments (ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), (SELECT id FROM users WHERE email = 'support3@demo.com'), '<p>It turned out to be the notification configuration.</p>', (date_trunc('day', now()) - interval '69 days' + interval '11 hours 15 minutes'));

-- INT-CS-2026-0023: el ticket interno que cuelga de CS-2026-0023
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0023'), 'INT-CS-2026-0023', 'cerrado', 'The escalation notification was missing because of an assignment configuration error.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev2@demo.com'),
        (date_trunc('day', now()) - interval '66 days' + interval '12 hours 0 minutes'), (date_trunc('day', now()) - interval '65 days' + interval '15 hours 0 minutes'),
        (date_trunc('day', now()) - interval '68 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '67 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'The escalation notification was missing because of an assignment configuration error.', (date_trunc('day', now()) - interval '68 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'dev2@demo.com'), 'asignado', NULL, NULL, 'Developer Two', (date_trunc('day', now()) - interval '68 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'dev2@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '66 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '65 days' + interval '15 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '68 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0023'), (SELECT id FROM users WHERE email = 'dev2@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '67 days' + interval '10 hours 0 minutes'));

-- CS-2026-0024 · Dos tickets con el mismo número
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0024', 2026, 24, 'Two tickets have the same number', '<p>I saw two consecutive tickets with the same number.</p>', 'cerrado',
        (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'user4@demo.com'), (SELECT id FROM users WHERE email = 'support1@demo.com'),
        (date_trunc('day', now()) - interval '71 days' + interval '13 hours 0 minutes'), (date_trunc('day', now()) - interval '69 days' + interval '16 hours 0 minutes'),
        (date_trunc('day', now()) - interval '74 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '73 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'software'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '69 days' + interval '16 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '71 days' + interval '13 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '73 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '74 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'asignado', NULL, NULL, 'Support One', (date_trunc('day', now()) - interval '74 days' + interval '10 hours 0 minutes'));

-- El archivo de este adjunto lo copia `scripts/dev-seed-en.sh` desde `config/seed/en/screenshot.png`.
INSERT INTO ticket_attachments (ticket_id, uploaded_by_id, filename, stored_name, content_type, size_bytes, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0024'), (SELECT id FROM users WHERE email = 'user4@demo.com'), 'screenshot.png', '2026/CS-2026-0024/04cd1e6b4d81a038f95ec51a6438abba.png', 'image/png', 1024, (date_trunc('day', now()) - interval '74 days' + interval '9 hours 30 minutes'));

-- CS-2026-0025 · Pedir la baja de una cuenta antigua
INSERT INTO tickets (number, number_year, number_seq, subject, description, state, requester_id, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at, category_id)
VALUES ('CS-2026-0025', 2026, 25, 'Request deactivation of an old account', '<p>A person left the organization and their account is still active.</p>', 'escalado',
        (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'user5@demo.com'), (SELECT id FROM users WHERE email = 'support2@demo.com'),
        NULL, NULL,
        (date_trunc('day', now()) - interval '80 days' + interval '9 hours 0 minutes'), (date_trunc('day', now()) - interval '79 days' + interval '17 hours 0 minutes'), (SELECT id FROM ticket_categories WHERE normalized = 'licenses'));

INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), NULL, 'estado', 'escalado', 'en progreso', 'Development returned the ticket to Support.', (date_trunc('day', now()) - interval '76 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'escalado', 'en progreso', 'escalado', 'Account removal is not implemented; the account must be deactivated manually.', (date_trunc('day', now()) - interval '78 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'estado', 'nuevo', 'en progreso', NULL, (date_trunc('day', now()) - interval '79 days' + interval '10 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), (SELECT id FROM users WHERE email = 'user5@demo.com'), 'creado', NULL, 'nuevo', NULL, (date_trunc('day', now()) - interval '80 days' + interval '9 hours 0 minutes'));
INSERT INTO ticket_history (ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), (SELECT id FROM users WHERE email = 'support2@demo.com'), 'asignado', NULL, NULL, 'Support Two', (date_trunc('day', now()) - interval '80 days' + interval '10 hours 0 minutes'));

-- INT-CS-2026-0025: el ticket interno que cuelga de CS-2026-0025
INSERT INTO internal_tickets (ticket_id, number, state, escalation_reason, created_by_id, assignee_id, resolved_at, closed_at, created_at, updated_at)
VALUES ((SELECT id FROM tickets WHERE number = 'CS-2026-0025'), 'INT-CS-2026-0025', 'cerrado', 'Account removal is not implemented; the account must be deactivated manually.', (SELECT id FROM users WHERE email = 'support1@demo.com'), (SELECT id FROM users WHERE email = 'dev1@demo.com'),
        (date_trunc('day', now()) - interval '76 days' + interval '12 hours 0 minutes'), (date_trunc('day', now()) - interval '75 days' + interval '15 hours 0 minutes'),
        (date_trunc('day', now()) - interval '78 days' + interval '11 hours 0 minutes'), (date_trunc('day', now()) - interval '77 days' + interval '16 hours 0 minutes'));

INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'creado', NULL, 'nuevo', 'Account removal is not implemented; the account must be deactivated manually.', (date_trunc('day', now()) - interval '78 days' + interval '11 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), 'asignado', NULL, NULL, 'Developer One', (date_trunc('day', now()) - interval '78 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), 'resuelto', 'en progreso', 'resuelto', NULL, (date_trunc('day', now()) - interval '76 days' + interval '12 hours 0 minutes'));
INSERT INTO ticket_history (internal_ticket_id, actor_id, event, from_state, to_state, detail, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'support1@demo.com'), 'cerrado', 'resuelto', 'cerrado', NULL, (date_trunc('day', now()) - interval '75 days' + interval '15 hours 0 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'support1@demo.com'), '<p>I am sending it over with everything I have tried so far.</p>', (date_trunc('day', now()) - interval '78 days' + interval '12 hours 20 minutes'));
INSERT INTO ticket_comments (internal_ticket_id, author_id, body, created_at)
VALUES ((SELECT id FROM internal_tickets WHERE number = 'INT-CS-2026-0025'), (SELECT id FROM users WHERE email = 'dev1@demo.com'), '<p>I will look at it today and let you know.</p>', (date_trunc('day', now()) - interval '77 days' + interval '10 hours 0 minutes'));

-- ============================================================================
-- 5. El contador de la numeración
-- ============================================================================

-- Los números que se han usado: el siguiente ticket de ejemplo que cree alguien será el 26.
INSERT INTO ticket_number_counters (year, last_number) VALUES (2026, 25);

-- Las etiquetas: **ahora viven en su propio catálogo** (`ticket_tag_names`) y los tickets las llevan
-- por clave ajena (docs/modules/tickets.md, decisión 72). Se dan de alta primero las que hacen falta
-- —distintas, que varias se repiten entre tickets— y después se enganchan. Siguen en minúsculas y con
-- guion medio en lugar de espacios, que es como se guardan.
--
-- Se dan de alta a nombre de Soporte Uno. `ON CONFLICT` deja el archivo aplicable dos veces por si
-- alguien lo lanza sin la limpieza de arriba.
INSERT INTO ticket_tag_names (tag, normalized, created_by_id)
SELECT v.tag, v.tag, u.id
FROM (VALUES
    ('access'),
    ('attachments'),
    ('user-creation'),
    ('notifications'),
    ('deactivation'),
    ('inbox'),
    ('deletion'),
    ('search'),
    ('expiration'),
    ('shared-folder'),
    ('closed'),
    ('comment'),
    ('password'),
    ('email'),
    ('account'),
    ('duplicate'),
    ('escalation'),
    ('profile'),
    ('date-format'),
    ('language'),
    ('printer'),
    ('report'),
    ('installation'),
    ('slowness'),
    ('logo'),
    ('morning'),
    ('mobile'),
    ('name'),
    ('numbering'),
    ('number'),
    ('sorting'),
    ('permissions'),
    ('reopen'),
    ('second-floor'),
    ('spam'),
    ('size'),
    ('phone'),
    ('tickets'),
    ('zip')
) AS v(tag)
JOIN users u ON u.email = 'support1@demo.com'
ON CONFLICT (normalized) DO NOTHING;

-- Y ahora cada ticket con las suyas: **dos por ticket**, enganchadas por el número, que es como se
-- lee aquí.
INSERT INTO ticket_tags (ticket_id, tag_id, created_by_id)
SELECT t.id, n.id, u.id
FROM (VALUES
    ('CS-2026-0001', 'password'),
    ('CS-2026-0001', 'access'),
    ('CS-2026-0002', 'report'),
    ('CS-2026-0002', 'date-format'),
    ('CS-2026-0003', 'user-creation'),
    ('CS-2026-0003', 'permissions'),
    ('CS-2026-0004', 'printer'),
    ('CS-2026-0004', 'second-floor'),
    ('CS-2026-0005', 'attachments'),
    ('CS-2026-0005', 'size'),
    ('CS-2026-0006', 'email'),
    ('CS-2026-0006', 'account'),
    ('CS-2026-0007', 'slowness'),
    ('CS-2026-0007', 'morning'),
    ('CS-2026-0008', 'notifications'),
    ('CS-2026-0008', 'email'),
    ('CS-2026-0009', 'profile'),
    ('CS-2026-0009', 'phone'),
    ('CS-2026-0010', 'comment'),
    ('CS-2026-0010', 'deletion'),
    ('CS-2026-0011', 'name'),
    ('CS-2026-0011', 'installation'),
    ('CS-2026-0012', 'search'),
    ('CS-2026-0012', 'number'),
    ('CS-2026-0013', 'shared-folder'),
    ('CS-2026-0013', 'permissions'),
    ('CS-2026-0014', 'email'),
    ('CS-2026-0014', 'spam'),
    ('CS-2026-0015', 'duplicate'),
    ('CS-2026-0015', 'printer'),
    ('CS-2026-0016', 'password'),
    ('CS-2026-0016', 'expiration'),
    ('CS-2026-0017', 'inbox'),
    ('CS-2026-0017', 'sorting'),
    ('CS-2026-0018', 'closed'),
    ('CS-2026-0018', 'reopen'),
    ('CS-2026-0019', 'zip'),
    ('CS-2026-0019', 'attachments'),
    ('CS-2026-0020', 'language'),
    ('CS-2026-0020', 'email'),
    ('CS-2026-0021', 'report'),
    ('CS-2026-0021', 'tickets'),
    ('CS-2026-0022', 'logo'),
    ('CS-2026-0022', 'mobile'),
    ('CS-2026-0023', 'escalation'),
    ('CS-2026-0023', 'notifications'),
    ('CS-2026-0024', 'numbering'),
    ('CS-2026-0024', 'duplicate'),
    ('CS-2026-0025', 'deactivation'),
    ('CS-2026-0025', 'account')
) AS v(numero, tag)
JOIN tickets t ON t.number = v.numero
JOIN ticket_tag_names n ON n.normalized = v.tag
JOIN users u ON u.email = 'support1@demo.com'
ON CONFLICT (ticket_id, tag_id) DO NOTHING;

-- ============================================================================
-- 6. Los dos caminos de directorio, configurados pero apagados
-- ============================================================================

-- Las dos configuraciones quedan puestas y **el método se queda en `local`**: el entorno arranca
-- entrando con las once cuentas de ejemplo, y cambiar de método se hace desde Configuración, que es
-- justo lo que hay que poder probar a mano (docs/ambientes.md, sección 3.3).
--
-- Los dos servicios viven **cada uno en su archivo** (`active-directory.yml` y `keycloak.yml`) y se
-- levantan con su propio comando. Con ellos levantados, «Probar la conexión» contesta que sí y se
-- puede cambiar el método; sin ellos, contesta que no, que es la verdad.
--
-- La contraseña de la cuenta de servicio es la del LDIF de pruebas (`config/ldap/`): es un directorio
-- de mentira en un contenedor de desarrollo, no un secreto.

-- Los ajustes de una instalación sellada pertenecen al usuario: no se reescriben.
DO $$
BEGIN
IF EXISTS (SELECT 1 FROM installation_settings WHERE id = 1 AND installed_at IS NULL) THEN
UPDATE directory_settings SET
    host = 'ldap',
    port = '389',
    use_tls = false,
    bind_dn = 'cn=admin,dc=ejemplo,dc=com',
    bind_password = 'admin-directorio',
    search_base = 'ou=personas,dc=ejemplo,dc=com',
    user_filter = '(objectClass=inetOrgPerson)',
    attr_username = 'uid',
    attr_email = 'mail',
    attr_name = 'givenName',
    attr_last_name = 'sn',
    attr_id = 'uid',
    updated_at = now()
WHERE id = 1;

UPDATE keycloak_settings SET
    issuer = 'http://127.0.0.1:11006/sso/realms/catalina-support',
    internal_issuer = 'http://keycloak:8080/sso/realms/catalina-support',
    client_id = 'catalina-support',
    client_secret = 'el-secreto-de-desarrollo',
    redirect_uri = 'http://127.0.0.1:11001/api/auth/keycloak/callback',
    updated_at = now()
WHERE id = 1;

UPDATE installation_settings
   SET installation_name = 'Catalina Support',
       language = 'en',
       entry_method = 'local',
       -- **El correo saliente de desarrollo vive aquí**, no en el entorno (decisión del responsable,
       -- 2026-09-30): es el buzón de pruebas, y sin él no saldría ningún correo en el entorno.
       ai_url          = 'http://catalina_support_ai:8080',
       ai_model        = 'qwen2.5-1.5b-instruct',
       smtp_host       = 'mail',
       smtp_port       = '1025',
       smtp_security_mode = 'plain',
       smtp_insecure_ack  = true,
       smtp_user       = '',
       smtp_password   = '',
       smtp_from_name  = 'Catalina Support',
       smtp_from_email = 'no-responder@catalina-support.local',
       time_zone      = 'America/Guayaquil',
       public_app_url = 'http://127.0.0.1:11001',
       -- La demo incompleta queda instalada; las pruebas recorren su propia instalación vacía.
       installed_at   = COALESCE(installed_at, now()),
       updated_at     = now()
 WHERE id = 1;

END IF;
END $$;

-- Las cuatro prioridades aparecen en la demo. La marca evita que volver a sembrar simule un aviso:
-- los seeders cargan datos; los correos se prueban mediante la aplicación.
UPDATE tickets
SET priority = CASE
    WHEN number_seq IN (4, 17) THEN 'critical'
    WHEN number_seq IN (2, 7, 11, 16, 21, 24) THEN 'high'
    WHEN number_seq IN (3, 8, 13, 18, 22, 25) THEN 'low'
    ELSE 'medium'
END,
critical_notified_at = CASE WHEN number_seq IN (4, 17) THEN created_at ELSE NULL END;

COMMIT;
