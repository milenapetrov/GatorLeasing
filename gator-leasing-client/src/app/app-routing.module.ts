import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthHttpInterceptor } from '@auth0/auth0-angular';
import { DisplayLeasesComponent } from './components/display-leases/display-leases.component';
import { HomeComponent } from './home/home.component';
//import { MyLeasesComponent } from './my-leases/my-leases.component';
import { AuthHttpInterceptorExtendedService } from './services/auth-http-interceptor-extended.service';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'display-leases',
    component: DisplayLeasesComponent,
  }
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
export class AppRoutingModule { }
