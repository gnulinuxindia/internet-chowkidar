import { config } from "./config";
import type { Category, SiteSuggestion, SiteInput, ResolveSuggestionInput } from "./admin-types";

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
