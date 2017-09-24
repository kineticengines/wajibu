import { Component } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { Field } from '../../models/field.interface';
import { FieldConfig } from '../../models/field-config.interface';

@Component({
  selector: 'form-text', 
  styleUrls: ['../../dynamic-form.css'], 
  template: `
    <div [formGroup]="group">      
    </div>
  `
})
export class FormTextComponent implements Field {
  config: FieldConfig;
  group: FormGroup;
  
}
