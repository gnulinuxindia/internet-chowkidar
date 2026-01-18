"use client";

import { useEffect, useState } from "react";
import { useAdminAuth } from "@/components/admin-auth-provider";
import { createSite, getCategories, getSites, deleteSite } from "@/lib/admin-api";
import type { Category, Site } from "@/lib/admin-types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";

export default function SitesPage() {
  const { apiKey } = useAdminAuth();
  const [categories, setCategories] = useState<Category[]>([]);
  const [sites, setSites] = useState<Site[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [deletingId, setDeletingId] = useState<number | null>(null);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const [form, setForm] = useState({
    domain: "",
    ping_url: "",
    categories: [] as string[],
  });

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const [categoriesData, sitesData] = await Promise.all([
        getCategories(),
        getSites(500),
      ]);
      setCategories(categoriesData);
      setSites(sitesData);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load data");
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!apiKey) return;

    if (!form.domain.trim()) {
      setError("Domain is required");
      return;
    }

    if (form.categories.length === 0) {
      setError("Select at least one category");
      return;
    }

    try {
      setSubmitting(true);
      setError("");
      setSuccess("");

      await createSite(apiKey, {
        domain: form.domain.trim(),
        ping_url: form.ping_url.trim() || undefined,
        categories: form.categories,
      });

      setSuccess(`Site "${form.domain}" created successfully!`);
      setForm({ domain: "", ping_url: "", categories: [] });
      await loadData();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create site");
    } finally {
      setSubmitting(false);
    }
  };

  const toggleCategory = (name: string) => {
    setForm((prev) => ({
      ...prev,
      categories: prev.categories.includes(name)
        ? prev.categories.filter((c) => c !== name)
        : [...prev.categories, name],
    }));
  };

  const handleDelete = async (id: number, domain: string) => {
    if (!apiKey) return;
    if (!confirm(`Are you sure you want to delete "${domain}"?`)) return;

    try {
      setDeletingId(id);
      setError("");
      await deleteSite(apiKey, id);
      await loadData();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to delete site");
    } finally {
      setDeletingId(null);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-12">
        <div className="h-8 w-8 rounded-full border-2 border-primary/30 border-t-primary animate-spin" />
      </div>
    );
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold mb-2">Sites</h1>
        <p className="text-muted-foreground">
          Manage tracked sites and their categories
        </p>
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Add Site Form */}
        <Card className="bg-card border-border lg:col-span-1">
          <CardHeader>
            <CardTitle>Add Site</CardTitle>
            <CardDescription>
              Add a new site to track
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              {error && (
                <div className="p-3 rounded-lg bg-destructive/10 border border-destructive/30 text-destructive text-sm">
                  {error}
                  <button onClick={() => setError("")} className="ml-2 underline">
                    Dismiss
                  </button>
                </div>
              )}

              {success && (
                <div className="p-3 rounded-lg bg-green-500/10 border border-green-500/30 text-green-500 text-sm">
                  {success}
                </div>
              )}

              <div>
                <label className="text-sm font-medium mb-1.5 block">
                  Domain <span className="text-destructive">*</span>
                </label>
                <Input
                  value={form.domain}
                  onChange={(e) => setForm({ ...form, domain: e.target.value })}
                  placeholder="example.com"
                  className="bg-secondary"
                />
              </div>

              <div>
                <label className="text-sm font-medium mb-1.5 block">
                  Ping URL
                </label>
                <Input
                  value={form.ping_url}
                  onChange={(e) => setForm({ ...form, ping_url: e.target.value })}
                  placeholder="https://example.com/health"
                  className="bg-secondary"
                />
              </div>

              <div>
                <label className="text-sm font-medium mb-1.5 block">
                  Categories <span className="text-destructive">*</span>
                </label>
                <div className="flex flex-wrap gap-1.5">
                  {categories.map((cat) => (
                    <button
                      key={cat.id}
                      type="button"
                      onClick={() => toggleCategory(cat.name)}
                      className={`px-2 py-1 rounded text-xs transition-colors ${
                        form.categories.includes(cat.name)
                          ? "bg-primary text-primary-foreground"
                          : "bg-secondary text-muted-foreground hover:text-foreground"
                      }`}
                    >
                      {cat.name}
                    </button>
                  ))}
                </div>
              </div>

              <Button type="submit" disabled={submitting} className="w-full">
                {submitting ? "Creating..." : "Create Site"}
              </Button>
            </form>
          </CardContent>
        </Card>

        {/* Sites List */}
        <Card className="bg-card border-border lg:col-span-2">
          <CardHeader>
            <CardTitle>All Sites ({sites.length})</CardTitle>
            <CardDescription>
              Sites being tracked for blocking status
            </CardDescription>
          </CardHeader>
          <CardContent>
            {sites.length === 0 ? (
              <p className="text-muted-foreground text-sm">No sites yet</p>
            ) : (
              <div className="space-y-2 max-h-[600px] overflow-y-auto pr-2">
                {sites.map((site) => (
                  <div
                    key={site.id}
                    className="p-3 rounded-lg bg-secondary/50 border border-border hover:border-primary/30 transition-colors"
                  >
                    <div className="flex items-start justify-between gap-3">
                      <div className="min-w-0 flex-1">
                        <div className="flex items-center gap-2">
                          <span className="font-mono text-sm font-medium truncate">
                            {site.domain}
                          </span>
                          <span className="text-xs text-muted-foreground">
                            #{site.id}
                          </span>
                        </div>
                        <div className="flex flex-wrap gap-1 mt-1.5">
                          {site.categories.map((cat) => (
                            <span
                              key={cat}
                              className="px-1.5 py-0.5 rounded bg-primary/10 text-primary text-xs"
                            >
                              {cat}
                            </span>
                          ))}
                        </div>
                      </div>
                      <div className="flex items-center gap-3 shrink-0">
                        <div className="text-right text-xs text-muted-foreground">
                          <div className="flex items-center gap-1">
                            <svg viewBox="0 0 24 24" className="h-3 w-3 text-red-500" fill="none" stroke="currentColor" strokeWidth="2">
                              <path d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
                            </svg>
                            <span>{site.block_reports}</span>
                          </div>
                          <div className="flex items-center gap-1 mt-0.5">
                            <svg viewBox="0 0 24 24" className="h-3 w-3 text-green-500" fill="none" stroke="currentColor" strokeWidth="2">
                              <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                            <span>{site.unblock_reports}</span>
                          </div>
                        </div>
                        <button
                          onClick={() => handleDelete(site.id, site.domain)}
                          disabled={deletingId === site.id}
                          className="p-1.5 rounded hover:bg-destructive/20 text-muted-foreground hover:text-destructive transition-colors disabled:opacity-50"
                          title="Delete site"
                        >
                          {deletingId === site.id ? (
                            <div className="h-4 w-4 rounded-full border-2 border-current border-t-transparent animate-spin" />
                          ) : (
                            <svg viewBox="0 0 24 24" className="h-4 w-4" fill="none" stroke="currentColor" strokeWidth="2">
                              <path d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                            </svg>
                          )}
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
