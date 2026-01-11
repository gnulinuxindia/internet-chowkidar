import { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";
import { getISP } from "@/lib/api";
import { Counter } from "@/components/counter";
import { Badge } from "@/components/ui/badge";

interface PageProps {
  params: Promise<{ id: string }>;
}

export async function generateMetadata({ params }: PageProps): Promise<Metadata> {
  const { id } = await params;
  try {
    const isp = await getISP(id);
    return {
      title: isp.name,
      description: `View websites blocked by ${isp.name} in India.`,
    };
  } catch {
    return {
      title: "ISP Not Found",
    };
  }
}

export default async function ISPDetailPage({ params }: PageProps) {
  const { id } = await params;
  let isp;

  try {
    isp = await getISP(id);
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
            <Link href="/isps" className="hover:text-primary transition-colors">
              ISPs
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
            <span className="text-foreground">{isp.name}</span>
          </nav>

          {/* Title */}
          <div className="flex items-start gap-4">
            <div className="h-16 w-16 rounded-xl bg-primary/10 border border-primary/30 flex items-center justify-center text-primary text-2xl font-bold shrink-0">
              {isp.name.charAt(0).toUpperCase()}
            </div>
            <div>
              <h1 className="text-3xl md:text-4xl font-bold mb-2">
                <span className="text-primary">{isp.name}</span>
              </h1>

              <div className="flex flex-wrap items-center gap-4 text-sm text-muted-foreground">
                <div className="flex items-center gap-2">
                  <svg
                    viewBox="0 0 24 24"
                    className="h-4 w-4"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                  >
                    <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z" />
                    <circle cx="12" cy="10" r="3" />
                  </svg>
                  {isp.latitude.toFixed(4)}, {isp.longitude.toFixed(4)}
                </div>
                {isp.updated_at && (
                  <div className="flex items-center gap-2">
                    <svg
                      viewBox="0 0 24 24"
                      className="h-4 w-4"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                    >
                      <circle cx="12" cy="12" r="10" />
                      <path d="M12 6v6l4 2" />
                    </svg>
                    Updated: {formatDate(isp.updated_at)}
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Stats */}
      <section className="py-8 border-b border-border/50">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <Counter
              endValue={isp.blocks?.length || 0}
              label="Sites Blocked"
              variant="blocked"
            />
            <Counter
              endValue={isp.blocks?.reduce((acc, b) => acc + b.block_reports, 0) || 0}
              label="Total Block Reports"
              variant="blocked"
            />
            <Counter
              endValue={isp.blocks?.reduce((acc, b) => acc + b.unblock_reports, 0) || 0}
              label="Total Unblock Reports"
              variant="unblocked"
            />
          </div>
        </div>
      </section>

      {/* Blocked Sites List */}
      <section className="py-8">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <h2 className="text-xl font-bold mb-6 flex items-center gap-3">
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
            Blocked Sites ({isp.blocks?.length || 0})
          </h2>

          {isp.blocks && isp.blocks.length > 0 ? (
            <div className="grid gap-4">
              {isp.blocks.map((block, index) => (
                <Link
                  key={block.id}
                  href={`/sites/${block.site_id}`}
                  className="group p-4 rounded-xl bg-card border border-border hover:border-primary/50 transition-all duration-300 card-hover animate-slide-up opacity-0"
                  style={{ animationDelay: `${index * 50}ms` }}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-4">
                      <div className="h-10 w-10 rounded-full bg-[oklch(0.7_0.22_25)]/10 border border-[oklch(0.7_0.22_25)]/30 flex items-center justify-center text-[oklch(0.8_0.2_25)] font-bold text-sm">
                        {index + 1}
                      </div>
                      <div>
                        <h3 className="font-mono font-semibold group-hover:text-primary transition-colors">
                          {block.domain}
                        </h3>
                        <p className="text-sm text-muted-foreground">
                          Last reported: {formatDate(block.last_reported_at)}
                        </p>
                      </div>
                    </div>

                    <div className="flex items-center gap-6">
                      <div className="text-right">
                        <p className="font-mono font-bold text-[oklch(0.8_0.2_25)] tabular-nums">
                          {block.block_reports.toLocaleString()}
                        </p>
                        <p className="text-xs text-muted-foreground">Blocks</p>
                      </div>
                      <div className="text-right">
                        <p className="font-mono font-bold text-[oklch(0.8_0.18_145)] tabular-nums">
                          {block.unblock_reports.toLocaleString()}
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
                <circle cx="12" cy="12" r="10" />
                <path d="M9 12l2 2 4-4" />
              </svg>
              <p>No blocked sites data available.</p>
            </div>
          )}
        </div>
      </section>
    </div>
  );
}
