import { ColDef} from "ag-grid-community";

export interface CustomColumnDef extends ColDef {
    dataType : string
    filterable : boolean
}