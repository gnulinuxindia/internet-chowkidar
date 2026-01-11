"use client";

import { useEffect, useState } from "react";
import { useAdminAuth } from "@/components/admin-auth-provider";
import { getSuggestions, resolveSuggestion, getCategories } from "@/lib/admin-api";
import type { SiteSuggestion, Category } from "@/lib/admin-types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default function SuggestionsPage() {
  const { apiKey } = useAdminAuth();
  const [suggestions, setSuggestions] = useState<SiteSuggestion[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [resolving, setResolving] = useState<number | null>(null);
  const [resolveForm, setResolveForm] = useState<{
    id: number;
    reason: string;
    categories: string[];
  } | null>(null);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [suggestionsData, categoriesData] = await Promise.all([
        getSuggestions(),
        getCategories(),
      ]);
      setSuggestions(suggestionsData);
      setCategories(categoriesData);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load data");
    } finally {
      setLoading(false);
    }
  };

  const handleResolve = async (id: number, status: "accepted" | "rejected") => {
    if (!apiKey || !resolveForm) return;

    try {
      setResolving(id);
      await resolveSuggestion(apiKey, id, {
        status,
        resolve_reason: resolveForm.reason,
        categories: status === "accepted" ? resolveForm.categories : undefined,
      });
      await loadData();
      setResolveForm(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to resolve");
    } finally {
      setResolving(null);
    }
  };

  const pendingSuggestions = suggestions.filter((s) => s.status === "pending");
  const resolvedSuggestions = suggestions.filter((s) => s.status !== "pending");

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
        <h1 className="text-2xl font-bold mb-2">Site Suggestions</h1>
        <p className="text-muted-foreground">
          Review and approve site suggestions from users
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

      {/* Pending Suggestions */}
      <div className="mb-8">
        <h2 className="text-lg font-semibold mb-4 flex items-center gap-2">
          <span className="h-2 w-2 rounded-full bg-yellow-500" />
          Pending ({pendingSuggestions.length})
        </h2>

        {pendingSuggestions.length === 0 ? (
          <Card className="bg-card border-border">
            <CardContent className="py-8 text-center text-muted-foreground">
              No pending suggestions
            </CardContent>
          </Card>
        ) : (
          <div className="space-y-4">
            {pendingSuggestions.map((suggestion) => (
              <Card key={suggestion.id} className="bg-card border-border">
                <CardHeader>
                  <div className="flex items-start justify-between">
                    <div>
                      <CardTitle className="font-mono text-lg">
                        {suggestion.domain}
                      </CardTitle>
                      <CardDescription className="mt-1">
                        {suggestion.reason}
                      </CardDescription>
                    </div>
                    <Badge variant="secondary" className="bg-yellow-500/10 text-yellow-500 border-yellow-500/30">
                      Pending
                    </Badge>
                  </div>
                </CardHeader>
                <CardContent>
                  {suggestion.categories && suggestion.categories.length > 0 && (
                    <div className="flex flex-wrap gap-2 mb-4">
                      {suggestion.categories.map((cat) => (
                        <Badge key={cat} variant="outline" className="text-xs">
                          {cat}
                        </Badge>
                      ))}
                    </div>
                  )}

                  {resolveForm?.id === suggestion.id ? (
                    <div className="space-y-4 p-4 rounded-lg bg-secondary/50 border border-border">
                      <div>
                        <label className="text-sm font-medium mb-2 block">
                          Resolution Reason
                        </label>
                        <Input
                          value={resolveForm.reason}
                          onChange={(e) =>
                            setResolveForm({ ...resolveForm, reason: e.target.value })
                          }
                          placeholder="Reason for accepting/rejecting..."
                          className="bg-card"
                        />
                      </div>
                      <div>
                        <label className="text-sm font-medium mb-2 block">
                          Categories (for acceptance)
                        </label>
                        <div className="flex flex-wrap gap-2">
                          {categories.map((cat) => (
                            <button
                              key={cat.id}
                              type="button"
                              onClick={() => {
                                const cats = resolveForm.categories.includes(cat.name)
                                  ? resolveForm.categories.filter((c) => c !== cat.name)
                                  : [...resolveForm.categories, cat.name];
                                setResolveForm({ ...resolveForm, categories: cats });
                              }}
                              className={`px-3 py-1 rounded-md text-sm transition-colors ${
                                resolveForm.categories.includes(cat.name)
                                  ? "bg-primary text-primary-foreground"
                                  : "bg-secondary text-muted-foreground hover:text-foreground"
                              }`}
                            >
                              {cat.name}
                            </button>
                          ))}
                        </div>
                      </div>
                      <div className="flex gap-2">
                        <Button
                          onClick={() => handleResolve(suggestion.id, "accepted")}
                          disabled={resolving === suggestion.id || !resolveForm.reason}
                          className="bg-green-600 hover:bg-green-700"
                        >
                          {resolving === suggestion.id ? "..." : "Accept"}
                        </Button>
                        <Button
                          onClick={() => handleResolve(suggestion.id, "rejected")}
                          disabled={resolving === suggestion.id || !resolveForm.reason}
                          variant="destructive"
                        >
                          {resolving === suggestion.id ? "..." : "Reject"}
                        </Button>
                        <Button
                          onClick={() => setResolveForm(null)}
                          variant="ghost"
                        >
                          Cancel
                        </Button>
                      </div>
                    </div>
                  ) : (
                    <Button
                      onClick={() =>
                        setResolveForm({
                          id: suggestion.id,
                          reason: "",
                          categories: suggestion.categories || [],
                        })
                      }
                      variant="outline"
                      size="sm"
                    >
                      Review
                    </Button>
                  )}

                  <p className="text-xs text-muted-foreground mt-4">
                    Submitted {new Date(suggestion.created_at).toLocaleDateString()}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </div>

      {/* Resolved Suggestions */}
      {resolvedSuggestions.length > 0 && (
        <div>
          <h2 className="text-lg font-semibold mb-4">
            Resolved ({resolvedSuggestions.length})
          </h2>
          <div className="space-y-2">
            {resolvedSuggestions.slice(0, 10).map((suggestion) => (
              <div
                key={suggestion.id}
                className="flex items-center justify-between p-4 rounded-lg bg-card border border-border"
              >
                <div>
                  <span className="font-mono text-sm">{suggestion.domain}</span>
                  <p className="text-xs text-muted-foreground mt-1">
                    {suggestion.resolve_reason}
                  </p>
                </div>
                <Badge
                  variant="secondary"
                  className={
                    suggestion.status === "accepted"
                      ? "bg-green-500/10 text-green-500 border-green-500/30"
                      : "bg-red-500/10 text-red-500 border-red-500/30"
                  }
                >
                  {suggestion.status}
                </Badge>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
