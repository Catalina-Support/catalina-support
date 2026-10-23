import { Component, OnDestroy, computed, inject, input, output, signal } from '@angular/core';

import { TranslationService } from '../../core/i18n/translation.service';
import { Boton } from '../../shared/components/boton';
import { claveDelError } from '../../shared/errores';
import { TicketsService, type EditorDeRedaccion } from './tickets.service';

/** Revisa una propuesta junto al editor sin sustituir nunca un borrador de forma automática. */
@Component({
  selector: 'app-mejorar-redaccion',
  imports: [Boton],
  template: `
    <div class="flex flex-col gap-3">
      <div>
        <app-boton
          forma="secundario"
          nombreDeIcono="ia"
          [texto]="propuesta() ? t().tickets.iaRedaccionRegenerar : t().tickets.iaRedaccionAbrir"
          [textoTrabajando]="t().tickets.iaRedaccionGenerando"
          [trabajando]="generando()"
          [deshabilitado]="!borrador().trim()"
          (pulsado)="generar()"
        />
      </div>

      @if (propuesta()) {
        <section
          class="flex flex-col gap-3 rounded-lg border border-borde bg-superficie-suave p-4"
          role="region"
          [attr.aria-label]="t().tickets.iaRedaccionTitulo"
        >
          <h3 class="text-sm font-semibold text-texto">{{ t().tickets.iaRedaccionTitulo }}</h3>
          <p class="select-text whitespace-pre-line text-sm text-texto">{{ propuesta() }}</p>

          @if (obsoleta()) {
            <p class="text-sm text-aviso" role="status">
              {{ t().tickets.iaRedaccionDesactualizada }}
            </p>
          }

          <div class="flex flex-wrap gap-2">
            <app-boton
              [texto]="t().tickets.iaRedaccionUsar"
              [deshabilitado]="obsoleta()"
              (pulsado)="usar()"
            />
            <app-boton
              forma="secundario"
              [texto]="t().tickets.iaRedaccionConservar"
              (pulsado)="conservarOriginal()"
            />
          </div>
        </section>
      }
    </div>
  `,
})
export class MejorarRedaccion implements OnDestroy {
  readonly numero = input.required<string>();
  readonly editor = input.required<EditorDeRedaccion>();
  readonly borrador = input.required<string>();
  readonly aplicado = output<string>();
  readonly fallido = output<string>();

  private readonly tickets = inject(TicketsService);
  private readonly textos = inject(TranslationService);
  protected readonly t = this.textos.textos;

  protected readonly generando = signal(false);
  protected readonly propuesta = signal('');
  private readonly borradorDeLaPropuesta = signal('');
  private generacion = 0;

  protected readonly obsoleta = computed(
    () => this.propuesta() !== '' && this.borrador().trim() !== this.borradorDeLaPropuesta(),
  );

  ngOnDestroy(): void {
    this.generacion++;
  }

  protected async generar(): Promise<void> {
    const borrador = this.borrador().trim();
    if (!borrador || this.generando()) return;

    const turno = ++this.generacion;
    const anterior = this.propuesta().trim() || undefined;
    this.generando.set(true);
    try {
      const respuesta = await this.tickets.mejorarRedaccion(
        this.numero(),
        this.editor(),
        borrador,
        anterior,
      );
      if (turno !== this.generacion) return;
      this.propuesta.set(respuesta.text);
      this.borradorDeLaPropuesta.set(borrador);
    } catch (error) {
      if (turno !== this.generacion) return;
      this.propuesta.set('');
      this.borradorDeLaPropuesta.set('');
      this.fallido.emit(this.textos.error(claveDelError(error)));
    } finally {
      if (turno === this.generacion) this.generando.set(false);
    }
  }

  protected usar(): void {
    if (!this.propuesta() || this.obsoleta()) return;
    this.aplicado.emit(this.propuesta());
    this.conservarOriginal();
  }

  protected conservarOriginal(): void {
    this.generacion++;
    this.generando.set(false);
    this.propuesta.set('');
    this.borradorDeLaPropuesta.set('');
  }
}
