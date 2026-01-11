"use client";

import { useEffect } from "react";
import { Button } from "@/components/ui/button";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center px-4">
        <div className="mb-8">
          <div className="h-20 w-20 mx-auto rounded-full bg-destructive/10 border border-destructive/30 flex items-center justify-center mb-4">
            <svg
              viewBox="0 0 24 24"
              className="h-10 w-10 text-destructive"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
            >
              <circle cx="12" cy="12" r="10" />
              <path d="M12 8v4M12 16h.01" />
            </svg>
          </div>
        </div>

        <h2 className="text-2xl font-bold mb-4">Something went wrong</h2>
        <p className="text-muted-foreground mb-8 max-w-md mx-auto">
          An error occurred while loading this page. Please try again.
        </p>

        <Button onClick={reset} className="glow-cyber">
          <svg
            viewBox="0 0 24 24"
            className="h-4 w-4 mr-2"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
          >
            <path d="M1 4v6h6M23 20v-6h-6" />
            <path d="M20.49 9A9 9 0 0 0 5.64 5.64L1 10m22 4l-4.64 4.36A9 9 0 0 1 3.51 15" />
          </svg>
          Try Again
        </Button>
      </div>
    </div>
  );
}
