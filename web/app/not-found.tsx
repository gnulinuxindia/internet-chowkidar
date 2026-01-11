import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center px-4">
        <div className="mb-8">
          <h1 className="text-8xl font-bold text-primary mb-4">404</h1>
          <div className="h-1 w-24 mx-auto bg-gradient-to-r from-primary via-accent to-primary rounded-full" />
        </div>

        <h2 className="text-2xl font-bold mb-4">Page Not Found</h2>
        <p className="text-muted-foreground mb-8 max-w-md mx-auto">
          The page you&apos;re looking for doesn&apos;t exist or has been moved.
        </p>

        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Button asChild className="glow-cyber">
            <Link href="/">
              <svg
                viewBox="0 0 24 24"
                className="h-4 w-4 mr-2"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
              >
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
                <path d="M9 22V12h6v10" />
              </svg>
              Go Home
            </Link>
          </Button>
          <Button asChild variant="outline">
            <Link href="/sites">
              <svg
                viewBox="0 0 24 24"
                className="h-4 w-4 mr-2"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
              >
                <circle cx="12" cy="12" r="10" />
                <path d="M12 6v6l4 2" />
              </svg>
              View Sites
            </Link>
          </Button>
        </div>
      </div>
    </div>
  );
}
