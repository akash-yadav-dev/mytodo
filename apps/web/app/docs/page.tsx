import type { Metadata } from "next";
import Link from "next/link";
import type { LucideIcon } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { ROUTES } from "@/lib/constants/routes";
import {
  BookOpen,
  Zap,
  ShieldCheck,
  Layers,
  GitBranch,
  Database,
  Bell,
  Search,
  Puzzle,
  Terminal,
  ArrowRight,
  Code2,
  Users,
  Kanban,
  FolderKanban,
  CircuitBoard,
} from "lucide-react";

export const metadata: Metadata = {
  title: "Documentation",
  description:
    "Learn how MyTodo works — architecture, features, API reference, and more.",
};

// ─── Doc sections sidebar ───────────────────────────────────────────────────

const sections = [
  { id: "overview", label: "Overview" },
  { id: "getting-started", label: "Getting started" },
  { id: "architecture", label: "Architecture" },
  { id: "features", label: "Features" },
  { id: "api", label: "API reference" },
  { id: "auth", label: "Authentication" },
  { id: "self-hosting", label: "Self-hosting" },
  { id: "sdk", label: "SDK & integrations" },
  { id: "roadmap", label: "Roadmap" },
];

// ─── Feature groups ──────────────────────────────────────────────────────────

const featureGroups = [
  {
    icon: Users,
    title: "User & Organization Management",
    items: [
      "JWT-based authentication with refresh-token rotation",
      "OAuth 2.0 — Google and GitHub social login",
      "Multi-tenant organizations with role-based access (Owner, Admin, Member, Viewer)",
      "Invitations, member management, and audit logs",
    ],
  },
  {
    icon: FolderKanban,
    title: "Projects & Issues",
    items: [
      "Unlimited projects per organization with custom identifiers",
      "Rich issue model: title, description (Markdown), type, status, priority, assignee, due date, story points, labels, and custom fields",
      "Parent–child issue hierarchy (sub-tasks)",
      "Bulk operations and CSV/JSON export",
    ],
  },
  {
    icon: Kanban,
    title: "Boards & Sprints",
    items: [
      "Kanban boards with drag-and-drop column management",
      "Sprint planning with backlog, active sprint, and velocity tracking",
      "Customizable workflow statuses per project",
      "Burn-down and cumulative-flow charts",
    ],
  },
  {
    icon: Bell,
    title: "Notifications",
    items: [
      "Real-time in-app notifications via WebSockets",
      "Email notifications with configurable digest schedules",
      "Fine-grained per-user notification preferences",
      "Mention (@user) and watch support",
    ],
  },
  {
    icon: Zap,
    title: "Automation Rules",
    items: [
      "Event-driven triggers: issue created/updated/transitioned, comment added, due-date passed",
      "Actions: assign, change status, add label, post comment, send webhook",
      "Execution history and rule audit trail",
    ],
  },
  {
    icon: Search,
    title: "Full-Text Search",
    items: [
      "Elasticsearch-powered search across all entities",
      "Faceted filters: project, assignee, status, priority, label, date range",
      "Saved search views and quick-find command palette",
    ],
  },
  {
    icon: Puzzle,
    title: "Integrations",
    items: [
      "GitHub — link PRs and commits to issues",
      "Slack — post updates and receive slash commands",
      "Outbound webhooks with retry logic and HMAC signatures",
      "Zapier / n8n compatible REST endpoints",
    ],
  },
];

// ─── Architecture layers ─────────────────────────────────────────────────────

const archLayers = [
  {
    label: "Presentation",
    detail: "Next.js 14 App Router · React Server Components · Tailwind CSS",
    color: "bg-blue-500",
  },
  {
    label: "API Gateway",
    detail: "Go HTTP/2 REST API + gRPC · OpenAPI 3 spec · Rate limiting · Auth middleware",
    color: "bg-violet-500",
  },
  {
    label: "Application",
    detail: "Use-case services · Command / Query segregation · DTO mapping",
    color: "bg-accent",
  },
  {
    label: "Domain",
    detail: "Entities · Value objects · Domain events · Repository interfaces (DDD)",
    color: "bg-emerald-500",
  },
  {
    label: "Infrastructure",
    detail: "PostgreSQL · Redis · Elasticsearch · NATS messaging · S3 attachments",
    color: "bg-yellow-500",
  },
];

