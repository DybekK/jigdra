import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-user-nav',
  template: `
    <nz-layout>
      <nz-sider nzCollapsible nzWidth="200px">
        <div class="logo"></div>

        <ul nz-menu nzTheme="dark" nzMode="inline">
          <!--User block-->
          <ul nz-submenu nzTitle="User" nzIcon="user" nzOpen="{{open}}">
            <li nz-menu-item>
              <i nz-icon nzType="profile"></i>
              <span>Profile</span>
            </li>
            <li nz-menu-item>
              <i nz-icon nzType="setting"></i>
              <span>Settings</span>
            </li>
          </ul>
          <!--Team block-->
          <li nz-submenu nzTitle="Team" nzIcon="team" nzOpen="{{open}}">
            <ul>
              <li nz-menu-item *ngFor="let user of users">
                <i nz-icon nzType="{{user.work_state}}" [nzTheme]="user.work_state_theme"
                   [nzTwotoneColor]="user.work_state_color"></i>
                <span>{{user.user_name}}</span></li>
            </ul>
          </li>
          <!--Project block-->
          <li nz-submenu nzTitle="Project" nzIcon="desktop" nzOpen="{{open}}">
            <ul>

              <li nz-submenu nzTitle="Select project" nzIcon="unordered-list" nzOpen="{{open}}">
                <ul>
                  <li nz-menu-item *ngFor="let project of projects">
                    <i nz-icon nzType="{{project.icon}}" nzTheme="outline"></i>
                    <span><a routerLink="project-main">{{project.title}}</a></span>
                  </li>
                </ul>
              </li>

              <li nz-submenu nzTitle="Your tasks" nzIcon="clock-circle" nzOpen="{{open}}">
                <ul>
                  <li nz-menu-item>
                    <i nz-icon nzType="unordered-list" nzTheme="outline"></i>
                    <span><a routerLink="user-tasks">Tasks</a></span>
                  </li>
                  <li nz-menu-item>
                    <i nz-icon nzType="plus" nzTheme="outline"></i>
                    <span><a routerLink="add-task">New Task</a></span>
                  </li>
                </ul>
              </li>

              <li nz-menu-item>
                <i nz-icon nzType="calendar" nzTheme="outline"></i>
                <span><a routerLink="calendar">Calendar</a></span>
              </li>

              <li nz-menu-item>
                <i nz-icon nzType="file-add" nzTheme="outline"></i>
                <span><a routerLink="project-create">Create project</a></span>
              </li>
              <li nz-menu-item>
                <i nz-icon nzType="setting" nzTheme="outline"></i>
                <span><a routerLink="project-settings">Project settings</a></span>
              </li>
            </ul>
          </li>

          <li nz-menu-item>
            <i nz-icon nzType="video-camera"></i>
            <span>nav 2</span>
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
            <router-outlet></router-outlet>
          </div>
        </nz-content>
        <app-footer></app-footer>
      </nz-layout>
    </nz-layout>
  `,
  styleUrls: ['./user-nav.component.scss']
})
export class UserNavComponent implements OnInit {
  open: Boolean = true;

  users: { user_name: string, work_state: string, work_state_theme: 'fill' | 'outline' | 'twotone', work_state_color: string }[] = [
    {
      user_name: 'Bob',
      work_state: 'bulb',
      work_state_theme: 'twotone',
      work_state_color: '#52c41a'
    }, {
      user_name: 'Alex',
      work_state: 'bulb',
      work_state_theme: 'twotone',
      work_state_color: '#e0b21d'
    }, {
      user_name: 'Cassandra',
      work_state: 'bulb',
      work_state_theme: 'twotone',
      work_state_color: '#bd0404'
    }, {
      user_name: 'Kotlin',
      work_state: 'bulb',
      work_state_theme: 'twotone',
      work_state_color: ''
    }
  ]
  projects = [
    {
      title: 'Project 1',
      icon: 'project'
    },
    {
      title: 'Project 2',
      icon: 'project'
    }
  ]

  constructor() {
  }

  ngOnInit(): void {
  }

}
