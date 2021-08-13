import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserViewComponent} from "./components/views/user-view/user-view.component";
import {CalendarComponent} from "./components/organisms/calendar/calendar.component";
import {UserAddTaskComponent} from "./components/organisms/user-add-task/user-add-task.component";
import {ProjectCreateComponent} from "./components/organisms/project-create/project-create.component";
import {ProjectMainComponent} from "./components/organisms/project-main/project-main.component";
import {UserTasksComponent} from "./components/organisms/user-tasks/user-tasks.component";
import {UserProfileComponent} from "./components/organisms/user-profile/user-profile.component";

const routes: Routes = [
  {
    path: '', component: UserViewComponent,
    children: [
      {
        path: 'calendar',
        component: CalendarComponent
      },
      {
        path: 'project-create',
        component: ProjectCreateComponent
      },
      {
        path: 'project-main',
        component: ProjectMainComponent
      },
      {
        path: 'add-task',
        component: UserAddTaskComponent
      },
      {
        path: 'user-profile',
        component: UserProfileComponent
      },
      {
        path: 'user-tasks',
        component: UserTasksComponent
      }],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule {
}
