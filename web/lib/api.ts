import { config } from "./config";
import type { Site, ISP } from "./types";

const API_URL = config.apiUrl;

function listUrl(path: string, limit = 0): string {
  const params = new URLSearchParams({ limit: String(limit) });
  return API_URL + path + "?" + params.toString();
}

export async function getSites(limit = 0): Promise<Site[]> {
  const res = await fetch(listUrl("/sites", limit), {
    next: { revalidate: 60 },
  });
  if (!res.ok) throw new Error("Failed to fetch sites");
  return res.json();
}

export async function getSite(id: string): Promise<Site> {
  const res = await fetch(`${API_URL}/sites/${id}`, {
    next: { revalidate: 60 },
  });
  if (!res.ok) throw new Error("Failed to fetch site");
  return res.json();
}

export async function getISPs(limit = 0): Promise<ISP[]> {
  const res = await fetch(listUrl("/isps", limit), {
    next: { revalidate: 60 },
  });
  if (!res.ok) throw new Error("Failed to fetch ISPs");
  return res.json();
}

export async function getISP(id: string): Promise<ISP> {
  const res = await fetch(`${API_URL}/isps/${id}`, {
    next: { revalidate: 60 },
  });
  if (!res.ok) throw new Error("Failed to fetch ISP");
  return res.json();
}
