import { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";
import { getSite } from "@/lib/api";
import { Counter } from "@/components/counter";
import { SiteMap } from "@/components/site-map";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";

interface PageProps {
  params: Promise<{ id: string }>;
}

export async function generateMetadata({ params }: PageProps): Promise<Metadata> {
  const { id } = await params;
  try {
    const site = await getSite(id);
    return {
      title: site.domain,
      description: `View blocking reports for ${site.domain} across Indian ISPs.`,
    };
  } catch {
    return {
      title: "Site Not Found",
    };
  }
}

export default async function SiteDetailPage({ params }: PageProps) {
  const { id } = await params;
  let site;

  try {
    site = await getSite(id);
  } catch {
    notFound();
  }

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("en-IN", {
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <div className="min-h-screen">
      {/* Header */}
      <section className="py-12 border-b border-border/50 bg-card/30">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          {/* Breadcrumb */}
          <nav className="flex items-center gap-2 text-sm text-muted-foreground mb-6">
            <Link href="/sites" className="hover:text-primary transition-colors">
              Sites
            </Link>
            <svg
              viewBox="0 0 24 24"
              className="h-4 w-4"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
            >
              <path d="M9 18l6-6-6-6" />
            </svg>
            <span className="text-foreground">{site.domain}</span>
          </nav>

          {/* Title */}
          <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-6">
            <div>
              <h1 className="text-3xl md:text-4xl font-bold mb-4 font-mono">
                <span className="text-primary">{site.domain}</span>
              </h1>

              {/* Categories */}
              {site.categories && site.categories.length > 0 && (
                <div className="flex flex-wrap gap-2 mb-4">
                  {site.categories.map((cat) => (
                    <Badge key={cat} className="badge-category">
                      {cat}
                    </Badge>
                  ))}
                </div>
              )}

              <p className="text-muted-foreground">
                Last reported: {formatDate(site.last_reported_at)}
              </p>
            </div>

            <Button asChild variant="outline" className="shrink-0">
              <a
                href={`https://${site.domain}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                <svg
                  viewBox="0 0 24 24"
                  className="h-4 w-4 mr-2"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6" />
                  <path d="M15 3h6v6" />
                  <path d="M10 14L21 3" />
                </svg>
                Visit Site
              </a>
            </Button>
          </div>
        </div>
      </section>

      {/* Stats */}
      <section className="py-8 border-b border-border/50">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Counter
              endValue={site.block_reports}
              label="Block Reports"
              variant="blocked"
            />
            <Counter
              endValue={site.unblock_reports}
              label="Unblock Reports"
              variant="unblocked"
            />
            <Counter
              endValue={site.blocked_by_isps?.length || 0}
              label="ISPs Blocking"
              variant="neutral"
            />
          </div>
        </div>
      </section>

      {/* Map */}
      {site.blocked_by_isps && site.blocked_by_isps.length > 0 && (
        <section className="py-8 border-b border-border/50">
          <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <h2 className="text-xl font-bold mb-6 flex items-center gap-3">
              <svg
                viewBox="0 0 24 24"
                className="h-6 w-6 text-primary"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
              >
                <circle cx="12" cy="12" r="10" />
                <path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z" />
              </svg>
              Geographic Distribution
            </h2>
            <SiteMap isps={site.blocked_by_isps} />
          </div>
        </section>
      )}

      {/* ISP List */}
      <section className="py-8">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <h2 className="text-xl font-bold mb-6 flex items-center gap-3">
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
            Blocked by ISPs ({site.blocked_by_isps?.length || 0})
          </h2>

          {site.blocked_by_isps && site.blocked_by_isps.length > 0 ? (
            <div className="grid gap-4">
              {site.blocked_by_isps.map((isp, index) => (
                <Link
                  key={isp.id}
                  href={`/isps/${isp.id}`}
                  className="group p-4 rounded-xl bg-card border border-border hover:border-primary/50 transition-all duration-300 card-hover animate-slide-up opacity-0"
                  style={{ animationDelay: `${index * 50}ms` }}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-4">
                      <div className="h-10 w-10 rounded-full bg-primary/10 border border-primary/30 flex items-center justify-center text-primary font-bold text-sm">
                        {index + 1}
                      </div>
                      <div>
                        <h3 className="font-semibold group-hover:text-primary transition-colors">
                          {isp.name}
                        </h3>
                        <p className="text-sm text-muted-foreground">
                          {isp.latitude.toFixed(4)}, {isp.longitude.toFixed(4)}
                        </p>
                      </div>
                    </div>

                    <div className="flex items-center gap-6">
                      <div className="text-right">
                        <p className="font-mono font-bold text-[oklch(0.8_0.2_25)] tabular-nums">
                          {isp.block_reports.toLocaleString()}
                        </p>
                        <p className="text-xs text-muted-foreground">Blocks</p>
                      </div>
                      <div className="text-right">
                        <p className="font-mono font-bold text-[oklch(0.8_0.18_145)] tabular-nums">
                          {isp.unblock_reports.toLocaleString()}
                        </p>
                        <p className="text-xs text-muted-foreground">Unblocks</p>
                      </div>
                      <svg
                        viewBox="0 0 24 24"
                        className="h-5 w-5 text-muted-foreground group-hover:text-primary transition-colors"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                      >
                        <path d="M9 18l6-6-6-6" />
                      </svg>
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          ) : (
            <div className="text-center py-12 text-muted-foreground">
              <svg
                viewBox="0 0 24 24"
                className="h-12 w-12 mx-auto mb-4 opacity-50"
                fill="none"
                stroke="currentColor"
                strokeWidth="1.5"
              >
                <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
              </svg>
              <p>No ISP blocking data available.</p>
            </div>
          )}
        </div>
      </section>
    </div>
  );
}
