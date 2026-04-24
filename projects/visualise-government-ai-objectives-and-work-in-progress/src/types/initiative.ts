export type InitiativeStatus = "objective" | "in-progress" | "completed" | "paused";

export type Department =
  | "DSIT"
  | "Cabinet Office"
  | "DHSC"
  | "DfE"
  | "Home Office"
  | "MoD"
  | "HMRC"
  | "DWP"
  | "MHCLG"
  | "FCDO"
  | "GDS"
  | "i.AI"
  | "Other";

export type Category =
  | "Healthcare"
  | "Education"
  | "Defence & Security"
  | "Public Services"
  | "Economic Growth"
  | "Infrastructure"
  | "Research & Innovation"
  | "Regulation & Safety"
  | "Digital Government"
  | "Other";

export type SourceType = "announcement" | "parliamentary-statement" | "blog" | "news" | "policy-document";

export interface Source {
  url: string;
  title: string;
  date: string;
  type: SourceType;
}

export interface Initiative {
  id: string;
  title: string;
  description: string;
  status: InitiativeStatus;
  department: Department;
  category: Category;
  tags: string[];
  announcedDate: string;
  lastUpdated: string;
  sources: Source[];
  estimatedCompletion?: string;
  budget?: string;
}

export interface InitiativesData {
  version: string;
  lastUpdated: string;
  initiatives: Initiative[];
}

export interface FilterState {
  department: Department | "All";
  category: Category | "All";
  status: InitiativeStatus | "All";
  dateFrom: string;
  dateTo: string;
  search: string;
}
