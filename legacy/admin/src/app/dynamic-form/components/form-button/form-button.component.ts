import { Component } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-button', 
  styleUrls: ['../../dynamic-form.css'], 
  template: `
    <div [formGroup]="group">
      <button [disabled]="config.disabled" type="submit"
       *ngIf="group.valid" md-raised-button class="save-button" >{{ config.label | uppercase}}</button>
    </div>
  `
})
export class FormButtonComponent implements Field {
  config: FieldConfig;
  group: FormGroup;
  
}
