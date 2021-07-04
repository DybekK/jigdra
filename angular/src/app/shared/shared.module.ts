import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {FooterComponent} from "./components/atoms/footer/footer.component";
import {NzLayoutModule} from "ng-zorro-antd/layout";

@NgModule({
  declarations: [
    FooterComponent
  ],
  exports: [
    FooterComponent
  ],
  imports: [
    CommonModule,
    NzLayoutModule
  ]
})
export class SharedModule { }
