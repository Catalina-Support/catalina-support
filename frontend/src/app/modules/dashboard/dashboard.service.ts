import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { firstValueFrom } from 'rxjs';

export interface Metric { readonly value:number; readonly previous:number; readonly available:boolean }
export interface Point { readonly key:string; readonly value:number; readonly secondary?:number }
export interface DashboardSummary {
  readonly open:Metric; readonly openCritical:Metric; readonly overdue:Metric;
  readonly averageFirstResponseMinutes:Metric; readonly averageResolutionMinutes:Metric; readonly reopens:Metric;
  readonly unanswered:number; readonly grouping:string;
  readonly createdClosed:readonly Point[]; readonly byState:readonly Point[]; readonly byPriority:readonly Point[];
  readonly byCategory:readonly Point[]; readonly byAssignee:readonly Point[];
}
@Injectable({providedIn:'root'})
export class DashboardService {
  private readonly http=inject(HttpClient);
  summary(params:HttpParams){return firstValueFrom(this.http.get<DashboardSummary>('/api/dashboard/summary',{params}))}
  export(params:HttpParams){return firstValueFrom(this.http.get('/api/dashboard/export',{params,responseType:'blob'}))}
}
