import { HttpParams } from '@angular/common/http';
import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { DashboardService, type DashboardSummary, type Point } from './dashboard.service';
import { TranslationService } from '../../core/i18n/translation.service';

@Component({selector:'app-dashboard-page',imports:[FormsModule],templateUrl:'./dashboard-page.html',host:{class:'block h-full min-h-0'}})
export class DashboardPage {
  private readonly api=inject(DashboardService); private readonly route=inject(ActivatedRoute); private readonly router=inject(Router); private readonly i18n=inject(TranslationService);
  protected readonly summary=signal<DashboardSummary|null>(null); protected readonly loading=signal(true); protected readonly error=signal(false);
  protected from=''; protected to=''; protected type=''; protected state=''; protected priority='';
  constructor(){const q=this.route.snapshot.queryParamMap;const now=new Date();const month=new Date(now);month.setDate(now.getDate()-30);this.from=q.get('from')??month.toISOString().slice(0,10);this.to=q.get('to')??now.toISOString().slice(0,10);this.type=q.get('type')??'';this.state=q.get('state')??'';this.priority=q.get('priority')??'';void this.refresh(false)}
  protected t(){return this.i18n.textos().dashboard}
  private params(){let p=new HttpParams().set('from',this.from).set('to',this.to);for(const [k,v] of Object.entries({type:this.type,state:this.state,priority:this.priority})){if(v)p=p.set(k,v)}return p}
  protected async refresh(updateUrl=true){this.loading.set(true);this.error.set(false);if(updateUrl)await this.router.navigate([],{relativeTo:this.route,queryParams:{from:this.from,to:this.to,type:this.type||null,state:this.state||null,priority:this.priority||null},replaceUrl:true});try{this.summary.set(await this.api.summary(this.params()))}catch{this.error.set(true)}finally{this.loading.set(false)}}
  protected async exportXlsx(){try{const blob=await this.api.export(this.params());const url=URL.createObjectURL(blob);const link=document.createElement('a');link.href=url;link.download=`report_${this.from}_${this.to}.xlsx`;link.click();URL.revokeObjectURL(url)}catch{this.error.set(true)}}
  protected duration(metric:{value:number;available:boolean}){if(!metric.available)return '—';const m=Math.round(metric.value);return `${Math.floor(m/60)} h ${m%60} min`}
  protected compare(metric:{value:number;previous:number}){const delta=metric.value-metric.previous;return `${delta>=0?'+':''}${Math.round(delta)}`}
  protected max(points:readonly Point[]){return Math.max(1,...points.map(p=>Math.max(p.value,p.secondary??0)))}
  protected hasClosedTickets(points:readonly Point[]){return points.some(p=>(p.secondary??0)>0)}
  protected showTimelineLabel(index:number,total:number,maximumLabels:number){if(total<=maximumLabels||index===0||index===total-1)return true;const step=Math.ceil((total-1)/(maximumLabels-1));return index%step===0}
  protected label(kind:string,key:string){if(kind==='state')return ({nuevo:this.t().new,'en progreso':this.t().inProgress,'en espera':this.t().waiting,escalado:this.t().escalated,resuelto:this.t().resolved,cerrado:this.t().closed} as Record<string,string>)[key]??key;if(kind==='priority')return ({critical:this.t().priorityCritical,high:this.t().priorityHigh,medium:this.t().priorityMedium,low:this.t().priorityLow} as Record<string,string>)[key]??key;if(kind==='type')return key==='main'?this.t().main:key==='internal'?this.t().internal:key;if(key==='unassigned')return this.t().unassigned;if(key==='other')return this.t().other;return key}
}
