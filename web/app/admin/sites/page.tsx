"use client";

import { useEffect, useState } from "react";
import { useAdminAuth } from "@/components/admin-auth-provider";
import { createSite, getCategories } from "@/lib/admin-api";
import type { Category } from "@/lib/admin-types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";

export default function AddSitePage() {
  const { apiKey } = useAdminAuth();
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const [form, setForm] = useState({
    domain: "",
    ping_url: "",
    categories: [] as string[],
  });

  useEffect(() => {
    loadCategories();
  }, []);

  const loadCategories = async () => {
    try {
      const data = await getCategories();
      setCategories(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load categories");
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
        <h1 className="text-2xl font-bold mb-2">Add Site</h1>
        <p className="text-muted-foreground">
          Manually add a new site to track for blocking
        </p>
      </div>

      <Card className="max-w-xl bg-card border-border">
        <CardHeader>
          <CardTitle>New Site</CardTitle>
          <CardDescription>
            Enter the site details below
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            {error && (
              <div className="p-3 rounded-lg bg-destructive/10 border border-destructive/30 text-destructive text-sm">
                {error}
              </div>
            )}

            {success && (
              <div className="p-3 rounded-lg bg-green-500/10 border border-green-500/30 text-green-500 text-sm">
                {success}
              </div>
            )}

            <div>
              <label className="text-sm font-medium mb-2 block">
                Domain <span className="text-destructive">*</span>
              </label>
              <Input
                value={form.domain}
                onChange={(e) => setForm({ ...form, domain: e.target.value })}
                placeholder="example.com"
                className="bg-secondary"
              />
              <p className="text-xs text-muted-foreground mt-1">
                Enter the domain without http:// or https://
              </p>
            </div>

            <div>
              <label className="text-sm font-medium mb-2 block">
                Ping URL (optional)
              </label>
              <Input
                value={form.ping_url}
                onChange={(e) => setForm({ ...form, ping_url: e.target.value })}
                placeholder="https://example.com/health"
                className="bg-secondary"
              />
              <p className="text-xs text-muted-foreground mt-1">
                URL to check if the site is accessible
              </p>
            </div>

            <div>
              <label className="text-sm font-medium mb-2 block">
                Categories <span className="text-destructive">*</span>
              </label>
              <div className="flex flex-wrap gap-2">
                {categories.map((cat) => (
                  <button
                    key={cat.id}
                    type="button"
                    onClick={() => toggleCategory(cat.name)}
                    className={`px-3 py-1.5 rounded-md text-sm transition-colors ${
                      form.categories.includes(cat.name)
                        ? "bg-primary text-primary-foreground"
                        : "bg-secondary text-muted-foreground hover:text-foreground"
                    }`}
                  >
                    {cat.name}
                  </button>
                ))}
              </div>
              {form.categories.length > 0 && (
                <p className="text-xs text-muted-foreground mt-2">
                  Selected: {form.categories.join(", ")}
                </p>
              )}
            </div>

            <Button type="submit" disabled={submitting} className="w-full">
              {submitting ? "Creating..." : "Create Site"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
