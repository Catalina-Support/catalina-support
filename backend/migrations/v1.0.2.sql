-- ============================================================================
-- v1.0.2 — las subcategorías
--
-- Una categoría puede tener subcategorías o no tenerlas. **La subcategoría es opcional en el
-- catálogo**: una categoría sin ninguna subcategoría activa sigue siendo válida. En el ticket,
-- en cambio, **es obligatoria cuando la categoría elegida tiene al menos una activa**, y esa
-- regla la comprueba el servicio al crear y al cambiar la clasificación.
--
-- Se guarda en su propia tabla y no como una categoría con padre a propósito: así el esquema
-- impide por construcción más de un nivel, y la clave ajena compuesta de `tickets` garantiza que
-- **la subcategoría pertenece a la categoría del ticket**.
--
-- No se reescribe `v1.0.0.sql` ni `v1.0.1.sql`, que ya están aplicadas. El archivo es idempotente.
-- docs/modules/tickets.md, sección 2.3.3.
-- ============================================================================

CREATE TABLE IF NOT EXISTS ticket_subcategories (
    id             bigserial   PRIMARY KEY,
    -- La categoría de la que cuelga. No se borra una categoría con subcategorías: se retira.
    category_id    bigint      NOT NULL REFERENCES ticket_categories (id) ON DELETE RESTRICT,
    -- El nombre tal cual se lee en la pantalla.
    name           text        NOT NULL,
    -- El nombre en minúsculas y sin acentos. Único **dentro de su categoría**: dos categorías
    -- distintas pueden tener cada una su «Web».
    normalized     text        NOT NULL,
    -- Retirada deja de ofrecerse; los tickets que la tienen la siguen enseñando.
    active         boolean     NOT NULL DEFAULT true,
    -- Nulo cuando no la creó nadie (los datos de ejemplo, por ejemplo).
    created_by_id  bigint      REFERENCES users (id) ON DELETE SET NULL,
    created_at     timestamptz NOT NULL DEFAULT now(),
    -- Cuándo se retiró. Volver a activarla lo limpia.
    deactivated_at timestamptz,
    CONSTRAINT ticket_subcategories_normalized_key UNIQUE (category_id, normalized),
    -- Pareja única que hace falta para la clave ajena compuesta de `tickets`.
    CONSTRAINT ticket_subcategories_id_category_key UNIQUE (id, category_id)
);

COMMENT ON TABLE ticket_subcategories IS
    'El catálogo de subcategorías: cuelgan de una categoría y son opcionales. Retirar es desactivar, y no se borra ninguna fila.';

COMMENT ON COLUMN ticket_subcategories.normalized IS
    'El nombre en minúsculas y sin acentos. Único dentro de su categoría.';

COMMENT ON COLUMN ticket_subcategories.category_id IS
    'La categoría a la que pertenece. Una categoría con subcategorías activas obliga a elegir una al crear el ticket.';

-- La columna del ticket nace nula: los tickets que ya existían no tienen subcategoría y siguen
-- siendo válidos. La obligación se aplica al crear y al cambiar la clasificación.
ALTER TABLE tickets ADD COLUMN IF NOT EXISTS subcategory_id bigint;

COMMENT ON COLUMN tickets.subcategory_id IS
    'La subcategoría del ticket: opcional, y obligatoria sólo cuando su categoría tiene subcategorías activas.';

-- **La clave ajena compuesta** es la que impide que un ticket tenga una subcategoría de otra
-- categoría, incluso escribiendo a mano en la base. Con `subcategory_id` nulo no se comprueba
-- nada, que es justo lo que hace falta cuando la categoría no tiene subcategorías.
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'tickets_subcategory_id_fkey') THEN
        ALTER TABLE tickets
            ADD CONSTRAINT tickets_subcategory_id_fkey
            FOREIGN KEY (subcategory_id, category_id)
            REFERENCES ticket_subcategories (id, category_id) ON DELETE RESTRICT;
    END IF;
END $$;

-- El índice de la subcategoría: es como se filtra la lista y como se cuenta cuántos tickets la usan.
CREATE INDEX IF NOT EXISTS tickets_subcategory_id_idx ON tickets (subcategory_id);
CREATE INDEX IF NOT EXISTS ticket_subcategories_category_id_idx ON ticket_subcategories (category_id);
