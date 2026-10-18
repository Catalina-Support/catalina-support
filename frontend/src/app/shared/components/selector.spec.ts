import { TestBed } from '@angular/core/testing';

import { Selector } from './selector';

describe('Selector', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({ imports: [Selector] }).compileComponents();
  });

  it('usa 42 px en toda la aplicación sin agrandar la tipografía', () => {
    const fixture = TestBed.createComponent(Selector);
    fixture.componentRef.setInput('identificador', 'tema');
    fixture.componentRef.setInput('etiqueta', 'Tema');
    fixture.componentRef.setInput('opciones', [
      { valor: 'claro', etiqueta: 'Claro', grupo: 'Temas' },
    ]);
    fixture.componentRef.setInput('valor', 'claro');
    fixture.detectChanges();

    const selector = (fixture.nativeElement as HTMLElement).querySelector('select');
    expect(selector?.classList.contains('h-[42px]')).toBe(true);
    expect(selector?.classList.contains('px-3')).toBe(true);
    expect(selector?.classList.contains('py-2')).toBe(true);
    expect(selector?.classList.contains('text-sm')).toBe(true);
    expect(selector?.classList.contains('min-h-12')).toBe(false);
    expect(selector?.classList.contains('min-h-[52px]')).toBe(false);
    expect(selector?.classList.contains('min-h-11')).toBe(false);
    expect(selector?.classList.contains('min-h-9')).toBe(false);
  });
});
