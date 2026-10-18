import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { type ComponentFixture, TestBed } from '@angular/core/testing';

import { TranslationService } from '../../core/i18n/translation.service';
import { MejorarRedaccion } from './mejorar-redaccion';

describe('MejorarRedaccion', () => {
  let fixture: ComponentFixture<MejorarRedaccion>;
  let http: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MejorarRedaccion],
      providers: [provideHttpClient(), provideHttpClientTesting()],
    }).compileComponents();

    TestBed.inject(TranslationService).cambiar('es');
    http = TestBed.inject(HttpTestingController);
    fixture = TestBed.createComponent(MejorarRedaccion);
    fixture.componentRef.setInput('numero', 'CS-2026-0001');
    fixture.componentRef.setInput('editor', 'comment');
    fixture.componentRef.setInput('borrador', 'hola necesito datos');
    fixture.detectChanges();
  });

  afterEach(() => http.verify());

  function boton(texto: string): HTMLButtonElement {
    const encontrado = [
      ...(fixture.nativeElement as HTMLElement).querySelectorAll<HTMLButtonElement>('button'),
    ].find((candidato) => candidato.textContent?.trim() === texto);
    expect(encontrado).toBeTruthy();
    return encontrado!;
  }

  it('genera sin modal ni tono y sólo aplica la propuesta cuando se confirma', async () => {
    let aplicado = '';
    fixture.componentInstance.aplicado.subscribe((texto) => (aplicado = texto));

    boton('Mejorar con IA').click();
    const peticion = http.expectOne('/api/tickets/CS-2026-0001/writing/improve');
    expect(peticion.request.body).toEqual({ editor: 'comment', draft: 'hola necesito datos' });
    peticion.flush({ text: 'Hola, ¿podría compartir los datos necesarios?', qualityWarning: false });

    await fixture.whenStable();
    fixture.detectChanges();
    const host = fixture.nativeElement as HTMLElement;
    expect(host.querySelector('dialog')).toBeNull();
    expect(host.querySelector('[role="region"]')?.textContent).toContain(
      'Hola, ¿podría compartir los datos necesarios?',
    );
    expect(aplicado).toBe('');

    boton('Usar propuesta').click();
    fixture.detectChanges();
    expect(aplicado).toBe('Hola, ¿podría compartir los datos necesarios?');
    expect(host.querySelector('[role="region"]')).toBeNull();
  });

  it('marca la propuesta como obsoleta y regenera desde el borrador actual evitando la última', async () => {
    boton('Mejorar con IA').click();
    http.expectOne('/api/tickets/CS-2026-0001/writing/improve').flush({
      text: 'Primera propuesta',
      qualityWarning: false,
    });
    await fixture.whenStable();
    fixture.detectChanges();

    fixture.componentRef.setInput('borrador', 'borrador actualizado');
    fixture.detectChanges();
    expect(boton('Usar propuesta').disabled).toBe(true);
    expect((fixture.nativeElement as HTMLElement).textContent).toContain('El borrador cambió');

    boton('Generar otra propuesta').click();
    const regeneracion = http.expectOne('/api/tickets/CS-2026-0001/writing/improve');
    expect(regeneracion.request.body).toEqual({
      editor: 'comment',
      draft: 'borrador actualizado',
      previousSuggestion: 'Primera propuesta',
    });
    regeneracion.flush({ text: 'Propuesta diferente', qualityWarning: false });
    await fixture.whenStable();
    fixture.detectChanges();
    expect((fixture.nativeElement as HTMLElement).textContent).toContain('Propuesta diferente');
    expect(boton('Usar propuesta').disabled).toBe(false);
  });

  it('descarta la propuesta o comunica el error fuera del panel', async () => {
    let error = '';
    fixture.componentInstance.fallido.subscribe((texto) => (error = texto));

    boton('Mejorar con IA').click();
    http.expectOne('/api/tickets/CS-2026-0001/writing/improve').flush({
      text: 'Una propuesta',
      qualityWarning: false,
    });
    await fixture.whenStable();
    fixture.detectChanges();
    boton('Conservar original').click();
    fixture.detectChanges();
    expect((fixture.nativeElement as HTMLElement).querySelector('[role="region"]')).toBeNull();

    boton('Mejorar con IA').click();
    http.expectOne('/api/tickets/CS-2026-0001/writing/improve').flush(
      { error: 'tickets.writing.unavailable' },
      { status: 503, statusText: 'Unavailable' },
    );
    await fixture.whenStable();
    fixture.detectChanges();
    expect(error).toContain('motor de IA no está disponible');
    expect((fixture.nativeElement as HTMLElement).querySelector('[role="region"]')).toBeNull();
  });

  it('avisa dentro del panel si la segunda propuesta cambia muy poco y permite usarla', async () => {
    boton('Mejorar con IA').click();
    http.expectOne('/api/tickets/CS-2026-0001/writing/improve').flush({
      text: 'Hola necesito datos.',
      qualityWarning: true,
    });
    await fixture.whenStable();
    fixture.detectChanges();

    const host = fixture.nativeElement as HTMLElement;
    expect(host.textContent).toContain('Esta propuesta cambió muy poco');
    expect(boton('Usar propuesta').disabled).toBe(false);
  });
});
