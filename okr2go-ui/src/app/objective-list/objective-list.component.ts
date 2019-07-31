import { Component, OnInit } from '@angular/core';
import { ObjectiveService } from '../objective.service';
import { Objective } from '../objective';

@Component({
  selector: 'app-objective-list',
  templateUrl: './objective-list.component.html',
  styleUrls: ['./objective-list.component.scss']
})
export class ObjectiveListComponent implements OnInit {

  objectives: Objective[];

  constructor(private objectiveService: ObjectiveService) {

  }

  ngOnInit() {
    this.objectiveService.getObjectives().subscribe(objectives => this.objectives = objectives);
  }

}
