import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoginModalService {
  private _isShow$ = new BehaviorSubject(false);

  constructor() {}

  get isShow$() {
    return this._isShow$.asObservable();
  }

  public onOpen() {
    this._isShow$.next(true);
  }

  public onClose() {
    this._isShow$.next(false);
  }
}
