"use client";

import { useEffect, useState } from "react";
import { useAdminAuth } from "@/components/admin-auth-provider";
import { getCategories, createCategory, deleteCategory } from "@/lib/admin-api";
import type { Category } from "@/lib/admin-types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function CategoriesPage() {
  const { apiKey } = useAdminAuth();
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [deletingId, setDeletingId] = useState<number | null>(null);
  const [error, setError] = useState("");
  const [newCategory, setNewCategory] = useState("");

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

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!apiKey || !newCategory.trim()) return;

    try {
      setSubmitting(true);
      setError("");
      await createCategory(apiKey, newCategory.trim());
      setNewCategory("");
      await loadCategories();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create category");
    } finally {
      setSubmitting(false);
    }
  };

  const handleDelete = async (id: number) => {
    if (!apiKey) return;
    if (!confirm("Are you sure you want to delete this category?")) return;

    try {
      setDeletingId(id);
      setError("");
      await deleteCategory(apiKey, id);
      await loadCategories();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to delete category");
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
        <h1 className="text-2xl font-bold mb-2">Categories</h1>
        <p className="text-muted-foreground">
          Manage site categories for organization
        </p>
      </div>

      {error && (
        <div className="mb-6 p-4 rounded-lg bg-destructive/10 border border-destructive/30 text-destructive">
          {error}
          <button onClick={() => setError("")} className="ml-2 underline">
            Dismiss
          </button>
        </div>
      )}

      <div className="grid gap-6 lg:grid-cols-2">
        {/* Add Category */}
        <Card className="bg-card border-border">
          <CardHeader>
            <CardTitle>Add Category</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleCreate} className="flex gap-3">
              <Input
                value={newCategory}
                onChange={(e) => setNewCategory(e.target.value)}
                placeholder="Category name"
                className="bg-secondary"
              />
              <Button type="submit" disabled={submitting || !newCategory.trim()}>
                {submitting ? "..." : "Add"}
              </Button>
            </form>
          </CardContent>
        </Card>

        {/* Categories List */}
        <Card className="bg-card border-border">
          <CardHeader>
            <CardTitle>All Categories ({categories.length})</CardTitle>
          </CardHeader>
          <CardContent>
            {categories.length === 0 ? (
              <p className="text-muted-foreground text-sm">No categories yet</p>
            ) : (
              <div className="flex flex-wrap gap-2">
                {categories.map((cat) => (
                  <div
                    key={cat.id}
                    className="px-3 py-1.5 rounded-md bg-secondary text-sm flex items-center gap-2 group"
                  >
                    <svg
                      viewBox="0 0 24 24"
                      className="h-3.5 w-3.5 text-primary"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                    >
                      <path d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                    </svg>
                    <span>{cat.name}</span>
                    <span className="text-xs text-muted-foreground">#{cat.id}</span>
                    <button
                      onClick={() => handleDelete(cat.id)}
                      disabled={deletingId === cat.id}
                      className="ml-1 p-0.5 rounded hover:bg-destructive/20 text-muted-foreground hover:text-destructive transition-colors disabled:opacity-50"
                      title="Delete category"
                    >
                      {deletingId === cat.id ? (
                        <div className="h-3 w-3 rounded-full border border-current border-t-transparent animate-spin" />
                      ) : (
                        <svg viewBox="0 0 24 24" className="h-3 w-3" fill="none" stroke="currentColor" strokeWidth="2">
                          <path d="M18 6L6 18M6 6l12 12" />
                        </svg>
                      )}
                    </button>
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
