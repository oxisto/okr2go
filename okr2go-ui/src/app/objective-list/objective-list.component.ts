import { Component, OnInit } from '@angular/core';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';

import { ObjectiveService } from '../objective.service';
import { Objective, KeyResult } from '../objective';

@Component({
  selector: 'app-objective-list',
  templateUrl: './objective-list.component.html',
  styleUrls: ['./objective-list.component.scss']
})
export class ObjectiveListComponent implements OnInit {

  objectives: Objective[];
  faPlus = faPlus;
  faMinus = faMinus;

  constructor(private objectiveService: ObjectiveService) {

  }

  ngOnInit() {
    this.refreshObjectives();
  }

  onPlus(objectiveId: number, keyResult: KeyResult) {
    this.objectiveService.resultPlusOne(objectiveId, keyResult.id).subscribe(_ => {
      this.refreshObjectives();
    })
  }

  onMinus(objectiveId: number, keyResult: KeyResult) {
    this.objectiveService.resultMinusOne(objectiveId, keyResult.id).subscribe(_ => {
      this.refreshObjectives();
    })
  }

  refreshObjectives() {
    this.objectiveService.getObjectives().subscribe(objectives => this.objectives = objectives);
  }

  getTypeForKeyResult(keyResult: KeyResult) {
    const percentage = keyResult.current / keyResult.target;

    if (percentage >= 0.7) {
      return 'success';
    } else if (percentage >= 0.3) {
      return 'warning';
    } else {
      return 'danger';
    }
  }
}
