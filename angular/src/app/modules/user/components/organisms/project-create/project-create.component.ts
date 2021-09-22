import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {NzMessageService} from "ng-zorro-antd/message";
import {NzUploadFile} from "ng-zorro-antd/upload";
import {BehaviorSubject, Observable, Observer, of} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {catchError, debounceTime, map, switchMap} from "rxjs/operators";
import {hexColorRegExp} from "../../../../../shared/regexps/regexps";


@Component({
  selector: 'app-project-create',
  template: `
    <form [formGroup]="validateForm" class="ant-advanced-search-form" (ngSubmit)="submitTask()">

      <div nz-row nzJustify="start" nzAlign="middle" [nzGutter]="[vGutter, hGutter]">
        <div nz-col nzFlex="100px">
          <nz-form-label>
            <span>Set picture</span>
          </nz-form-label>
          <nz-form-item>
            <nz-upload
              class="avatar-uploader"
              nzAction="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              nzName="avatar"
              nzListType="picture-card"
              [nzShowUploadList]="false"
              [nzBeforeUpload]="beforeUpload"
              (nzChange)="handleChange($event)"
            >
              <ng-container *ngIf="!avatarUrl">
                <i class="upload-icon" nz-icon [nzType]="loading ? 'loading' : 'plus'"></i>
                <div class="ant-upload-text">Upload</div>
              </ng-container>
              <img *ngIf="avatarUrl" [src]="avatarUrl" style="width: 100%"/>
            </nz-upload>
          </nz-form-item>
        </div>

        <div nz-col nzFlex="auto">

          <div nz-row nzJustify="center" nzAlign="middle" [nzGutter]="[vGutter, hGutter]">
            <div nz-col [nzSpan]="24">
              <nz-form-item>
                <nz-form-control nzErrorTip="You need to give project name">
                  <nz-input-group nzAddOnBefore="Project">
                    <input formControlName="projectName" type="text" nz-input placeholder="Project name"/>
                  </nz-input-group>
                </nz-form-control>
              </nz-form-item>
            </div>
          </div>

          <div nz-row nzJustify="center" nzAlign="middle" [nzGutter]="[vGutter, hGutter]">
            <div nz-col [nzSpan]="24">
              <nz-form-item>
                <nz-form-label>
                  <span>Select color</span>
                </nz-form-label>

                <button nz-button nzShape="circle" class="color-sample" style="background-color: rgba(255,0,0,0.5)">R
                </button>
                <button nz-button nzShape="circle" class="color-sample" style="background-color: rgba(0,255,0,0.5)">G
                </button>
                <button nz-button nzShape="circle" class="color-sample" style="background-color: rgba(0,0,255,0.5)">B
                </button>
                <button nz-button nzShape="circle" class="color-sample" style="background-color: rgba(255,255,0,0.5)">Y
                </button>
                <button nz-button id="personalColor" nzShape="circle" class="color-sample"
                        [ngStyle]="{'background-color': personalColor}">
                  P
                </button>
                <nz-form-control nzErrorTip="Incorrect hex number">
                  <nz-input-group style="width: 100px" nzAddOnBefore="Personal">
                    <input nz-input style="width: 100px" formControlName="hexColor" placeholder="#00000000"
                           (ngModelChange)="setPersonalColor($event)">
                  </nz-input-group>
                </nz-form-control>
              </nz-form-item>
            </div>
          </div>

          <!--tag block-->
          <div nz-row nzJustify="center" nzAlign="middle">
            <div nz-col [nzSpan]="24">
              <nz-form-item class="box-holder">
                <nz-tag *ngFor="let tag of tags"
                        nzMode="checkable"
                        [nzChecked]="tag.checked"
                        (nzCheckedChange)="checkChange($event)">
                  {{tag.tagName}}
                </nz-tag>
              </nz-form-item>
            </div>
          </div>
          <!--tag block-->

          <div nz-row nzJustify="center" nzAlign="middle" [nzGutter]="[vGutter, hGutter]">
            <div nz-col [nzSpan]="24">
              <nz-form-item>
                <nz-select
                  nzMode="multiple"
                  nzPlaceHolder="Select users"
                  nzAllowClear
                  nzShowSearch
                  nzServerSearch
                  formControlName="selectedUser"
                  (nzOnSearch)="onSearch($event)"
                >
                  <ng-container *ngFor="let o of optionList">
                    <nz-option *ngIf="!isLoading" [nzValue]="o" [nzLabel]="o"></nz-option>
                  </ng-container>
                  <nz-option *ngIf="isLoading" nzDisabled nzCustomContent>
                    <i nz-icon nzType="loading" class="loading-icon"></i>
                    Loading Data...
                  </nz-option>
                </nz-select>
              </nz-form-item>
            </div>
          </div>

        </div>
      </div>

    </form>
  `,
  styleUrls: ['./project-create.component.scss'],
  providers: [NzMessageService]
})
export class ProjectCreateComponent implements OnInit {

