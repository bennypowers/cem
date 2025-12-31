type SubmitType = 'submit';
type ResetType = 'reset';
type ButtonType = SubmitType | ResetType | 'button';

export class MyButton extends HTMLElement {
  type: ButtonType = 'submit';
}
