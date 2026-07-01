"use client";

import Link from "next/link";
import { useEffect, useState, useMemo } from "react";
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
import type { ISP } from "@/lib/types";

interface ISPsTableProps {
  isps: ISP[];
}

export function ISPsTable({ isps }: ISPsTableProps) {
  const [search, setSearch] = useState("");
  const [sortBy, setSortBy] = useState<"name" | "blocks" | "date">("blocks");
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("desc");
  const [currentPage, setCurrentPage] = useState(1);
  const pageSize = 25;

  const filteredAndSortedISPs = useMemo(() => {
    let filtered = isps.filter((isp) =>
      isp.name.toLowerCase().includes(search.toLowerCase())
    );

    filtered.sort((a, b) => {
      let comparison = 0;
      switch (sortBy) {
        case "name":
          comparison = a.name.localeCompare(b.name);
          break;
        case "blocks":
          comparison = a.block_reports - b.block_reports;
          break;
        case "date":
          comparison =
            new Date(a.last_reported_at || 0).getTime() -
            new Date(b.last_reported_at || 0).getTime();
          break;
      }
      return sortOrder === "asc" ? comparison : -comparison;
    });

    return filtered;
  }, [isps, search, sortBy, sortOrder]);

  const totalPages = Math.max(1, Math.ceil(filteredAndSortedISPs.length / pageSize));
  const currentPageForView = Math.min(currentPage, totalPages);
  const pageStart = (currentPageForView - 1) * pageSize;
  const paginatedISPs = filteredAndSortedISPs.slice(pageStart, pageStart + pageSize);

  useEffect(() => {
    setCurrentPage(1);
  }, [search, sortBy, sortOrder]);

  useEffect(() => {
    setCurrentPage((page) => Math.min(page, totalPages));
  }, [totalPages]);

  const toggleSort = (column: "name" | "blocks" | "date") => {
    if (sortBy === column) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortBy(column);
      setSortOrder("desc");
    }
  };

  const formatDate = (dateString?: string) => {
    if (!dateString) return "N/A";
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
            placeholder="Search ISPs"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-10 bg-card border-border"
          />
        </div>
        <p className="text-sm text-muted-foreground">
          Showing {filteredAndSortedISPs.length === 0 ? 0 : pageStart + 1}-
          {Math.min(pageStart + pageSize, filteredAndSortedISPs.length)} of {filteredAndSortedISPs.length} ISPs
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
                  onClick={() => toggleSort("name")}
                >
                  <div className="flex items-center gap-2">
                    ISP name
                    {sortBy === "name" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead
                  className="cursor-pointer hover:text-primary transition-colors text-right"
                  onClick={() => toggleSort("blocks")}
                >
                  <div className="flex items-center justify-end gap-2">
                    Block reports
                    {sortBy === "blocks" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead className="text-right">Unblock reports</TableHead>
                <TableHead
                  className="cursor-pointer hover:text-primary transition-colors"
                  onClick={() => toggleSort("date")}
                >
                  <div className="flex items-center gap-2">
                    Last report
                    {sortBy === "date" && (
                      <span className="text-primary">
                        {sortOrder === "asc" ? "↑" : "↓"}
                      </span>
                    )}
                  </div>
                </TableHead>
                <TableHead className="text-right">Details</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {paginatedISPs.map((isp, index) => (
                <TableRow
                  key={isp.id}
                  className="border-border hover:bg-secondary/50 transition-colors animate-slide-up opacity-0"
                  style={{ animationDelay: `${Math.min(index * 30, 300)}ms` }}
                >
                  <TableCell>
                    <div className="flex items-center gap-3">
                      <div className="h-8 w-8 rounded-lg bg-primary/10 border border-primary/30 flex items-center justify-center text-primary font-bold text-xs">
                        {isp.name.charAt(0).toUpperCase()}
                      </div>
                      <span className="font-medium">{isp.name}</span>
                    </div>
                  </TableCell>
                  <TableCell className="text-right">
                    <span className="font-mono font-bold text-[oklch(0.8_0.2_25)] tabular-nums">
                      {isp.block_reports.toLocaleString()}
                    </span>
                  </TableCell>
                  <TableCell className="text-right">
                    <span className="font-mono font-bold text-[oklch(0.8_0.18_145)] tabular-nums">
                      {isp.unblock_reports.toLocaleString()}
                    </span>
                  </TableCell>
                  <TableCell className="text-muted-foreground text-sm">
                    {formatDate(isp.last_reported_at)}
                  </TableCell>
                  <TableCell className="text-right">
                    <Button
                      asChild
                      variant="ghost"
                      size="sm"
                      className="hover:bg-primary/10 hover:text-primary"
                    >
                      <Link href={`/isps/${isp.id}`}>
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

      {filteredAndSortedISPs.length > pageSize && (
        <div className="flex flex-col sm:flex-row items-center justify-between gap-3">
          <p className="text-sm text-muted-foreground">
            Page {currentPageForView} of {totalPages}
          </p>
          <div className="flex items-center gap-2">
            <Button
              type="button"
              variant="outline"
              size="sm"
              onClick={() => setCurrentPage((page) => Math.max(1, page - 1))}
              disabled={currentPageForView === 1}
            >
              Previous
            </Button>
            <Button
              type="button"
              variant="outline"
              size="sm"
              onClick={() => setCurrentPage((page) => Math.min(totalPages, page + 1))}
              disabled={currentPageForView === totalPages}
            >
              Next
            </Button>
          </div>
        </div>
      )}

      {filteredAndSortedISPs.length === 0 && (
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
          <p>No ISPs match your search.</p>
        </div>
      )}
    </div>
  );
}
