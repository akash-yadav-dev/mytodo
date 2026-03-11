import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import Link from "next/link";
import { ROUTES } from "@/lib/constants/routes";

const features = [
  {
    title: "Project clarity",
    description:
      "Organize work into projects with clear ownership, milestones, and priorities.",
    icon: "���",
  },
  {
    title: "Focused execution",
    description:
      "Kanban boards and sprint planning keep the team aligned without overhead.",
    icon: "���",
  },
  {
    title: "Smart notifications",
    description:
      "Stay informed on what matters. Quiet the noise, amplify the signal.",
    icon: "���",
  },
  {
    title: "Team collaboration",
    description:
      "Comments, mentions, and real-time updates keep everyone in the loop.",
    icon: "���",
  },
  {
    title: "Powerful search",
    description:
      "Find any issue, project, or team member instantly with full-text search.",
    icon: "���",
  },
  {
    title: "Developer-friendly",
    description:
      "REST + gRPC APIs, webhooks, and SDK for seamless integrations.",
    icon: "⚙️",
  },
];

const testimonials = [
  {
    quote: "MyTodo replaced three tools for us. Everything we need in one place.",
    author: "Sarah K.",
    role: "Engineering Lead",
  },
  {
    quote: "Our sprint velocity went up 30% after switching. The boards are amazing.",
    author: "James R.",
    role: "Product Manager",
  },
  {
    quote: "Clean UI, fast API, and no bloat. Exactly what a dev team needs.",
    author: "Priya S.",
    role: "Staff Engineer",
  },
];

export default function HomePage() {
  return (
    <>
      {/* Hero */}
      <section className="relative overflow-hidden px-6 py-24 text-center">
        <div className="absolute inset-0 -z-10 bg-[radial-gradient(ellipse_at_top,_rgba(249,115,22,0.12),_transparent_60%),radial-gradient(ellipse_at_bottom-left,_rgba(59,130,246,0.10),_transparent_60%)]" />
        <div className="mx-auto max-w-3xl">
          <Badge variant="default" className="mb-4">
            Now in public beta
          </Badge>
          <h1 className="mb-6 text-5xl font-semibold leading-tight tracking-tight md:text-6xl">
            Tickets managed.{" "}
            <span className="text-accent">Teams unblocked.</span>
          </h1>
          <p className="mb-8 text-lg text-fg/70 md:text-xl">
            MyTodo is a production-grade project workspace — boards, issues,
            sprints, and automations — without the enterprise bloat.
          </p>
          <div className="flex flex-wrap items-center justify-center gap-3">
            <Button asChild href={ROUTES.register}>
              Start for free
            </Button>
            <Button asChild href={ROUTES.login} variant="secondary">
              Sign in
            </Button>
          </div>
          <p className="mt-4 text-xs text-fg/50">
            No credit card required. Free for teams up to 5.
          </p>
        </div>
      </section>

      {/* Features */}
      <section id="features" className="px-6 py-20">
        <div className="mx-auto max-w-6xl">
          <div className="mb-12 text-center">
            <h2 className="text-3xl font-semibold md:text-4xl">
              Everything your team needs
            </h2>
            <p className="mt-3 text-fg/70">
              From first issue to shipped feature — all in one workspace.
            </p>
          </div>
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {features.map((feat) => (
              <div
                key={feat.title}
                className="group rounded-3xl border border-border bg-card p-6 shadow-soft transition hover:-translate-y-1 hover:border-accent/40 hover:shadow-lg"
              >
                <span className="mb-3 block text-3xl">{feat.icon}</span>
                <h3 className="mb-2 text-lg font-medium">{feat.title}</h3>
                <p className="text-sm leading-relaxed text-fg/70">
                  {feat.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Testimonials */}
      <section className="bg-card/50 px-6 py-20">
        <div className="mx-auto max-w-5xl">
          <h2 className="mb-12 text-center text-3xl font-semibold">
            Loved by teams
          </h2>
          <div className="grid gap-6 md:grid-cols-3">
            {testimonials.map((t) => (
              <blockquote
                key={t.author}
                className="rounded-3xl border border-border bg-card p-6 shadow-soft"
              >
                <p className="mb-4 leading-relaxed text-fg/80">
                  &ldquo;{t.quote}&rdquo;
                </p>
                <footer className="text-sm">
                  <span className="font-medium text-fg">{t.author}</span>
                  <span className="ml-1 text-fg/50">· {t.role}</span>
                </footer>
              </blockquote>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section id="pricing" className="px-6 py-24 text-center">
        <div className="mx-auto max-w-2xl">
          <h2 className="mb-4 text-3xl font-semibold md:text-4xl">
            Simple pricing, no surprises
          </h2>
          <p className="mb-8 text-fg/70">
            Free for individuals and small teams. Scale as you grow.
          </p>
          <Link
            href={ROUTES.register}
            className="inline-flex items-center gap-2 rounded-full bg-accent px-8 py-3 text-sm font-semibold text-white shadow-soft transition hover:-translate-y-0.5 hover:shadow-lg"
          >
            Get started — it&apos;s free
          </Link>
        </div>
      </section>
    </>
  );
}