// ─── Quick-start steps ───────────────────────────────────────────────────────

const quickStart = [
  {
    step: "01",
    title: "Clone & configure",
    code: "git clone https://github.com/yourusername/mytodo.git\ncd mytodo\ncp .env.example .env   # fill in DB / Redis / SMTP",
  },
  {
    step: "02",
    title: "Start services",
    code: "docker compose up -d    # postgres · redis · elastic · nats\nmake migrate            # run DB migrations\nmake seed               # optional: load demo data",
  },
  {
    step: "03",
    title: "Run the API",
    code: "cd apps/api\ngo run ./cmd/server/main.go\n# → listening on :8080 (REST) and :9090 (gRPC)",
  },
  {
    step: "04",
    title: "Run the web app",
    code: "cd apps/web\nnpm install && npm run dev\n# → http://localhost:3000",
  },
];

// ─── REST API overview ───────────────────────────────────────────────────────

const apiGroups = [
  {
    tag: "Auth",
    endpoints: [
      { method: "POST", path: "/api/v1/auth/signup", desc: "Register a new account" },
      { method: "POST", path: "/api/v1/auth/login", desc: "Obtain JWT access + refresh tokens" },
      { method: "POST", path: "/api/v1/auth/refresh", desc: "Rotate refresh token" },
      { method: "DELETE", path: "/api/v1/auth/logout", desc: "Revoke session" },
    ],
  },
  {
    tag: "Issues",
    endpoints: [
      { method: "GET", path: "/api/v1/issues", desc: "List issues (pagination, filters)" },
      { method: "POST", path: "/api/v1/issues", desc: "Create issue" },
      { method: "GET", path: "/api/v1/issues/:id", desc: "Get issue detail" },
      { method: "PATCH", path: "/api/v1/issues/:id", desc: "Update issue fields" },
      { method: "DELETE", path: "/api/v1/issues/:id", desc: "Delete issue" },
    ],
  },
  {
    tag: "Projects",
    endpoints: [
      { method: "GET", path: "/api/v1/projects", desc: "List projects for org" },
      { method: "POST", path: "/api/v1/projects", desc: "Create project" },
      { method: "GET", path: "/api/v1/projects/:id", desc: "Project details & stats" },
      { method: "PATCH", path: "/api/v1/projects/:id", desc: "Update project settings" },
    ],
  },
  {
    tag: "Boards",
    endpoints: [
      { method: "GET", path: "/api/v1/boards/:id", desc: "Board columns & issues" },
      { method: "POST", path: "/api/v1/boards/:id/move", desc: "Move issue between columns" },
    ],
  },
];

// ─── Method color ────────────────────────────────────────────────────────────

function methodColor(method: string) {
  switch (method) {
    case "GET":
      return "text-emerald-600 dark:text-emerald-400";
    case "POST":
      return "text-blue-600 dark:text-blue-400";
    case "PATCH":
      return "text-yellow-600 dark:text-yellow-400";
    case "DELETE":
      return "text-red-500 dark:text-red-400";
    default:
      return "text-fg/70";
  }
}

// ─── Page ────────────────────────────────────────────────────────────────────

