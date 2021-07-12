import {Directive, ElementRef, Renderer2} from '@angular/core';
import {environment} from '../environments/environment';

@Directive({
  // tslint:disable-next-line:directive-selector
  selector: `[data-angular]`
})

export class DataCypressDirective {
  constructor(private el: ElementRef, private renderer: Renderer2) {
    if (environment.production) {
      renderer.removeAttribute(el.nativeElement, 'data-cy');
    }
  }
}
