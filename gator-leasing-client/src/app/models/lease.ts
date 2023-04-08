import { NumberMatcher } from "cypress/types/net-stubbing";
import {Address} from './address'

export interface Lease{
    id: number,
    name: string,
    createdAt: Date,
    ownerID: number,
    address: Address,
    startDate: Date,
    endDate: Date,
    rent: number,
    utilites: number,
    parkingCost: number,
    totalCost: number,
    squareFootage: number,
    furnished: boolean,
    parking: boolean,
    beds: number,
    baths: number,
    amenities: string,
    appliances: string,
    description: string,
    contacts: object,
}