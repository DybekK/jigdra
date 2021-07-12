import {Component, OnInit} from '@angular/core';
import {NzMarks, NzSliderValue} from "ng-zorro-antd/slider";
import {FormBuilder, FormGroup} from "@angular/forms";

@Component({
  selector: 'app-user-content',
  template: `
    <form [formGroup]="createTask" class="ant-advanced-search-form">
      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <nz-form-item>
            <nz-form-control nzErrorTip="You need to specify task name">
              <nz-input-group nzAddOnBefore="Task">
                <input formControlName="taskName" type="text" nz-input placeholder="Task name"/>
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>
      <div nz-row [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="24">
          <nz-form-item>
            <nz-form-control>
              <nz-slider formControlName="time" [nzMarks]="timeStamps" [nzStep]="1" nzMax="120"
                         (nzOnAfterChange)="getTimeValue($event)"></nz-slider>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>
      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="10">
          <nz-form-item>
            <nz-form-control nzErrorTip="Number of brakes">
              <nz-input-group nzAddOnBefore="Breaks">
                <input formControlName="breaksNumber" type="number" nz-input placeholder="Task name"/>
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>
        <div nz-col [nzSpan]="10">
          <nz-form-item>
            <nz-form-control nzErrorTip="You need to specify task name">
              <nz-input-group nzAddOnBefore="Lorem">
                <input type="text" nz-input placeholder="Lorem ipsum"/>
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>
      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <nz-form-item>
              <textarea formControlName="comment" nz-input placeholder="Write your comment here!"
                        [nzAutosize]="{ minRows: 2, maxRows: 5 }"></textarea>
          </nz-form-item>
        </div>
      </div>
      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <button nz-button [nzType]="'primary'">Crate task</button>
        </div>
      </div>
    </form>
  `,
  styleUrls: ['./user-content.component.scss']
})
export class UserContentComponent implements OnInit {

  createTask!: FormGroup;

  hGutter = {xs: 8, sm: 16, md: 24};
  vGutter = {xs: 8, sm: 16, md: 24};

  timeStamps: NzMarks = {
    0: '0 min',
    10: '10 min',
    25: '25 min',
    50: '50 min'
  }

  constructor(private fb: FormBuilder) {
  }

  /*
   * Get time value from slider as **number**
   */
  getTimeValue(value: NzSliderValue) {
    console.log(value);
  }

  ngOnInit(): void {
    this.createTask = this.fb.group({
      taskName: [null],
      time: [null],
      breaksNumber: [null],
      comment: [null]
    });
  }

}
