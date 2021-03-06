import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Objective, KeyResult } from './objective';
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

  resultPlusOne(objectiveId: number, resultId: string): Observable<KeyResult> {
    return this.http.get<Objective[]>('/api/objectives/' + objectiveId + '/results/' + resultId + '/plus').pipe(map(data => {
      return Object.assign(new KeyResult, data);
    }));
  }

  resultMinusOne(objectiveId: number, resultId: string): Observable<KeyResult> {
    return this.http.get<Objective[]>('/api/objectives/' + objectiveId + '/results/' + resultId + '/minus').pipe(map(data => {
      return Object.assign(new KeyResult, data);
    }));
  }

  postKeyResult(objectiveId: number, result: KeyResult): Observable<KeyResult> {
    return this.http.post<KeyResult>('/api/objectives/' + objectiveId + '/results', result).pipe(map(data => {
      return Object.assign(new KeyResult, data);
    }));
  }

}
