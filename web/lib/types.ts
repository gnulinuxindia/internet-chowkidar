export interface Site {
  id: number;
  domain: string;
  categories: string[];
  block_reports: number;
  unblock_reports: number;
  last_reported_at: string;
  created_at?: string;
  updated_at?: string;
  blocked_by_isps?: ISP[];
}

export interface ISP {
  id: number;
  name: string;
  latitude: number;
  longitude: number;
  block_reports: number;
  unblock_reports: number;
  last_reported_at?: string;
  created_at?: string;
  updated_at?: string;
  blocks?: BlockedSite[];
}

export interface BlockedSite {
  id: number;
  domain: string;
  site_id: number;
  block_reports: number;
  unblock_reports: number;
  last_reported_at: string;
  created_at: string;
  updated_at: string;
}
