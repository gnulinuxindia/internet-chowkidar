import { Metadata } from "next";
import { getISPs } from "@/lib/api";
import { ISPsTable } from "@/components/isps-table";

export const metadata: Metadata = {
  title: "ISPs",
  description:
    "View all Internet Service Providers in India and their website blocking patterns.",
};

export default async function ISPsPage() {
  const isps = await getISPs();

  return (
    <div className="min-h-screen">
      {/* Header */}
      <section className="py-16 border-b border-border/50 bg-card/30">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex items-center gap-4 mb-4">
            <div className="h-12 w-12 rounded-xl bg-primary/10 border border-primary/30 flex items-center justify-center">
              <svg
                viewBox="0 0 24 24"
                className="h-6 w-6 text-primary"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
              >
                <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
                <path d="M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12" />
              </svg>
            </div>
            <div>
              <h1 className="text-3xl md:text-4xl font-bold">
                <span className="gradient-text">Internet Service Providers</span>
              </h1>
              <p className="text-muted-foreground mt-1">
                ISPs across India and their blocking patterns
              </p>
            </div>
          </div>

          {/* Stats */}
          <div className="flex flex-wrap gap-6 mt-8">
            <div className="flex items-center gap-2">
              <span className="text-2xl font-bold font-mono text-primary tabular-nums">
                {isps.length.toLocaleString()}
              </span>
              <span className="text-sm text-muted-foreground">Total ISPs</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-2xl font-bold font-mono text-[oklch(0.8_0.2_25)] tabular-nums">
                {isps
                  .reduce((acc, isp) => acc + isp.block_reports, 0)
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
          <ISPsTable isps={isps} />
        </div>
      </section>
    </div>
  );
}
