import { Component, effect, input, output } from '@angular/core';

import { Aviso, type MensajeDePantalla } from './aviso';

/**
 * Un resultado transitorio que permanece visible aunque la acción ocurra al final de una pantalla.
 *
 * Todos los resultados admiten cierre manual y desaparecen como máximo a los diez segundos
 * (`docs/interfaz-y-experiencia.md`, secciones 15 y 24).
 */
@Component({
  selector: 'app-toast',
  imports: [Aviso],
  host: {
    class:
      'fixed inset-x-4 top-4 z-50 block sm:left-auto sm:right-4 sm:w-full sm:max-w-md',
  },
  template: `
    <app-aviso
      [forma]="mensaje().forma"
      [texto]="mensaje().texto"
      [etiquetaCerrar]="etiquetaCerrar()"
      (cerrado)="cerrar()"
    />
  `,
})
export class Toast {
  readonly mensaje = input.required<MensajeDePantalla>();
  readonly etiquetaCerrar = input.required<string>();
  readonly cerrado = output<void>();

  constructor() {
    /** Cada mensaje nuevo sustituye al anterior y estrena sus diez segundos. */
    effect((limpiar) => {
      this.mensaje();
      const temporizador = window.setTimeout(() => this.cerrar(), 10_000);
      limpiar(() => window.clearTimeout(temporizador));
    });
  }

  protected cerrar(): void {
    this.cerrado.emit();
  }
}
