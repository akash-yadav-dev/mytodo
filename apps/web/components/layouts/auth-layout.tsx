import Link from "next/link";
import { APP_NAME } from "@/lib/constants/config";

export function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-bg text-fg">
      <div className="mx-auto flex min-h-screen max-w-4xl flex-col justify-center gap-8 px-6 py-12">
        <div className="flex items-center justify-between">
          <Link href="/" className="text-lg font-semibold">
            {APP_NAME}
          </Link>
          <span className="text-xs uppercase tracking-[0.3em] text-fg/60">
            Secure access
          </span>
        </div>
        <div className="rounded-3xl border border-border bg-card p-8 shadow-soft">
          {children}
        </div>
      </div>
    </div>
  );
}
