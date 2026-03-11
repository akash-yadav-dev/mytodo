import { cn } from "@/lib/utils/cn";
import Link from "next/link";
import React from "react";

type Variant = "primary" | "secondary" | "ghost" | "danger";
type Size = "sm" | "md" | "lg";

type BaseProps = {
  variant?: Variant;
  size?: Size;
  className?: string;
};

type ButtonProps = BaseProps &
  (
    | ({ asChild?: false } & React.ButtonHTMLAttributes<HTMLButtonElement>)
    | ({ asChild: true; href: string } & Omit<React.AnchorHTMLAttributes<HTMLAnchorElement>, "href">)
  );

const variantClasses: Record<Variant, string> = {
  primary:
    "bg-accent text-white hover:bg-accent/90 active:scale-[0.98]",
  secondary:
    "border border-border bg-card text-fg hover:bg-accentMuted hover:border-accent/40",
  ghost:
    "text-fg/70 hover:bg-accentMuted hover:text-fg",
  danger:
    "bg-red-500 text-white hover:bg-red-600",
};

const sizeClasses: Record<Size, string> = {
  sm: "px-3 py-1.5 text-xs",
  md: "px-4 py-2 text-sm",
  lg: "px-6 py-3 text-base",
};

export function Button({
  variant = "primary",
  size = "md",
  className,
  ...props
}: ButtonProps) {
  const base =
    "inline-flex items-center justify-center gap-2 rounded-xl font-medium transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent/50 disabled:pointer-events-none disabled:opacity-50";

  const classes = cn(base, variantClasses[variant], sizeClasses[size], className);

  if ("asChild" in props && props.asChild) {
    const { asChild: _, href, ...rest } = props as { asChild: true; href: string; className?: string };
    return <Link href={href} className={classes} {...(rest as Record<string, unknown>)} />;
  }

  const { asChild: _, ...rest } = props as { asChild?: false } & React.ButtonHTMLAttributes<HTMLButtonElement>;
  return <button className={classes} {...rest} />;
}
