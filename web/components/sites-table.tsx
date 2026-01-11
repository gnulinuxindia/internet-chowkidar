"use client";

import Link from "next/link";
import { useState, useMemo } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import type { Site } from "@/lib/types";

interface SitesTableProps {
  sites: Site[];
}

export function SitesTable({ sites }: SitesTableProps) {
  const [search, setSearch] = useState("");
  const [sortBy, setSortBy] = useState<"domain" | "blocks" | "date">("blocks");
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("desc");

  const filteredAndSortedSites = useMemo(() => {
    let filtered = sites.filter(
      (site) =>
        site.domain.toLowerCase().includes(search.toLowerCase()) ||
        site.categories?.some((cat) =>
          cat.toLowerCase().includes(search.toLowerCase())
        )
    );

    filtered.sort((a, b) => {
      let comparison = 0;
      switch (sortBy) {
        case "domain":
          comparison = a.domain.localeCompare(b.domain);
          break;
        case "blocks":
          comparison = a.block_reports - b.block_reports;
          break;
        case "date":
          comparison =
            new Date(a.last_reported_at).getTime() -
            new Date(b.last_reported_at).getTime();
          break;
      }
      return sortOrder === "asc" ? comparison : -comparison;
    });

    return filtered;
  }, [sites, search, sortBy, sortOrder]);

  const toggleSort = (column: "domain" | "blocks" | "date") => {
    if (sortBy === column) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortBy(column);
      setSortOrder("desc");
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("en-IN", {
      year: "numeric",
      month: "short",
      day: "numeric",
    });
  };

  return (
    <div className="space-y-4">
      {/* Search and filters */}
      <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
        <div className="relative w-full sm:w-96">
          <svg
            viewBox="0 0 24 24"
            className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
          >
            <circle cx="11" cy="11" r="8" />
            <path d="M21 21l-4.35-4.35" />
          </svg>
          <Input
            placeholder="Search domains or categories..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-10 bg-card border-border"
          />
        </div>
        <p className="text-sm text-muted-foreground">
          {filteredAndSortedSites.length} sites found
        </p>
      </div>

      {/* Table */}
      <div className="rounded-xl border border-border bg-card overflow-hidden cyber-border">
        <div className="overflow-x-auto">
          <Table className="data-table">
            <TableHeader>
              <TableRow className="border-border hover:bg-transparent">
                <TableHead
                  className="cursor-pointer hover:text-primary transition-colors"
                  onClick={() => toggleSort("domain")}
                >
                  <div className="flex items-center gap-2">
                    Domain
                    {sortBy === "domain" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead>Categories</TableHead>
                <TableHead
                  className="cursor-pointer hover:text-primary transition-colors text-right"
                  onClick={() => toggleSort("blocks")}
                >
                  <div className="flex items-center justify-end gap-2">
                    Blocks
                    {sortBy === "blocks" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead className="text-right">Unblocks</TableHead>
                <TableHead
                  className="cursor-pointer hover:text-primary transition-colors"
                  onClick={() => toggleSort("date")}
                >
                  <div className="flex items-center gap-2">
                    Last Reported
                    {sortBy === "date" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredAndSortedSites.map((site, index) => (
                <TableRow
                  key={site.id}
                  className="border-border hover:bg-secondary/50 transition-colors animate-slide-up opacity-0"
                  style={{ animationDelay: `${Math.min(index * 30, 300)}ms` }}
                >
                  <TableCell className="font-mono text-sm">
                    <span className="text-foreground">{site.domain}</span>
                  </TableCell>
                  <TableCell>
                    <div className="flex flex-wrap gap-1">
                      {site.categories?.slice(0, 3).map((cat) => (
                        <Badge
                          key={cat}
                          variant="secondary"
                          className="badge-category text-xs"
                        >
                          {cat}
                        </Badge>
                      ))}
                      {site.categories && site.categories.length > 3 && (
                        <Badge variant="secondary" className="text-xs">
                          +{site.categories.length - 3}
                        </Badge>
                      )}
                    </div>
                  </TableCell>
                  <TableCell className="text-right">
                    <span className="font-mono font-bold text-[oklch(0.8_0.2_25)] tabular-nums">
                      {site.block_reports.toLocaleString()}
                    </span>
                  </TableCell>
                  <TableCell className="text-right">
                    <span className="font-mono font-bold text-[oklch(0.8_0.18_145)] tabular-nums">
                      {site.unblock_reports.toLocaleString()}
                    </span>
                  </TableCell>
                  <TableCell className="text-muted-foreground text-sm">
                    {formatDate(site.last_reported_at)}
                  </TableCell>
                  <TableCell className="text-right">
                    <Button
                      asChild
                      variant="ghost"
                      size="sm"
                      className="hover:bg-primary/10 hover:text-primary"
                    >
                      <Link href={`/sites/${site.id}`}>
                        <svg
                          viewBox="0 0 24 24"
                          className="h-4 w-4 mr-1"
                          fill="none"
                          stroke="currentColor"
                          strokeWidth="2"
                        >
                          <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                          <circle cx="12" cy="12" r="3" />
                        </svg>
                        View
                      </Link>
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </div>

      {filteredAndSortedSites.length === 0 && (
        <div className="text-center py-12 text-muted-foreground">
          <svg
            viewBox="0 0 24 24"
            className="h-12 w-12 mx-auto mb-4 opacity-50"
            fill="none"
            stroke="currentColor"
            strokeWidth="1.5"
          >
            <circle cx="11" cy="11" r="8" />
            <path d="M21 21l-4.35-4.35" />
          </svg>
          <p>No sites found matching your search.</p>
        </div>
      )}
    </div>
  );
}
