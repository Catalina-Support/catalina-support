import { HttpParams, provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { DashboardService } from './dashboard.service';

describe('DashboardService', () => {
  let service: DashboardService;
  let http: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({ providers: [provideHttpClient(), provideHttpClientTesting()] });
    service = TestBed.inject(DashboardService);
    http = TestBed.inject(HttpTestingController);
  });

  afterEach(() => http.verify());

  it('uses the same statistical filters for the summary and XLSX export', async () => {
    const params = new HttpParams().set('from', '2026-10-01').set('to', '2026-10-18').set('type', 'internal');

    void service.summary(params);
    const summary = http.expectOne((request) => request.url === '/api/dashboard/summary');
    expect(summary.request.params.get('type')).toBe('internal');
    summary.flush({});

    const exported = service.export(params);
    const xlsx = http.expectOne((request) => request.url === '/api/dashboard/export');
    expect(xlsx.request.responseType).toBe('blob');
    expect(xlsx.request.params.get('type')).toBe('internal');
    xlsx.flush(new Blob([], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' }));
    await expect(exported).resolves.toBeInstanceOf(Blob);
  });
});
