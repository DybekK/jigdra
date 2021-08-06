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
import {
  AppstoreOutline,
  BarChartOutline,
  BulbTwoTone,
  CalendarOutline,
  CloudOutline,
  DesktopOutline,
  FileAddOutline,
  FileOutline,
  LoadingOutline,
  LockOutline,
  PlusOutline,
  ProfileOutline,
  ProjectOutline,
  SettingOutline,
  ShopOutline,
  TeamOutline,
  UnorderedListOutline,
  UploadOutline,
  UserOutline,
  VideoCameraOutline
} from "@ant-design/icons-angular/icons";
import {NzGridModule} from "ng-zorro-antd/grid";
import {NzInputModule} from "ng-zorro-antd/input";
import {NzSliderModule} from "ng-zorro-antd/slider";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NzFormModule} from "ng-zorro-antd/form";
import {NzButtonModule} from "ng-zorro-antd/button";

const icons: IconDefinition[] = [
  UserOutline,
  LockOutline,
  TeamOutline,
  FileOutline,
  VideoCameraOutline,
  UploadOutline,
  BarChartOutline,
  CloudOutline,
  AppstoreOutline,
  ShopOutline,
  DesktopOutline,
  CalendarOutline,
  ProjectOutline,
  SettingOutline,
  FileAddOutline,
  ProfileOutline,
  PlusOutline,
  UnorderedListOutline,
  LoadingOutline,
  BulbTwoTone
];

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
        NzIconModule.forChild(icons),
        NzGridModule,
        NzInputModule,
        NzSliderModule,
        FormsModule,
        ReactiveFormsModule,
        NzFormModule,
        NzButtonModule
    ]
})
export class UserModule {
}
