import { config } from "./config";
import type { Category, Site, SiteSuggestion, SiteInput, ResolveSuggestionInput } from "./admin-types";

const API_URL = config.apiUrl;

// Categories
export async function getCategories(): Promise<Category[]> {
  const res = await fetch(`${API_URL}/categories`);
  if (!res.ok) throw new Error("Failed to fetch categories");
  return res.json();
}

export async function createCategory(apiKey: string, name: string): Promise<Category> {
  const res = await fetch(`${API_URL}/categories`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Api-Key": apiKey,
    },
    body: JSON.stringify({ name }),
  });
  if (!res.ok) {
    const error = await res.text();
    throw new Error(error || "Failed to create category");
  }
  return res.json();
}

export async function deleteCategory(apiKey: string, id: number): Promise<void> {
  const res = await fetch(`${API_URL}/categories/${id}`, {
    method: "DELETE",
    headers: {
      "X-Api-Key": apiKey,
    },
  });
  if (!res.ok) {
    if (res.status === 404) {
      throw new Error("Category not found");
    }
    throw new Error("Failed to delete category");
  }
}

// Sites
export async function getSites(limit = 100): Promise<Site[]> {
  const res = await fetch(`${API_URL}/sites?limit=${limit}`);
  if (!res.ok) throw new Error("Failed to fetch sites");
  return res.json();
}

export async function deleteSite(apiKey: string, id: number): Promise<void> {
  const res = await fetch(`${API_URL}/sites/${id}`, {
    method: "DELETE",
    headers: {
      "X-Api-Key": apiKey,
    },
  });
  if (!res.ok) {
    if (res.status === 404) {
      throw new Error("Site not found");
    }
    throw new Error("Failed to delete site");
  }
}

// Site Suggestions
export async function getSuggestions(): Promise<SiteSuggestion[]> {
  const res = await fetch(`${API_URL}/sites/suggestions`);
  if (!res.ok) throw new Error("Failed to fetch suggestions");
  return res.json();
}

export async function getSuggestion(id: number): Promise<SiteSuggestion> {
  const res = await fetch(`${API_URL}/sites/suggestions/${id}`);
  if (!res.ok) throw new Error("Failed to fetch suggestion");
  return res.json();
}

export async function resolveSuggestion(
  apiKey: string,
  id: number,
  input: ResolveSuggestionInput
): Promise<SiteSuggestion> {
  const res = await fetch(`${API_URL}/sites/suggestions/${id}/resolve`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Api-Key": apiKey,
    },
    body: JSON.stringify(input),
  });
  if (!res.ok) {
    const error = await res.text();
    throw new Error(error || "Failed to resolve suggestion");
  }
  return res.json();
}

// Sites
export async function createSite(apiKey: string, input: SiteInput): Promise<unknown> {
  const res = await fetch(`${API_URL}/sites`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Api-Key": apiKey,
    },
    body: JSON.stringify(input),
  });
  if (!res.ok) {
    const error = await res.text();
    throw new Error(error || "Failed to create site");
  }
  return res.json();
}

export async function createSuggestion(input: {
  domain: string;
  reason: string;
  ping_url?: string;
  categories?: string[];
}): Promise<SiteSuggestion> {
  const res = await fetch(`${API_URL}/sites/suggestions`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(input),
  });
  if (!res.ok) {
    const error = await res.text();
    throw new Error(error || "Failed to create suggestion");
  }
  return res.json();
}
