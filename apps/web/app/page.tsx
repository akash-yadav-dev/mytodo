import { Button } from "@/components/ui/button";

export default function HomePage() {
  return (
    <main className="min-h-screen bg-bg text-fg">
      <div className="relative overflow-hidden">
        <div className="absolute inset-0 -z-10 bg-[radial-gradient(circle_at_top,_rgba(251,191,36,0.18),_transparent_40%),radial-gradient(circle_at_10%_20%,_rgba(59,130,246,0.16),_transparent_50%)]" />
        <section className="mx-auto flex max-w-5xl flex-col gap-8 px-6 py-20">
          <div className="flex flex-col gap-4">
            <span className="w-fit rounded-full border border-border bg-card px-3 py-1 text-xs uppercase tracking-[0.2em] text-fg/70">
              MyTodo Workspace
            </span>
            <h1 className="text-4xl font-semibold leading-tight md:text-5xl">
              Keep projects steady, tasks simple, and progress visible.
            </h1>
            <p className="max-w-2xl text-base text-fg/80 md:text-lg">
              A clean base that gives you room to grow. Add boards, issues, and
              workflows on your schedule.
            </p>
          </div>
          <div className="flex flex-wrap gap-3">
            <Button asChild href="/auth/register">
              Create account
            </Button>
            <Button asChild href="/auth/login" variant="ghost">
              Sign in
            </Button>
          </div>
          <div className="grid gap-4 md:grid-cols-3">
            {[
              "Project clarity",
              "Focused execution",
              "Quiet notifications",
            ].map((item) => (
              <div
                key={item}
                className="rounded-2xl border border-border bg-card p-5 shadow-soft"
              >
                <h3 className="text-lg font-medium">{item}</h3>
                <p className="mt-2 text-sm text-fg/70">
                  Add only what you need. Everything else stays out of the way.
                </p>
              </div>
            ))}
          </div>
        </section>
      </div>
    </main>
  );
}