  validateForm!: FormGroup;
  dropDownStyle: { [key: string]: string; } | null = {
    'background-color': '#ff91af'
  };
  hGutter = {xs: 8, sm: 16, md: 24};
  vGutter = {xs: 8, sm: 16, md: 24};

  loading = false;
  avatarUrl?: string;

  listOfOption: Array<{ label: string; value: string }> = [];

  randomUserUrl = 'https://api.randomuser.me/?results=5';
  searchChange$ = new BehaviorSubject('');
  optionList: string[] = [];
  selectedUser?: string;
  isLoading = false;
  personalColor: string;

  checkChange(e: boolean): void {
    console.log(e);
  }

  tags = [
    {
      nzColor: 'magenta',
      tagName: 'Work',
      checked: false
    },
    {
      nzColor: 'gold',
      tagName: 'Bonus',
      checked: false
    },
    {
      nzColor: 'red',
      tagName: 'Deadline',
      checked: false
    }
  ]

  onSearch(value: string): void {
    this.isLoading = true;
    this.searchChange$.next(value);
  }

  constructor(private fb: FormBuilder, private msg: NzMessageService, private http: HttpClient) {
    this.personalColor = "#00000000"
  }

  ngOnInit(): void {
    /* eslint-disable @typescript-eslint/no-explicit-any */
    //todo: make selection hide for the user, as it is picked
    /*code for simulating user selection for project*/
    const getRandomNameList = (name: string) =>
      this.http
        .get(`${this.randomUserUrl}`)
        .pipe(
          catchError(() => of({results: []})),
          map((res: any) => res.results)
        )
        .pipe(map((list: any) => list.map((item: any) => `${item.name.first} ${name}`)));
    const optionList$: Observable<string[]> = this.searchChange$
      .asObservable()
      .pipe(debounceTime(500))
      .pipe(switchMap(getRandomNameList));
    optionList$.subscribe(data => {
      this.optionList = data;
      this.isLoading = false;
    });

    /*=============================================================================*/
    /*code for tags*/
    const children: Array<{ label: string; value: string }> = [];

    for (let i = 10; i < 36; i++) {
      children.push({label: i.toString(36) + i, value: i.toString(36) + i});
    }

    this.listOfOption = children;
    /*=============================================================================*/

    /*validation for form*/
    this.validateForm = this.fb.group({
      projectName: [null, [Validators.required]],
      tagName: [null],
      selectedUser: [null],
      hexColor: [null, [Validators.pattern(hexColorRegExp)]]
    })

  }

  /*Project picture*/
  beforeUpload = (file: NzUploadFile, _fileList: NzUploadFile[]) =>
    new Observable((observer: Observer<boolean>) => {
      const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
      if (!isJpgOrPng) {
        this.msg.error('You can only upload JPG file!');
        observer.complete();
        return;
      }
      const isLt2M = file.size! / 1024 / 1024 < 2;
      if (!isLt2M) {
        this.msg.error('Image must smaller than 2MB!');
        observer.complete();
        return;
      }
      observer.next(isJpgOrPng && isLt2M);
      observer.complete();
    });

  private getBase64(img: File, callback: (img: string) => void): void {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result!.toString()));
    reader.readAsDataURL(img);
  }

  handleChange(info: { file: NzUploadFile }): void {
    switch (info.file.status) {
      case 'uploading':
        this.loading = true;
        break;
      case 'done':
        // Get this url from response in real world.
        this.getBase64(info.file!.originFileObj!, (img: string) => {
          this.loading = false;
          this.avatarUrl = img;
          // this.avatarUrl = 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png'
        });
        break;
      case 'error':
        this.msg.error('Network error');
        this.loading = false;
        break;
    }
  }

  submitTask() {
  }

  setPersonalColor(color: any) {
    if (color.length > 8) {

      this.personalColor = color;
    }

    console.log(color);
  }
}
