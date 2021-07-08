import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserViewComponent} from "./components/views/user-view/user-view.component";

const routes: Routes = [
  {
    path: '', component: UserViewComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule {
}
