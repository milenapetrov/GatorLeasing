import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthHttpInterceptor } from '@auth0/auth0-angular';
import { HomeComponent } from './home/home.component';
import { MyLeasesComponent } from './my-leases/my-leases.component';
import { LeaseListingsComponent } from './modules/lease-listings/lease-listings.component';
import { AuthHttpInterceptorExtendedService } from './services/auth-http-interceptor-extended.service';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'my-leases',
    component: MyLeasesComponent,
  },
  {
    path: 'listings',
    component: LeaseListingsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthHttpInterceptorExtendedService,
      multi: true,
    },
  ],
})
export class AppRoutingModule {}
