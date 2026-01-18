"use client";

import { useState, useEffect } from "react";
import { createSuggestion, getCategories, getSuggestions } from "@/lib/admin-api";
import type { Category, SiteSuggestion } from "@/lib/admin-types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Badge } from "@/components/ui/badge";

export default function SuggestPage() {
  const [domain, setDomain] = useState("");
  const [pingUrl, setPingUrl] = useState("");
  const [reason, setReason] = useState("");
  const [selectedCategories, setSelectedCategories] = useState<string[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [suggestions, setSuggestions] = useState<SiteSuggestion[]>([]);
  const [loading, setLoading] = useState(false);
  const [dataLoading, setDataLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setDataLoading(true);
      const [categoriesData, suggestionsData] = await Promise.all([
        getCategories(),
        getSuggestions(),
      ]);
      setCategories(categoriesData);
      setSuggestions(suggestionsData);
    } catch (err) {
      console.error("Failed to load data:", err);
    } finally {
      setDataLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess(false);

    if (!domain || !reason) {
      setError("Domain and reason are required");
      return;
    }

    // Check for duplicate domain
    const normalizedDomain = domain.toLowerCase().trim();
    const duplicate = suggestions.find(
      (s) => s.domain.toLowerCase() === normalizedDomain
    );

    if (duplicate) {
      setError(
        `This domain has already been suggested and is currently ${duplicate.status}. Please check the existing suggestions below.`
      );
      return;
    }

    try {
      setLoading(true);
      await createSuggestion({
        domain,
        reason,
        ping_url: pingUrl || undefined,
        categories: selectedCategories.length > 0 ? selectedCategories : undefined,
      });
      setSuccess(true);
      setDomain("");
      setPingUrl("");
      setReason("");
      setSelectedCategories([]);
      // Reload suggestions to show the new one
      await loadData();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to submit suggestion");
    } finally {
      setLoading(false);
    }
  };

  const toggleCategory = (categoryName: string) => {
    setSelectedCategories((prev) =>
      prev.includes(categoryName)
        ? prev.filter((c) => c !== categoryName)
        : [...prev, categoryName]
    );
  };

  return (
    <div className="relative min-h-screen">
      {/* Background Effects */}
      <div className="absolute inset-0 overflow-hidden pointer-events-none">
        <div className="absolute top-1/4 right-1/4 w-96 h-96 bg-primary/10 rounded-full blur-3xl animate-pulse-glow" />
        <div className="absolute bottom-1/3 left-1/3 w-80 h-80 bg-accent/10 rounded-full blur-3xl animate-pulse-glow delay-500" />
      </div>

      {/* Grid overlay */}
      <div className="absolute inset-0 grid-bg opacity-30 pointer-events-none" />

      <div className="relative z-10 mx-auto max-w-4xl px-4 sm:px-6 lg:px-8 py-12">
        {/* Header */}
        <div className="text-center mb-12 animate-slide-up opacity-0 delay-100">
          <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary/10 border border-primary/30 text-sm text-primary font-medium mb-6">
            <span className="relative flex h-2 w-2">
              <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-primary opacity-75" />
              <span className="relative inline-flex rounded-full h-2 w-2 bg-primary" />
            </span>
            Help Us Track Censorship
          </div>

          <h1 className="text-4xl md:text-5xl font-extrabold tracking-tight mb-4">
            <span className="text-primary glow-text-cyber">Suggest</span>
            <br />
            <span className="text-foreground">a Blocked Site</span>
          </h1>

          <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
            Found a website being blocked by your ISP? Help us build a comprehensive
            database of internet censorship in India.
          </p>
        </div>

        {/* Form Card */}
        <div className="animate-slide-up opacity-0 delay-200">
          <div className="relative p-8 md:p-10 rounded-2xl bg-card/80 backdrop-blur-sm border border-border cyber-border">
            {/* Success Message */}
            {success && (
              <div className="mb-6 p-4 rounded-lg bg-unblocked/10 border border-unblocked/30 text-unblocked flex items-start gap-3 animate-slide-up">
                <svg
                  viewBox="0 0 24 24"
                  className="h-5 w-5 mt-0.5 flex-shrink-0"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
                  <polyline points="22 4 12 14.01 9 11.01" />
                </svg>
                <div>
                  <p className="font-semibold mb-1">Suggestion submitted successfully!</p>
                  <p className="text-sm opacity-90">
                    Thank you for contributing. We'll review your suggestion soon.
                  </p>
                </div>
              </div>
            )}

            {/* Error Message */}
            {error && (
              <div className="mb-6 p-4 rounded-lg bg-destructive/10 border border-destructive/30 text-destructive flex items-start gap-3 animate-slide-up">
                <svg
                  viewBox="0 0 24 24"
                  className="h-5 w-5 mt-0.5 flex-shrink-0"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <circle cx="12" cy="12" r="10" />
                  <line x1="12" y1="8" x2="12" y2="12" />
                  <line x1="12" y1="16" x2="12.01" y2="16" />
                </svg>
                <div>
                  <p className="font-semibold">{error}</p>
                </div>
              </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-6">
              {/* Domain Field */}
              <div>
                <label className="block text-sm font-semibold mb-2 text-foreground">
                  Website Domain
                  <span className="text-primary ml-1">*</span>
                </label>
                <Input
                  type="text"
                  value={domain}
                  onChange={(e) => setDomain(e.target.value)}
                  placeholder="example.com"
                  className="bg-secondary/50 border-border focus:border-primary/50 focus:ring-primary/20 transition-all duration-200 font-mono"
                  required
                />
                <p className="text-xs text-muted-foreground mt-2">
                  Enter the domain name of the blocked website
                </p>
              </div>

              {/* Ping URL Field */}
              <div>
                <label className="block text-sm font-semibold mb-2 text-foreground">
                  Ping URL
                  <span className="text-muted-foreground ml-1 font-normal">(Optional)</span>
                </label>
                <Input
                  type="url"
                  value={pingUrl}
                  onChange={(e) => setPingUrl(e.target.value)}
                  placeholder="https://example.com/health"
                  className="bg-secondary/50 border-border focus:border-primary/50 focus:ring-primary/20 transition-all duration-200 font-mono"
                />
                <p className="text-xs text-muted-foreground mt-2">
                  A specific URL we can use to check if the site is accessible
                </p>
              </div>

              {/* Categories */}
              <div>
                <label className="block text-sm font-semibold mb-3 text-foreground">
                  Categories
                  <span className="text-muted-foreground ml-1 font-normal">(Optional)</span>
                </label>
                {categories.length > 0 ? (
                  <div className="flex flex-wrap gap-2">
                    {categories.map((cat) => (
                      <button
                        key={cat.id}
                        type="button"
                        onClick={() => toggleCategory(cat.name)}
                        className={`px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 ${
                          selectedCategories.includes(cat.name)
                            ? "bg-primary text-primary-foreground glow-cyber"
                            : "bg-secondary/50 text-muted-foreground hover:text-foreground hover:bg-secondary border border-border"
                        }`}
                      >
                        {cat.name}
                      </button>
                    ))}
                  </div>
                ) : (
                  <p className="text-sm text-muted-foreground">Loading categories...</p>
                )}
                <p className="text-xs text-muted-foreground mt-3">
                  Select all categories that apply to this website
                </p>
              </div>

              {/* Reason Field */}
              <div>
                <label className="block text-sm font-semibold mb-2 text-foreground">
                  Why do you think it's blocked?
                  <span className="text-primary ml-1">*</span>
                </label>
                <Textarea
                  value={reason}
                  onChange={(e) => setReason(e.target.value)}
                  placeholder="Describe why you believe this site is being blocked by your ISP. Include details like error messages, your ISP name, location, etc."
                  className="bg-secondary/50 border-border focus:border-primary/50 focus:ring-primary/20 transition-all duration-200 min-h-[120px] resize-none"
                  required
                />
                <p className="text-xs text-muted-foreground mt-2">
                  Provide as much detail as possible to help us verify
                </p>
              </div>

              {/* Submit Button */}
              <div className="pt-4">
                <Button
                  type="submit"
                  disabled={loading}
                  className="w-full bg-primary hover:bg-primary/90 text-primary-foreground text-base py-6 glow-cyber disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
                >
                  {loading ? (
                    <span className="flex items-center gap-2">
                      <span className="h-4 w-4 rounded-full border-2 border-primary-foreground/30 border-t-primary-foreground animate-spin" />
                      Submitting...
                    </span>
                  ) : (
                    <span className="flex items-center gap-2">
                      <svg
                        viewBox="0 0 24 24"
                        className="h-5 w-5"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                      >
                        <path d="M22 2L11 13" />
                        <path d="M22 2L15 22L11 13L2 9L22 2Z" />
                      </svg>
                      Submit Suggestion
                    </span>
                  )}
                </Button>
              </div>
            </form>
          </div>
        </div>

        {/* Existing Suggestions Section */}
        <div className="mt-12 animate-slide-up opacity-0 delay-300">
          <div className="flex items-center gap-3 mb-6">
            <h2 className="text-2xl font-bold">
              <span className="text-primary">Recent</span> Suggestions
            </h2>
            <Badge variant="secondary" className="bg-primary/10 text-primary border-primary/30">
              {suggestions.length} total
            </Badge>
          </div>

          {dataLoading ? (
            <div className="flex items-center justify-center py-12">
              <div className="h-8 w-8 rounded-full border-2 border-primary/30 border-t-primary animate-spin" />
            </div>
          ) : suggestions.length === 0 ? (
            <div className="p-8 rounded-xl bg-card/50 border border-border text-center">
              <p className="text-muted-foreground">No suggestions yet. Be the first!</p>
            </div>
          ) : (
            <div className="space-y-3">
              {suggestions.slice(0, 20).map((suggestion, idx) => (
                <div
                  key={suggestion.id}
                  className="p-4 md:p-5 rounded-xl bg-card/50 border border-border hover:border-primary/30 transition-all duration-200 animate-slide-up"
                  style={{ animationDelay: `${idx * 30}ms` }}
                >
                  <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-3">
                    <div className="flex-1 min-w-0">
                      <div className="flex items-center gap-3 mb-2">
                        <h3 className="font-mono text-base font-semibold truncate">
                          {suggestion.domain}
                        </h3>
                        <Badge
                          variant="secondary"
                          className={
                            suggestion.status === "pending"
                              ? "bg-yellow-500/10 text-yellow-500 border-yellow-500/30"
                              : suggestion.status === "accepted"
                              ? "bg-green-500/10 text-green-500 border-green-500/30"
                              : "bg-red-500/10 text-red-500 border-red-500/30"
                          }
                        >
                          {suggestion.status}
                        </Badge>
                      </div>
                      <p className="text-sm text-muted-foreground line-clamp-2 mb-2">
                        {suggestion.reason}
                      </p>
                      {suggestion.categories && suggestion.categories.length > 0 && (
                        <div className="flex flex-wrap gap-1.5">
                          {suggestion.categories.map((cat) => (
                            <Badge
                              key={cat}
                              variant="outline"
                              className="text-xs bg-secondary/50"
                            >
                              {cat}
                            </Badge>
                          ))}
                        </div>
                      )}
                    </div>
                    <div className="flex-shrink-0 text-xs text-muted-foreground">
                      {new Date(suggestion.created_at).toLocaleDateString("en-US", {
                        month: "short",
                        day: "numeric",
                        year: "numeric",
                      })}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