export default function DocsPage() {
  return (
    <div className="mx-auto flex max-w-7xl gap-10 px-4 py-12 sm:px-6 lg:px-8">
      {/* Sticky sidebar — hidden on mobile, visible on lg */}
      <aside className="hidden w-52 shrink-0 lg:block">
        <div className="sticky top-24">
          <p className="mb-3 text-xs font-semibold uppercase tracking-wider text-fg/40">
            On this page
          </p>
          <nav aria-label="Documentation sections">
            <ul className="space-y-1 text-sm">
              {sections.map((s) => (
                <li key={s.id}>
                  <a
                    href={`#${s.id}`}
                    className="block rounded-lg px-2 py-1.5 text-fg/60 transition hover:bg-accentMuted hover:text-fg"
                  >
                    {s.label}
                  </a>
                </li>
              ))}
            </ul>
          </nav>
        </div>
      </aside>

      {/* Main content */}
      <article className="min-w-0 flex-1 space-y-20">
        {/* Hero ────────────────────────────────── */}
        <section id="overview">
          <div className="mb-4 flex flex-wrap items-center gap-2">
            <Badge variant="default">v1.0 — Public Beta</Badge>
            <Badge variant="success">REST + gRPC API</Badge>
            <Badge variant="warning">Open source · MIT</Badge>
          </div>

          <h1 className="mb-4 text-4xl font-semibold tracking-tight md:text-5xl">
            MyTodo Documentation
          </h1>
          <p className="max-w-2xl text-lg leading-relaxed text-fg/70">
            MyTodo is a <strong className="font-medium text-fg">production-grade project
            management workspace</strong> — think Jira, minus the bloat. It handles
            everything from single-person task lists to multi-team enterprise sprints,
            with a clean API-first architecture that developers love.
          </p>

          <div className="mt-6 flex flex-wrap gap-3">
            <Link
              href="#getting-started"
              className="inline-flex items-center gap-2 rounded-xl bg-accent px-5 py-2.5 text-sm font-semibold text-white shadow-sm transition hover:-translate-y-0.5 hover:shadow-md"
            >
              Get started <ArrowRight size={14} />
            </Link>
            <Link
              href={ROUTES.register}
              className="inline-flex items-center gap-2 rounded-xl border border-border bg-card px-5 py-2.5 text-sm font-semibold text-fg transition hover:bg-accentMuted"
            >
              Try the app free
            </Link>
          </div>

          {/* Key metrics */}
          <div className="mt-10 grid grid-cols-2 gap-4 sm:grid-cols-4">
            {[
              { label: "Scalability", value: "5M+ users" },
              { label: "Issue capacity", value: "1B+ issues" },
              { label: "API protocols", value: "REST + gRPC" },
              { label: "Test coverage", value: "Unit / Int / E2E" },
            ].map((m) => (
              <div
                key={m.label}
                className="rounded-2xl border border-border bg-card p-4 text-center shadow-sm"
              >
                <p className="text-xl font-semibold text-fg">{m.value}</p>
                <p className="mt-0.5 text-xs text-fg/50">{m.label}</p>
              </div>
            ))}
          </div>
        </section>

        {/* Getting started ─────────────────────── */}
        <section id="getting-started">
          <SectionHeading icon={Terminal} title="Getting started" />
          <p className="mb-8 text-fg/70">
            MyTodo runs on Docker Compose for local development. All services — database,
            cache, search, and messaging — spin up with a single command.
          </p>

          <div className="grid gap-6 md:grid-cols-2">
            {quickStart.map((qs) => (
              <div
                key={qs.step}
                className="rounded-2xl border border-border bg-card p-5 shadow-sm"
              >
                <div className="mb-3 flex items-center gap-3">
                  <span className="rounded-lg bg-accentMuted px-2 py-0.5 font-mono text-xs font-bold text-accent">
                    {qs.step}
                  </span>
                  <h3 className="font-medium text-fg">{qs.title}</h3>
                </div>
                <pre className="overflow-x-auto rounded-xl bg-[var(--color-bg)] p-4 font-mono text-xs leading-relaxed text-fg/80">
                  <code>{qs.code}</code>
                </pre>
              </div>
            ))}
          </div>

          <div className="mt-6 rounded-2xl border border-border bg-card p-5 shadow-sm">
            <h3 className="mb-2 font-medium text-fg">Prerequisites</h3>
            <ul className="grid gap-1 text-sm text-fg/70 sm:grid-cols-2">
              {[
                "Go 1.21+",
                "Node.js 18+",
                "Docker & Docker Compose",
                "PostgreSQL 15+ (via Docker)",
                "Redis 7+ (via Docker)",
                "Elasticsearch 8+ (via Docker)",
              ].map((p) => (
                <li key={p} className="flex items-center gap-2">
                  <span className="h-1.5 w-1.5 rounded-full bg-accent" />
                  {p}
                </li>
              ))}
            </ul>
          </div>
        </section>

        {/* Architecture ────────────────────────── */}
        <section id="architecture">
          <SectionHeading icon={Layers} title="Architecture" />
          <p className="mb-8 text-fg/70">
            MyTodo is built on <strong className="font-medium text-fg">Clean Architecture</strong>{" "}
            and <strong className="font-medium text-fg">Domain-Driven Design (DDD)</strong>. Each
            layer has a clear responsibility and depends only on inner layers — making
            the codebase easy to test, extend, and maintain at any scale.
          </p>

          <div className="space-y-2">
            {archLayers.map((layer) => (
              <div
                key={layer.label}
                className="flex items-stretch gap-3 overflow-hidden rounded-2xl border border-border bg-card shadow-sm"
              >
                <div className={`${layer.color} w-1.5 shrink-0 rounded-l-2xl`} />
                <div className="py-4 pr-4">
                  <p className="font-medium text-fg">{layer.label}</p>
                  <p className="mt-0.5 text-sm text-fg/60">{layer.detail}</p>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-8 rounded-2xl border border-border bg-card p-6 shadow-sm">
            <h3 className="mb-4 flex items-center gap-2 font-medium text-fg">
              <CircuitBoard size={16} className="text-accent" />
              Infrastructure components
            </h3>
            <div className="grid gap-3 text-sm sm:grid-cols-2">
              {[
                { name: "PostgreSQL 15", role: "Primary relational store — all business data" },
                { name: "Redis 7", role: "Session cache, distributed locks, rate limiting" },
                { name: "Elasticsearch 8", role: "Full-text search and faceted filters" },
                { name: "NATS JetStream", role: "Async event streaming between services" },
                { name: "S3-compatible", role: "File attachments and media storage" },
                { name: "SMTP / SendGrid", role: "Transactional email delivery" },
              ].map((c) => (
                <div key={c.name} className="flex gap-3">
                  <Database size={14} className="mt-0.5 shrink-0 text-accent" />
                  <div>
                    <p className="font-medium text-fg">{c.name}</p>
                    <p className="text-fg/60">{c.role}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Features ────────────────────────────── */}
        <section id="features">
          <SectionHeading icon={Zap} title="Features" />
          <p className="mb-8 text-fg/70">
            Every feature is production-ready — not a prototype. Below is a breakdown
            of what ships out of the box.
          </p>

          <div className="grid gap-6 md:grid-cols-2">
            {featureGroups.map((group) => {
              const Icon = group.icon;
              return (
                <div
                  key={group.title}
                  className="rounded-2xl border border-border bg-card p-6 shadow-sm"
                >
                  <div className="mb-4 flex items-center gap-3">
                    <span className="flex h-9 w-9 items-center justify-center rounded-xl bg-accentMuted text-accent">
                      <Icon size={18} />
                    </span>
                    <h3 className="font-medium text-fg">{group.title}</h3>
                  </div>
                  <ul className="space-y-2 text-sm text-fg/70">
                    {group.items.map((item) => (
                      <li key={item} className="flex gap-2">
                        <span className="mt-1.5 h-1.5 w-1.5 shrink-0 rounded-full bg-accent" />
                        {item}
                      </li>
                    ))}
                  </ul>
                </div>
              );
            })}
          </div>
        </section>

        {/* API Reference ──────────────────────── */}
        <section id="api">
          <SectionHeading icon={Code2} title="API reference" />
          <p className="mb-2 text-fg/70">
            All endpoints follow REST conventions and return JSON. Authentication uses a
            Bearer token in the{" "}
            <code className="rounded bg-accentMuted px-1.5 py-0.5 font-mono text-xs text-accent">
              Authorization
            </code>{" "}
            header.
          </p>
          <p className="mb-8 text-sm text-fg/50">
            Base URL:{" "}
            <code className="rounded bg-accentMuted px-1.5 py-0.5 font-mono text-xs text-accent">
              https://api.yourdomain.com
            </code>
          </p>

          <div className="space-y-6">
            {apiGroups.map((group) => (
              <div key={group.tag} className="rounded-2xl border border-border bg-card shadow-sm">
                <div className="border-b border-border px-5 py-3">
                  <span className="rounded-lg bg-accentMuted px-2 py-0.5 text-xs font-bold text-accent">
                    {group.tag}
                  </span>
                </div>
                <div className="divide-y divide-border">
                  {group.endpoints.map((ep) => (
                    <div
                      key={`${ep.method}-${ep.path}`}
                      className="flex flex-wrap items-center gap-x-4 gap-y-1 px-5 py-3 text-sm"
                    >
                      <span
                        className={`w-14 shrink-0 font-mono font-bold ${methodColor(ep.method)}`}
                      >
                        {ep.method}
                      </span>
                      <code className="font-mono text-xs text-fg/80">{ep.path}</code>
                      <span className="text-fg/50">{ep.desc}</span>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>

          <p className="mt-6 text-sm text-fg/60">
            Full OpenAPI 3 specification available at{" "}
            <code className="rounded bg-accentMuted px-1.5 py-0.5 font-mono text-xs text-accent">
              /docs/api/openapi.yaml
            </code>{" "}
            in the repository.
          </p>
        </section>

        {/* Authentication ─────────────────────── */}
        <section id="auth">
          <SectionHeading icon={ShieldCheck} title="Authentication" />
          <p className="mb-6 text-fg/70">
            MyTodo uses a dual-token strategy for secure, stateless authentication.
          </p>

          <div className="grid gap-4 sm:grid-cols-2">
            {[
              {
                title: "Access token",
                detail:
                  "Short-lived JWT (15 min). Sent as Authorization: Bearer <token> with every API request. Contains user ID, org ID, and role claims.",
              },
              {
                title: "Refresh token",
                detail:
                  "Long-lived opaque token (30 days). Stored in an HttpOnly cookie. Used to silently obtain a new access token via POST /auth/refresh.",
              },
              {
                title: "OAuth 2.0",
                detail:
                  "Google and GitHub sign-in via authorization-code flow. Social accounts are linked to the same user record as password-based logins.",
              },
              {
                title: "RBAC",
                detail:
                  "Role-Based Access Control at org and project level. Roles: Owner > Admin > Member > Viewer. Enforced in middleware before reaching use-case services.",
              },
            ].map((c) => (
              <div key={c.title} className="rounded-2xl border border-border bg-card p-5 shadow-sm">
                <h3 className="mb-2 font-medium text-fg">{c.title}</h3>
                <p className="text-sm text-fg/60">{c.detail}</p>
              </div>
            ))}
          </div>
        </section>

        {/* Self-hosting ──────────────────────── */}
        <section id="self-hosting">
          <SectionHeading icon={GitBranch} title="Self-hosting" />
          <p className="mb-6 text-fg/70">
            MyTodo ships with first-class self-hosting support. Choose the deployment
            strategy that fits your infrastructure.
          </p>

          <div className="grid gap-4 sm:grid-cols-3">
            {[
              {
                title: "Docker Compose",
                desc: "Best for small teams and local evaluation. All services in one compose file.",
                cmd: "docker compose -f docker-compose.prod.yml up -d",
              },
              {
                title: "Kubernetes",
                desc: "Production-grade. Helm chart with HPA, PDB, and Ingress included.",
                cmd: "helm install mytodo ./infrastructure/kubernetes",
              },
              {
                title: "Terraform / AWS",
                desc: "Managed cloud deployment on ECS / RDS / ElastiCache / OpenSearch.",
                cmd: "cd infrastructure/terraform/aws && terraform apply",
              },
            ].map((opt) => (
              <div key={opt.title} className="rounded-2xl border border-border bg-card p-5 shadow-sm">
                <h3 className="mb-2 font-medium text-fg">{opt.title}</h3>
                <p className="mb-3 text-sm text-fg/60">{opt.desc}</p>
                <pre className="overflow-x-auto rounded-xl bg-[var(--color-bg)] p-3 font-mono text-xs text-fg/70">
                  <code>{opt.cmd}</code>
                </pre>
              </div>
            ))}
          </div>
        </section>

        {/* SDK & Integrations ─────────────────── */}
        <section id="sdk">
          <SectionHeading icon={Puzzle} title="SDK & integrations" />
          <p className="mb-6 text-fg/70">
            A typed TypeScript SDK ships in{" "}
            <code className="rounded bg-accentMuted px-1.5 py-0.5 font-mono text-xs text-accent">
              packages/sdk
            </code>
            . Install it in any JS/TS project:
          </p>

          <pre className="mb-6 overflow-x-auto rounded-2xl border border-border bg-card p-5 font-mono text-sm text-fg/80">
            <code>{`import { MyTodoClient } from "@mytodo/sdk";

const client = new MyTodoClient({ baseUrl: "https://api.yourdomain.com" });

// Authenticate
const { accessToken } = await client.auth.login({ email, password });

// Create an issue
const issue = await client.issues.create({
  projectId: "proj_abc123",
  title: "Fix login redirect on mobile",
  priority: "high",
  assigneeId: "usr_xyz789",
});`}</code>
          </pre>

          <div className="grid gap-4 sm:grid-cols-2">
            {[
              {
                title: "Webhooks",
                desc: "Subscribe to any domain event (issue.created, sprint.completed, …). Payloads are signed with HMAC-SHA256.",
              },
              {
                title: "GitHub integration",
                desc: "Link pull requests and commits to issues. Status transitions on PR merge.",
              },
              {
                title: "Slack integration",
                desc: "Post issue updates to Slack channels. Use /mytodo slash commands.",
              },
              {
                title: "gRPC API",
                desc: "High-throughput binary protocol for service-to-service communication. Proto definitions in apps/api/proto/.",
              },
            ].map((item) => (
              <div key={item.title} className="rounded-2xl border border-border bg-card p-5 shadow-sm">
                <h3 className="mb-1 font-medium text-fg">{item.title}</h3>
                <p className="text-sm text-fg/60">{item.desc}</p>
              </div>
            ))}
          </div>
        </section>

        {/* Roadmap ────────────────────────────── */}
        <section id="roadmap">
          <SectionHeading icon={BookOpen} title="Roadmap" />
          <p className="mb-6 text-fg/70">
            MyTodo is under active development. Upcoming features:
          </p>

          <div className="grid gap-3 sm:grid-cols-2">
            {[
              { label: "AI-powered issue summarization", status: "planned" },
              { label: "Time tracking & reports", status: "planned" },
              { label: "Gantt / timeline view", status: "planned" },
              { label: "Mobile apps (iOS & Android)", status: "planned" },
              { label: "Self-hosted SSO (SAML / OIDC)", status: "planned" },
              { label: "Two-factor authentication (TOTP)", status: "in-progress" },
              { label: "Advanced audit logs + SIEM export", status: "in-progress" },
              { label: "Custom fields on issues", status: "in-progress" },
            ].map((item) => (
              <div
                key={item.label}
                className="flex items-center justify-between gap-3 rounded-2xl border border-border bg-card px-4 py-3 shadow-sm"
              >
                <span className="text-sm text-fg/80">{item.label}</span>
                <Badge variant={item.status === "in-progress" ? "warning" : "default"}>
                  {item.status === "in-progress" ? "In progress" : "Planned"}
                </Badge>
              </div>
            ))}
          </div>

          <div className="mt-8 rounded-2xl border border-border bg-accentMuted/30 p-6 text-center">
            <p className="text-sm text-fg/70">
              Missing a feature?{" "}
              <Link
                href={ROUTES.register}
                className="font-semibold text-accent hover:underline"
              >
                Sign up and let us know
              </Link>{" "}
              — we prioritize based on real feedback.
            </p>
          </div>
        </section>
      </article>
    </div>
  );
}

// ─── Section heading helper ──────────────────────────────────────────────────

function SectionHeading({
  icon: Icon,
  title,
}: {
  icon: LucideIcon;
  title: string;
}) {
  return (
    <div className="mb-6 flex items-center gap-3">
      <span className="flex h-9 w-9 items-center justify-center rounded-xl bg-accentMuted text-accent">
        <Icon size={18} />
      </span>
      <h2 className="text-2xl font-semibold text-fg">{title}</h2>
    </div>
  );
}
