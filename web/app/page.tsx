import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function HomePage() {
  return (
    <div className="relative">
      {/* Hero Section */}
      <section className="hero-gradient grid-bg relative min-h-[90vh] flex items-center justify-center overflow-hidden">
        {/* Decorative elements */}
        <div className="absolute inset-0 overflow-hidden">
          <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-primary/10 rounded-full blur-3xl animate-pulse-glow" />
          <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-accent/10 rounded-full blur-3xl animate-pulse-glow delay-500" />
        </div>

        <div className="relative z-10 mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 text-center">
          {/* Badge */}
          <div className="animate-slide-up opacity-0 delay-100">
            <span className="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary/10 border border-primary/30 text-sm text-primary font-medium mb-8">
              <span className="relative flex h-2 w-2">
                <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-primary opacity-75" />
                <span className="relative inline-flex rounded-full h-2 w-2 bg-primary" />
              </span>
              Monitoring Internet Freedom in India
            </span>
          </div>

          {/* Title */}
          <h1 className="hero-title font-extrabold tracking-tight mb-6 animate-slide-up opacity-0 delay-200">
            <span className="gradient-text glitch-hover">Internet</span>
            <br />
            <span className="text-foreground">Chowkidar</span>
          </h1>

          {/* Subtitle */}
          <p className="text-xl md:text-2xl text-muted-foreground max-w-2xl mx-auto mb-4 animate-slide-up opacity-0 delay-300">
            We notice when{" "}
            <span className="text-[oklch(0.8_0.2_25)] font-semibold relative">
              ISPs block
              <span className="absolute bottom-0 left-0 right-0 h-0.5 bg-[oklch(0.8_0.2_25)]" />
            </span>
            .
          </p>

          <p className="text-lg text-muted-foreground max-w-xl mx-auto mb-12 animate-slide-up opacity-0 delay-400">
            Tracking censorship across Indian internet service providers.
            Transparent data for a free internet.
          </p>

          {/* CTA Buttons */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center animate-slide-up opacity-0 delay-500">
            <Button
              asChild
              size="lg"
              className="text-lg px-8 py-6 bg-primary hover:bg-primary/90 text-primary-foreground glow-cyber"
            >
              <Link href="/sites">
                <svg
                  viewBox="0 0 24 24"
                  className="h-5 w-5 mr-2"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <circle cx="12" cy="12" r="10" />
                  <path d="M12 6v6l4 2" />
                </svg>
                See What&apos;s Blocked
              </Link>
            </Button>
            <Button
              asChild
              variant="outline"
              size="lg"
              className="text-lg px-8 py-6 border-border hover:bg-secondary"
            >
              <Link href="/isps">
                <svg
                  viewBox="0 0 24 24"
                  className="h-5 w-5 mr-2"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z" />
                  <path d="M22 6l-10 7L2 6" />
                </svg>
                View ISPs
              </Link>
            </Button>
          </div>

          {/* Tech Stack */}
          <div className="mt-20 animate-slide-up opacity-0 delay-600">
            <p className="text-xs uppercase tracking-widest text-muted-foreground mb-4">
              Built With
            </p>
            <div className="flex flex-wrap justify-center gap-4">
              {[
                { name: "Next.js", url: "https://nextjs.org" },
                { name: "Tailwind", url: "https://tailwindcss.com" },
                { name: "Go", url: "https://go.dev" },
                { name: "PostgreSQL", url: "https://postgresql.org" },
              ].map((tech) => (
                <a
                  key={tech.name}
                  href={tech.url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="px-4 py-2 rounded-lg bg-card border border-border text-sm text-muted-foreground hover:text-primary hover:border-primary/50 transition-all duration-200"
                >
                  {tech.name}
                </a>
              ))}
            </div>
          </div>
        </div>

        {/* Scroll indicator */}
        <div className="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
          <svg
            viewBox="0 0 24 24"
            className="h-6 w-6 text-muted-foreground"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
          >
            <path d="M12 5v14M19 12l-7 7-7-7" />
          </svg>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-24 bg-card/50">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              <span className="gradient-text">Why This Matters</span>
            </h2>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              Internet censorship affects millions. We provide transparency.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                icon: (
                  <svg
                    viewBox="0 0 24 24"
                    className="h-8 w-8"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  >
                    <circle cx="12" cy="12" r="10" />
                    <path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z" />
                  </svg>
                ),
                title: "Track Blocked Sites",
                description:
                  "See which websites are being blocked by ISPs across India in real-time.",
              },
              {
                icon: (
                  <svg
                    viewBox="0 0 24 24"
                    className="h-8 w-8"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  >
                    <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
                    <path d="M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12" />
                  </svg>
                ),
                title: "ISP Analysis",
                description:
                  "Compare different ISPs and understand their blocking patterns.",
              },
              {
                icon: (
                  <svg
                    viewBox="0 0 24 24"
                    className="h-8 w-8"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  >
                    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
                    <path d="M9 12l2 2 4-4" />
                  </svg>
                ),
                title: "Community Reports",
                description:
                  "Crowdsourced data from users across the country ensures accuracy.",
              },
            ].map((feature, i) => (
              <div
                key={i}
                className="p-8 rounded-xl bg-card border border-border cyber-border card-hover group"
              >
                <div className="w-14 h-14 rounded-xl bg-primary/10 border border-primary/30 flex items-center justify-center text-primary mb-6 group-hover:glow-cyber transition-all duration-300">
                  {feature.icon}
                </div>
                <h3 className="text-xl font-bold mb-3">{feature.title}</h3>
                <p className="text-muted-foreground">{feature.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-24 relative overflow-hidden">
        <div className="absolute inset-0 grid-bg opacity-50" />
        <div className="relative z-10 mx-auto max-w-4xl px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            Notice Something <span className="text-[oklch(0.8_0.2_25)]">Blocked</span>?
          </h2>
          <p className="text-lg text-muted-foreground mb-8 max-w-2xl mx-auto">
            Help us track internet censorship. Report blocked sites and
            contribute to a more transparent internet.
          </p>
          <Button
            asChild
            size="lg"
            variant="outline"
            className="text-lg px-8 py-6 border-primary/50 hover:bg-primary/10 hover:border-primary"
          >
            <a
              href="https://github.com/gnulinuxindia/internet-chowkidar/issues"
              target="_blank"
              rel="noopener noreferrer"
            >
              <svg
                viewBox="0 0 24 24"
                className="h-5 w-5 mr-2"
                fill="currentColor"
              >
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
              </svg>
              Report on GitHub
            </a>
          </Button>
        </div>
      </section>
    </div>
  );
}
