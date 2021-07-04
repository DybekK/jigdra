import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {NotFoundComponent} from "./not-found/not-found.component";
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {NzTypographyModule} from "ng-zorro-antd/typography";
import {NzButtonModule} from "ng-zorro-antd/button";
import {AppRoutingModule} from "../../app-routing.module";


@NgModule({
  declarations: [
    NotFoundComponent
  ],
  exports: [
    NotFoundComponent
  ],
  imports: [
    CommonModule,
    NzLayoutModule,
    NzTypographyModule,
    NzButtonModule,
    AppRoutingModule
  ]
})
export class ErrorsModule { }
