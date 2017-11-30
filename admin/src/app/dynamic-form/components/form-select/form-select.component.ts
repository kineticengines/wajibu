import { Component } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-select',
  styleUrls: ['../../dynamic-form.css'],
  template: `
    <div class="dynamic-field form-select" [formGroup]="group">
      <md-select class="general-form-input"  [formControlName]="config.name" [placeholder]="config.placeholder | titlecase">
            <md-option *ngFor="let option of config.options" [value]="option"  class="the-input">{{option | titlecase}}</md-option>
      </md-select><br><br>
    </div>
  `
})
export class FormSelectComponent implements Field {
  config: FieldConfig;
  group: FormGroup;  
}
