import { Component, ViewContainerRef } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-input',
  styleUrls: ['../../dynamic-form.css'],
  template: `   
    <div [formGroup]="group">
      <!--<label>{{ config.label }}</label>-->
      <md-input-container class="general-form-input">
        <input mdInput [placeholder]="config.placeholder | titlecase" [formControlName]="config.name" type="text">
      </md-input-container>
    </div>
  `
})
export class FormInputComponent implements Field {
  config: FieldConfig;
  group: FormGroup;
}
