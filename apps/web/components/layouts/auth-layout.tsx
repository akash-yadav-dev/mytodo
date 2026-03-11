export function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-screen items-center justify-center bg-bg px-4">
      <div className="w-full max-w-sm">
        <div className="mb-8 text-center">
          <div className="mx-auto mb-3 flex h-10 w-10 items-center justify-center rounded-xl bg-accent text-sm font-bold text-white">
            MT
          </div>
          <p className="text-xs text-fg/50">MyTodo — manage smarter</p>
        </div>
        <div className="rounded-2xl border border-border bg-card p-8 shadow-soft">
          {children}
        </div>
      </div>
    </div>
  );
}
