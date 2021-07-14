import {Component, OnInit} from '@angular/core';
import {NavigationEnd, Router} from "@angular/router";
import {filter} from "rxjs/operators";

@Component({
  selector: 'app-auth-content',
  template: `
    <nz-content class="auth-content">
      <nz-card class="box-shadow" id="cardTitle" [nzTitle]="cardTitle">
        <router-outlet></router-outlet>
      </nz-card>
    </nz-content>
  `,
  styleUrls: ["auth-content.component.scss"]
})
export class AuthContentComponent implements OnInit {
  cardTitle: string = "";

  setUrl(url: string) {
    this.cardTitle = url === '/register' ? 'Register' : 'Sign in';
  }

  constructor(private router: Router) {
    this.setUrl(router.url);
  }

  ngOnInit(): void {
    this.router.events.pipe(filter(event => event instanceof NavigationEnd))
      .subscribe(() => {
        this.setUrl(this.router.url);
    });
  }
}
