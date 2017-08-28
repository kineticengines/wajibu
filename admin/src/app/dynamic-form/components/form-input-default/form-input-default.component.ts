import { Component, ViewContainerRef } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-input',
  styleUrls: ['../../dynamic-form.css'],
  template: `   
    <div [formGroup]="group">      
      <md-input-container class="general-form-input">
        <input mdInput [placeholder]="config.placeholder | titlecase"
         [formControlName]="config.name" type="text" [value]="config.value | titlecase" >
      </md-input-container>
    </div>
  `
})
export class FormInputDefaultComponent implements Field {
  config: FieldConfig;
  group: FormGroup;
}
