import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { DisplayLeasesComponent } from './display-leases/display-leases.component';
import { AgGridModule } from 'ag-grid-angular';

@NgModule({
  declarations: [
    LoginComponent,
    DisplayLeasesComponent,
  ],
  imports: [
    CommonModule,
    AgGridModule
  ],
  exports: [
    LoginComponent,
    DisplayLeasesComponent
  ]
})
export class ComponentsModule { }
