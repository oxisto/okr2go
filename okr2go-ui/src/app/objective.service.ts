import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Objective } from './objective';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ObjectiveService {

  constructor(private http: HttpClient) {
  }

  getObjectives(): Observable<Objective[]> {
    return this.http.get<Objective[]>('/api/objectives').pipe(map(data => {
      return Object.assign(new Objective, data);
    }));
  }

}
