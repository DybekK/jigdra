import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-user-profile',
  template: `
    <nz-layout>
      <nz-sider nzCollapsible nzWidth="200px">
        <div class="logo"></div>
        <ul nz-menu nzTheme="dark" nzMode="inline">
          <li nz-submenu nzTitle="User" nzIcon="user">
            <ul>
              <li nz-menu-item>Tom</li>
              <li nz-menu-item>Bill</li>
              <li nz-menu-item>Alex</li>
            </ul>
          </li>
          <li nz-submenu nzTitle="Team" nzIcon="team">
            <ul>
              <li nz-menu-item>Team 1</li>
              <li nz-menu-item>Team 2</li>
            </ul>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="file"></i>
            <span>nav 1</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="video-camera"></i>
            <span>nav 2</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="upload"></i>
            <span>nav 3</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="bar-chart"></i>
            <span>nav 4</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="cloud-o"></i>
            <span>nav 5</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="appstore-o"></i>
            <span>nav 6</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="team"></i>
            <span>nav 7</span>
          </li>
          <li nz-menu-item>
            <i nz-icon nzType="shop"></i>
            <span>nav 8</span>
          </li>
        </ul>
      </nz-sider>
      <nz-layout class="right-layout">
        <nz-header></nz-header>
        <nz-content>
          <nz-breadcrumb>
            <nz-breadcrumb-item>User</nz-breadcrumb-item>
            <nz-breadcrumb-item>Bill</nz-breadcrumb-item>
          </nz-breadcrumb>
          <div class="inner-content">
            <app-user-content></app-user-content>
          </div>
        </nz-content>
        <app-footer></app-footer>
      </nz-layout>
    </nz-layout>
  `,
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent implements OnInit {

  constructor() {
  }

  ngOnInit(): void {
  }

}
