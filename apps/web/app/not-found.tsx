import Link from "next/link";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";

export const metadata = {
  title: "Page not found",
};

export default function NotFound() {
  return (
    <div className="flex min-h-[calc(100vh-8rem)] flex-col items-center justify-center px-6 text-center">
      {/* Decorative blob */}
      <div
        aria-hidden
        className="pointer-events-none absolute inset-0 -z-10 overflow-hidden"
      >
        <div className="absolute left-1/2 top-1/3 h-[600px] w-[600px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-accent/10 blur-[120px]" />
      </div>

      {/* 404 number */}
      <p className="mb-2 font-mono text-8xl font-bold tracking-tighter text-accent md:text-9xl">
        404
      </p>

      <h1 className="mb-3 text-2xl font-semibold md:text-3xl">
        This page doesn&apos;t exist
      </h1>
      <p className="mb-8 max-w-md text-base text-fg/60">
        The page you&apos;re looking for may have been moved, deleted, or never
        existed. Double-check the URL or head back to safety.
      </p>

      <div className="flex flex-wrap items-center justify-center gap-3">
        <Button asChild href={ROUTES.home}>
          Go home
        </Button>
        <Button asChild href={ROUTES.dashboard} variant="secondary">
          Open dashboard
        </Button>
      </div>

      <div className="mt-12 flex items-center gap-6 text-sm text-fg/50">
        <Link href={ROUTES.issues} className="transition hover:text-fg">
          Issues
        </Link>
        <span aria-hidden>·</span>
        <Link href={ROUTES.projects} className="transition hover:text-fg">
          Projects
        </Link>
        <span aria-hidden>·</span>
        <Link href={ROUTES.settings} className="transition hover:text-fg">
          Settings
        </Link>
      </div>
    </div>
  );
}
