import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { DisplayLeasesComponent } from './display-leases/display-leases.component';
import { AgGridModule } from 'ag-grid-angular';
import { GridCellComponent } from './grid-cell/grid-cell.component';
import { LeaseListingsComponent } from './lease-listings/lease-listings.component';


@NgModule({
  declarations: [LoginComponent, DisplayLeasesComponent, GridCellComponent],
  imports: [CommonModule, AgGridModule],
  exports: [LoginComponent, DisplayLeasesComponent],
  providers: [GridCellComponent, DisplayLeasesComponent, LeaseListingsComponent]
})
export class ComponentsModule {}
