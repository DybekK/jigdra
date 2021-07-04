import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-not-found',
  template:`
    <nz-layout>
      <nz-header></nz-header>
      <nz-content>
        <h1 nz-typography>Page not found</h1>
        <h2 nz-typography>Error 404</h2>
        <button nz-button nzType="primary" routerLink="/">Return to home page</button>
      </nz-content>
      <nz-footer></nz-footer>
    </nz-layout>
  `,
  styleUrls: ['./not-found.component.css']
})
export class NotFoundComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
