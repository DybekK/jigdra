import {FormGroup} from "@angular/forms";

export type ValidateStatus = (controlName: string) => string;

export default class StatusValidator {
  static validateStatus(formGroup: FormGroup): ValidateStatus {
    return (controlName: string): string => {
      const control = formGroup.get(controlName);
      if(!control?.dirty) return "";
      return control?.errors ? "error" : "success";
    }
  }
}
