import {Address} from './address'
export interface Post {
    name: string;
    address: Address;
    startDate: Date;
    endDate: Date;
    rent: number;
    utilities: number;
    parkingCost: number;
    squareFootage: number;
    furnished: boolean;
    parking: boolean;
    beds: number;
    baths: number;
    amenities: string;
    appliances: string;
    description: string;
}
