import { Address } from './address';

export interface Contact {
  id: number;
  lastName: string;
  firstName: string;
  salutation: string;
  leaseID: number;
  phoneNumber: string;
  email: string;
  address: Address;
}
