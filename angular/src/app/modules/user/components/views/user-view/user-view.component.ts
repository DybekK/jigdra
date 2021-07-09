import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-user-view',
  template: `
    <nz-layout class="user-view">
      <nz-layout class="user-view__inner-content">
        <app-user-profile></app-user-profile>

      </nz-layout>
    </nz-layout>
  `,
  styleUrls: ['./user-view.component.scss']
})
export class UserViewComponent implements OnInit {

  constructor() {
  }

  ngOnInit(): void {
  }

}
