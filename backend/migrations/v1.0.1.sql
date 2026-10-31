-- ============================================================================
-- v1.0.1 — correo de cuenta creada desde un ticket
--
-- La plantilla número 14 avisa a la persona a la que se acaba de crear cuenta desde AD al abrirle
-- un ticket. No lleva contraseña ni token: entra con su cuenta de la organización.
--
-- No se reescribe v1.0.0.sql, que ya está aplicada. Las filas son idempotentes.
-- docs/modules/mail.md, sección 19; docs/modules/users.md, sección 18.
-- ============================================================================

INSERT INTO mail_templates (key, language, subject, body, default_subject, default_body) VALUES
('auth.accountCreated', 'es',
 'Tu cuenta en la mesa de ayuda',
 '<p>Hola {{nombre}}:</p><p>Se ha creado tu cuenta en la mesa de ayuda. Ya puedes entrar con tu cuenta de la organización.</p><p><a href="{{enlace}}">Entrar en la mesa de ayuda</a></p>',
 'Tu cuenta en la mesa de ayuda',
 '<p>Hola {{nombre}}:</p><p>Se ha creado tu cuenta en la mesa de ayuda. Ya puedes entrar con tu cuenta de la organización.</p><p><a href="{{enlace}}">Entrar en la mesa de ayuda</a></p>'),

('auth.accountCreated', 'en',
 'Your help desk account',
 '<p>Hi {{nombre}},</p><p>Your help desk account has been created. You can now sign in with your organization account.</p><p><a href="{{enlace}}">Open the help desk</a></p>',
 'Your help desk account',
 '<p>Hi {{nombre}},</p><p>Your help desk account has been created. You can now sign in with your organization account.</p><p><a href="{{enlace}}">Open the help desk</a></p>')
ON CONFLICT (key, language) DO NOTHING;
