import { Component, ViewContainerRef } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-text-are',
  styleUrls: ['../../dynamic-form.css'],
  template: `   
    <div [formGroup]="group">
      <!--<label>{{ config.label }}</label>-->
      <md-input-container class="general-form-input">
        <textarea mdInput [placeholder]="config.placeholder | titlecase" [formControlName]="config.name" type="text"></textarea>
      </md-input-container>
    </div>
  `
})
export class FormTextAreaComponent implements Field {
  config: FieldConfig;
  group: FormGroup;
}