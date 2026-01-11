import { Metadata } from "next";
import { getSites } from "@/lib/api";
import { SitesTable } from "@/components/sites-table";

export const metadata: Metadata = {
  title: "Blocked Sites",
  description:
    "View all websites blocked by ISPs in India. Track internet censorship in real-time.",
};

export default async function SitesPage() {
  const sites = await getSites();

  return (
    <div className="min-h-screen">
      {/* Header */}
      <section className="py-16 border-b border-border/50 bg-card/30">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex items-center gap-4 mb-4">
            <div className="h-12 w-12 rounded-xl bg-[oklch(0.7_0.22_25)]/10 border border-[oklch(0.7_0.22_25)]/30 flex items-center justify-center">
              <svg
                viewBox="0 0 24 24"
                className="h-6 w-6 text-[oklch(0.8_0.2_25)]"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
              >
                <circle cx="12" cy="12" r="10" />
                <path d="M4.93 4.93l14.14 14.14" />
              </svg>
            </div>
            <div>
              <h1 className="text-3xl md:text-4xl font-bold">
                <span className="gradient-text">Blocked Sites</span>
              </h1>
              <p className="text-muted-foreground mt-1">
                Websites blocked by ISPs across India
              </p>
            </div>
          </div>

          {/* Stats */}
          <div className="flex flex-wrap gap-6 mt-8">
            <div className="flex items-center gap-2">
              <span className="text-2xl font-bold font-mono text-primary tabular-nums">
                {sites.length.toLocaleString()}
              </span>
              <span className="text-sm text-muted-foreground">Total Sites</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-2xl font-bold font-mono text-[oklch(0.8_0.2_25)] tabular-nums">
                {sites
                  .reduce((acc, site) => acc + site.block_reports, 0)
                  .toLocaleString()}
              </span>
              <span className="text-sm text-muted-foreground">
                Total Block Reports
              </span>
            </div>
          </div>
        </div>
      </section>

      {/* Table */}
      <section className="py-8">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <SitesTable sites={sites} />
        </div>
      </section>
    </div>
  );
}
