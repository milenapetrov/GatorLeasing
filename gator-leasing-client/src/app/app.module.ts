import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { AgGridModule } from 'ag-grid-angular';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { PostComponent } from './post/post.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatRadioModule } from '@angular/material/radio';
import { MatDialog } from '@angular/material/dialog';

import { MatGridListModule } from '@angular/material/grid-list';
import { MatCardModule } from '@angular/material/card';
import { MatMenuModule } from '@angular/material/menu';
import { LayoutModule } from '@angular/cdk/layout';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';

import { HomeComponent } from './home/home.component';

import { ComponentsModule } from './components/components.module';
import { AuthModule } from '@auth0/auth0-angular';
import { environment as env } from 'src/environments/environment';
import { MyLeasesComponent } from './my-leases/my-leases.component';
import { NavigationComponent } from './navigation/navigation.component';
import { LoginComponent } from './components/login/login.component';
import { ProfileComponent } from './profile/profile.component';
import { MessagesComponent } from './messages/messages.component';
import { CreateComponent } from './post/create/create.component';
import { UpdateComponent } from './post/update/update.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { LeaseListingsComponent } from './components/lease-listings/lease-listings.component';
import { ViewComponent } from './components/view/view.component';
import { GridCellComponent } from 'src/app/components/grid-cell/grid-cell.component';
import { DisplayLeasesComponent } from './components/display-leases/display-leases.component';

@NgModule({
  declarations: [
    AppComponent,
    PostComponent,
    HomeComponent,
    MyLeasesComponent,
    NavigationComponent,
    ProfileComponent,
    MessagesComponent,
    CreateComponent,
    UpdateComponent,
    LeaseListingsComponent,
    ViewComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot([
      { path: 'home', component: HomeComponent },
      { path: 'post', component: PostComponent },
      { path: 'login', component: LoginComponent },
      { path: 'create', component: CreateComponent },
      { path: 'update', component: UpdateComponent },
      { path: 'view', component: ViewComponent},
      { path: 'display', component: DisplayLeasesComponent}
    ]),
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
    MatSlideToggleModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatRadioModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    LayoutModule,
    MatToolbarModule,
    MatSidenavModule,
    MatCheckboxModule,
    MatListModule,
    ComponentsModule,
    MatDatepickerModule,
    MatNativeDateModule,
    AuthModule.forRoot({
      ...env.auth,
      httpInterceptor: {
        ...env.httpInterceptor,
      },
    }),
    AgGridModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
