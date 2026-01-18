export interface Category {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Site {
  id: number;
  domain: string;
  ping_url: string;
  categories: string[];
  block_reports: number;
  unblock_reports: number;
  last_reported_at: string;
  created_at: string;
  updated_at: string;
}

export interface SiteSuggestion {
  id: number;
  domain: string;
  ping_url?: string;
  categories?: string[];
  reason: string;
  status: "pending" | "accepted" | "rejected";
  resolve_reason?: string;
  linked_site?: number;
  resolved_at?: string;
  created_at: string;
  updated_at: string;
}

export interface SiteInput {
  domain: string;
  ping_url?: string;
  categories: string[];
}

export interface ResolveSuggestionInput {
  resolve_reason: string;
  status: "accepted" | "rejected";
  domain?: string;
  ping_url?: string;
  categories?: string[];
}
