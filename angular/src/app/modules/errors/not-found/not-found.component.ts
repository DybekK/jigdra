import { Component, OnInit } from '@angular/core';

@Component({
  selector:  'app-not-found',
  template:`
    <div class="error-content">
      <nz-result nzStatus="404" nzTitle="404" nzSubTitle="Sorry, the page you visited does not exist.">
        <div nz-result-extra>
          <button [routerLink]="'/login'" nz-button nzType="primary">Back Home</button>
        </div>
      </nz-result>
    </div>
  `,
  styleUrls: ['./not-found.component.scss']
})
export class NotFoundComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
