import { Component } from '@angular/core';
import { Address } from '../models/address';
import { Contact } from '../models/contact';
import { LeaseService } from '../services/lease.service';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  addy: Address = {
    street: '',
    roomNumber: '',
    city: '',
    state: '',
    zipCode: ''
  };

  contact: Contact = {
    id: 0,
    lastName: '',
    firstName: '',
    salutation: '',
    leaseID: 0,
    phoneNumber: '',
    email: '',
    address: this.addy
  };


  constructor(private leaseService:LeaseService, private formBuilder: FormBuilder){
  }

  onSubmit(contact:Contact) {
    //this.leaseService.createContact(this.contact);
  };

}
