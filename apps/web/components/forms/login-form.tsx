"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useLogin } from "@/hooks/useAuth/useLogin";
import { ROUTES } from "@/lib/constants/routes";

export function LoginForm() {
  const router = useRouter();
  const { submit, loading, error } = useLogin();
  const [form, setForm] = useState({ email: "", password: "" });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const ok = await submit(form);
    if (ok) router.push(ROUTES.dashboard);
  };

  return (
    <div className="space-y-6">
      <div className="space-y-1">
        <h1 className="text-2xl font-semibold">Welcome back</h1>
        <p className="text-sm text-fg/60">Sign in to keep building momentum.</p>
      </div>

      <form className="space-y-4" onSubmit={handleSubmit}>
        <Input
          type="email"
          placeholder="Email address"
          autoComplete="email"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.target.value })}
          required
        />
        <div className="space-y-1">
          <Input
            type="password"
            placeholder="Password"
            autoComplete="current-password"
            value={form.password}
            onChange={(e) => setForm({ ...form, password: e.target.value })}
            required
          />
          <div className="text-right">
            <Link href={ROUTES.forgotPassword} className="text-xs text-fg/50 hover:text-fg">
              Forgot password?
            </Link>
          </div>
        </div>

        {error && (
          <p className="rounded-lg border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-600">
            {error}
          </p>
        )}

        <Button className="w-full" type="submit" disabled={loading}>
          {loading ? "Signing in…" : "Sign in"}
        </Button>
      </form>

      <p className="text-center text-sm text-fg/60">
        No account yet?{" "}
        <Link href={ROUTES.register} className="font-medium text-fg hover:text-accent">
          Create one free
        </Link>
      </p>
    </div>
  );
}
