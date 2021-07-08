import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserRoutingModule} from './user-routing.module';
import {UserViewComponent} from './components/views/user-view/user-view.component';
import {UserProfileComponent} from './components/templates/user-profile/user-profile.component';
import {UserContentComponent} from './components/organisms/user-content/user-content.component';
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {NzBreadCrumbModule} from "ng-zorro-antd/breadcrumb";
import {NzMenuModule} from "ng-zorro-antd/menu";
import {NzIconModule} from "ng-zorro-antd/icon";
import {SharedModule} from "../../shared/shared.module";
import {IconDefinition} from "@ant-design/icons-angular";
import {LockOutline, UserOutline} from "@ant-design/icons-angular/icons";

const icons: IconDefinition[] = [ UserOutline, LockOutline ];

@NgModule({
  declarations: [
    UserViewComponent,
    UserProfileComponent,
    UserContentComponent
  ],
  imports: [
    CommonModule,
    UserRoutingModule,
    NzLayoutModule,
    NzBreadCrumbModule,
    NzMenuModule,
    NzIconModule,
    SharedModule,
    NzIconModule.forChild(icons)
  ]
})
export class UserModule {
}
