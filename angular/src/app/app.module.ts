import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { en_US } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {AuthModule} from "./modules/auth/auth.module";
import {ErrorsModule} from "./modules/errors/errors.module";
import { FooterComponent } from './shared/components/atoms/footer/footer.component';
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {SharedModule} from "./shared/shared.module";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzIconModule} from "ng-zorro-antd/icon";


registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    SharedModule,
    AuthModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ErrorsModule,
    AppRoutingModule,
    NzLayoutModule,
  ],
  providers: [{provide: NZ_I18N, useValue: en_US}],
  bootstrap: [AppComponent]
})
export class AppModule { }
