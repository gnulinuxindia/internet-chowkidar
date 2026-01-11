"use client";

import { useEffect, useState, useRef } from "react";

interface CounterProps {
  endValue: number;
  label: string;
  variant?: "blocked" | "unblocked" | "neutral";
  duration?: number;
}

export function Counter({
  endValue,
  label,
  variant = "neutral",
  duration = 1000,
}: CounterProps) {
  const [count, setCount] = useState(0);
  const [isVisible, setIsVisible] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
          observer.disconnect();
        }
      },
      { threshold: 0.1 }
    );

    if (ref.current) {
      observer.observe(ref.current);
    }

    return () => observer.disconnect();
  }, []);

  useEffect(() => {
    if (!isVisible) return;

    const startTime = performance.now();
    const startValue = 0;

    const animate = (currentTime: number) => {
      const elapsed = currentTime - startTime;
      const progress = Math.min(elapsed / duration, 1);

      // Cubic ease-out
      const easeOut = 1 - Math.pow(1 - progress, 3);
      const currentValue = Math.floor(startValue + (endValue - startValue) * easeOut);

      setCount(currentValue);

      if (progress < 1) {
        requestAnimationFrame(animate);
      }
    };

    requestAnimationFrame(animate);
  }, [isVisible, endValue, duration]);

  const variantClasses = {
    blocked: "text-[oklch(0.8_0.2_25)] glow-text-blocked",
    unblocked: "text-[oklch(0.8_0.18_145)] glow-text-unblocked",
    neutral: "text-primary glow-text-cyber",
  };

  return (
    <div
      ref={ref}
      className="flex flex-col items-center gap-2 p-6 rounded-xl bg-card border border-border cyber-border card-hover"
    >
      <span
        className={`text-4xl md:text-5xl font-bold tabular-nums font-mono ${variantClasses[variant]} animate-count-up`}
        style={{ fontFamily: "var(--font-mono)" }}
      >
        {count.toLocaleString()}
      </span>
      <span className="text-sm uppercase tracking-widest text-muted-foreground font-semibold">
        {label}
      </span>
    </div>
  );
}
