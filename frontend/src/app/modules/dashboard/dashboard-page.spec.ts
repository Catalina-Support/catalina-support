import { TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';

import { TranslationService } from '../../core/i18n/translation.service';
import { DashboardPage } from './dashboard-page';
import { DashboardService, type DashboardSummary } from './dashboard.service';

const metric = { value: 0, previous: 0, available: true } as const;

function summary(closed: number): DashboardSummary {
  return {
    open: metric,
    openCritical: metric,
    overdue: metric,
    averageFirstResponseMinutes: metric,
    averageResolutionMinutes: metric,
    reopens: metric,
    unanswered: 0,
    grouping: 'day',
    createdClosed: [{ key: '2026-10-19', value: 3, secondary: closed }],
    byState: [],
    byPriority: [],
    byCategory: [],
    byAssignee: [],
  };
}

describe('DashboardPage timeline', () => {
  async function mount(closed: number) {
    await TestBed.configureTestingModule({
      imports: [DashboardPage],
      providers: [
        { provide: DashboardService, useValue: { summary: () => Promise.resolve(summary(closed)), export: () => Promise.resolve(new Blob()) } },
        { provide: ActivatedRoute, useValue: { snapshot: { queryParamMap: { get: () => null } } } },
        { provide: Router, useValue: { navigate: () => Promise.resolve(true) } },
      ],
    }).compileComponents();
    TestBed.inject(TranslationService).cambiar('en');
    const fixture = TestBed.createComponent(DashboardPage);
    fixture.detectChanges();
    await fixture.whenStable();
    fixture.detectChanges();
    return fixture;
  }

  afterEach(() => TestBed.resetTestingModule());

  it('shows both values without interaction and explains a period with no closures', async () => {
    const fixture = await mount(0);
    const root = fixture.nativeElement as HTMLElement;

    expect(root.querySelector('[aria-label="2026-10-19 · Created: 3"]')?.textContent).toContain('3');
    expect(root.querySelector('[aria-label="2026-10-19 · Closed: 0"]')?.textContent).toContain('0');
    expect(root.textContent).toContain('No tickets closed in this period.');
    expect(root.querySelector('.bg-dashboard-created')).not.toBeNull();
    expect(root.querySelector('.bg-dashboard-closed')).not.toBeNull();
  });

  it('does not show the no-closures explanation when the period contains a closure', async () => {
    const fixture = await mount(2);
    const root = fixture.nativeElement as HTMLElement;

    expect(root.querySelector('[aria-label="2026-10-19 · Closed: 2"]')?.textContent).toContain('2');
    expect(root.textContent).not.toContain('No tickets closed in this period.');
  });
});
