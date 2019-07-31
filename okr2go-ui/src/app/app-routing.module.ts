import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ObjectiveListComponent } from './objective-list/objective-list.component';


const routes: Routes = [
  { path: 'objectives', component: ObjectiveListComponent },
  { path: '', redirectTo: '/objectives', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
