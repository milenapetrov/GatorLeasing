import { Component } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular'

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {
  constructor(public auth: AuthService) {}
}
