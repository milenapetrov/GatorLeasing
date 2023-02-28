import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { DisplayLeasesComponent } from './display-leases/display-leases.component';



@NgModule({
  declarations: [
    LoginComponent,
    DisplayLeasesComponent,
  ],
  imports: [
    CommonModule
  ],
  exports: [
    LoginComponent,
    DisplayLeasesComponent
  ]
})
export class ComponentsModule { }
