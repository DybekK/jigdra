import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-footer',
  template: `
    <nz-footer class="footer">Jigdra @2021 Implemented By Angular</nz-footer>
  `,
  styleUrls: ['./footer.component.scss']
})
export class FooterComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
